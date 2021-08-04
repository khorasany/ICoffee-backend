package orders

import "github.com/khorasany/coffee/api/backend/database"

func DeleteCard(ownerID string) error {
	db := database.CreateCon()
	_, err := db.Exec("delete from ico_card where customer_id=" + ownerID + ";")
	if err != nil {
		return err
	}

	return nil
}

func DeleteOrder(orderID string) error {
	db := database.CreateCon()
	_, err := db.Exec("delete from ico_order where id=" + orderID + ";")
	if err != nil {

		return err
	}

	return nil
}
