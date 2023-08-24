package main

import (
	"context"
	"gopher/src/coreplugins"
	"gopher/src/handler"
	"gopher/src/middlewere"
	"gopher/src/repository"
	"gopher/src/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	coreplugins.InitConfig()
	// config, _ := coreplugins.Env()
	db := coreplugins.InitDatabase()
	mg := coreplugins.InitMongo()

	pg, err := db.DB()
	if err != nil {
		panic(err)
	}

	pg.SetMaxIdleConns(10)
	pg.SetMaxOpenConns(100)
	pg.SetConnMaxIdleTime(time.Hour)

	defer func() {
		if err = pg.Close(); err != nil {
			panic(err)
		}
	}()
	defer func() {
		if err = mg.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	app := fiber.New()
	// app.Use(recover.New())
	// app.Use("/", func(c *fiber.Ctx) error {
	// 	err := c.Next()
	// 	defer coreplugins.WebhookSend(coreplugins.NewDiscord(), "hello")
	// 	return err
	// })
	app.Use(middlewere.Recover)
	app.Use(logger.New())
	api := app.Group("/api")

	userRepo := repository.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepo)
	// _ = userService
	authHandler := handler.NewAuthHandler(userService)

	api.Post("/login", authHandler.Login)

	app.Listen(":8000")
	// hashPassword, err := coreplugins.HashPassword("1234")
	// if err != nil {
	// 	panic(err)
	// }
	// nk := "bnmbnm"
	// age := 12
	// user := model.NewUserRequest{
	// 	Email:    "beam@gmail.com",
	// 	Password: "5678",
	// 	Name:     "beam",
	// 	Surname:  "test",
	// 	Nickname: &nk,
	// 	Age:      &age,
	// 	Gender:   model.Unspecified,
	// }

	// res, err := userService.NewUser(user)
	// users, err := userService.GetUsers()
	// res, err := userService.Login(model.LoginRequest{Email: "test@gmail.com", Password: "5678"})
	// if err != nil {
	// 	log.Println(err)
	// }
	// jsonData, err := json.Marshal(res)
	// if err != nil {
	// 	log.Fatalf("Error encoding JSON: %v", err)
	// }
	// coreplugins.WebhookSend(coreplugins.NewDiscord(), string(jsonData))

}

type Tea struct {
	Type   string
	Rating int32
	Vendor []string `bson:"vendor,omitempty" json:"vendor,omitempty"`
}

// func GetGenders() {
// 	genders := []Gender{}
// 	tx := db.Find(&genders)
// 	if tx.Error != nil {
// 		fmt.Println(tx.Error)
// 		return
// 	}
// 	fmt.Println(genders)
// 	jsonData, err := json.Marshal(genders)
// 	if err != nil {
// 		log.Fatalf("Error encoding JSON: %v", err)
// 	}
// 	dc := coreplugins.NewDiscord()
// 	coreplugins.WebhookSend(dc, string(jsonData))
// }
