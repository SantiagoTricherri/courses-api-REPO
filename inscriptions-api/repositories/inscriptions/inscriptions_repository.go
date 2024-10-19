package repositories

import (
	"fmt"
	"os"

	dao "inscriptions-api/DAOs/inscriptions"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Función auxiliar para obtener variables de entorno con valores predeterminados
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func Connect() (*gorm.DB, error) {
	dbHost := getEnv("DB_HOST", "mysql")
	dbPort := getEnv("DB_PORT", "3306")
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "rootpassword")
	dbName := getEnv("DB_NAME", "inscriptions")

	fmt.Printf("Connecting to MySQL: Host=%s, Port=%s, User=%s, DBName=%s\n", dbHost, dbPort, dbUser, dbName)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=30s",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to MySQL: %v", err)
	}

	err = db.AutoMigrate(&dao.InscriptionModel{})
	if err != nil {
		return nil, fmt.Errorf("error migrating database: %v", err)
	}

	return db, nil
}
