package main

type Core struct {
	Name    string
	Config  string
	Key     string
	Secret  string
	Version string
}

func SelectCore() (*Core, error) {
	var core Core
	rows, err := db.Query("SELECT * FROM core")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&core.Name, &core.Config, &core.Key, &core.Secret, &core.Version)
		if err != nil {
			return nil, err
		}
	}
	return &core, err
}