package connections

import (
	"database/sql"
	"fmt"
)

type PropertyData struct {
	PropID int64
}

func PropsByClientID(db *sql.DB, clientID int64) ([]PropertyData, error) {
	var properties []PropertyData

	rows, err := db.Query("SELECT PropID FROM `Properties` WHERE `ClientID`=?", clientID)
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
		var property PropertyData
		if err := rows.Scan(&property.PropID); err != nil {
			return nil, fmt.Errorf("propsByClientID %q: %v", clientID, err)
		}
		properties = append(properties, property)
	}

	return properties, nil
}

func DeleteProperty(db *sql.DB, propID int64) error {
	if _, err := db.Exec("DELETE FROM `Properties` WHERE `PropID`=?", propID); err != nil {
		return fmt.Errorf("DeleteProperty %q: %v", propID, err)
	}
	return nil
}
