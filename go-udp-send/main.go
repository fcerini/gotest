package main

import (
	"fmt"
	"time"
	"net"
	"os"
	"strconv"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Especificar el puerto")
		return
	}
	puerto, _ := strconv.Atoi(arguments[1])

	Conn, _ := net.DialUDP("udp", nil, &net.UDPAddr{IP: []byte{127, 0, 0, 1}, Port: puerto, Zone: ""})
	defer Conn.Close()
	for index := 0; index < 10000; index++ {
		Conn.Write([]byte("TEST"))
		time.Sleep( 1 * time.Millisecond )
	}
	
}
