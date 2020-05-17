package driver

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/pucsd2020-pp/rest-api/config"
	"github.com/pucsd2020-pp/rest-api/model"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MYSQL_DRIVER_NAME   = "mysql"
	CONN_MAX_LIFETIME   = 30 * 60 * 60 // 30 day
	COLUMN_INGNORE_FLAG = "1"
	COLUMN_PRIMARY      = "primary"
)

func NewMysqlConnection(cfg config.MysqlConnection) (*sql.DB, error) {
	db, err := sql.Open(MYSQL_DRIVER_NAME, cfg.ConnString())
	if err != nil {
		log.Fatalf("Failed to open mysql connection: %v", err)
		return nil, err
	}

	if cfg.IdleConnection > 0 {
		db.SetMaxIdleConns(cfg.IdleConnection)
	}
	if cfg.MaxConnection > 0 {
		db.SetMaxOpenConns(cfg.MaxConnection)
	}
	db.SetConnMaxLifetime(time.Second * CONN_MAX_LIFETIME)

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping mysql: %v", err)
	}

	return db, err
}

// return the placeholder string with given count
func GetPlaceHolder(count int) string {
	if count > 0 {
		str := strings.Repeat("?, ", count)
		return str[:len(str)-2]
	}

	return ""
}

/**
 * Insert new row
 */
