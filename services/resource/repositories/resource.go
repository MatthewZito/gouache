package repositories

import (
	"time"

	"github.com/exbotanical/gouache/models"
	"github.com/lib/pq"
)

// GetResources retrieves all resources from the database.
// @todo Paginate/batch.
func (db *DB) GetResources() ([]models.Resource, error) {
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
		return []models.Resource{}, err
	}

	defer rows.Close()

	var r []models.Resource

	for rows.Next() {
		resource := models.Resource{}

		if err = rows.Scan(
			&resource.Id,
			&resource.Title,
			&resource.CreatedAt,
			&resource.UpdatedAt,
			&resource.Tags,
		); err != nil {
			return []models.Resource{}, err
		}

		r = append(r, resource)
	}

	if err = rows.Err(); err != nil {
		return []models.Resource{}, err
	}

	return r, nil
}

// GetResource retrieves the resource that corresponds to a given ID.
func (db *DB) GetResource(id string) (models.Resource, error) {
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

	var r models.Resource

	if err := row.Scan(
		&r.Id,
		&r.Title,
		&r.CreatedAt,
		&r.UpdatedAt,
		&r.Tags,
	); err != nil {
		return models.Resource{}, err
	}

	return r, nil
}

// CreateResource creates a new resource in the database and returns its ID.
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

// UpdateResource updates a given resource.
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
