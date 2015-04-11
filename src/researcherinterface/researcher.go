package main

type Name struct {
	First string
	Last  string
}

type Researcher struct {
	Id int64
	Name
	Email string
}

func NewResearcher(name Name, email, password string) (r *Researcher, err error) {
	stmt, err := DB.Prepare("INSERT INTO researchers (fname, lname, email, pword) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(name.First, name.Last, email, password).Scan(&id)
	if err != nil {
		return
	}

	r = &Researcher{
		Id:    id,
		Name:  name,
		Email: email,
	}

	return
}

func GetResearcherById(id int64) (r *Researcher, err error) {
	stmt, err := DB.Prepare("SELECT id, fname, lname, email FROM researchers WHERE id = $1")
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var fname string
	var lname string
	var email string

	err = row.Scan(&id)
	err = row.Scan(&fname)
	err = row.Scan(&lname)
	err = row.Scan(&email)

	name := Name{
		First: fname,
		Last:  lname,
	}

	r = &Researcher{
		Id:    id,
		Name:  name,
		Email: email,
	}

	return
}