func Create(conn *sql.DB, object model.IModel) (sql.Result, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	var params []interface{}
	var sessionKey string
	count := 0
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		value := rValue.Elem().Field(idx)

		// if value.IsNil() || COLUMN_INGNORE_FLAG == field.Tag.Get("autoincr") ||
		// 	COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
		// 	continue
		// }

		column := field.Tag.Get("column")
		// fmt.Println(column)
		if column == "sessionKey" {
			sessionKey = value.String()
		}
		if column != "sessionKey" {
			columns = append(columns, column)
			params = append(params, value.Interface())
			count++
		}
	}
	tag := model.NotPermit{}
	// fmt.Println(sessionKey)
	err := conn.QueryRow("select sessionKey from usersKey where sessionKey ='" + sessionKey + "';").Scan(&tag.Msg)
	// log.Printf(tag.Msg)
	if err != nil {
		return nil, errors.New("session Expire")
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("INSERT INTO ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString("(")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(") VALUES(")
	queryBuffer.WriteString(GetPlaceHolder(count))
	queryBuffer.WriteString(");")

	query := queryBuffer.String()
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Insert Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(params...)
	if nil != err {
		log.Printf("Insert Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, err
	}

	return result, nil

	// tag.Msg = "session Expire"

}

func CreateFileFolder(conn *sql.DB, object model.IModel) (sql.Result, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	var params []interface{}
	columns1 := []string{}
	var params1 []interface{}
	count := 0
	cnt := 0
	var userid, filefolderPath, sessionKey string
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		value := rValue.Elem().Field(idx)
		column := field.Tag.Get("column")
		// fmt.Println(column)
		if column != "sessionKey" && column != "userId" {
			columns = append(columns, column)
			params = append(params, value.Interface())
			if column == "filefolderPath" {
				filefolderPath = value.String()
			}
			columns1 = append(columns1, column)
			params1 = append(params1, value.Interface())
			count++
			cnt++
		}
		if column == "sessionKey" {
			sessionKey = value.String()
		}
		if column == "userId" {
			columns1 = append(columns1, column)
			params1 = append(params1, value.Interface())
			cnt++
			userid = value.String()
		}
		// fmt.Println(params, columns)
	}
	tag := model.NotPermit{}
	// fmt.Println(sessionKey)
	err := conn.QueryRow("select sessionKey from usersKey where sessionKey ='" + sessionKey + "';").Scan(&tag.Msg)
	// log.Printf(tag.Msg)
	if err != nil {
		return nil, errors.New("session Expire")
	}
	// var writepermission int
	obj1 := new(model.ScanData)
	_ = conn.QueryRow("select count(*) as count from userPermission where userId='" + userid + "'and filefolderPath= '" + filefolderPath + "' and permissionValue='w';").Scan(&obj1.WritePermissionusr)
	// fmt.Println("write permission :", obj1.WritePermissionusr, userid, filefolderPath)
	_ = conn.QueryRow("select count(*) as count from (select permissionValue from groupPermission where groupName IN (select groupName from userGroupMap where userId='" + userid + "') AND filefolderPath='" + filefolderPath + "' AND permissionValue='w') AS writePermission;").Scan(&obj1.WritePermissiongrp)
	// fmt.Println("write permission :", obj1.WritePermissiongrp, userid, filefolderPath)
	_ = conn.QueryRow("select userType from users where userId='" + userid + "';").Scan(&obj1.CheckUserType)
	fmt.Println(obj1.CheckUserType)
	if obj1.WritePermissionusr > 0 || obj1.WritePermissiongrp > 0 || obj1.CheckUserType == "s" {
		var queryBuffer bytes.Buffer
		queryBuffer.WriteString("INSERT INTO ")
		queryBuffer.WriteString(object.Table())
		queryBuffer.WriteString("(")
		queryBuffer.WriteString(strings.Join(columns, ", "))
		queryBuffer.WriteString(") VALUES(")
		queryBuffer.WriteString(GetPlaceHolder(count))
		queryBuffer.WriteString(");")

		query := queryBuffer.String()
		stmt, err := conn.Prepare(query)
		if nil != err {
			log.Printf("Insert Syntax Error: %s\n\tError Query: %s : %s\n",
				err.Error(), object.String(), query)
			return nil, err
		}

		defer stmt.Close()

		result, err := stmt.Exec(params...)
		if nil != err {
			log.Printf("Insert Execute Error: %s\nError Query: %s : %s\n",
				err.Error(), object.String(), query)
			return nil, err
		}
		var queryBuffer1 bytes.Buffer
		queryBuffer1.WriteString("INSERT INTO ")
		queryBuffer1.WriteString("userPermission")
		queryBuffer1.WriteString("(")
		queryBuffer1.WriteString(strings.Join(columns1, ", "))
		queryBuffer1.WriteString(") VALUES(")
		queryBuffer1.WriteString(GetPlaceHolder(cnt))
		queryBuffer1.WriteString(");")
		query1 := queryBuffer1.String()
		stmt1, err := conn.Prepare(query1)
		if nil != err {
			log.Printf("Insert Syntax Error: %s\n\tError Query: %s : %s\n",
				err.Error(), object.String(), query1)
			return nil, err
		}

		defer stmt1.Close()

		_, err = stmt1.Exec(params1...)
		if nil != err {
			log.Printf("Insert Execute Error: %s\nError Query: %s : %s\n",
				err.Error(), object.String(), query1)
			return nil, err
		}
		if rValue.Elem().Field(2).String() == "d" {
			_, fileerr := os.Stat(model.RootDir + rValue.Elem().Field(0).String() + rValue.Elem().Field(1).String())

			if os.IsNotExist(fileerr) {
				errDir := os.MkdirAll(model.RootDir+rValue.Elem().Field(0).String()+rValue.Elem().Field(1).String(), 0755)
				if errDir != nil {
					log.Fatal(fileerr)
				}

			}
		} else {

			f, err := os.Create(model.RootDir + rValue.Elem().Field(0).String() + rValue.Elem().Field(1).String())

			if err != nil {
				fmt.Println(err)
				f.Close()
			}

		}

		// fmt.Println(value.Interface())

		return result, nil
	} else {
		return nil, errors.New("You dont have permission of write") //create error massage
	}
}
func DeleteFileFolder(conn *sql.DB, object model.IModel) (sql.Result, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)
	var userid, filefolderPath, filefolderName, sessionKey string

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		value := rValue.Elem().Field(idx)
		column := field.Tag.Get("column")
		if column == "filefolderPath" {
			filefolderPath = value.String()
		}
		if column == "filefolderName" {
			filefolderName = value.String()
		}
		if column == "sessionKey" {
			sessionKey = value.String()
		}
		// if column == "filesOrFolderId" {
		// 	filesOrFolderId = value.String()
		// }
		if column == "userId" {
			userid = value.String()
		}
	}
	tag := model.NotPermit{}
	// fmt.Println(sessionKey)
	err := conn.QueryRow("select sessionKey from usersKey where sessionKey ='" + sessionKey + "';").Scan(&tag.Msg)
	// log.Printf(tag.Msg)
	if err != nil {
		return nil, errors.New("session Expire")
	}
	// fmt.Println(filefolderName)
	obj1 := new(model.ScanData)
	_ = conn.QueryRow("select count(*) as count from userPermission where userId='" + userid + "'and filefolderPath= '" + filefolderPath + "' and permissionValue='w';").Scan(&obj1.WritePermissionusr)
	// fmt.Println("write permission :", obj1.WritePermissionusr, userid, filefolderPath)
	_ = conn.QueryRow("select count(*) as count from (select permissionValue from groupPermission where groupName IN (select groupName from userGroupMap where userId='" + userid + "') AND filefolderPath='" + filefolderPath + "' AND permissionValue='w') AS writePermission;").Scan(&obj1.WritePermissiongrp)
	// fmt.Println("write permission :", obj1.WritePermissiongrp, userid, filefolderPath)
	_ = conn.QueryRow("select userType from users where userId='" + userid + "';").Scan(&obj1.CheckUserType)

	if obj1.WritePermissionusr > 0 || obj1.WritePermissiongrp > 0 || obj1.CheckUserType == "s" {
		var queryBuffer bytes.Buffer
		var str = " WHERE filefolderPath='" + filefolderPath + "' AND filefolderName='" + filefolderName + "' ; "
		queryBuffer.WriteString("DELETE FROM userPermission " + str)
		query := queryBuffer.String()
		stmt, err := conn.Prepare(query)
		if nil != err {
			log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n",
				err.Error(), object.String(), query)
			return nil, err
		}
		defer stmt.Close()
		result, err := stmt.Exec()
		if nil != err {
			log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
				err.Error(), object.String(), query)
		}
		var queryBuffer2 bytes.Buffer

		queryBuffer2.WriteString("DELETE FROM groupPermission " + str)
		query2 := queryBuffer2.String()
		stmt2, err := conn.Prepare(query2)
		if nil != err {
			log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n",
				err.Error(), object.String(), query2)
			return nil, err
		}
		defer stmt.Close()
		result, err2 := stmt2.Exec()
		if nil != err2 {
			log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
				err2.Error(), object.String(), query2)
		}
		var queryBuffer1 bytes.Buffer
		queryBuffer1.WriteString("DELETE FROM filesfolder " + str)
		query1 := queryBuffer1.String()
		stmt1, err := conn.Prepare(query1)
		if nil != err {
			log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n",
				err.Error(), object.String(), query1)
			return nil, err
		}
		defer stmt.Close()
		result, err1 := stmt1.Exec()
		if nil != err1 {
			log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
				err.Error(), object.String(), query1)
		}
		var filefold = model.RootDir + filefolderPath + filefolderName
		// if filesOrFolderId == "d" {
		_, fileerr := os.Stat(filefold)
		if os.IsNotExist(fileerr) {
			errDir := os.RemoveAll(filefold)
			if errDir != nil {
				log.Fatal(fileerr)
			}

		}
		// } else {

		// 	err = os.Remove(filefold)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}

		// }

		// fmt.Println(value.Interface())

		return result, nil
	} else {
		return nil, errors.New("You dont have permission of Delete") //create error massage
	}
}

