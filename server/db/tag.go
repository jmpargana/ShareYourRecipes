package db

import (
	"context"
	"fmt"
	"server/models"
	"time"
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
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		stmt, err := w.db.PrepareContext(ctx, insertTagQuery)
		if err != nil {
			return fmt.Errorf("failed preparing query: %s with: %s", insertTagQuery, err)
		}
		defer stmt.Close()

		if _, err := stmt.ExecContext(ctx, tag); err != nil {
			return fmt.Errorf("failed inserting tag with: %s", err)
		}

		tagID, err = w.FindTag(tag)
	}
	if err != nil {
		return fmt.Errorf("nothing returned from select tag query: %s", err)
	}

	return w.InsertRecipeTag(id, tagID)
}

/// Insert Tag if non existent with NULL id and retrieve ID for recipetags table.
func (w *DBWrapper) FindTag(tag models.Tag) (id int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = w.db.QueryRowContext(ctx, selectTagByName, tag).
		Scan(&id)

	if err != nil {
		return 0, ErrNoID
	}

	return
}
