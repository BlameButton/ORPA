package main

import (
	"github.com/blamebutton/orpa/gui"
	"github.com/blamebutton/orpa/parser"
	"log"
)

func main() {
	replay, e := parser.GetFileReplay("replays/featurive - Drop - Granat [Yuzu's Insane] (2017-05-14) Osu.osr")
	if e != nil {
		log.Fatal(e)
	}
	window := gui.NewOrpaWindow(replay)
	window.Start()
}