//delete File and folder
/**
 * Update existing row with key column
 */
func UpdateById(conn *sql.DB, object model.IModel) error {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	var params []interface{}

	keyColumns := []string{}
	var keyParams []interface{}
	var sessionKey string
	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		value := rValue.Elem().Field(idx)

		// if value.IsNil() ||
		// 	COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
		// 	continue
		// }

		column := field.Tag.Get("column")
		if column != "sessionKey" {
			sessionKey = value.String()
		}
		if column != "sessionKey" {
			if COLUMN_PRIMARY == field.Tag.Get("key") {
				keyColumns = append(keyColumns, column+" = ?")
				keyParams = append(keyParams, value.Interface())

			} else {
				columns = append(columns, column+" = ?")
				params = append(params, value.Interface())
			}
		}
	}

	for _, param := range keyParams {
		params = append(params, param)
	}
	tag := model.NotPermit{}
	// fmt.Println(sessionKey)
	err := conn.QueryRow("select sessionKey from usersKey where sessionKey ='" + sessionKey + "';").Scan(&tag.Msg)
	// log.Printf(tag.Msg)
	if err != nil {
		return errors.New("session Expire")
	}
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("UPDATE ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" SET ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" WHERE ")
	queryBuffer.WriteString(strings.Join(keyColumns, ", "))
	queryBuffer.WriteString(";")

	query := queryBuffer.String()
	//	log.Println("Update statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Update Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(params...)
	if nil != err {
		log.Printf("Update Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}

	return err
}
func ChangePermission(conn *sql.DB, object model.IModel) (interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	var whocallToChange, sessionKey, permissionValue, filesOrFolderId, useridOrGroupId, filefolderName, filefolderPath string

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		value := rValue.Elem().Field(idx)
		column := field.Tag.Get("column")
		if column == "whocallToChange" {
			whocallToChange = value.String()
		}
		if column == "filefolderPath" {
			filefolderPath = value.String()
		}
		if column == "filefolderName" {
			filefolderName = value.String()
		}
		if column == "useridOrGroupId" {
			useridOrGroupId = value.String()
		}
		if column == "filesOrFolderId" {
			filesOrFolderId = value.String()
		}
		if column == "permissionValue" {
			permissionValue = value.String()
		}
		if column == "sessionKey" {
			sessionKey = value.String()
		}
	}
	// tag := model.NotPermit{}
	tag := model.NotPermit{}
	// fmt.Println(sessionKey)
	err := conn.QueryRow("select sessionKey from usersKey where sessionKey ='" + sessionKey + "';").Scan(&tag.Msg)
	// log.Printf(tag.Msg)
	if err != nil {
		return nil, errors.New("session Expire")
	}
	_ = conn.QueryRow("select userType from users where userId='" + whocallToChange + "';").Scan(&tag.Msg)
	fmt.Println(tag.Msg, whocallToChange)
	if tag.Msg == "s" {
		// fmt.Println(whocallToChange, permissionValue)
		var queryBuffer bytes.Buffer
		queryBuffer.WriteString("UPDATE ")
		queryBuffer.WriteString("userPermission")
		queryBuffer.WriteString(" SET ")
		queryBuffer.WriteString("permissionValue='" + permissionValue + "' ")
		queryBuffer.WriteString(" WHERE ")
		queryBuffer.WriteString("filefolderPath='" + filefolderPath + "' AND filefolderName='" + filefolderName + "' AND userId='" + useridOrGroupId)
		queryBuffer.WriteString("' AND filesOrFolderId='" + filesOrFolderId + "'")
		queryBuffer.WriteString(";")

		query := queryBuffer.String()
		// fmt.Println(query)
		//	log.Println("Update statement is: %s", query)

		stmt, err := conn.Prepare(query)
		if nil != err {
			log.Printf("Update Syntax Error: %s\n\tError Query: %s : %s\n",
				err.Error(), object.String(), query)
			return nil, err
		}

		defer stmt.Close()
		_, err = stmt.Exec()
		if nil != err {
			log.Printf("Update Execute Error: %s\nError Query: %s : %s\n",
				err.Error(), object.String(), query)
		}

		var queryBuffer1 bytes.Buffer
		queryBuffer1.WriteString("UPDATE ")
		queryBuffer1.WriteString("groupPermission")
		queryBuffer1.WriteString(" SET ")
		queryBuffer1.WriteString("permissionValue='" + permissionValue + "' ")
		queryBuffer1.WriteString(" WHERE ")
		queryBuffer1.WriteString("filefolderPath='" + filefolderPath + "' AND filefolderName='" + filefolderName + "' AND groupName='" + useridOrGroupId)
		queryBuffer1.WriteString("' AND filesOrFolderId='" + filesOrFolderId + "'")
		queryBuffer1.WriteString(";")

		query1 := queryBuffer1.String()
		// fmt.Println(query)
		//	log.Println("Update statement is: %s", query)

		stmt1, err := conn.Prepare(query1)
		if nil != err {
			log.Printf("Update Syntax Error: %s\n\tError Query: %s : %s\n",
				err.Error(), object.String(), query1)
			return nil, err
		}

		defer stmt.Close()
		_, err = stmt1.Exec()
		if nil != err {
			log.Printf("Update Execute Error: %s\nError Query: %s : %s\n",
				err.Error(), object.String(), query1)
		}
		tag.Msg = "update success"
		return tag, err
	}
	tag.Msg = "you are Not Authorized user to Change Permission"
	return tag, nil
}
func GetGroupById(conn *sql.DB, object model.IModel, id string) ([]interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	//pointers := make([]interface{}, 0)

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}
		column := field.Tag.Get("column")
		if column != "sessionKey" {
			columns = append(columns, column)
		}
		//pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}

	var queryBuffer bytes.Buffer
	// var params []interface{}
	tag := model.NotPermit{}

	_ = conn.QueryRow("select userType from users where userId='" + id + "';").Scan(&tag.Msg)
	// fmt.Println(tag.Msg, whocallToChange)
	if tag.Msg == "s" {
		queryBuffer.WriteString("SELECT ")
		queryBuffer.WriteString(strings.Join(columns, ", "))
		queryBuffer.WriteString(" FROM ")
		queryBuffer.WriteString(object.Table())
		queryBuffer.WriteString(" ;")
	} else {
		queryBuffer.WriteString("SELECT ")
		queryBuffer.WriteString(strings.Join(columns, ", "))
		queryBuffer.WriteString(" FROM ")
		queryBuffer.WriteString(object.Table())
		queryBuffer.WriteString(" WHERE userId= '" + id + "';")
	}
	query := queryBuffer.String()
	row, err := conn.Query(query)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	objects := make([]interface{}, 0)
	recds, err := row.Columns()
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		values := make([]interface{}, len(recds))
		recdsWrite := make([]string, len(recds))
		for index, _ := range recds {
			values[index] = &recdsWrite[index]
		}
		err = row.Scan(values...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}

		objects = append(objects, values)

	}

	return objects, nil

}

