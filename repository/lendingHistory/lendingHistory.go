package lendingHistoryRepo

import (
	"database/sql"
	"librarysystem/structs"
)

func GetAllLendingHistory(db *sql.DB) (result []structs.History, err error) {
	sql := "SELECT  h.id, s.name, b.title, a.name, h.duration, h.created_at FROM lendinghistory h INNER JOIN students s ON s.id = h.student_id INNER JOIN book b ON b.id = h.book_id INNER JOIN admin a ON a.id = h.admin_id"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var history = structs.History{}

		err = rows.Scan(&history.ID, &history.Student, &history.Book, &history.Admin, &history.Duration, &history.Created_at)
		if err != nil {
			panic(err)
		}

		result = append(result, history)
	}
	return
}

func InsertLendingHistory(db *sql.DB, lendingHistory structs.LendingHistory) (err error) {
	sql := "INSERT INTO lendinghistory (student_id, admin_id, book_id, duration, created_at) VALUES ($1, $2, $3, $4, $5)"
	errs := db.QueryRow(sql, lendingHistory.Student_id, lendingHistory.Admin_id, lendingHistory.Book_id, lendingHistory.Duration, lendingHistory.Created_at)
	return errs.Err()
}

func DeleteLendingHistory(db *sql.DB, lendingHistory structs.LendingHistory) (err error) {
	sql := "DELETE FROM lendinghistory where id = $1"
	errs := db.QueryRow(sql, lendingHistory.ID)
	return errs.Err()
}
