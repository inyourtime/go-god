package main

import (
	"context"
	"fmt"
	"gopher/src/coreplugins"
	"gopher/src/database"
	"gopher/src/middlewere"
	"gopher/src/router"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func main() {
	coreplugins.InitConfig()
	database.ConnectSqlDatabase()
	// database.ConnectNoSqlDatabase()

	cleanup := databaseSetting(database.SqlDB, database.NoSqlDB)
	defer cleanup()

	coreplugins.NewDiscord()

	app := fiber.New()

	app.Use(middlewere.Recover())
	app.Use(logger.New())

	// set context for use in service logger
	app.Use(func(c *fiber.Ctx) error {
		coreplugins.SetContext(c)
		return c.Next()
	})

	router.SetupRoutes(app)

	fmt.Println("Server is running on Port: " + coreplugins.Config.ServerPort)
	// logs.Error("test")
	app.Listen(":" + coreplugins.Config.ServerPort)
}

func databaseSetting(sql *gorm.DB, nosql *mongo.Client) func() {
	pg, err := sql.DB()
	if err != nil {
		panic(err)
	}

	pg.SetMaxIdleConns(10)
	pg.SetMaxOpenConns(100)
	pg.SetConnMaxIdleTime(time.Hour)

	return func() {
		if err = pg.Close(); err != nil {
			panic(err)
		}
		if err = database.NoSqlDB.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}
}
