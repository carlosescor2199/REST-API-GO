package router

import (
	"github.com/gorilla/mux"
)

var Router = mux.NewRouter().StrictSlash(true)
