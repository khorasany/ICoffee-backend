package products

import (
	"fmt"
	"github.com/khorasany/coffee/api/backend/database"
)

func DeleteProduct(productID string) error {
	db := database.CreateCon()
	_, err := db.Exec(fmt.Sprintf("delete from ico_product where id=%v;", productID))
	if err != nil {
		return err
	}

	return nil
}
