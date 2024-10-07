package main

import (
	"database/sql"
	"fmt"
	"github.com/ekou123/blog/internal/config"
	"github.com/ekou123/blog/internal/database"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Println(cfg)

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

	db, err := sql.Open("postgres", programState.Cfg.DbURL)
	if err != nil {
		fmt.Println("Unable to open sql connection")
		return
	}

	dbQueries := database.New(db)
}
