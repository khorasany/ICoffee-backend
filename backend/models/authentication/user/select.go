package user

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/khorasany/coffee/api/backend/database"
	"github.com/khorasany/coffee/api/backend/models"
	"strconv"
)

func getAllUsers() {

}

func GetAdminInfo(adminID int64) (*models.Admin, error) {
	db := database.CreateCon()
	var admin models.Admin
	err := db.QueryRow(fmt.Sprintf("select id,firstname,lastname,username,email,role_id,created_at,status from ico_user_admin where id=%v;", strconv.Itoa(int(adminID)))).Scan(&admin.ID, &admin.FirstName, &admin.LastName, &admin.UserName, &admin.Email, &admin.Role.ID, &admin.CreatedAt, &admin.Status)
	strAdminID := strconv.Itoa(int(adminID))
	res, _ := db.Query(fmt.Sprintf("select meta_key,meta_value from `ico_admin_meta` where meta_key in ('address','avatar','birth_date','phone','mobile') and admin_id=%v;", strAdminID))
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var meta models.MetaInfo
		_ = res.Scan(&meta.MetaKey, &meta.MetaValue)
		if meta.MetaKey == "address" {
			admin.Meta.Address = meta.MetaValue
		}
		if meta.MetaKey == "phone" {
			admin.Meta.Phone = meta.MetaValue
		}
		if meta.MetaKey == "avatar" {
			admin.Meta.Avatar = meta.MetaValue
		}
		if meta.MetaKey == "birth_date" {
			admin.Meta.BirthDate = meta.MetaValue
		}
		if meta.MetaKey == "mobile" {
			admin.Meta.Mobile = meta.MetaValue
		}
	}
	role := getRoleByID(admin.Role.ID)
	admin.Role.RoleName = role.RoleName
	admin.Role.Status = role.Status

	return &admin, nil
}

func login(params map[string]string) {

}

func LoginAdmin(admin models.Admin) (bool, *models.Admin) {
	stringToHash := []byte(admin.Password)
	hashPassword := sha1.Sum(stringToHash)

	db := database.CreateCon()
	err := db.QueryRow(fmt.Sprintf("select id,firstname,lastname,username,email,role_id,created_at,status from ico_user_admin where role_id=1 and username='%v' and password='%v';", admin.UserName, hex.EncodeToString(hashPassword[:]))).Scan(&admin.ID, &admin.FirstName, &admin.LastName, &admin.UserName, &admin.Email, &admin.Role.ID, &admin.CreatedAt, &admin.Status)
	if err != nil {
		return false, nil
	}
	if admin.Status == 0 {
		return false, nil
	}
	role := getRoleByID(admin.Role.ID)
	admin.Role.RoleName = role.RoleName
	admin.Role.Status = role.Status

	return true, &admin
}

func getRoleByID(roleID int64) models.Role {
	var role models.Role
	db := database.CreateCon()
	strRoleID := strconv.Itoa(int(roleID))
	_ = db.QueryRow(fmt.Sprintf("select role_name,status from `ico_roles` where id=%v;", strRoleID)).Scan(&role.RoleName, &role.Status)
	role.ID = roleID
	return role
}

func LoginSuperUserAdmin(admin *models.Admin) (bool, *models.Admin) {
	stringToHash := []byte(admin.Password)
	hashPassword := sha1.Sum(stringToHash)

	db := database.CreateCon()
	err := db.QueryRow(fmt.Sprintf("select id,firstname,lastname,username,email,role_id,created_at,status from `ico_user_admin` where role_id=2 and username='%v' and password='%v';", admin.UserName, hex.EncodeToString(hashPassword[:]))).
		Scan(&admin.ID, &admin.FirstName, &admin.LastName, &admin.UserName, &admin.Email, &admin.Role.ID, &admin.CreatedAt, &admin.Status)
	adminID := strconv.Itoa(int(admin.ID))
	rows, _ := db.Query(fmt.Sprintf("select meta_key,meta_value from `ico_admin_meta` where admin_id=%v and meta_key in ('birth_date','address','phone','mobile','avatar')", adminID))
	if err != nil {
		return false, nil
	}

	for rows.Next() {
		var meta models.MetaInfo
		_ = rows.Scan(&meta.MetaKey, &meta.MetaValue)
		if meta.MetaKey == "birth_date" {
			admin.Meta.BirthDate = meta.MetaValue
		}
		if meta.MetaKey == "address" {
			admin.Meta.Address = meta.MetaValue
		}
		if meta.MetaKey == "phone" {
			admin.Meta.Phone = meta.MetaValue
		}
		if meta.MetaKey == "mobile" {
			admin.Meta.Mobile = meta.MetaValue
		}
		if meta.MetaKey == "avatar" {
			admin.Meta.Avatar = meta.MetaValue
		}
	}
	role := getRoleByID(admin.Role.ID)
	admin.Role.RoleName = role.RoleName
	admin.Role.Status = role.Status

	return true, admin
}

func GetAllAdminUsers() ([]*models.Admin, error) {
	db := database.CreateCon()
	res, err := db.Query("select id,firstname,lastname,username,email,role_id,created_at,status from `ico_user_admin` where role_id != 2;")
	if err != nil {
		return nil, err
	}

	admins := []*models.Admin{}
	for res.Next() {
		var admin models.Admin
		_ = res.Scan(&admin.ID, &admin.FirstName, &admin.LastName, &admin.UserName, &admin.Email, &admin.Role.ID, &admin.CreatedAt, &admin.Status)
		adminID := strconv.Itoa(int(admin.ID))
		metas, _ := db.Query("select meta_key,meta_value from ico_admin_meta where meta_key in ('address','phone','mobile','avatar','birth_date') and admin_id=" + adminID + ";")
		for metas.Next() {
			var meta models.MetaInfo
			_ = metas.Scan(&meta.MetaKey, &meta.MetaValue)
			if meta.MetaKey == "address" {
				admin.Meta.Address = meta.MetaValue
			}
			if meta.MetaKey == "phone" {
				admin.Meta.Phone = meta.MetaValue
			}
			if meta.MetaKey == "mobile" {
				admin.Meta.Mobile = meta.MetaValue
			}
			if meta.MetaKey == "avatar" {
				admin.Meta.Avatar = meta.MetaValue
			}
			if meta.MetaKey == "birth_date" {
				admin.Meta.BirthDate = meta.MetaValue
			}
		}
		role := getRoleByID(admin.Role.ID)
		admin.Role.RoleName = role.RoleName
		admin.Role.Status = role.Status
		admins = append(admins, &admin)
	}

	return admins, nil
}

func getUserMeta(userID string) {

}

func getAdminMeta(adminID string) {

}
