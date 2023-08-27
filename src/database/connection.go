package database

import (
	"context"
	"fmt"
	"gopher/src/coreplugins"
	"gopher/src/model"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var SqlDB *gorm.DB
var NoSqlDB *mongo.Client

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	go coreplugins.WebhookSqlLogSend(sql)
}

func ConnectSqlDatabase() {
	var err error
	dsn := coreplugins.Dsn()
	SqlDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: &SqlLogger{},
		// Logger: logger.Default.LogMode(logger.Silent),
		DryRun: false,
	})
	if err != nil {
		panic(err)
	}

	SqlDB.AutoMigrate(&model.User{})
	fmt.Println("Postgres Database has been initialize")
}

func ConnectNoSqlDatabase() {
	var err error
	NoSqlDB, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(coreplugins.MongoUri()))
	if err != nil {
		panic(err)
	}

	err = NoSqlDB.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("MongoDB has been initialize")
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("my_go_dev").Collection(collectionName)
}
