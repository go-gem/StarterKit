package main

import (
	"flag"
	"log"

	"github.com/go-gem/StarterKit/src/basic/controllers"
	"github.com/go-gem/gem"
)

var (
	app *gem.Application
)

func main() {
	var err error

	var config string
	flag.StringVar(&config, "c", "", "config")
	flag.Parse()
	if config == "" {
		panic("no configuration file specified")
	}

	// initialize application
	if app, err = gem.NewApplication(config); err != nil {
		panic("failed to create an application: " + err.Error())
	}
	if err = app.Init(); err != nil {
		panic("failed to initialize application: " + err.Error())
	}

	// close application
	defer app.Close()

	// register controllers
	registerControllers()

	// start server
	if app.ServerOpt.CertFile != "" || app.ServerOpt.KeyFile != "" {
		log.Fatal(gem.ListenAndServeTLS(app.ServerOpt.Addr, app.ServerOpt.CertFile, app.ServerOpt.KeyFile, app.Router().Handler()))
		return
	}

	log.Fatal(gem.ListenAndServe(app.ServerOpt.Addr, app.Router().Handler()))
}

func registerControllers() {
	var err error

	// create base controller
	c := controllers.Controller{}
	if err = c.Init(app); err != nil {
		panic(err)
	}

	app.SetController("/", &controllers.Index{Controller: c})

	if err = app.InitControllers(); err != nil {
		panic("failed to initialize controllers: " + err.Error())
	}
}
