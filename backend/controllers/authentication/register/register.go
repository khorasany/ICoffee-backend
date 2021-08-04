package register

import (
	"encoding/json"
	"github.com/khorasany/coffee/api/backend/helpers/jwtToken"
	"io/ioutil"
	"net/http"

	"github.com/khorasany/coffee/api/backend/controllers/authentication/users"
	"github.com/khorasany/coffee/api/backend/models/authentication/user"
)

func CreateAdminUser(w http.ResponseWriter, request *http.Request) {
	result, err := ioutil.ReadAll(request.Body)
	r := make(map[string]string)
	_ = json.Unmarshal(result, &r)
	token, err := request.Cookie("AuthenticationToken")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, _, _, err = jwtToken.AuthenticationJwtToken(token.Value)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	adminModel := users.UserAdminParamToModel(r)
	admin, err := user.RegisterAdmin(adminModel)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(admin)
	return
}

func CreateSuperAdminUser(w http.ResponseWriter, r *http.Request) {
	result, _ := ioutil.ReadAll(r.Body)
	request := make(map[string]string)
	_ = json.Unmarshal(result, &request)

	token, err := r.Cookie("AuthenticationToken")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, _, _, err = jwtToken.AuthenticationJwtToken(token.Value)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	superAdminModel := users.UserAdminParamToModel(request)
	superAdmin, err := user.RegisterSuperUser(superAdminModel)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(superAdmin)
	return
}
