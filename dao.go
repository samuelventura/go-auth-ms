package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type daoDso struct {
	db *gorm.DB
}

type Dao interface {
	Close() error
	GetApps() (*[]AppDro, error)
	AddApp(name string) error
	GetApp(name string) (*AppDro, error)
	DelApp(name string) error
}

func dialector(args Args) (gorm.Dialector, error) {
	driver := args.Get("driver").(string)
	source := args.Get("source").(string)
	switch driver {
	case "sqlite":
		return sqlite.Open(source), nil
	case "postgres":
		return postgres.Open(source), nil
	}
	return nil, fmt.Errorf("unknown driver %s", driver)
}

func NewDao(args Args) (Dao, error) {
	mode := logger.Default.LogMode(logger.Silent)
	config := &gorm.Config{Logger: mode}
	dialector, err := dialector(args)
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(dialector, config)
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&AppDro{}, &MessageDro{}, &AttemptDro{})
	if err != nil {
		return nil, err
	}
	return &daoDso{db}, nil
}

func (dso *daoDso) Close() error {
	sqlDB, err := dso.db.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		return err
	}
	return nil
}

func (dso *daoDso) GetApps() (*[]AppDro, error) {
	var dros []AppDro
	result := dso.db.Find(&dros)
	return &dros, result.Error
}

func (dso *daoDso) AddApp(name string) error {
	dro := &AppDro{Name: name}
	result := dso.db.Create(dro)
	return result.Error
}

func (dso *daoDso) GetApp(name string) (*AppDro, error) {
	dro := &AppDro{}
	result := dso.db.Where("name = ?", name).First(dro)
	return dro, result.Error
}

func (dso *daoDso) DelApp(name string) error {
	dro := &AppDro{}
	result := dso.db.Where("name = ?", name).Delete(dro)
	if result.Error == nil && result.RowsAffected != 1 {
		return fmt.Errorf("row not found")
	}
	return result.Error
}
