package connections

import (
	"database/sql"
	"fmt"
)

type JobData struct {
	ID    int64
	SubID int64
	SubID2 int64
}

func GetJobsFromPropertyID(db *sql.DB, jobId int64) ([]JobData, error) {
	var jobs []JobData

	rows, err := db.Query("SELECT JobID, SubID, SubID2 FROM `Jobs` WHERE `PropID`=? AND NOT `SubID` IS NULL", jobId)
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
		var job JobData

		if err = rows.Scan(&job.ID, &job.SubID, &job.SubID2); err != nil {
			return nil, fmt.Errorf("jobsByPropID %q: %v", jobId, err)
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}

func DeleteJobs(db *sql.DB, jobId int64) error {
	if _, err := db.Exec("DELETE FROM `Jobs` WHERE `JobID`=?", jobId); err != nil {
		return fmt.Errorf("jobsByPropID %q: %v", jobId, err)
	}
	return nil
}
