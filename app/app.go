package app

import (
	"log"
	"net/http"

	"github.com/danikg/go-todo-rest-api/app/pg"
	"github.com/danikg/go-todo-rest-api/config"
	"github.com/gorilla/mux"

	controllers "github.com/danikg/go-todo-rest-api/controllers/http"
	repos "github.com/danikg/go-todo-rest-api/repositories/pg"
	services "github.com/danikg/go-todo-rest-api/services/web"
)

// App ...
type App struct {
	config *config.Config
}

// NewApp ...
func NewApp() *App {
	return &App{
		config: config.GetConfig(),
	}
}

// Run ...
func (a *App) Run() {
	db := pg.GetDB(a.config)
	router := mux.NewRouter()

	userRepo := repos.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	controllers.SetupUserRoutes(router, userController)

	todoListRepo := repos.NewTodoListRepository(db)
	todoListService := services.NewTodoListService(userRepo, todoListRepo)
	todoListController := controllers.NewTodoListController(todoListService)
	controllers.SetupTodoListRoutes(router, todoListController)

	todoItemRepo := repos.NewTodoItemRepository(db)
	todoItemService := services.NewTodoItemService(todoItemRepo, todoListRepo)
	todoItemController := controllers.NewTodoItemController(todoItemService)
	controllers.SetupTodoItemRoutes(router, todoItemController)

	tagRepo := repos.NewTagRepository(db)
	tagService := services.NewTagService(tagRepo, todoItemRepo)
	tagController := controllers.NewTagController(tagService)
	controllers.SetupTagRoutes(router, tagController)

	addr := a.config.AppHost + ":" + a.config.AppPort
	log.Printf("starting at %s...", addr)
	http.ListenAndServe(addr, router)
}
