package main

import (
	"fmt"
	"time"
	"log"
	"net"
	"os"

	"github.com/go-redis/redis/v7"
)

var redisClient *redis.Client

func main() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Conectado a Redis: ", redisClient.Options().Addr)

	sip, err := net.ListenPacket("udp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	defer sip.Close()
	fmt.Println("Escuchando SIP: ", sip.LocalAddr())

	rdp, err := net.ListenPacket("udp", "localhost:9001")
	if err != nil {
		log.Fatal(err)
	}
	defer rdp.Close()
	fmt.Println("Escuchando RDP: ", rdp.LocalAddr())

	go listenSIP( sip )

	listenRDP(rdp)
}

func listenSIP(sip net.PacketConn) {	
	for {
		buf := make([]byte, 1024)
		n, addr, err := sip.ReadFrom(buf)
		if err != nil {
			continue
		}
		go saveSIP(sip, addr, buf[:n])
	}
}

func saveSIP(sip net.PacketConn, addr net.Addr, buf []byte) {
	time.Sleep(5 * time.Millisecond)
	redisClient.RPush("SIP", buf)
}


func listenRDP(rdp net.PacketConn) {
	for {
		buf := make([]byte, 1024)
		n, addr, err := rdp.ReadFrom(buf)
		if err != nil {
			continue
		}
		go saveRDP(rdp, addr, buf[:n])
	}
}

func saveRDP(rdp net.PacketConn, addr net.Addr, buf []byte) {
	time.Sleep(5 * time.Millisecond)
	redisClient.RPush("RDP", buf)
}