func GetUserByGroupId(conn *sql.DB, object model.IModel, id string) ([]interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	//pointers := make([]interface{}, 0)

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}
		column := field.Tag.Get("column")
		if column != "sessionKey" {
			columns = append(columns, column)
		}
		//pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}

	var queryBuffer bytes.Buffer
	// // var params []interface{}
	// tag := model.NotPermit{}

	// _ = conn.QueryRow("select userType from users where userId='" + id + "';").Scan(&tag.Msg)
	// fmt.Println(tag.Msg, whocallToChange)
	// if tag.Msg == "s" {
	// 	queryBuffer.WriteString("SELECT ")
	// 	queryBuffer.WriteString(strings.Join(columns, ", "))
	// 	queryBuffer.WriteString(" FROM ")
	// 	queryBuffer.WriteString(object.Table())
	// 	queryBuffer.WriteString(" ;")
	// } else {
	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE groupName= '" + id + "';")
	// }
	query := queryBuffer.String()
	row, err := conn.Query(query)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	objects := make([]interface{}, 0)
	recds, err := row.Columns()
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		values := make([]interface{}, len(recds))
		recdsWrite := make([]string, len(recds))
		for index, _ := range recds {
			values[index] = &recdsWrite[index]
		}
		err = row.Scan(values...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}

		objects = append(objects, values)

	}

	return objects, nil

}

