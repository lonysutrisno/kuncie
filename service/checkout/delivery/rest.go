package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lonysutrisno/kuncie/pkg"
	"github.com/lonysutrisno/kuncie/service/checkout/usecase"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"/checkout", "POST", Checkout},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
	return router
}

func Checkout(w http.ResponseWriter, req *http.Request) {
	var FormOrder []pkg.FormOrder
	decode := json.NewDecoder(req.Body)
	err := decode.Decode(&FormOrder)
	if err != nil {
		pkg.JSONResponse(w, 400, nil, err)
		return
	}

	res, err := usecase.Checkout(FormOrder)
	if err != nil {
		pkg.JSONResponse(w, 500, nil, err)
		return
	}
	pkg.JSONResponse(w, 200, res, nil)
}
