package DataBase

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgresql://neondb_owner:npg_BsP7YaURK4pE@ep-frosty-queen-a8ljml87-pooler.eastus2.azure.neon.tech/Unotes?sslmode=require"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
