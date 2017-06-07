package cg

import "fmt"

type Player struct {
	Name  string "name"
	Level int    "level"
	Exp   int    "exp"
	mq    chan *Message
}

func NewPlayer() *Player {
	mq := make(chan *Message, 1024)
	player := &Player{"", 0, 0, mq}
	go func(player *Player) {
		for {
			msg := <-player.mq
			fmt.Println(player.Name, " Recevied Message:", msg.Content)
		}

	}(player)
	return player
}
