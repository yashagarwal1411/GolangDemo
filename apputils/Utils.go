package apputils

import (
	"fmt"
	"database/sql"
	"yash_agarwal/models"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func UpdateUserStruct(row *sql.Row, user *models.UserInfo) {
	var tempAddress, tempTelephone, tempFullname sql.NullString
	err := row.Scan(&user.Id, &user.Email, &user.Password, &tempAddress, &tempTelephone, &tempFullname)
	fmt.Println("EMAIL:" + user.Email + " Id: " + " " + " address: " + tempAddress.String, user.Id)
	if err == sql.ErrNoRows {
		user.Id = 0
		return
	}
	CheckErr(err)
	if tempAddress.Valid {
		user.Address = tempAddress.String
	}
	if tempTelephone.Valid {
		user.Telephone = tempTelephone.String
	}
	if tempFullname.Valid {
		user.FullName = tempFullname.String
	}
}
