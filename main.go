package main

import (
	"Gator/Internal/Config"
	"fmt"
)

func main() {
	cfg, err := Config.Read()
	if err != nil {
		fmt.Println(err)
	}

	cfg.SetUser("Strus")

	cfg, err = Config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg)
}
