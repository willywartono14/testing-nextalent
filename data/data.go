package skeleton

import (
	"context"
	"database/sql"
	"testing-nextalent/model"
)

type (
	// Data ...
	Data struct {
		db   *sql.DB
		stmt map[string]*sql.Stmt
	}
)

// New ...
func New(db *sql.DB) Data {
	d := Data{
		db: db,
	}

	return d
}

func (d Data) GetCountry(ctx context.Context, person string) (string, error) {
	var (
		country string
		err     error
	)

	query := `SELECT country FROM person where name = $1`

	err = d.db.QueryRowContext(ctx, query, person).Scan(&country)

	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	return country, err
}

func (d Data) GetCountryAll(ctx context.Context) ([]model.Person, error) {
	var (
		results []model.Person
		err     error
	)

	query := `SELECT name,country FROM person`

	rows, err := d.db.QueryContext(ctx, query)

	if err != nil && err != sql.ErrNoRows {
		return results, err
	}

	defer func() {
		err = rows.Close()
	}()

	for rows.Next() {
		var result model.Person

		err := rows.Scan(
			&result.Name,
			&result.Country,
		)
		if err != nil {
			return results, err
		}

		results = append(results, result)
	}

	return results, err
}

func (d Data) ScriptInsertData(ctx context.Context, person model.Person) error {

	var id int

	query := `INSERT INTO person(
		name,
  		country
	)VALUES (
		$1,
		$2
	) returning id`
	err := d.db.QueryRowContext(
		ctx,
		query,
		person.Name,
		person.Country,
	).Scan(&id)

	if err != nil {
		return err
	}

	return err
}
