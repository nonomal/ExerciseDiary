package main

import (
	"flag"

	_ "time/tzdata"

	"github.com/aceberg/ExerciseDiary/internal/check"
	"github.com/aceberg/ExerciseDiary/internal/web"
)

const confPath = "/data/ExerciseDiary/config.yaml"
const yamlPath = "/data/ExerciseDiary/board.yaml"
const nodePath = ""

func main() {
	yamlPtr := flag.String("b", yamlPath, "Path to board yaml file")
	confPtr := flag.String("c", confPath, "Path to config yaml file")
	nodePtr := flag.String("n", nodePath, "Path to node modules")
	flag.Parse()

	check.Path(*confPtr)
	check.Path(*yamlPtr)

	web.Gui(*confPtr, *yamlPtr, *nodePtr) // webgui.go
}