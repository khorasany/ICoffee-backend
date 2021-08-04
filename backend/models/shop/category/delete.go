package category

import (
	"fmt"
	"github.com/khorasany/coffee/api/backend/database"
)

func DeleteCategory(catID string) error {
	db := database.CreateCon()
	_, err := db.Exec(fmt.Sprintf("delete from ico_category where id=%v", catID))
	if err != nil {
		return err
	}
	return nil
}
