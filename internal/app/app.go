package app

import (
	"log"
	"net/http"

	"github.com/danikg/go-todo-rest-api/internal/config"
	"github.com/danikg/go-todo-rest-api/internal/models"
	"github.com/gorilla/mux"

	taghttp "github.com/danikg/go-todo-rest-api/internal/tag/controller/http"
	tagrepo "github.com/danikg/go-todo-rest-api/internal/tag/repository"
	tagservice "github.com/danikg/go-todo-rest-api/internal/tag/service"
	tihttp "github.com/danikg/go-todo-rest-api/internal/todoitem/controller/http"
	tirepo "github.com/danikg/go-todo-rest-api/internal/todoitem/repository"
	tiservice "github.com/danikg/go-todo-rest-api/internal/todoitem/service"
	tlhttp "github.com/danikg/go-todo-rest-api/internal/todolist/controller/http"
	tlrepo "github.com/danikg/go-todo-rest-api/internal/todolist/repository"
	tlservice "github.com/danikg/go-todo-rest-api/internal/todolist/service"
	userhttp "github.com/danikg/go-todo-rest-api/internal/user/controller/http"
	userrepo "github.com/danikg/go-todo-rest-api/internal/user/repository"
	userservice "github.com/danikg/go-todo-rest-api/internal/user/service"
)

// App ...
type App struct {
	Config *config.Config
}

// NewApp ...
func NewApp() *App {
	return &App{
		Config: config.GetConfig(),
	}
}

// Run ...
func (a *App) Run() {
	db := models.GetDB(a.Config)
	router := mux.NewRouter()

	userRepo := userrepo.NewUserRepository(db)
	userService := userservice.NewUserService(userRepo)
	userController := userhttp.NewUserController(userService)
	userhttp.SetupRoutes(router, userController)

	todoListRepo := tlrepo.NewTodoListRepository(db)
	todoListService := tlservice.NewTodoListService(userRepo, todoListRepo)
	todoListController := tlhttp.NewTodoListController(todoListService)
	tlhttp.SetupRoutes(router, todoListController)

	todoItemRepo := tirepo.NewTodoItemRepository(db)
	todoItemService := tiservice.NewTodoItemService(todoItemRepo, todoListRepo)
	todoItemController := tihttp.NewTodoItemController(todoItemService)
	tihttp.SetupRoutes(router, todoItemController)

	tagRepo := tagrepo.NewTagRepository(db)
	tagService := tagservice.NewTagService(tagRepo, todoItemRepo)
	tagController := taghttp.NewTagController(tagService)
	taghttp.SetupRoutes(router, tagController)

	addr := a.Config.AppHost + ":" + a.Config.AppPort
	log.Printf("starting at %s...", addr)
	http.ListenAndServe(addr, router)
}
