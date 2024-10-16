package repositories

import (
	"fmt"
	"os"

	dao "inscriptions-api/DAOs/inscriptions"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	// Primero, conecta sin especificar una base de datos
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to MySQL: %v", err)
	}

	// Crea la base de datos si no existe
	dbName := os.Getenv("DB_NAME")
	err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)).Error
	if err != nil {
		return nil, fmt.Errorf("error creating database: %v", err)
	}

	// Conecta a la base de datos recién creada
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), dbName)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to created database: %v", err)
	}
	// Migración automática de los modelos
	err = db.AutoMigrate(&dao.InscriptionModel{})
	if err != nil {
		return nil, fmt.Errorf("error migrating database: %v", err)
	}

	return db, nil
}
