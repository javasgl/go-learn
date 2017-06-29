package main

import (
	_ "learnGo/searcher/src/matchers"
	"learnGo/searcher/src/search"
	"log"
	"os"
)

func main() {

}

func init() {
	log.SetOutput(os.Stdout)
	search.Run("run")
}
