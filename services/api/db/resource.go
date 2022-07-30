package db // @todo docs

import (
	"time"

	"github.com/MatthewZito/gouache/models"
	"github.com/lib/pq"
)

func (db *DB) GetResources() ([]models.ResourceV2, error) {
	rows, err := db.Query(`
		SELECT
			id,
			title,
			created_at,
			COALESCE(updated_at, created_at),
			tags
		FROM resource_record
		ORDER BY created_at DESC
	`)

	if err != nil {
		return []models.ResourceV2{}, err
	}

	defer rows.Close()

	var r []models.ResourceV2

	for rows.Next() {
		resource := models.ResourceV2{}

		if err = rows.Scan(
			&resource.Id,
			&resource.Title,
			&resource.CreatedAt,
			&resource.UpdatedAt,
			&resource.Tags,
		); err != nil {
			return []models.ResourceV2{}, err
		}

		r = append(r, resource)
	}

	if err = rows.Err(); err != nil {
		return []models.ResourceV2{}, err
	}

	return r, nil
}

func (db *DB) GetResource(id string) (models.ResourceV2, error) {
	sql := `
		SELECT
			id,
			title,
			created_at,
			COALESCE(updated_at, created_at),
			tags
		FROM resource_record
		WHERE id=$1
		LIMIT 1;
	`

	row := db.QueryRow(sql, id)

	var r models.ResourceV2

	if err := row.Scan(
		&r.Id,
		&r.Title,
		&r.CreatedAt,
		&r.UpdatedAt,
		&r.Tags,
	); err != nil {
		return models.ResourceV2{}, err
	}

	return r, nil
}

// @todo return id
func (db *DB) CreateResource(t *models.NewResourceTemplate) (string, error) {
	sql := `
    	INSERT INTO resource_record
			(
				title,
				tags
			)
			VALUES ($1, $2)
			RETURNING id;
  	`

	var id string
	row := db.QueryRow(sql, t.Title,
		pq.StringArray(t.Tags))
	if err := row.Scan(&id); err != nil { // scan will release the connection
		return "", err
	}

	return id, nil
}

func (db *DB) UpdateResource(t *models.UpdateResourceTemplate) error {
	sql := `
    	UPDATE resource_record
		SET
			title = $1,
			tags = $2,
			updated_at = $3
		WHERE id = $4
  	`

	if _, err := db.Exec(sql,
		t.Title,
		pq.StringArray(t.Tags),
		time.Now().UTC(),
		t.Id,
	); err != nil {
		return err
	}

	return nil
}
