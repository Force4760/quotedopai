package main

import (
	"math/rand"
	"time"

	"github.com/force4760/quotedopai/cmd"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	cmd.Execute()
}
