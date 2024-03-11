package repositories

import (
	"errors"
	"fmt"
	"github.com/djfemz/rave/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"reflect"
	"strings"
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
	_ = addEntities(models.Entities, db)
	return db
}

func addEntities(m map[string]any, db *gorm.DB) error {
	for _, v := range m {
		err := db.AutoMigrate(v)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func GetId(T any) (any, error) {
	obj := reflect.ValueOf(T)
	numberOfFields := obj.NumField()
	for index := 0; index < numberOfFields; index++ {
		tag := obj.Type().Field(index).Tag
		fmt.Println(tag)
		if strings.Contains(string(tag), "id") {
			idField := obj.Field(index)
			if idField.CanConvert(reflect.TypeOf(uint64(5))) {
				v := idField.Uint()
				return v, nil
			} else if idField.CanConvert(reflect.TypeOf("")) {
				return idField.String(), nil
			}
			break
		}
	}

	return nil, errors.New("could not retrieve id")
}
