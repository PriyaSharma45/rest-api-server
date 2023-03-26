package postgres_client

import (
	"database/sql"
	"fmt"
	"rest-api-server/pkg/model"
)

func InsertBankDetails(postgresConn *sql.DB, cardDetails model.CardDetails) error {
	row := postgresConn.QueryRow("INSERT INTO card_details(card_number, ccv, expiry_date, name, user_id) VALUES($1, $2, $3, $4, $5);", cardDetails.CardNumber, cardDetails.CVV, cardDetails.ExpiryDate, cardDetails.Name, cardDetails.UserID)

	if row.Err() != nil {
		fmt.Println("error inserting card details ", row.Err())
		return row.Err()
	}

	return nil
}
