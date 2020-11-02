package app

import (
	"log"
	"net/http"

	"github.com/danikg/go-todo-rest-api/app/pg"
	"github.com/danikg/go-todo-rest-api/config"
	"github.com/gorilla/mux"

	tagHttp "github.com/danikg/go-todo-rest-api/tag/controller/http"
	tagRepo "github.com/danikg/go-todo-rest-api/tag/repository/pg"
	tagService "github.com/danikg/go-todo-rest-api/tag/service"
	todoItemHttp "github.com/danikg/go-todo-rest-api/todoitem/controller/http"
	todoItemRepo "github.com/danikg/go-todo-rest-api/todoitem/repository/pg"
	todoItemService "github.com/danikg/go-todo-rest-api/todoitem/service"
	todoListHttp "github.com/danikg/go-todo-rest-api/todolist/controller/http"
	todoListRepo "github.com/danikg/go-todo-rest-api/todolist/repository/pg"
	todoListService "github.com/danikg/go-todo-rest-api/todolist/service"
	userHttp "github.com/danikg/go-todo-rest-api/user/controller/http"
	userRepo "github.com/danikg/go-todo-rest-api/user/repository/pg"
	userService "github.com/danikg/go-todo-rest-api/user/service"
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

	userRepo := userRepo.NewUserRepository(db)
	userService := userService.NewUserService(userRepo)
	userController := userHttp.NewUserController(userService)
	userHttp.SetupRoutes(router, userController)

	todoListRepo := todoListRepo.NewTodoListRepository(db)
	todoListService := todoListService.NewTodoListService(userRepo, todoListRepo)
	todoListController := todoListHttp.NewTodoListController(todoListService)
	todoListHttp.SetupRoutes(router, todoListController)

	todoItemRepo := todoItemRepo.NewTodoItemRepository(db)
	todoItemService := todoItemService.NewTodoItemService(todoItemRepo, todoListRepo)
	todoItemController := todoItemHttp.NewTodoItemController(todoItemService)
	todoItemHttp.SetupRoutes(router, todoItemController)

	tagRepo := tagRepo.NewTagRepository(db)
	tagService := tagService.NewTagService(tagRepo, todoItemRepo)
	tagController := tagHttp.NewTagController(tagService)
	tagHttp.SetupRoutes(router, tagController)

	addr := a.config.AppHost + ":" + a.config.AppPort
	log.Printf("starting at %s...", addr)
	http.ListenAndServe(addr, router)
}