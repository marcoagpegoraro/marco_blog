package initializers

import (
	"fmt"
	"os"

	"github.com/marcoagpegoraro/marco_blog/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error

	// dsn := os.Getenv("DB_URL")
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	dsn := os.Getenv("DSN")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		fmt.Println("Failed to connect to database")
	}
}

func SyncDB() {
	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.Tag{})
}
