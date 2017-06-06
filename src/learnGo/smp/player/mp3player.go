package player

import (
	"fmt"
	"time"
)

type MP3Player struct {
	stat     int
	progress int
}

func (player *MP3Player) Play(source string) {
	fmt.Println("Playing mp3 music", source)
	player.progress = 0

	for player.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(".")
		player.progress += 10
	}
	fmt.Println("\nFinished playing", source)
}
