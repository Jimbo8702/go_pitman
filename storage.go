package main

import (
	"database/sql"
)

type Storage interface {
	SaveData(string) error
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgressStore, error){
	psqlInfo := "host=localhost port=55000 user=postgres password=postgrespw dbname=postgres sslmode=disable"
    
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgressStore{
		db: db,
	}, nil
}
func (s *PostgressStore) Init() error {
	return s.createCrawledDataTable()
}

func (s *PostgressStore) createCrawledDataTable() error {
	query := `create table if not exists crawled_data (
	   id serial primary key,
	   data varchar(50),
	)`

	_, err := s.db.Exec(query)

	return err
}

// SaveData saves fetched data to the database.
func (s *PostgressStore) SaveData(data string) error {
	query :=  `"INSERT INTO crawled_data (data) VALUES ($1)"`

	_, err := s.db.Query(query, data)
	if err != nil {
		return err
	}

	return nil
}
