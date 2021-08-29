package main

import "net"

type User struct {
	Name    string      `json:"name"`
	Address string      `json:"address"`
	C       chan string `json:"channel"`
	conn    net.Conn    `json:"connect"`
}

// 创建用户
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:    userAddr,
		Address: userAddr,
		C:       make(chan string),
		conn:    conn,
	}

	// 启动监听当前 user channel 消息的 goroutine
	go user.ListenMessage()

	return user
}

// 监听当前用户的 channel，一旦有消息，发送给对端客户端
func (u *User) ListenMessage() {
	for {
		msg := <-u.C
		u.conn.Write([]byte(msg + "\n"))
	}
}
