package routers

import (
	"github.com/khorasany/coffee/api/backend/controllers/index"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/khorasany/coffee/api/backend/controllers/authentication/login"
	"github.com/khorasany/coffee/api/backend/controllers/authentication/permissions"
	"github.com/khorasany/coffee/api/backend/controllers/authentication/register"
	"github.com/khorasany/coffee/api/backend/controllers/authentication/users"
	"github.com/khorasany/coffee/api/backend/controllers/shop"
	"github.com/khorasany/coffee/api/backend/services"
)

func Router() {
	r := services.RunnerICoffeeService()

	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch})

	// Super user admin section
	r.HandleFunc("/super-admin/admins", users.GetAllAdminUsers).Methods("GET")           // checked
	r.HandleFunc("/super-admin/login", login.LoginSuperUserAdmin).Methods("POST")        // checked
	r.HandleFunc("/super-admin/delete/{id}", users.DeleteAdmin).Methods("DELETE")        // checked
	r.HandleFunc("/super-admin/{id}", users.GetAdminInfo).Methods("GET")                 // checked
	r.HandleFunc("/super-admin/register", register.CreateSuperAdminUser).Methods("POST") // checked

	r.HandleFunc("/super-admin/role/register", permissions.SetRole).Methods("POST") // checked
	r.HandleFunc("/super-admin/role", permissions.GetRole).Methods("POST")          // checked
	r.HandleFunc("/super-admin/roles", permissions.GetRoles).Methods("POST")        // checked

	r.HandleFunc("/super-admin/orders/shop/{shop_id}", shop.GetAllOrdersOfSpecificShop).Methods("POST")
	r.HandleFunc("/super-admin/orders/owner/{owner_id}", shop.GetAllOrdersOfOwnerShops).Methods("POST")
	r.HandleFunc("/super-admin/orders/{shop_id}", shop.GetOrdersShop).Methods("POST")

	// category
	r.HandleFunc("/super-admin/category/register", shop.RegisterCategory).Methods("POST")             // checked
	r.HandleFunc("/super-admin/category/id/{cat_id}", shop.GetCategoryByID).Methods("GET")            // checked
	r.HandleFunc("/super-admin/category/name", shop.GetCategoryByCatName).Methods("POST")             // checked
	r.HandleFunc("/super-admin/category/delete/{cat_id}", shop.DeleteCategory).Methods("DELETE")      // checked
	r.HandleFunc("/super-admin/categories/select/{type}", shop.GetAllCategoriesByType).Methods("GET") // checked

	r.HandleFunc("/super-admin/shop/register", shop.RegisterShop).Methods("POST")            // checked
	r.HandleFunc("/super-admin/shop/owner/{owner_id}", shop.GetShopByOwnerID).Methods("GET") // checked
	r.HandleFunc("/super-admin/shop/{shop_id}", shop.GetShopByShopID).Methods("GET")         // checked
	r.HandleFunc("/super-admin/shops/select", shop.GetShops).Methods("GET")                  // checked
	r.HandleFunc("/super-admin/shop/delete/{shop_id}", shop.DeleteShop).Methods("DELETE")    // checked
	r.HandleFunc("/super-admin/shop/update", shop.UpdateShop).Methods("PATCH")               // checked

	r.HandleFunc("/super-admin/product/register", shop.RegisterProduct).Methods("POST")             // checked
	r.HandleFunc("/super-admin/product/upload/avatar", shop.UploadProductImage).Methods("POST")     // checked
	r.HandleFunc("/super-admin/product/{product_id}", shop.GetProductByProductID).Methods("GET")    //checked
	r.HandleFunc("/super-admin/product/owner/{owner_id}", shop.GetProductsByOwnerID).Methods("GET") // checked
	r.HandleFunc("/super-admin/product/update", shop.UpdateProduct).Methods("PATCH")                // checked
	r.HandleFunc("/super-admin/product/delete/{product_id}", shop.DeleteProduct).Methods("DELETE")  // checked

	// Admin section
	r.HandleFunc("/admin/login", login.LoginAdmin).Methods("POST")
	r.HandleFunc("/admin/register", register.CreateAdminUser).Methods("POST")
	r.HandleFunc("/admin/card/update", shop.UpdateCard).Methods("PATCH") // need to check
	r.HandleFunc("/admin/delete/{owner_id}", shop.DeleteCard).Methods("GET")
	r.HandleFunc("/admin/create", shop.CreateCard).Methods("POST")

	r.HandleFunc("/admin/update", shop.UpdateOrder).Methods("PATCH")
	r.HandleFunc("/admin/delete/{order_id}", shop.DeleteOrder).Methods("DELETE")
	r.HandleFunc("/admin/create", shop.CreateOrder).Methods("POST")
	r.HandleFunc("/admin/{order_id}", shop.GetCustomersOrder).Methods("POST")
	r.HandleFunc("/admin/{customer_id}", shop.GetCustomersOrders).Methods("POST")

	r.HandleFunc("/", index.GetConnectedMessage).Methods("GET")

	log.Fatal(http.ListenAndServe(":9006", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
