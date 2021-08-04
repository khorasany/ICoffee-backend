package orders

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/khorasany/coffee/api/backend/database"
	"github.com/khorasany/coffee/api/backend/models"
)

func CreateOrder(order models.Order) (*models.Order, error) {
	customerID := strconv.Itoa(int(order.CustomerID))
	db := database.CreateCon()
	orderResult, err := db.Exec("insert into ico_order (customer_id, products, ref_id, order_token, authority, transport_price, total_price, discount_amount, discount_coupon, wallet_reduction, created_at, date_paid, description, status) " +
		"values (" + customerID + ",'" + order.Products + "','" + order.RefID + "','" + order.OrderToken + "','" + order.Authority + "','" + order.TransportPrice + "','" + order.TotalPrice + "','" + order.DiscountAmount + "','" + order.DiscountCoupon +
		"','" + order.WalletReduction + "','" + order.CreatedAt.String() + "','" + order.DatePaid.String() + "','" + order.Description + "','" + order.Status + "');")
	if err != nil {
		return nil, err
	}

	orderID, _ := orderResult.LastInsertId()
	order.ID = orderID

	return &order, nil
}

func RegisterCard(cards []*models.Card, customer_id string) (*models.FullCard, error) {
	db := database.CreateCon()
	fullCard, _ := json.Marshal(cards)
	cardResult, err := db.Exec("insert into ico_card (customer_id, card, status) values (" + customer_id + ",'" + string(fullCard) + "',1);")
	if err != nil {
		return nil, err
	}

	cardID, _ := cardResult.LastInsertId()
	cardIDString := strconv.Itoa(int(cardID))
	createAt := strconv.Itoa(int(time.Now().Unix()))
	status := strconv.Itoa(1)

	return &models.FullCard{
		CardID:     cardIDString,
		CustomerID: customer_id,
		Card:       string(fullCard),
		CreatedAt:  createAt,
		Status:     status,
	}, nil
}
