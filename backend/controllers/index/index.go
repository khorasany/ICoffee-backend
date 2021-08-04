package index

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetConnectedMessage(w http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("ICoffee service started successfully..."))
}
