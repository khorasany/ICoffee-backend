package permissions

import (
	"github.com/khorasany/coffee/api/backend/database"
	"github.com/khorasany/coffee/api/backend/models"
	"strconv"
)

func InsertToken(jwtToken string) error {
	db := database.CreateCon()
	_, err := db.Exec("insert into ico_jwt_token (token) value ('" + jwtToken + "');")
	if err != nil {
		return err
	}

	return nil
}

func SetRole(role models.Admin) (int64, error) {
	db := database.CreateCon()
	roleStatus := strconv.Itoa(int(role.Role.Status))
	roleQuery, err := db.Exec("insert into ico_roles (role_name,status) values (" + role.Role.RoleName + "," + roleStatus + ");")
	if err != nil {
		return 0, err
	}
	roleID, _ := roleQuery.LastInsertId()

	return roleID, nil
}

func GetRole(roleName string) (string, error) {
	db := database.CreateCon()
	var role models.Role
	err := db.QueryRow("select id from ico_roles where role_name='" + roleName + "';").Scan(&role.ID)
	if err != nil {
		return "", err
	}
	roleID := strconv.Itoa(int(role.ID))

	return roleID, nil
}

func GetRoleInfo(roleName string) (*models.Role, error) {
	db := database.CreateCon()
	var role models.Role
	err := db.QueryRow("select * from ico_roles where role_name='"+roleName+"';").Scan(&role.ID, &role.RoleName, &role.Status)
	if err != nil {
		return &models.Role{}, err
	}

	return &models.Role{
		ID:       role.ID,
		RoleName: role.RoleName,
		Status:   role.Status,
	}, nil
}

func GetRoles() ([]*models.Role, error) {
	roles := []*models.Role{}
	db := database.CreateCon()
	result, err := db.Query("select * from ico_roles where role_name not in ('super-admin');")
	if err != nil {
		return nil, err
	}
	for result.Next() {
		role := models.Role{}
		_ = result.Scan(&role.ID, &role.RoleName, &role.Status)
		roles = append(roles, &role)
	}

	return roles, nil
}

func RegisterRole(role *models.Role) error {
	status := strconv.Itoa(int(role.Status))
	db := database.CreateCon()
	_, err := db.Exec("insert into ico_roles (role_name,status) values ('" + role.RoleName + "'," + status + ");")
	if err != nil {
		return err
	}

	return nil
}

func CheckValidRole(roleName string) error {
	var role models.Role
	var err error
	if roleName == "super-user" || roleName == "admin" {
		return err
	}

	db := database.CreateCon()
	err = db.QueryRow("select role_name from ico_roles where role_name='" + roleName + "';").Scan(&role.RoleName)
	if len(role.RoleName) > 1 && role.RoleName == roleName {
		return err
	}

	return nil
}
