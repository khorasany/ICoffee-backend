package user

import (
	"github.com/khorasany/coffee/api/backend/database"
	"strconv"
)

func deleteUser() {

}

func deleteUserMeta() {

}

func deleteAdminMeta() {

}

func DeleteAdmin(userID int64) error {
	id := strconv.Itoa(int(userID))
	db := database.CreateCon()
	_, err := db.Exec("delete from ico_user_admin where id=" + id + ";")
	if err != nil {
		return err
	}

	return nil
}
