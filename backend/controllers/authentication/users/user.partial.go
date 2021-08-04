package users

import (
	"github.com/khorasany/coffee/api/backend/models"
	"strconv"
	"time"
)

func UserAdminParamToModel(param map[string]string) *models.Admin {
	id, _ := strconv.Atoi(param["id"])
	roleID, _ := strconv.Atoi(param["role_id"])
	status, _ := strconv.Atoi(param["status"])
	roleStatus, _ := strconv.Atoi(param["role_status"])
	return &models.Admin{
		ID:        int64(id),
		FirstName: param["firstname"],
		LastName:  param["lastname"],
		UserName:  param["username"],
		Password:  param["password"],
		Email:     param["email"],
		CreatedAt: time.Now().String(),
		Status:    int64(status),
		Role: models.Role{
			ID:       int64(roleID),
			RoleName: param["role_name"],
			Status:   int64(roleStatus),
		},
		Meta: models.Meta{
			BirthDate: CheckValue(param["birth_date"]),
			Address:   CheckValue(param["address"]),
			Phone:     CheckValue(param["phone"]),
			Mobile:    CheckValue(param["mobile"]),
			Avatar:    CheckValue(param["avatar"]),
		},
	}
}

func UsersAdminParamToModel(param map[string]string) []*models.Admin {
	admins := []*models.Admin{}
	for _, item := range param {
		admin := models.Admin{}
		id, _ := strconv.Atoi(item)
		admin.ID = int64(id)
		admins = append(admins, &admin)
	}

	return admins
}

func RoleParamToModel(param map[string]string) *models.Role {
	roleID, _ := strconv.Atoi(param["id"])
	status, _ := strconv.Atoi(param["status"])
	return &models.Role{
		ID:       int64(roleID),
		RoleName: param["role"],
		Status:   int64(status),
	}
}

func CheckValue(input string) string {
	if len(input) < 1 {
		input = ""
		return input
	}
	return input
}
