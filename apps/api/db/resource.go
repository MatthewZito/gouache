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
		FROM resource
		ORDER BY created_at ASC
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
		FROM resource
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
func (db *DB) CreateResource(t *models.NewResourceTemplate) error {
	sql := `
    	INSERT INTO resource
			(
				title,
				tags
			)
			VALUES ($1, $2);
  	`

	if _, err := db.Exec(sql,
		t.Title,
		pq.StringArray(t.Tags),
	); err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateResource(t *models.UpdateResourceTemplate) error {
	sql := `
    	UPDATE resource
		SET
			(
				title = $1,
				tags = $2,
				updated_at = $3
			)
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
