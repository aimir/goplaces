package main

import (
	"log"
	"github.com/aimir/goplaces/database"
	"github.com/aimir/goplaces/config"
	"github.com/spf13/viper"
	"github.com/aimir/goplaces/api"
)

func main() {
	//Init configuration
	config.InitConfig()
	//Create database connection
	err := database.OpenDB(viper.GetString("database"))
	if err != nil {
		log.Fatal(err)
	}
	defer database.GetDB().Close()

	log.Fatal(api.StartAPI())
}