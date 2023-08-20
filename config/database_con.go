package config

//using direct MSQL
// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/joho/godotenv"
// )

// // InitDB initializes the database connection.
// func InitDB() (*sql.DB, error) {
//     // Load environment variables from .env file
//     err := godotenv.Load()
//     if err != nil {
//         log.Fatalf("Error loading .env file: %v", err)
//     }

//     // Read database configuration from environment variables
//     dbHost := os.Getenv("DB_HOST")
//     dbPort := os.Getenv("DB_PORT")
//     dbUser := os.Getenv("DB_USER")
//     dbPassword := os.Getenv("DB_PASSWORD")
//     dbName := os.Getenv("DB_NAME")

//     // Construct the MySQL connection string
//     connectionString := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

//     // Establish a database connection
//     db, err := sql.Open("mysql", connectionString)
//     if err != nil {
//         return nil, err
//     }

//     // Check for database connectivity by executing a simple query
//     err = db.Ping()
//     if err != nil {
//         return nil, err
//     }

//     // Print a success message if connected
//     fmt.Println("Connected to the database")

//     return db, nil
// }

//using connection GORM like moongose
import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func InitDB() (*gorm.DB, error) {
    
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
    
	dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    // Construct the MySQL connection string
    connectionString := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
        dbUser, dbPassword, dbHost, dbPort, dbName,
    )

    // Initialize the database connection
    db, err := gorm.Open("mysql", connectionString)
    if err != nil {
        return nil, err
    }

    // Check if the database connection is successful
    if err := db.DB().Ping(); err != nil {
        return nil, err
    }

    // Print a success message if connected
    fmt.Println("Connected to the database")

    return db, nil
}
