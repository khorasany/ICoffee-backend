package services

import "github.com/gorilla/mux"

func RunnerICoffeeService() *mux.Router {
	return mux.NewRouter().StrictSlash(true)
}
