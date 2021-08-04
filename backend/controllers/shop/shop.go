package shop

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/khorasany/coffee/api/backend/helpers/jwtToken"
	"github.com/khorasany/coffee/api/backend/models/shop/category"
	"github.com/khorasany/coffee/api/backend/models/shop/shop"
	"io/ioutil"
	"net/http"
)

func RegisterShop(w http.ResponseWriter, r *http.Request) {
	result, _ := ioutil.ReadAll(r.Body)
	request := make(map[string]string)
	_ = json.Unmarshal(result, &request)
	w.Header().Set("Content-Type", "application/json")
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
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	shopModel := shopMapToMapParmaToModel(request)
	shopInfo, err := shop.RegisterShop(shopModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(shopInfo)
	return
}

func GetShopByShopID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
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
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	shopInfo, err := shop.GetShopByShopID(param["shop_id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(shopInfo)
	return
}

func GetShopByOwnerID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
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
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	shopInfo, err := shop.GetShopByOwnerID(param["owner_id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(shopInfo)
	return
}

func GetShops(w http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)
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
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	shopsInfo, err := shop.GetShops()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(shopsInfo)
	return
}

func DeleteShop(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
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
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	err = shop.DeleteShop(param["shop_id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

func UpdateShop(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	shopModel := shopMapToMapParmaToModel(request)
	shopInfo, err := shop.UpdateShop(shopModel)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(shopInfo)
	return
}

func RegisterCategory(w http.ResponseWriter, r *http.Request) {
	result, _ := ioutil.ReadAll(r.Body)
	request := make(map[string]string)
	_ = json.Unmarshal(result, &request)
	w.Header().Set("Content-Type", "application/json")
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
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	categoryModel := categoryMapToMapParamToModel(request)
	categoryInfo, err := category.RegisterCategory(categoryModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(categoryInfo)
	return
}

func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
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
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	catModel := categoryMapToMapParamToModel(param)
	catInfo, err := category.GetCategoryByID(catModel.CatID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(catInfo)
	return
}

func GetCategoryByCatName(w http.ResponseWriter, r *http.Request) {
	param, _ := ioutil.ReadAll(r.Body)
	request := make(map[string]string)
	_ = json.Unmarshal(param, &request)
	w.Header().Set("Content-Type", "application/json")
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
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	catInfo, err := category.GetCategoryByCatName(request["cat_name"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(catInfo)
	return
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
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
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	err = category.DeleteCategory(param["cat_id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("record has been deleted"))
	return
}

func GetAllCategoriesByType(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	token, err := r.Cookie("AuthenticationToken")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	_, _, _, err = jwtToken.AuthenticationJwtToken(token.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	categories, err := category.GetAllCategoriesByType(param["type"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	_ = json.NewEncoder(w).Encode(categories)
}
