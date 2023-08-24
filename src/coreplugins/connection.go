package coreplugins

import (
	"context"
	"fmt"
	"gopher/src/model"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	// fmt.Printf("%v\n=======================================\n", sql)
	WebhookSqlLogSend(NewDiscord(), sql)
}

func InitDatabase() *gorm.DB {
	dsn := Dsn()
	dial := postgres.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{
		// Logger: &SqlLogger{},
		Logger: logger.Default.LogMode(logger.Silent),
		DryRun: false,
	})
	if err != nil {
		panic(err)
	}

	// migration
	db.AutoMigrate(&model.User{})

	fmt.Println("Database has been initialize")
	return db
}

func InitMongo() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoUri()))
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("MongoDB has been initialize")
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("my_go_dev").Collection(collectionName)
}
