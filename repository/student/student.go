package studentRepo

import (
	"database/sql"
	"fmt"
	"librarysystem/structs"
)

func GetAllStudents(db *sql.DB) (result []structs.Student, err error) {
	sql := "SELECT * FROM students"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var student = structs.Student{}

		err = rows.Scan(&student.ID, &student.Name, &student.Major)
		if err != nil {
			fmt.Println("DB Connection Failed")
		}

		result = append(result, student)
	}
	return
}

func InsertStudent(db *sql.DB, student structs.Student) (err error) {
	sql := "INSERT INTO students (name, major) VALUES ($1, $2)"
	errs := db.QueryRow(sql, student.Name, student.Major)
	return errs.Err()
}

func UpdateStudent(db *sql.DB, student structs.Student) (err error) {
	sql := "UPDATE students SET name = $2, major = $3 WHERE id = $1"
	errs := db.QueryRow(sql, student.ID, student.Name, student.Major)
	return errs.Err()
}

func DeleteStudent(db *sql.DB, student structs.Student) (err error) {
	sql := "DELETE FROM students WHERE id = $1"
	errs := db.QueryRow(sql, student.ID)
	return errs.Err()
}
