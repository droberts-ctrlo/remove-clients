package connections

import (
	"database/sql"
	"fmt"
)

func PropsByClientID(db *sql.DB, clientID int64) ([]int64, error) {
	var properties []int64

	rows, err := db.Query("SELECT PropID FROM `Properties` WHERE `ClientID`=? LIMIT 10", clientID)
	if err != nil {
		return nil, fmt.Errorf("propsByClientID %q: %v", clientID, err)
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			fmt.Printf("propsByClientID.close: %v", err)
		}
	}()

	for rows.Next() {
		var property int64
		if err := rows.Scan(&property); err != nil {
			return nil, fmt.Errorf("propsByClientID %q: %v", clientID, err)
		}
		properties = append(properties, property)
	}

	return properties, nil
}
