package cg

import (
	"encoding/json"
	"errors"
	"fmt"
	"learnGo/chatroom/src/ipc"
)

type CenterClient struct {
	*ipc.IpcClient
}

func (client *CenterClient) AddPlayer(player *Player) error {
	b, err := json.Marshal(*player)
	if err != nil {
		return err
	}
	resp, err := client.Call("addplayer", string(b))
	fmt.Println(resp)

	if err == nil && resp.Code == "200" {
		return nil
	}
	return err
}

func (client *CenterClient) RemovePlayer(name string) error {
	ret, _ := client.Call("removeplayer", name)
	if ret.Code == "200" {

		return nil
	}
	return errors.New(ret.Code)
}
func (clent *CenterClient) ListPlayer(params string) (ps []*Player, err error) {
	fmt.Println("list players")
	resp, _ := clent.Call("listplayer", params)
	if resp.Code != "200" {
		err = errors.New(resp.Code)
		return
	}
	err = json.Unmarshal([]byte(resp.Body), &ps)
	return
}
func (client *CenterClient) Boardcast(message string) error {
	m := &Message{Content: message}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	resp, _ := client.Call("boardcast", string(b))
	if resp.Code == "200" {
		return nil
	}
	return errors.New(resp.Code)
}
