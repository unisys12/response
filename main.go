package main

import (
	"github.com/responserms/response/cmd"
	"math/rand"
	"time"
)

func main() {
	// seed random math
	rand.Seed(time.Now().UnixNano())

	cmd.Execute()
}
