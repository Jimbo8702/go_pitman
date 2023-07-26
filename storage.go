package main

type Storage interface {
	StoreData(*Parseable) error
	DeleteData(*Parseable) error
	// GetData(*Parseable) error
	// GetDataByURL(*Parseable) error
}