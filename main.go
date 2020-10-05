package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/tnqv/todo-app/app"
	"github.com/tnqv/todo-app/config"
	"log"
)

func main() {

	var s config.Config
	err := envconfig.Process("myapp", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	format := "Debug: HOST: %s\nUser: %s\nPASSWORD: %s\n NAME: %s\n"
	fmt.Printf(format, s.DB.Host, s.DB.Username, s.DB.Password, s.DB.Name)

	appInit := &app.App{}
	appInit.Initialize(&s)
	appInit.Run(":3000")
}
