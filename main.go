package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
	"todolist/app/middleware/auth"
	handlerUsers "todolist/app/presenter/users"
	"todolist/app/routes"
	serviceUsers "todolist/bussiness/users"
	mysqlDrivers "todolist/drivers/mysql"
	repositoryUsers "todolist/drivers/mysql/users"
)

func init() {
	viper.SetConfigName("test-config")
	viper.AddConfigPath("./app/config/")
	viper.AutomaticEnv()
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	configDB := mysqlDrivers.ConfigDB{
		DBUsername: viper.GetString(`database.user`),
		DBPassword: viper.GetString(`database.pass`),
		DBHost:     viper.GetString(`database.host`),
		DBPort:     viper.GetString(`database.port`),
		DBDatabase: viper.GetString(`database.name`),
	}

	configJWT := auth.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	//initial DB
	db := configDB.InitDB()
	//Migrate DB
	mysqlDrivers.MigrateDB(db)
	//Init Fiber Framework
	app := fiber.New()

	//factory of domain
	userRepository := repositoryUsers.NewRepositoryMySQL(db)
	userService := serviceUsers.NewService(userRepository, &configJWT)
	userHandler := handlerUsers.NewHandler(userService)

	routesInit := routes.HandlerList{
		UserHandler: *userHandler,
	}
	routesInit.Routes(app)
	log.Fatal(app.Listen(viper.GetString("server.address")))
}
