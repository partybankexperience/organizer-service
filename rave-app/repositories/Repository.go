package repositories

import (
	"errors"
	"fmt"
	"github.com/djfemz/rave/rave-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Repository[T, U any] interface {
	crudRepository[T, U]
}

type crudRepository[T, U any] interface {
	Save(t *T) (*T, error)
	FindById(id U) (*T, error)
	FindAll() ([]*T, error)
	FindAllBy(pageable Pageable) ([]*T, error)
	DeleteById(id U) error
}

type repositoryImpl[T, U any] struct {
}

var db *gorm.DB

func (r *repositoryImpl[T, U]) Save(t *T) (*T, error) {
	db = connect()
	err := db.Save(t).Error
	if err != nil {
		return nil, err
	}
	var id, _ = GetId(*t)
	err = db.First(t, id).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *repositoryImpl[T, U]) FindById(id U) (*T, error) {
	db = connect()
	var t = new(T)
	err := db.First(t, id).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *repositoryImpl[T, U]) FindAll() ([]*T, error) {
	db = connect()
	var organizations []*T
	err := db.Find(&organizations).Error
	if err != nil {
		return nil, err
	}
	return organizations, nil
}

func (r *repositoryImpl[T, U]) FindAllBy(pageable Pageable) ([]*T, error) {
	db = connect()
	page := getPage[T](db, pageable)
	if page == nil {
		return nil, errors.New("failed to find page")
	}
	return page.GetElements(), nil
}

func (r *repositoryImpl[T, U]) DeleteById(id U) error {
	db = connect()
	err := db.Delete(new(T), id).Error

	if err != nil {
		return err
	}
	return nil
}

func connect() *gorm.DB {
	port, err := strconv.ParseUint(os.Getenv("DATABASE_PORT"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Africa/Lagos", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"), port)
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
	for index := ZERO; index < numberOfFields; index++ {
		tag := obj.Type().Field(index).Tag
		isPrimaryKeyField := strings.Contains(string(tag), "id")
		if isPrimaryKeyField {
			id, err, done := getPrimaryKey(obj, index)
			if done {
				return id, err
			}
		}
	}
	return nil, errors.New("could not retrieve id")
}

func getPrimaryKey(obj reflect.Value, index int) (any, error, bool) {
	idField := obj.Field(index)
	if idField.CanConvert(reflect.TypeOf(uint64(5))) {
		v := idField.Uint()
		return v, nil, true
	} else if idField.CanConvert(reflect.TypeOf(5)) {
		v := idField.Int()
		return v, nil, true
	} else if idField.CanConvert(reflect.TypeOf("")) {
		return idField.String(), nil, true
	}
	return nil, nil, false
}
