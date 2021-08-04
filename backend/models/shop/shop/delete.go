package shop

import (
	"fmt"
	"github.com/khorasany/coffee/api/backend/database"
)

func DeleteShop(shopID string) error {
	db := database.CreateCon()
	_, err := db.Exec(fmt.Sprintf("delete from ico_shop where id=%v;", shopID))
	if err != nil {
		return err
	}
	return nil
}
