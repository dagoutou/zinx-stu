package main

import "zinx/zinx/znet"

func main() {
	s := znet.NewServer("zin-v0.1")
	s.Serve()
}
