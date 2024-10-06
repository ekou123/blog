package main

import (
	"fmt"
	"github.com/ekou123/blog/internal/config"
)

func main() {
	cfg := config.Config{
		User: "Ethan",
	}

	err := config.SetUser(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	newConfig, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(newConfig)

}
