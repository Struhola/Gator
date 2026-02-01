package main

import (
	"Gator/Internal/Commands"
	"Gator/Internal/Config"
	"fmt"
	"os"
)

func main() {
	cfg, err := Config.Read()
	if err != nil {
		fmt.Println(err)
	}
	App_State := &Commands.State{
		App_Config: &cfg,
	}
	Cmds := Commands.Commands{
		Cmd_List: make(map[string]func(*Commands.State, Commands.Command) error),
	}

	Cmds.Register("login", Commands.Handler_Login)

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		os.Exit(1)
	}
	Cmd_Name := os.Args[1]
	Cmd_Args := os.Args[2:]
	User_Cmd := Commands.Command{
		Name: Cmd_Name,
		Args: Cmd_Args,
	}

	err = Cmds.Run(App_State, User_Cmd)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
