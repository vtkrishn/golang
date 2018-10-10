package main

import (
	"fmt"
	"time"
)


func say(s string){
	for i := 0; i<10;i++ {
		fmt.Println(s, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func main(){
	go say("Hi")
	say("Helo")
}
