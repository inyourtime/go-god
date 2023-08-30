package coreplugins

import (
	"fmt"
	"gopher/src/configs"
	"log"

	"github.com/spf13/viper"
)

var Config *configs.Env

func Env() (*configs.Env, error) {
	configs := configs.Env{}
	err := viper.Unmarshal(&configs)
	if err != nil {
		log.Printf("Env struct error: %v", err)
		return nil, err
	}
	return &configs, nil
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// save configs variable
	Config, _ = Env()
}

func Dsn() string {
	if Config == nil {
		Config, _ = Env()
	}

	dsn := fmt.Sprintf("host=%v user=%v dbname=%v port=%v password=%v search_path=%v sslmode=%v",
		Config.Database.Postgres.Host,
		Config.Database.Postgres.User,
		Config.Database.Postgres.Dbname,
		Config.Database.Postgres.Port,
		Config.Database.Postgres.Password,
		Config.Database.Postgres.SearchPath,
		Config.Database.Postgres.Sslmode,
	)
	return dsn
}

func MongoUri() string {
	config, _ := Env()
	return config.Database.Mongo.Uri
}
