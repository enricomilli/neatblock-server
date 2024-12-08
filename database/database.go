package db

type Database interface {
	NewClient()
	AddPool()
}
