package main

import (
	"net"
	"strings"
)

type User struct {
	Name    string      `json:"name"`
	Address string      `json:"address"`
	C       chan string `json:"channel"`
	conn    net.Conn    `json:"connect"`
	server  *Server     `json"server"`
}

// 创建用户
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:    userAddr,
		Address: userAddr,
		C:       make(chan string),
		conn:    conn,
		server:  server,
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

// 用户上线
func (u *User) Online() {
	// 用户加入到 OnlineMap
	u.server.mapLock.Lock()
	u.server.OnlineMap[u.Name] = u
	u.server.mapLock.Unlock()
	// 广播用户上线消息
	u.server.BoardCase(u, "已上线")
}

// 用户下线
func (u *User) Offline() {
	// 用户从 OnlineMap 删除
	u.server.mapLock.Lock()
	delete(u.server.OnlineMap, u.Name)
	u.server.mapLock.Unlock()
	// 广播用户上线消息
	u.server.BoardCase(u, "已下线")
}

// 指定用户发送消息
func (u *User) SendMessage(msg string) {
	u.conn.Write([]byte(msg))
}

// 用户处理消息
func (u *User) DoMessage(msg string) {

	if msg == "who" {
		// 查询当前用户
		u.server.mapLock.Lock()

		for _, user := range u.server.OnlineMap {
			onlineMsg := "[" + user.Address + "]" + user.Name + ": 在线...\n"
			u.SendMessage(onlineMsg)
		}

		u.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 消息格式 'rename|章三'
		newName := strings.Split(msg, "|")[1]

		_, ok := u.server.OnlineMap[newName]

		if ok {
			u.SendMessage("当前用户名" + newName + "已被使用\n")
		} else {
			u.server.mapLock.Lock()
			delete(u.server.OnlineMap, u.Name)
			u.Name = newName
			u.server.OnlineMap[newName] = u
			u.server.mapLock.Unlock()

			u.SendMessage("您已经更新用户名为" + newName)
		}

	} else {
		u.server.BoardCase(u, msg)
	}

}