func GetById(conn *sql.DB, object model.IModel, id string) (model.IModel, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	pointers := make([]interface{}, 0)

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		if column != "sessionKey" {
			columns = append(columns, column)
			pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
		}
	}

	var queryBuffer bytes.Buffer

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE userId = ?")

	query := queryBuffer.String()
	fmt.Println("getByID method :- ", query)

	//	log.Printf("GetById sql: %s\n", query)
	row, err := conn.Query(query, id)

	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}

	defer row.Close()
	if row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}

		err = row.Scan(pointers...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}
	} else {
		return nil, errors.New(fmt.Sprintf("Entry not found for id: %d", id))
	}

	return object, nil
}

func GetByuserId(conn *sql.DB, object model.IModel, id string) (model.IModel, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	pointers := make([]interface{}, 0)

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
		pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}

	var queryBuffer bytes.Buffer

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE userId = ?")

	query := queryBuffer.String()
	//	log.Printf("GetById sql: %s\n", query)
	row, err := conn.Query(query, id)

	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}

	defer row.Close()
	if row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}

		err = row.Scan(pointers...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}
	} else {
		return nil, errors.New(fmt.Sprintf("Entry not found for id: %d", id))
	}

	return object, nil
}

func GetAllGroups(conn *sql.DB, object model.IModel, limit, offset int64) ([]interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	//pointers := make([]interface{}, 0)

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		if column != "sessionKey" {
			columns = append(columns, column)
		}
		//pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}

	var queryBuffer bytes.Buffer
	var params []interface{}

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.Table())
	if 0 != limit && 0 != offset {
		queryBuffer.WriteString(" LIMIT ? OFFSET ?")
		params = append(params, limit)
		params = append(params, offset)
	}

	query := queryBuffer.String()
	row, err := conn.Query(query)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	objects := make([]interface{}, 0)
	recds, err := row.Columns()
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		values := make([]interface{}, len(recds))
		recdsWrite := make([]string, len(recds))
		for index, _ := range recds {
			values[index] = &recdsWrite[index]
		}
		err = row.Scan(values...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}

		objects = append(objects, values)

	}

	return objects, nil
}
func GetAll(conn *sql.DB, object model.IModel, limit, offset int64) ([]interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	//pointers := make([]interface{}, 0)

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		if column != "sessionKey" {
			columns = append(columns, column)
		}
		//pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}

	var queryBuffer bytes.Buffer
	var params []interface{}

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.Table())
	if 0 != limit && 0 != offset {
		queryBuffer.WriteString(" LIMIT ? OFFSET ?")
		params = append(params, limit)
		params = append(params, offset)
	}

	query := queryBuffer.String()
	row, err := conn.Query(query)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	objects := make([]interface{}, 0)
	recds, err := row.Columns()
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		values := make([]interface{}, len(recds))
		recdsWrite := make([]string, len(recds))
		for index, _ := range recds {
			values[index] = &recdsWrite[index]
		}
		err = row.Scan(values...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}

		objects = append(objects, values)

	}

	return objects, nil
}

