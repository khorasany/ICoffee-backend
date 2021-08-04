package shop

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"

	"github.com/khorasany/coffee/api/backend/helpers/jwtToken"
	"github.com/khorasany/coffee/api/backend/models/shop/orders"
)

//func BuyRequest(w http.ResponseWriter, r *http.Request) {
//
//}
//
//func ReturnPaymentStatus(w http.ResponseWriter, r *http.Request) {
//
//}

//GetOrdersShop for super user access level
func GetOrdersShop(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	token, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.Header().Set("content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(err.Error()))

			return
		}
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	_, _, _, err = jwtToken.AuthenticationJwtToken(token.Value)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	ordersList, err := orders.GetOrdersShop(param["shop_id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(ordersList)

	return
}

// GetAllOrdersOfOwnerShops for super user access level
func GetAllOrdersOfOwnerShops(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	token, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.Header().Set("content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(err.Error()))

			return
		}
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	_, _, _, err = jwtToken.AuthenticationJwtToken(token.Value)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	ordersList, err := orders.GetAllOrdersOfOwnerShops(param["owner_id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(ordersList)

	return
}

//GetAllOrdersOfSpecificShop for admin users access level
func GetAllOrdersOfSpecificShop(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	token, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.Header().Set("content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(err.Error()))

			return
		}
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	_, _, _, err = jwtToken.AuthenticationJwtToken(token.Value)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	ordersInformation, err := orders.GetAllOrdersOfSpecificShop(param["shop_id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(ordersInformation)

	return
}

// GetCustomersOrders for admin users section
func GetCustomersOrders(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	token, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.Header().Set("content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(err.Error()))

			return
		}
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	_, _, _, err = jwtToken.AuthenticationJwtToken(token.Value)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	ordersInformation, err := orders.GetCustomersOrders(param["customer_id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(ordersInformation)

	return
}

//GetCustomersOrder for admin user section
func GetCustomersOrder(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	token, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.Header().Set("content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(err.Error()))

			return
		}
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	_, _, _, err = jwtToken.AuthenticationJwtToken(token.Value)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	orderInfo, err := orders.GetCustomersOrder(param["order_id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(orderInfo)

	return
}

// CreateOrder for admin user section
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	result, _ := ioutil.ReadAll(r.Body)
	request := make(map[string]string)
	_ = json.Unmarshal(result, &request)

	token, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.Header().Set("content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(err.Error()))

			return
		}
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	_, _, _, err = jwtToken.AuthenticationJwtToken(token.Value)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	orderModel := orderMapToMapParamToModel(request)
	orderInfo, err := orders.CreateOrder(orderModel)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(orderInfo)

	return
}

// DeleteOrder for admin user section
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	token, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.Header().Set("content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(err.Error()))

			return
		}
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	_, _, _, err = jwtToken.AuthenticationJwtToken(token.Value)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	err = orders.DeleteOrder(param["order_id"])
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

// UpdateOrder for admin user section
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	result, _ := ioutil.ReadAll(r.Body)
	request := make(map[string]string)
	_ = json.Unmarshal(result, &request)

	token, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.Header().Set("content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(err.Error()))

			return
		}
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	_, _, _, err = jwtToken.AuthenticationJwtToken(token.Value)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	orderModel := orderMapToMapParamToModel(request)
	orderInfo, err := orders.UpdateOrder(orderModel)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(orderInfo)

	return
}

// CreateCard for admin user section
func CreateCard(w http.ResponseWriter, r *http.Request) {
	result, _ := ioutil.ReadAll(r.Body)
	request := make(map[string][]map[string]string)
	_ = json.Unmarshal(result, &request)

	token, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.Header().Set("content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(err.Error()))

			return
		}
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	_, _, _, err = jwtToken.AuthenticationJwtToken(token.Value)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	cardModel := cardMapToMapParamToModel(request)
	cardsInfo, err := orders.RegisterCard(cardModel, request["customer"][0]["customer_id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(cardsInfo)

}

// DeleteCard for admin user section
func DeleteCard(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	token, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.Header().Set("content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(err.Error()))

			return
		}
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	_, _, _, err = jwtToken.AuthenticationJwtToken(token.Value)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	err = orders.DeleteCard(param["owner_id"])
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

// UpdateCard for admin user section
func UpdateCard(w http.ResponseWriter, r *http.Request) {
	result, _ := ioutil.ReadAll(r.Body)
	request := make(map[string][]map[string]string)
	_ = json.Unmarshal(result, &request)

	token, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.Header().Set("content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(err.Error()))

			return
		}
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	_, _, _, err = jwtToken.AuthenticationJwtToken(token.Value)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	cardModel := cardMapToMapParamToModel(request)
	cardsInfo, err := orders.UpdateCard(cardModel, request["customer"][0]["customer_id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(cardsInfo)

	return
}
