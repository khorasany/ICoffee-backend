package shop

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/khorasany/coffee/api/backend/helpers/jwtToken"
	"github.com/khorasany/coffee/api/backend/models/shop/products"
	"io/ioutil"
	"net/http"
	"os"
)

func RegisterProduct(w http.ResponseWriter, r *http.Request) {
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

	productModel := productMapToMapParamToModel(request)
	productInfo, err := products.RegisterProduct(productModel)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(productInfo)
	return
}

func GetProductByProductID(w http.ResponseWriter, r *http.Request) {
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

	productInfo, err := products.GetProductByProductID(param["product_id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(productInfo)

	return
}

func GetProductsByOwnerID(w http.ResponseWriter, r *http.Request) {
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

	productsInfo, err := products.GetProductByOwnerID(param["owner_id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(productsInfo)
	return
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
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

	productModel := productMapToMapParamToModel(request)
	productInfo, err := products.UpdateProduct(productModel)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(productInfo)

	return
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
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

	err = products.DeleteProduct(param["product_id"])
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

func UploadProductImage(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseMultipartForm(10 << 20)

	w.Header().Set("Content-Type", "application/json")
	file, handler, err := r.FormFile("product-image")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error Retrieving the File"))
		return
	}

	defer file.Close()
	_, _ = w.Write([]byte(fmt.Sprintf("Uploaded File: %+v\n", handler.Filename)))
	_, _ = w.Write([]byte(fmt.Sprintf("File Size: %+v\n", handler.Size)))
	_, _ = w.Write([]byte(fmt.Sprintf("MIME Header: %+v\n", handler.Header)))

	tempFile, err := ioutil.TempFile("temp", "upload-product-*.png")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	_, _ = tempFile.Write(fileBytes)
	newPath := fmt.Sprintf("img/product/%v", handler.Filename)
	err = os.Rename(tempFile.Name(), newPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("upload file has been stopped,please try again at the moment!"))
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(newPath)
	return
}
