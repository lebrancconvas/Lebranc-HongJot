package model

import (
	"log"

	"github.com/lebrancconvas/Lebranc-HongJot/db"
	"github.com/lebrancconvas/Lebranc-HongJot/form"
)

func GetExpenses() ([]form.Transaction, error) {
	db := db.GetDB()

	var expenses []form.Transaction

	stmt, err := db.Prepare(`
		SELECT id, date, amount, transaction_type, note, image_url, spender_id
		FROM transactions
		WHERE transaction_type = 2 AND used_flg = true
		ORDER BY id, date 
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return expenses, err
	}
	defer rows.Close()

	for rows.Next() {
		var expense form.Transaction
		err := rows.Scan(
			&expense.ID,
			&expense.Date,
			&expense.Amount,
			&expense.TransactionType,
			&expense.Note,
			&expense.ImageURL,
			&expense.SpenderID,
		)
		if err != nil {
			return expenses, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}
