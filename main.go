package main

import (
	"flag"
	"log/slog"
	"os"
	"rabbitmq-update-deffinitions/src"
)

func main() {
	flagRmqUrl := flag.String("host", "http://127.0.0.1:15672/api/definitions", "RMQ URL with endpoint /api/definitions")
	flagRmqUser := flag.String("user", "admin", "RMQ username")
	flagRmqPass := flag.String("password", "admin", "RMQ password")
	flagRmqPathToFile := flag.String("file", "", "RabbitMQ path to definitions file")
	flagRmqModeValidate := flag.Bool("validate", false, "RabbitMQ mode validate")
	flagRmqModeUpdate := flag.Bool("update", false, "RabbitMQ mode update")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	rmqUrl := *flagRmqUrl
	rmqUser := *flagRmqUser
	rmqPass := *flagRmqPass
	rmqPathToFile := *flagRmqPathToFile
	rmqModeValidate := *flagRmqModeValidate
	rmqModeUpdate := *flagRmqModeUpdate

	if rmqModeValidate {
		src.Validate(rmqPathToFile, logger)
	}
	if rmqModeUpdate {
		src.WorkWithApi(rmqPathToFile, rmqUrl, rmqPass, rmqUser, logger)
	}
}
