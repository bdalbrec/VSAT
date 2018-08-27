package models

type Entry struct {
	Timestamp string
	Location string
	Equipment string
	Date string
	Tech string
}

func AllEntries() ([]*Entry, error){
	rows, err := db.Query("SELECT * FROM audits")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ents := make([]*Entry, 0)
	for rows.Next() {
		e := new(Entry)
		err := rows.Scan(&e.Timestamp, &e.Location, &e.Equipment, &e.Date, &e.Tech)
		if err != nil {
			return nil, err
		}
		ents = append(ents, e)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ents, nil
}
