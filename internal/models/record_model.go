package models

import (
	"bitacora/internal/config"
	"bitacora/internal/core"
	"time"
)

func AddRecord(record core.Record) error {
	_, err := config.Database.Exec(`
		INSERT INTO records (name, lab, equipment, startDateTime, endDateTime,
			received, returned, comments, timestamp)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		record.Name,
		record.Lab,
		record.Equipment,
		record.StartDateTime,
		record.EndDateTime,
		record.Received,
		record.Returned,
		record.Comments,
		int(time.Now().Unix()),
	)
	return err
}

func GetRecordByID(id string) (core.Record, error) {
	var record core.Record

	row := config.Database.QueryRow(`
        SELECT id, name, lab, equipment, startDateTime, endDateTime,
               received, returned, comments, timestamp
        FROM records
        WHERE id = ?;
    `, id)

	err := row.Scan(
		&record.ID,
		&record.Name,
		&record.Lab,
		&record.Equipment,
		&record.StartDateTime,
		&record.EndDateTime,
		&record.Received,
		&record.Returned,
		&record.Comments,
		&record.Timestamp,
	)

	if err != nil {
		return core.Record{}, err
	}

	return record, nil
}

func GetRecordByMachine(machine string) ([]core.Record, error) {
	rows, err := config.Database.Query(`
		SELECT  id, name, lab, equipment, startDateTime, endDateTime,
			received, returned, comments, timestamp
		FROM records
		WHERE equipment = ?
		ORDER BY startDateTime DESC
		LIMIT 10
	`, machine)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var records []core.Record
	for rows.Next() {
		var r core.Record
		err := rows.Scan(
			&r.ID,
			&r.Name,
			&r.Lab,
			&r.Equipment,
			&r.StartDateTime,
			&r.EndDateTime,
			&r.Received,
			&r.Returned,
			&r.Comments,
			&r.Timestamp,
		)
		if err != nil {
			return nil, err
		}
		records = append(records, r)
	}
	return records, nil
}
