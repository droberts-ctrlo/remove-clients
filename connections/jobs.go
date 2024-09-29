package connections

import (
	"database/sql"
	"fmt"
)

type Property struct {
	PropID int64
}

func PropsByClientID(db *sql.DB, clientID int64) ([]Property, error) {
	var properties []Property

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
		var property Property
		if err := rows.Scan(&property.PropID); err != nil {
			return nil, fmt.Errorf("propsByClientID %q: %v", clientID, err)
		}
		properties = append(properties, property)
	}

	return properties, nil
}
