package main

import "zinx/zinx/znet"

func main() {
	s := znet.NewServer("zin-v0.2")
	s.Serve()
}
