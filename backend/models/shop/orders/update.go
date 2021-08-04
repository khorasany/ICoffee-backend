package orders

import (
	"encoding/json"
	"github.com/khorasany/coffee/api/backend/database"
	"github.com/khorasany/coffee/api/backend/models"
	"strconv"
)

func UpdateOrder(order models.Order) (*models.Order, error) {
	orderID := strconv.Itoa(int(order.ID))
	db := database.CreateCon()
	_, err := db.Exec("update ico_order set products='" + order.Products + "',ref_id='" + order.RefID + "'" +
		",authority='" + order.Authority + "',transport_price='" + order.TransportPrice + "',total_price='" + order.TotalPrice + "'" +
		",discount_amount='" + order.DiscountAmount + "',discount_coupon='" + order.DiscountCoupon + "'" +
		",wallet_reduction='" + order.WalletReduction + "',date_paid='" + order.DatePaid.String() + "',description='" + order.Description + "'" +
		",status='" + order.Status + "' where id=" + orderID + ";")
	if err != nil {

		return nil, err
	}

	err = db.QueryRow("select * from ico_order where id="+orderID+";").Scan(&order.ID, &order.CustomerID,
		&order.Products, &order.RefID, &order.OrderToken, &order.Authority, &order.TransportPrice, &order.TotalPrice,
		&order.DiscountAmount, &order.DiscountCoupon, &order.WalletReduction, &order.CreatedAt, &order.DatePaid, &order.Description, &order.Status)
	if err != nil {

		return nil, err
	}

	return &order, nil
}

func UpdateCard(cards []*models.Card, customerID string) (*models.FullCard, error) {
	fullCard, _ := json.Marshal(cards)
	db := database.CreateCon()
	_, err := db.Exec("update ico_card set card='" + string(fullCard) + "',status=1 where customer_id=" + customerID + ";")
	if err != nil {
		return nil, err
	}

	fullCardModel := models.FullCard{}
	err = db.QueryRow("select * from ico_card where customer_id="+customerID+";").
		Scan(&fullCardModel.CardID, &fullCardModel.CustomerID, &fullCardModel.Card, &fullCardModel.CreatedAt, &fullCardModel.Status)
	if err != nil {
		return nil, err
	}

	return &fullCardModel, nil
}
