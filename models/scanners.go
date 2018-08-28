package models

type Scanner struct {
	Name string
	Fab string
	Location string
}

func GetScanners() ([]*Scanner, error){
	rows, err := db.Query("SELECT * FROM scanners ORDER BY name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ents := make([]*Scanner, 0)
	for rows.Next() {
		e := new(Scanner)
		err := rows.Scan(&e.Name, &e.Fab, &e.Location)
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