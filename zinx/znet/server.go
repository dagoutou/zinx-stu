package znet

import (
	"fmt"
	"net"
	"zinx/zinx/ziface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

// CallBackToClient 可由用户实现
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("conn handel callback client")
	_, err := conn.Write(data[:cnt])
	if err != nil {
		return err
	}
	return nil
}
func (s *Server) Start() {
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve error", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen error", err)
			return
		}
		var cid uint32
		cid = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept error", err)
				continue
			}
			Conn := NewConnection(conn, cid, CallBackToClient)
			cid++
			go Conn.Start()
			//go func() {
			//	for {
			//		buf := make([]byte, 512)
			//		n, err := conn.Read(buf)
			//		if err != nil {
			//			fmt.Println("read error", err)
			//			continue
			//		}
			//		fmt.Printf("recev client: %s\n", string(buf[:]))
			//		_, err = conn.Write(buf[:n])
			//		if err != nil {
			//			fmt.Println("Write error", err)
			//			continue
			//		}
			//	}
			//}()

		}
	}()
}
func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()
	select {}
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
