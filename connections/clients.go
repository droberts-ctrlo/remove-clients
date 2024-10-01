package connections

import (
	"database/sql"
	"fmt"
)

type Client struct {
	Id   int64
	Name string
}

func ClientsByName(db *sql.DB, name string) ([]Client, error) {
	var clients []Client

	rows, err := db.Query("SELECT ClientID, Name FROM `Clients` WHERE `Name`=?", name)
	if err != nil {
		return nil, fmt.Errorf("clientsByName %q: %v", name, err)
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			fmt.Printf("clientsByName.close: %v", err)
		}
	}()

	for rows.Next() {
		var client Client
		if err := rows.Scan(&client.Id, &client.Name); err != nil {
			return nil, fmt.Errorf("clientsByName %q: %v", name, err)
		}
		clients = append(clients, client)
	}

	return clients, nil
}
