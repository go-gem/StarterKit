package main

import (
	"flag"
	"log"

	"github.com/go-gem/StarterKit/src/advanced/frontend/controllers"
	"github.com/go-gem/gem"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
		panic("invalid config file: " + config)
	}

	if app, err = gem.NewApplication(config); err != nil {
		panic("failed to create an application: " + err.Error())
	}

	if err = app.Init(); err != nil {
		panic("failed to initialize application: " + err.Error())
	}

	defer app.Close()

	registerComponents()

	registerControllers()

	if err = app.InitControllers(); err != nil {
		panic("failed to initialize controllers")
	}

	if app.ServerOpt.CertFile != "" || app.ServerOpt.KeyFile != "" {
		log.Fatal(gem.ListenAndServeTLS(app.ServerOpt.Addr, app.ServerOpt.CertFile, app.ServerOpt.KeyFile, app.Router().Handler()))
		return
	}

	log.Fatal(gem.ListenAndServe(app.ServerOpt.Addr, app.Router().Handler()))
}

func registerComponents() {
	var err error

	// session store component
	sessionsStore := sessions.NewCookieStore([]byte("something-very-secret"))
	if err = app.SetComponent("sessionsStore", sessionsStore); err != nil {
		panic(err)
	}

	// database component
	db, err := gorm.Open("mysql", "root:123456@/gem?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
	// enable log mode
	db.LogMode(true)
	// close database after closing application.
	app.SetCloseCallback(db.Close)
	if err = app.SetComponent("db", db); err != nil {
		panic(err)
	}
}

func registerControllers() {
	var err error

	c := controllers.Controller{}
	// Initialize the base controller, set component(such as database).
	if err = c.Init(app); err != nil {
		panic(err)
	}

	// homepage's controller
	app.SetController("/", &controllers.Index{Controller: c})

	// user's controller
	app.SetController("/join", &controllers.UserJoin{Controller: c})
	app.SetController("/login", &controllers.UserLogin{Controller: c})
	app.SetController("/logout", &controllers.UserLogout{Controller: c})
	app.SetController("/user/:name", &controllers.UserProfile{Controller: c})

	// post's controller
	postController := &controllers.Post{Controller: c}
	app.SetController("/posts", postController)
	app.SetController("/posts/:page", postController)
	app.SetController("/post/:id", &controllers.PostDetail{Controller: c})
	app.SetController("/new", &controllers.PostNew{Controller: c})
}
