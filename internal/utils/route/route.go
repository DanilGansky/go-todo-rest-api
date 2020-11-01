package route

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetRouteVar ...
func GetRouteVar(r *http.Request, routeVar string) (uint, error) {
	id, err := strconv.Atoi(mux.Vars(r)[routeVar])
	return uint(id), err
}
