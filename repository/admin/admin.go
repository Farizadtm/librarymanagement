package adminRepo

import (
	"database/sql"
	"fmt"
	"librarysystem/structs"
)

func GetAllAdmin(db *sql.DB) (result []structs.Admin, err error) {
	sql := "SELECT id, name, birth_date FROM admin"

	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var admin = structs.Admin{}

		err = rows.Scan(&admin.ID, &admin.Name, &admin.Birth_date)
		if err != nil {
			panic(err)
		}

		result = append(result, admin)
	}
	return
}

func InsertAdmin(db *sql.DB, admin structs.Admin) (err error) {
	sql := "INSERT INTO admin (name, username, pass, birth_date) VALUES ($1, $2, $3, $4)"
	errs := db.QueryRow(sql, admin.Name, admin.Username, admin.Password, admin.Birth_date)
	return errs.Err()
}

func UpdateAdmin(db *sql.DB, admin structs.Admin) (err error) {
	sql := "UPDATE admin SET name = $2, birth_date = $3 WHERE id = $1"
	errs := db.QueryRow(sql, admin.ID, admin.Name, admin.Birth_date)
	return errs.Err()
}

func DeleteAdmin(db *sql.DB, admin structs.Admin) (err error) {
	sql := "DELETE FROM admin where id = $1"
	errs := db.QueryRow(sql, admin.ID)
	return errs.Err()
}

func GetAdminByID(db *sql.DB) (result structs.Admin, err error) {
	sql := "SELECT id, name, birth_date FROM admin WHERE id = $1 LIMIT 1"
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var admin = structs.Admin{}

		err = rows.Scan(&admin.ID, &admin.Name, &admin.Birth_date)
		if err != nil {
			panic(err)
		}

		result = admin
	}
	return
}

func GetAdminByUserPass(db *sql.DB, user structs.User) (result structs.Admin, err error) {
	sql := "SELECT username, pass FROM admin WHERE username = $1 AND pass = $2"

	rows, err := db.Query(sql, user.Username, user.Password)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var admin = structs.Admin{}

		err = rows.Scan(&admin.Username, &admin.Password)
		if err != nil {
			panic(err)
		}

		result = admin
	}
	return
}
