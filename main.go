package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/softjech/go-fiber/database"
	"github.com/softjech/go-fiber/lead"
)

// Create routes
func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/leads", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	// Using gorm to open a sqlite3 connection named as leads.db
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Database connection successfully opened")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")

}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	// start the server at port 3000
	app.Listen(3000)
	// start a connection to the database
	defer database.DBConn.Close()
}
