package orders

import (
	"github.com/khorasany/coffee/api/backend/database"
	"github.com/khorasany/coffee/api/backend/models"
)

func GetCustomersOrder(orderID string) (*models.Order, error) {
	var order models.Order
	db := database.CreateCon()
	err := db.QueryRow("select * from ico_order where id="+orderID+";").Scan(&order.ID, &order.CustomerID,
		&order.Products, &order.RefID, &order.OrderToken, &order.Authority, &order.TransportPrice, &order.TotalPrice,
		&order.DiscountAmount, &order.DiscountCoupon, &order.WalletReduction, &order.CreatedAt, &order.DatePaid, &order.Description, &order.Status)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func GetCustomersOrders(customerID string) ([]*models.Order, error) {
	db := database.CreateCon()
	orderResult, err := db.Query("select * from ico_order where customer_id=" + customerID + ";")
	if err != nil {
		return nil, err
	}

	orders := []*models.Order{}
	for orderResult.Next() {
		var order models.Order
		_ = orderResult.Scan(&order.ID, &order.CustomerID, &order.Products, &order.RefID, &order.OrderToken,
			&order.Authority, &order.TransportPrice, &order.TotalPrice, &order.DiscountAmount, &order.DiscountCoupon,
			&order.WalletReduction, &order.CreatedAt, &order.DatePaid, &order.Description, &order.Status)
		orders = append(orders, &order)
	}

	return orders, nil
}

func GetAllOrdersOfSpecificShop(shopID string) ([]*models.Order, error) {
	db := database.CreateCon()
	ordersResult, err := db.Query("select * from ico_order where customer_id in (select user_id from ico_usermeta where meta_key='_customer_shop' and meta_value='" + shopID + "');")
	if err != nil {
		return nil, err
	}

	orders := []*models.Order{}
	for ordersResult.Next() {
		var order models.Order
		_ = ordersResult.Scan(&order.ID, &order.CustomerID, &order.Products, &order.RefID, &order.OrderToken,
			&order.Authority, &order.TransportPrice, &order.TotalPrice, &order.DiscountAmount, &order.DiscountCoupon,
			&order.WalletReduction, &order.CreatedAt, &order.DatePaid, &order.Description, &order.Status)
		orders = append(orders, &order)
	}

	return orders, nil
}

func GetAllOrdersOfOwnerShops(ownerID string) ([]*models.Order, error) {
	db := database.CreateCon()
	ordersResult, err := db.Query("select * from ico_order where customer_id in (select user_id from ico_usermeta where meta_key='_customer_shop' and meta_value in (select id from ico_shop where owner_id=" + ownerID + "));")
	if err != nil {
		return nil, err
	}

	orders := []*models.Order{}
	for ordersResult.Next() {
		var order models.Order
		_ = ordersResult.Scan(&order.ID, &order.CustomerID, &order.Products, &order.RefID, &order.OrderToken,
			&order.Authority, &order.TransportPrice, &order.TotalPrice, &order.DiscountAmount, &order.DiscountCoupon,
			&order.WalletReduction, &order.CreatedAt, &order.DatePaid, &order.Description, &order.Status)
		orders = append(orders, &order)
	}

	return orders, nil
}

func GetOrdersShop(shopID string) ([]*models.Order, error) {
	db := database.CreateCon()
	ordersList, err := db.Query("select * from ico_order where customer_id in (select user_id from ico_usermeta where meta_key='_customer_shop' and meta_value='" + shopID + "');")
	if err != nil {
		return nil, err
	}

	orders := []*models.Order{}
	for ordersList.Next() {
		var order models.Order
		_ = ordersList.Scan(&order.ID, &order.CustomerID, &order.Products, &order.RefID, &order.OrderToken,
			&order.Authority, &order.TransportPrice, &order.TotalPrice, &order.DiscountAmount, &order.DiscountCoupon,
			&order.WalletReduction, &order.CreatedAt, &order.DatePaid, &order.Description, &order.Status)
		orders = append(orders, &order)
	}

	return orders, nil
}
