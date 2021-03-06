package models

type Entry struct {
	ID string
	Timestamp string
	Location string
	Equipment string
	Date string
	Tech string
	Building string
}

func LastEntries() ([]*Entry, error){
	rows, err := db.Query("SELECT ID, entity, timestamp, date, tech, fab, location FROM (SELECT *, ROW_NUMBER() OVER (PARTITION BY entity ORDER BY DATE DESC) AS rn FROM audit2 a JOIN scanners s ON a.entity = s.name) t WHERE rn = 1 ORDER BY date")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ents := make([]*Entry, 0)
	for rows.Next() {
		e := new(Entry)
		err := rows.Scan(&e.ID, &e.Equipment, &e.Timestamp, &e.Date, &e.Tech, &e.Building, &e.Location)
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

func AllEntries() ([]*Entry, error){
	rows, err := db.Query("SELECT timestamp, entity, date, tech FROM audit2 ORDER BY entity, Date DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ents := make([]*Entry, 0)
	for rows.Next() {
		e := new(Entry)
		err := rows.Scan(&e.Timestamp, &e.Equipment, &e.Date, &e.Tech)
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

func LimitedEntries(date string) ([]*Entry, error){
	rows, err := db.Query("SELECT timestamp, entity, date, tech FROM audit2 WHERE CONVERT(datetime, date) >= CONVERT(datetime, $1) ORDER BY entity, date DESC", date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ents := make([]*Entry, 0)
	for rows.Next() {
		e := new(Entry)
		err := rows.Scan(&e.Timestamp, &e.Equipment, &e.Date, &e.Tech)
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