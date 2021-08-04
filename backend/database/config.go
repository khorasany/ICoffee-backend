package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/khorasany/coffee/api/backend/globals"
)

func CreateCon() *sql.DB {
	db, err := sql.Open(
		"mysql", globals.MYSQL_DATABASE_USERNAME+":"+globals.MYSQL_DATABASE_PASSWORD+"@tcp("+globals.MYSQL_IP_ADDRESS+":"+globals.MYSQL_PORT+")/"+globals.MYSQL_DATABASE_NAME)

	if err != nil {
		panic(err.Error())
	}

	return db
}

//func chooseDatabase() *sql.DB {
//	db := CreateCon()
//	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS `" + globals.MYSQL_DATABASE_NAME + "`;")
//	if err != nil {
//		panic(err.Error())
//	}
//
//	_, err = db.Exec("USE " + globals.MYSQL_DATABASE_NAME + ";")
//	if err != nil {
//		panic(err.Error())
//	}
//
//	return db
//}

//func createTable(tableName string, statement string) {
//	db := chooseDatabase()
//
//	_, checkTable := db.Query("SELECT * FROM information_schema.tables WHERE table_schema = " + globals.MYSQL_DATABASE_NAME + " AND table_name = " + tableName + " LIMIT 1;")
//	if checkTable != nil {
//		table, err := db.Prepare(statement)
//		if err != nil {
//			panic(err.Error())
//		}
//		table.Exec()
//	}
//}

//func CreateTables() {
//	createTable("ico_users",
//		"CREATE TABLE ico_users (id int NOT NULL AUTO_INCREMENT,firstname varchar(255)," +
//		"lastname varchar(255),username varchar(255),password varchar(50),email varchar(255)," +
//		"created_at varchar(255),status tinyint(2),PRIMARY KEY (id));")
//	createTable("ico_usermeta",
//		"CREATE TABLE ico_usermeta (id int NOT NULL AUTO_INCREMENT,user_id int NOT NULL," +
//		"meta_key varchar(255),meta_value longtext,PRIMARY KEY (id));")
//	createTable("ico_user_admin",
//		"CREATE TABLE ico_user_admin (id int NOT NULL AUTO_INCREMENT,firstname varchar(255)," +
//		"lastname varchar(255),username varchar(255),password varchar(255),email varchar(255)," +
//		"role_id int,created_at varchar(255),status tinyint(2),PRIMARY KEY (id));")
//	createTable("ico_admin_meta",
//		"CREATE TABLE ico_admin_meta (id int NOT NULL AUTO_INCREMENT,admin_id int NOT NULL," +
//		"meta_key varchar(255),meta_value longtext,PRIMARY KEY (id));")
//}

//func CheckRecordExists(recordID string, tableName string) bool {
//	db := CreateCon()
//	recordExist := db.QueryRow("select * from " + tableName + " where id=" + recordID + ";")
//	if recordExist != nil {
//		return true
//	}
//
//	return false
//}
//
//func metaDataExists(metaType string, typeID int64, metaKey string) (int64, error) {
//	db := CreateCon()
//	var scanStatement *sql.Row
//	var tbName string
//
//	stringID := strconv.Itoa(int(typeID))
//	err := db.QueryRow("select * from " + tbName + " where meta_key=" + metaKey + " and " + metaType + "_id=" + stringID + ";")
//	if err != nil {
//		return 0, err.Err()
//	}
//
//	if metaType == "admin" {
//		var adminMeta user.AdminMeta
//		_ = scanStatement.Scan(&adminMeta.ID, &adminMeta.AdminID, &adminMeta.MetaKey, &adminMeta.MetaValue)
//		tbName = "ico_admin_meta"
//
//		return adminMeta.ID, nil
//	}
//
//	if metaType == "user" {
//		var userMeta user.UserMeta
//		_ = scanStatement.Scan(&userMeta.ID, &userMeta.UserID, &userMeta.MetaKey, &userMeta.MetaValue)
//		tbName = "ico_usermeta"
//
//		return userMeta.ID, nil
//	}
//
//	return 0, nil
//}
