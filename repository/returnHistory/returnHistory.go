package returnHistoryRepo

import (
	"database/sql"
	"fmt"
	"librarysystem/structs"
)

func GetAllReturnHistory(db *sql.DB) (result []structs.ReturnHist, err error) {
	sql := "SELECT h.id, s.name, b.title, a.name FROM returnhistory h INNER JOIN students s ON s.id = h.student_id INNER JOIN admin a ON a.id = h.admin_id INNER JOIN book b ON b.id = h.book_id"

	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var returnhistory = structs.ReturnHist{}

		err = rows.Scan(&returnhistory.ID, &returnhistory.Student, &returnhistory.Book, &returnhistory.Admin)
		if err != nil {
			panic(err)
		}

		result = append(result, returnhistory)
	}
	return
}

func InsertReturnHistory(db *sql.DB, returnHist structs.ReturnHistory) (err error) {
	sql := "INSERT INTO returnhistory (student_id, admin_id, book_id, created_at) VALUES ($1, $2, $3, $4)"
	errs := db.QueryRow(sql, returnHist.Student_id, returnHist.Admin_id, returnHist.Book_id, returnHist.Created_at)
	return errs.Err()
}

func DeleteLendingHistory(db *sql.DB, lendingHistory structs.LendingHistory) (err error) {
	sql := "DELETE FROM lendinghistory where id = $1"
	errs := db.QueryRow(sql, lendingHistory.ID)
	return errs.Err()
}
