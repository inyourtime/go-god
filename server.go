package main

import (
	"context"
	"encoding/json"
	"gopher/src/coreplugins"
	"gopher/src/repository"
	"gopher/src/service"
	"log"
	"time"
)

func main() {
	coreplugins.InitConfig()

	db := coreplugins.InitDatabase()
	mg := coreplugins.InitMongo()

	mgDb := mg.Database("my_go_dev")
	_ = mgDb

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

	userRepo := repository.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepo)
	// hashPassword, err := coreplugins.HashPassword("1234")
	// if err != nil {
	// 	panic(err)
	// }
	// user := model.User{
	// 	Email:    "op@gmail.com",
	// 	Password: hashPassword,
	// 	Name:     "big",
	// 	Surname:  "test",
	// 	Nickname: "opopopo",
	// 	Age:      15,
	// 	Gender:   model.Female,
	// }
	// userAdded, err := userRepo.CreateOne(user)
	users, err := userService.GetUser(2)
	if err != nil {
		panic(err)
	}
	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}
	coreplugins.WebhookSend(coreplugins.NewDiscord(), string(jsonData))

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
