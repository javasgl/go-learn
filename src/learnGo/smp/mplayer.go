package main

import (
	"bufio"
	"fmt"
	"learnGo/smp/lib"
	"learnGo/smp/player"
	"os"
	"strconv"
	"strings"
)

var manager *lib.MusicManager
var id int = 1
var ctrl, signal chan int

func main() {

	fmt.Println(`

			Enter following commands to control the player:
			lib list --View the existing music lib
			lib add <name><author><source><type> --Add a music to the music lib
			lit remove <name> --Remove the specified music form the lib
			play <name> --Play the specified music

	`)
	manager = lib.NewMusicManager()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter command->")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommands(tokens)
		} else {
			fmt.Println("Unrecongnized command:", tokens[0])
		}
	}
}

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < manager.Len(); i++ {
			music, _ := manager.Get(i)
			fmt.Println(i+1, ":", music.Name, music.Author, music.Source, music.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			manager.Add(&lib.Music{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]})

		} else {
			fmt.Println("Usage: lib add <name><author><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			manager.Remove(tokens[2])
		} else {
			fmt.Println("Usage:lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}

func handlePlayCommands(tokens []string) {

	if len(tokens) != 2 {
		fmt.Println("Usage:play <name>")
		return
	}
	m := manager.Find(tokens[1])
	if m == nil {
		fmt.Println("The Music", tokens[1], " does not exists!")
		return
	}
	player.Play(m.Source, m.Type)
}