func DeleteById(conn *sql.DB, object model.IModel, id string) (sql.Result, error) {
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("DELETE FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE id = ?")

	query := queryBuffer.String()
	//	log.Println("Delete statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, err
	}

	defer stmt.Close()
	result, err := stmt.Exec(id)
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}

	return result, err
}

func SoftDeleteById(conn *sql.DB, object model.IModel, id string) error {
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString(" DELETE FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE userId = ?")

	query := queryBuffer.String()
	//	log.Println("Delete statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(id)
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}

	return err
}

func SetAll(conn *sql.DB, object model.IModel, limit, offset int64) ([]interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	//pointers := make([]interface{}, 0)

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
		//pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}

	var queryBuffer bytes.Buffer
	var params []interface{}

	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.Table())
	if 0 != limit && 0 != offset {
		queryBuffer.WriteString(" LIMIT ? OFFSET ?")
		params = append(params, limit)
		params = append(params, offset)
	}

	query := queryBuffer.String()
	row, err := conn.Query(query)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	objects := make([]interface{}, 0)
	recds, err := row.Columns()
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		values := make([]interface{}, len(recds))
		recdsWrite := make([]string, len(recds))
		for index, _ := range recds {
			values[index] = &recdsWrite[index]
		}
		err = row.Scan(values...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}

		objects = append(objects, values)

	}

	return objects, nil
}

