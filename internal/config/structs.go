package config

import (
	"Gator/internal/database"
)

type Config struct {
	Db_URL            string `json:"db_url"`
	Current_User_Name string `json:"current_user_name"`
}

type State struct {
	DB         *database.Queries
	App_Config *Config
}
