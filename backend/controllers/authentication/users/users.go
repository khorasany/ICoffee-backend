package users

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/khorasany/coffee/api/backend/helpers/jwtToken"
	"github.com/khorasany/coffee/api/backend/models/authentication/user"
	"net/http"
)

func GetAllAdminUsers(w http.ResponseWriter, r *http.Request) {
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

	admins, err := user.GetAllAdminUsers()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(admins)
	return
}

func GetAdminInfo(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
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
	}

	adminMap := UserAdminParamToModel(request)
	admin, err := user.GetAdminInfo(adminMap.ID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(admin)
}

func DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
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

	adminModel := UserAdminParamToModel(params)
	if err := user.DeleteAdmin(adminModel.ID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte("user requested has been deleted"))
	w.WriteHeader(http.StatusOK)
	return
}
