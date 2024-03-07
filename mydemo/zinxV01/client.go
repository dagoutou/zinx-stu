package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	dial, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("dial error", err)
		return
	}
	for {
		_, err := dial.Write([]byte("hello zinx v0.1"))
		if err != nil {
			fmt.Println("Write error", err)
			return
		}
		buf := make([]byte, 512)
		_, err = dial.Read(buf)
		if err != nil {
			fmt.Println("Read error", err)
			return
		}
		fmt.Printf("server call back:%s\n", buf)
		time.Sleep(1 * time.Second)
	}

}
