package src

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

var cfg Config

func MakeNameMap[T any](items []T, getName func(T) string) map[string]bool {
	mapNames := make(map[string]bool, len(items))
	for _, item := range items {
		mapNames[getName(item)] = true
	}
	return mapNames
}

func ValidateBindings(queueNames, exchangeNames, vhostNames map[string]bool, logger *slog.Logger) {
	for _, binding := range cfg.Bindings {
		switch binding.DestinationType {
		case "queue":
			if !queueNames[binding.Destination] {
				logger.Error("❌ Binding refers to non-existent queue", "destination", binding.Destination)
				os.Exit(1)
			}
		case "exchange":
			if !exchangeNames[binding.Destination] {
				logger.Error("❌ Binding refers to non-existent exchange", "destination", binding.Destination)
				os.Exit(1)
			}
		default:
			logger.Error("❌ Unknown binding.destination_type", "type", binding.DestinationType)
			os.Exit(1)
		}

		if !vhostNames[binding.Vhost] {
			logger.Error("❌ Binding refers to non-existent vhost", "vhost", binding.Vhost)
			os.Exit(1)
		}
	}
	logger.Info("✅ Bindings validated successfully")
}

func ValidateExchanges(vhostNames map[string]bool, logger *slog.Logger) {

	for _, exchange := range cfg.Exchanges {
		if !vhostNames[exchange.Vhost] {
			logger.Error("❌ Exchanges refers to non-existent vhost: " + exchange.Vhost)
			os.Exit(1)
		}
	}
	logger.Info("✅ Exchanges validated successfully")
}

func ValidatePermisions(userNames map[string]bool, vhostNames map[string]bool, logger *slog.Logger) {
	for _, permission := range cfg.Permissions {
		if !userNames[permission.User] {
			logger.Error("❌ Permissions refers to non-existent user: " + permission.User)
			os.Exit(1)
		}
		if !vhostNames[permission.Vhost] {
			logger.Error("❌ Permissions refers to non-existent vhost: " + permission.Vhost)
			os.Exit(1)
		}
	}
	logger.Info("✅ Permisions validated successfully")
}

func ValidatePolicies(vhostNames map[string]bool, logger *slog.Logger) {
	for _, policy := range cfg.Policies {
		if !vhostNames[policy.Vhost] {
			logger.Error("❌ Policies refers to non-existent vhost: " + policy.Vhost)
			os.Exit(1)
		}
	}
	logger.Info("✅ Policies validated successfully")
}

func ValidateQueues(vhostNames map[string]bool, logger *slog.Logger) {
	for _, queue := range cfg.Queues {
		if !vhostNames[queue.Vhost] {
			logger.Error("❌ Queue refers to non-existent vhost: " + queue.Vhost)
			os.Exit(1)
		}
	}
	logger.Info("✅ Queues validated successfully")
}

func ValidateSchema(pathToFile string, logger *slog.Logger) {
	schema, err := jsonschema.Compile("schema.json")
	if err != nil {
		logger.Error("❌ Failed to compile schema: ", "ERROR", err)
		panic(err)
	}

	data, err := os.ReadFile(pathToFile)
	if err != nil {
		logger.Error("❌ Failed to read config:", "ERROR", err)
		panic(err)
	}

	var jsonMap interface{}
	if err := json.Unmarshal(data, &jsonMap); err != nil {
		logger.Error("❌ Invalid JSON: ", "ERROR", err)
		panic(err)
	}

	if err := schema.Validate(jsonMap); err != nil {
		logger.Error("❌ Invalid config: ", "ERROR", err)
		panic(err)
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		logger.Error("❌ Failed to parse config: ", "ERROR", err)
		panic(err)
	}
}

func Validate(pathToFile string, logger *slog.Logger) {

	ValidateSchema(pathToFile, logger)

	vhostNames := MakeNameMap(cfg.Vhosts, func(vhost Vhost) string { return vhost.Name })
	userNames := MakeNameMap(cfg.Users, func(user User) string { return user.Name })
	queueNames := MakeNameMap(cfg.Queues, func(queue Queue) string { return queue.Name })
	exchangeNames := MakeNameMap(cfg.Exchanges, func(exchange Exchange) string { return exchange.Name })

	ValidateQueues(vhostNames, logger)
	ValidatePolicies(vhostNames, logger)
	ValidateBindings(queueNames, exchangeNames, vhostNames, logger)
	ValidateExchanges(vhostNames, logger)
	ValidatePermisions(userNames, vhostNames, logger)

	logger.Info("✅ Config JSON is valid")
}
