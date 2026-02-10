package main

import (
	"Gator/internal/commands"
	"Gator/internal/config"
	"Gator/internal/database"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	db, err := sql.Open("postgres", cfg.Db_URL)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	DB_Queries := database.New(db)
	App_State := &commands.State{
		DB:         DB_Queries,
		App_Config: &cfg,
	}

	Cmds := commands.Commands{
		Cmd_List: make(map[string]func(*commands.State, commands.Command) error),
	}

	Cmds.Register("login", commands.Handler_login)
	Cmds.Register("register", commands.Handler_register)
	Cmds.Register("reset", commands.Handler_reset)
	Cmds.Register("users", commands.Handler_users)
	Cmds.Register("agg", commands.Handler_agg)
	Cmds.Register("addfeed", commands.Handler_add_feed)
	Cmds.Register("feeds", commands.Handler_feeds)
	Cmds.Register("follow", commands.Handler_follow)
	Cmds.Register("following", commands.Handler_following)

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		os.Exit(1)
	}
	Cmd_Name := os.Args[1]
	Cmd_Args := os.Args[2:]
	User_Cmd := commands.Command{
		Name: Cmd_Name,
		Args: Cmd_Args,
	}

	err = Cmds.Run(App_State, User_Cmd)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

}
