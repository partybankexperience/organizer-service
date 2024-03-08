package repositories

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Repository[T, U any] interface {
	Save(t *T) *T
	FindById(id U) *T
	FindAll() []*T
	DeleteById(id U)
}

type RepositoryImpl[T, U any] struct {
}

var db = connect()

func (r *RepositoryImpl[T, U]) Save(t *T) *T {
	
	return nil
}

func (r *RepositoryImpl[T, U]) FindById(id U) *T {
	return nil
}

func (r *RepositoryImpl[T, U]) FindAll() []*T {
	return nil
}

func (r *RepositoryImpl[T, U]) DeleteById(id U) {

}

func connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=rave port=5432 sslmode=disable TimeZone=Africa/Lagos"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