func Authentication(conn *sql.DB, object model.IModel) (interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	pointers := make([]interface{}, 0)

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		columns = append(columns, column)
		pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT ")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(" FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE userId = '")
	queryBuffer.WriteString(rValue.Elem().Field(0).String())
	queryBuffer.WriteString("' AND password = '")
	queryBuffer.WriteString(rValue.Elem().Field(1).String())
	queryBuffer.WriteString("' ; ")

	query := queryBuffer.String()
	// log.Printf(query, id, pass)
	row, err := conn.Query(query)
	if nil != err {
		log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
		return nil, err
	}

	defer row.Close()
	if row.Next() {
		if nil != err {
			log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}

		err = row.Scan(pointers...)
		if nil != err {
			log.Printf("Error: row.Scan: %s\n", err.Error())
			return nil, err
		}
	} else {
		return nil, errors.New(fmt.Sprintf("Please Check userId: %s OR  Password", rValue.Elem().Field(0)))
	}
	var query1 string = "INSERT INTO usersKey (userId,sessionKey) VALUES ('" + rValue.Elem().Field(0).String() + "',replace(uuid(),'-',''));"
	stmt, err := conn.Prepare(query1)

	if nil != err {
		log.Printf("Insert Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Exec()
	if nil != err {
		log.Printf("Insert Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query, result)
		return nil, err
	}
	tag := model.Key{}
	err = conn.QueryRow("SELECT userId, sessionKey FROM usersKey where userId = ?  ORDER BY logDate DESC  LIMIT 1", rValue.Elem().Field(0).String()).Scan(&tag.UserId, &tag.SessionKey)
	log.Printf(tag.UserId, tag.SessionKey)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return tag, nil
}
func GetFilesFold(conn *sql.DB, object model.IModel, limit, offset int64) ([]interface{}, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	var filefolderPath, userId, sessionKey string

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		value := rValue.Elem().Field(idx)

		if COLUMN_INGNORE_FLAG == field.Tag.Get("ignore") {
			continue
		}

		column := field.Tag.Get("column")
		if column == "userId" {
			userId = value.String()
		}
		if column == "sessionKey" {
			sessionKey = value.String()
		}
		if column != "sessionKey" && column != "userId" {

			if column == "filefolderPath" {
				filefolderPath = value.String()
			}

			columns = append(columns, column)
			//pointers = append(pointers, rValue.Elem().Field(idx).Addr().Interface())
		}
	}
	tag := model.NotPermit{}
	// fmt.Println(sessionKey)
	err := conn.QueryRow("select sessionKey from usersKey where sessionKey ='" + sessionKey + "';").Scan(&tag.Msg)
	// log.Printf(tag.Msg)
	if err != nil {
		return nil, errors.New("session Expire")
	}

	var queryBuffer bytes.Buffer
	// var params []interface{}

	obj1 := new(model.SelFileFold)
	_ = conn.QueryRow("select userType from users where userId='" + userId + "';").Scan(&obj1.UserType)
	// fmt.Println("select userType from users where userId='" + userId + "';")
	// fmt.Println(obj1.UserType)
	if obj1.UserType == "s" {

		queryBuffer.WriteString("select * from filesfolder where filefolderPath='" + filefolderPath + "';")
		query := queryBuffer.String()
		// fmt.Println(query)

		row, err := conn.Query(query)
		if nil != err {
			log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		objects := make([]interface{}, 0)
		recds, err := row.Columns()
		if nil != err {
			log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		defer row.Close()
		for row.Next() {
			if nil != err {
				log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
				return nil, err
			}
			values := make([]interface{}, len(recds))
			recdsWrite := make([]string, len(recds))
			for index, _ := range recds {
				values[index] = &recdsWrite[index]
			}
			err = row.Scan(values...)
			if nil != err {
				log.Printf("Error: row.Scan: %s\n", err.Error())
				return nil, err
			}

			objects = append(objects, values)

		}

		return objects, nil

	} else {

		queryBuffer.WriteString("select * from groupPermission where  groupName IN (select groupName from userGroupMap where userId='" + userId + "') ;")
		// queryBuffer.WriteString("select * from groupPermission where filefolderPath='" + filefolderPath + "' AND groupName IN (select groupName from userGroupMap where userId='" + userId + "') ;")
		query := queryBuffer.String()
		// fmt.Println(query)
		row, err := conn.Query(query)
		if nil != err {
			log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		objects := make([]interface{}, 0)
		recds, err := row.Columns()
		if nil != err {
			log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query)
			return nil, err
		}
		defer row.Close()
		for row.Next() {
			if nil != err {
				log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query)
				return nil, err
			}
			values := make([]interface{}, len(recds))
			recdsWrite := make([]string, len(recds))
			for index, _ := range recds {
				values[index] = &recdsWrite[index]
			}
			err = row.Scan(values...)
			if nil != err {
				log.Printf("Error: row.Scan: %s\n", err.Error())
				return nil, err
			}
			objects = append(objects, values)

		}
		var queryBuffer1 bytes.Buffer
		queryBuffer1.WriteString("select * from userPermission where  userId='" + userId + "';")
		// queryBuffer1.WriteString("select * from userPermission where filefolderPath='" + filefolderPath + "' AND userId='" + userId + "';")
		query1 := queryBuffer1.String()
		// fmt.Println(query)
		row1, err := conn.Query(query1)
		if nil != err {
			log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query1)
			return nil, err
		}
		// objects1 := make([]interface{}, 0)
		recds1, err := row1.Columns()
		if nil != err {
			log.Printf("Error conn.Query: %s\n\tError Query: %s\n", err.Error(), query1)
			return nil, err
		}
		defer row1.Close()
		for row1.Next() {
			if nil != err {
				log.Printf("Error row.Columns(): %s\n\tError Query: %s\n", err.Error(), query1)
				return nil, err
			}
			values1 := make([]interface{}, len(recds))
			recdsWrite1 := make([]string, len(recds))
			for index, _ := range recds1 {
				values1[index] = &recdsWrite1[index]
			}
			err = row1.Scan(values1...)
			if nil != err {
				log.Printf("Error: row.Scan: %s\n", err.Error())
				return nil, err
			}

			objects = append(objects, values1)
		}

		return objects, nil

	}

}

func Logout(conn *sql.DB, object model.IModel) (interface{}, error) {
	rValue := reflect.ValueOf(object)
	// rType := reflect.TypeOf(object)
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("DELETE FROM ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString(" WHERE userId = '")
	queryBuffer.WriteString(rValue.Elem().Field(0).String())
	queryBuffer.WriteString("' AND sessionKey = '")
	queryBuffer.WriteString(rValue.Elem().Field(1).String())
	queryBuffer.WriteString("' ; ")

	query := queryBuffer.String()
	//	log.Println("Delete statement is: %s", query)
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Delete Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, err
	}

	defer stmt.Close()
	_, err = stmt.Exec()
	if nil != err {
		log.Printf("Delete Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
	}
	tag := model.NotPermit{}
	tag.Msg = "logout success"
	return tag, err
}

func CreateGroup(conn *sql.DB, object model.IModel) (sql.Result, error) {
	rValue := reflect.ValueOf(object)
	rType := reflect.TypeOf(object)

	columns := []string{}
	var params []interface{}
	columns1 := []string{}
	var params1 []interface{}
	count := 0
	var groupName, userId, sessionKey string

	for idx := 0; idx < rValue.Elem().NumField(); idx++ {
		field := rType.Elem().Field(idx)
		value := rValue.Elem().Field(idx)

		column := field.Tag.Get("column")
		if column == "userId" {
			userId = value.String()
		}
		if column == "groupName" {
			groupName = value.String()
		}
		if column == "sessionKey" {
			sessionKey = value.String()
		}
		if column != "userId" && column != "sessionKey" {
			// fmt.Println(column)

			columns = append(columns, column)
			params = append(params, value.Interface())
			// log.Printf("string of parameters:- ", value.Interface())
		}
		if column != "sessionKey" && column != "groupDescription" {
			columns1 = append(columns1, column)
			params1 = append(params1, value.Interface())
		}
		count++

	}
	tag := model.NotPermit{}
	// fmt.Println(sessionKey)
	err := conn.QueryRow("select sessionKey from usersKey where sessionKey ='" + sessionKey + "';").Scan(&tag.Msg)
	// log.Printf(tag.Msg)
	if err != nil {
		return nil, errors.New("session Expire")
	}
	// fmt.Println("length of params:--  ", params)
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("INSERT INTO ")
	queryBuffer.WriteString(object.Table())
	queryBuffer.WriteString("(")
	queryBuffer.WriteString(strings.Join(columns, ", "))
	queryBuffer.WriteString(") VALUES(")
	queryBuffer.WriteString(GetPlaceHolder(count - 2))
	queryBuffer.WriteString(");")

	query := queryBuffer.String()
	stmt, err := conn.Prepare(query)
	if nil != err {
		log.Printf("Insert Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(params...)
	if nil != err {
		log.Printf("Insert Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query)
		return nil, err
	}

	var queryBuffer1 bytes.Buffer
	queryBuffer1.WriteString("INSERT INTO ")
	queryBuffer1.WriteString("whoGroupCreated ")
	queryBuffer1.WriteString("(")
	queryBuffer1.WriteString(strings.Join(columns1, ", "))
	queryBuffer1.WriteString(") VALUES(")
	queryBuffer1.WriteString(GetPlaceHolder(count - 2))
	queryBuffer1.WriteString(");")

	query1 := queryBuffer1.String()
	stmt1, err := conn.Prepare(query1)
	if nil != err {
		log.Printf("Insert Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query1)
		return nil, err
	}

	defer stmt1.Close()
	// result1, err := stmt1.Exec(params1...)
	_, err = stmt1.Exec(params1...)
	// fmt.Println(result1)
	if nil != err {
		log.Printf("Insert Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query1)
		return nil, err
	}
	var queryBuffer2 bytes.Buffer
	queryBuffer2.WriteString("INSERT INTO ")
	queryBuffer2.WriteString("userGroupMap ")
	queryBuffer2.WriteString("( groupName,userId")
	// queryBuffer2.WriteString(strings.Join(columns1, ", "))
	queryBuffer2.WriteString(") VALUES('" + groupName + "','" + userId + "');")
	// queryBuffer2.WriteString()
	// queryBuffer2.WriteString()

	query2 := queryBuffer2.String()
	stmt2, err := conn.Prepare(query2)
	if nil != err {

		log.Printf("Insert Syntax Error: %s\n\tError Query: %s : %s\n",
			err.Error(), object.String(), query2)
		return nil, err
	}

	defer stmt1.Close()
	// result1, err := stmt1.Exec(params1...)
	_, err = stmt2.Exec()
	// fmt.Println(result2)
	if nil != err {
		log.Printf("Insert Execute Error: %s\nError Query: %s : %s\n",
			err.Error(), object.String(), query2)
		return nil, err
	}

	return result, nil
}
