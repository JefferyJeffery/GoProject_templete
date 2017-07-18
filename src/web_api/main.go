package main

import (
	"config"
	"fmt"
)

func main() {
	fmt.Println(
		"i,m backend/main",
	)

	config.Run()
}
