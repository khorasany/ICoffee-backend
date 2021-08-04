package login

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/khorasany/coffee/api/backend/controllers/authentication/users"
	"github.com/khorasany/coffee/api/backend/models/authentication/user"
	"github.com/khorasany/coffee/api/backend/services"
	"io/ioutil"
	"net/http"
)

func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	adminModel := users.UserAdminParamToModel(params)
	authenticate, admin := user.LoginAdmin(*adminModel)
	if authenticate == false {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	claims := mapToMapAdminToJwtToken(admin)
	cookie, err := services.SetToken(&claims)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	http.SetCookie(w, cookie)
	return
}

func LoginSuperUserAdmin(w http.ResponseWriter, r *http.Request) {
	result, _ := ioutil.ReadAll(r.Body)
	request := make(map[string]string)
	_ = json.Unmarshal(result, &request)

	superAdminModel := users.UserAdminParamToModel(request)
	authenticate, superAdmin := user.LoginSuperUserAdmin(superAdminModel)
	if authenticate == false {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	claims := mapToMapAdminToJwtToken(superAdmin)
	cookie, err := services.SetToken(&claims)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	return
}
