package connections

import (
	"database/sql"
	"fmt"
)

func GetJobsFromPropertyID(db *sql.DB, jobId int64) ([]int64, error) {
	var jobs []int64

	rows, err := db.Query("SELECT SubID FROM `Jobs` WHERE `PropID`=? AND NOT `SubID` IS NULL", jobId)
	if err != nil {
		return nil, fmt.Errorf("jobsByPropID %q: %v", jobId, err)
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			fmt.Printf("jobsByPropID.close: %v", err)
		}
	}()

	for rows.Next() {
		var subID int64

		if err = rows.Scan(&subID); err != nil {
			return nil, fmt.Errorf("jobsByPropID %q: %v", jobId, err)
		}

		jobs = append(jobs, subID)
	}

	return jobs, nil
}
