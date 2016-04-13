package main

import "database/sql"

type Storage interface {
	ShowStore() ShowStore
}

type storage struct {
	shows *showStore
}

func (s *storage) ShowStore() ShowStore {
	return s.shows
}

func NewStorage(driver, src string) *storage {
	db, err := sql.Open(driver, src)
	if err != nil {
		panic(err)
	}

	if _, err = db.Exec(showSchema); err != nil {
		panic(err)
	}

	return &storage{
		&showStore{db},
	}
}

const showSchema = `CREATE TABLE IF NOT EXISTS shows (
    id integer PRIMARY KEY,
    tvdb_id integer,
    series_id text,
    series_name text
);
CREATE UNIQUE INDEX IF NOT EXISTS duplicateTvdbShow ON shows(tvdb_id);`

type showStore struct {
	*sql.DB
}

func (db *showStore) Update(s *Show) error {
	result, err := db.Exec("INSERT INTO shows (tvdb_id, series_id, series_name) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	s.ID = uint64(id)
	return err
}

func (db *showStore) Get(id uint64) (s Show, err error) {
	row := db.QueryRow("SELECT id, tvbdb_id, series_id, series_name FROM shows WHERE id=?", id)
	if err = row.Scan(&s.ID, &s.TvdbID, &s.SeriesID, &s.SeriesName); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	return s, nil
}

func (db *showStore) GetByName(name string) (s Show, err error) {
	row := db.QueryRow("SELECT id, tvdb_id, series_id, series_name FROM shows WHERE series_name=?", name)
	if err = row.Scan(&s.ID, &s.TvdbID, &s.SeriesID, &s.SeriesName); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	return s, nil
}

func (db *showStore) Remove(id uint64) error {
	_, err := db.Exec("DELETE FROM shows WHERE id=?", id)
	return err
}
