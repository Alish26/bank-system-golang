package main

import (
	"os"
	"time"

	"github.com/Alish26/bank-system-golang/infrastructure"
	"github.com/Alish26/bank-system-golang/infrastructure/database"
	"github.com/Alish26/bank-system-golang/infrastructure/log"
	"github.com/Alish26/bank-system-golang/infrastructure/router"
	"github.com/Alish26/bank-system-golang/infrastructure/validation"
)

func main() {
	var app = infrastructure.NewConfig().
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		Logger(log.InstanceLogrusLogger).
		Validator(validation.InstanceGoPlayground).
		DbSQL(database.InstancePostgres).
		DbNoSQL(database.InstanceMongoDB)

	app.WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceGorillaMux).
		Start()
}
