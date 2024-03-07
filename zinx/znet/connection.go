package znet

import (
	"fmt"
	"net"
	"zinx/zinx/ziface"
)

type Connection struct {
	Conn      *net.TCPConn
	ConnID    uint32
	IsClosed  bool
	HandleAPI ziface.HandleFunc
	ExitChan  chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, callbackApi ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		IsClosed:  false,
		HandleAPI: callbackApi,
		ExitChan:  make(chan bool, 1),
	}
	return c
}
func (c *Connection) StartReader() {
	fmt.Println("reader goroutine is running")
	defer c.Stop()
	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("reciver buf err", err)
			continue
		}
		if err = c.HandleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("connID ", c.ConnID, "handle is error", err)
			break
		}
	}
}
func (c *Connection) Start() {
	fmt.Println("conn start() connID = ", c.ConnID)
	go c.StartReader()
}
func (c *Connection) Stop() {
	fmt.Println("conn stop() connID = ", c.ConnID)
	if c.IsClosed {
		return
	}
	c.IsClosed = true
	c.Conn.Close()
	close(c.ExitChan)
}
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
func (c *Connection) Send(data []byte) error {
	return nil
}
