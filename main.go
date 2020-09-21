package main

import (
	"github.com/tnqv/todo-app/app"
	"github.com/tnqv/todo-app/config"
	"github.com/kelseyhightower/envconfig"
	"fmt"
	"log"
)

func main() {

	var s config.Config
	err := envconfig.Process("myapp", &s)
	if err != nil {
			log.Fatal(err.Error())
	}

	format := "Debug: %v\nHOST: %d\nUser: %s\nPASSWORD: %f\nName: %s\n"
	_, err = fmt.Printf(format, s.DB.Host, s.DB.Username, s.DB.Password, s.DB.Name)

	// config := config.GetConfig()

	app := &app.App{}
	app.Initialize(&s)
	app.Run(":3000")
}
