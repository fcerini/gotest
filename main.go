package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v7"
)

func main() {
	fmt.Println("Para probar...")
	fmt.Println("$redis-cli")
	fmt.Println(">LPUSH queue zzzz")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("PING: %s.\n", pong)

	/*
	ReceiveMessage es online pero no acumula mensajes...
	subs := client.Subscribe("ch1")
	fmt.Println("$redis-cli")

	for {
		msg, _ := subs.ReceiveMessage()
		go func() {
			fmt.Println(msg.Payload)
			time.Sleep(5 * time.Second)
			fmt.Println(".")
		}()
	}
	*/
	lastErr :=""

	for {
		result, err := client.BRPop(10 * time.Second, "queue").Result()
		
		if err != nil {
			if err.Error()!= lastErr {
				fmt.Println(err.Error())
				lastErr = err.Error()
			}
			time.Sleep(100 * time.Millisecond)				
			
		} else {
			go func() {
				fmt.Println(result[1])
				time.Sleep(5 * time.Second)
				fmt.Println(".")
			}()	
		}
	}


}
