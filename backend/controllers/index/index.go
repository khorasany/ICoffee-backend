package index

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetConnectedMessage(w http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("ICoffee service started successfully..."))
}
