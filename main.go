package main

import (
	"fmt"
	"github.com/ekou123/blog/internal/config"
	"log"
	"os"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	err = cfg.SetUser("Benson")
	if err != nil {
		log.Fatalf("error setting Benson user: %v", err)
	}

	programState := &config.State{
		Cfg: &cfg,
	}

	cmds := config.Commands{
		Handler: make(map[string]func(*config.State, config.Command) error),
	}

	cmds.Register("login", config.HandlerLogin)

	programState.Cfg.SetUser("HelloWorld")

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.Run(programState, config.Command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
