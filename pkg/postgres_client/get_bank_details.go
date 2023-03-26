package postgres_client

import (
	"database/sql"
	"fmt"
	"pricing_details/pkg/model"
)

func GetBankDetails(postgresConn *sql.DB, userID int) (details model.CardDetails, err error) {

	row, err := postgresConn.Query("SELECT user_id, card_number, expiry_date, ccv, name FROM card_details where user_id = $1", userID)
	if err != nil {
		fmt.Println("error querying table ", err)
		return details, err
	}

	for row.Next() {
		err = row.Scan(&details.UserID, &details.CardNumber, &details.ExpiryDate, &details.CVV, &details.Name)
		if err != nil {
			fmt.Println("error scanning ", err)
			return details, err
		}
	}

	fmt.Println("row ", details)

	return details, nil
}
