package player

import "fmt"

type Player interface {
	Play(source string)
}

func Play(source, musicType string) {

	var player Player
	switch musicType {
	case "mp3":
		player = &MP3Player{}
	default:
		fmt.Println("Unsupported music type", musicType)
	}
	player.Play(source)
}
