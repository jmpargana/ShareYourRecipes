package db

import (
	"database/sql"
	"fmt"
	"server/models"
)

func (w *DBWrapper) InsertTags(id int, tags []models.Tag) error {
	for _, tag := range tags {
		if err := w.InsertTag(id, tag); err != nil {
			return fmt.Errorf("failed inserting tag: %s with: %s", tag, err)
		}
	}
	return nil
}

func (w *DBWrapper) InsertTag(id int, tag models.Tag) error {
	tagID, err := w.FindTag(tag)

	if err == ErrNoID {
		if err := w.execQueryWithPrepare(insertTagQuery, tag); err != nil {
			return err
		}
		tagID, err = w.FindTag(tag)
	}

	return w.InsertRecipeTag(id, tagID)
}

/// Insert Tag if non existent with NULL id and retrieve ID for recipetags table.
func (w *DBWrapper) FindTag(tag models.Tag) (id int, err error) {
	return w.queryID(
		selectTagByName,
		string(tag),
	)
}

func (w *DBWrapper) FindTagsByRecipeID(r *models.Recipe) error {
	return w.queryRows(
		selectTagsByRecipeID,

		func(rows *sql.Rows) error {
			for rows.Next() {
				var tag string
				if err := rows.Scan(&tag); err != nil {
					return err
				}
				r.Tags = append(r.Tags, models.Tag(tag))
			}
			return nil
		},
		r.ID,
	)
}
