package src

import (
	"encoding/json"
	"fmt"
	"log"
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

func ValidateBindings( logger *log.Logger) {
	queueNames := MakeNameMap(cfg.Queues, func(queue Queue) string { return queue.Name })
	exchangeNames := MakeNameMap(cfg.Exchanges, func(exchange Exchange) string { return exchange.Name })
	vhostNames := MakeNameMap(cfg.Vhosts, func(vhost Vhost) string { return vhost.Name })
	for _, binding := range cfg.Bindings {
		if binding.DestinationType == "queue" {
			if !queueNames[binding.Destination] {
				logger.Error("❌Binding refers to non-existent queue: " + binding.Destination)
			}
		}
		if binding.DestinationType == "exchange" {
			if !exchangeNames[binding.Destination] {
				logger.Error("❌Binding refers to non-existent exchange: " + binding.Destination)
			}
		}
		if !vhostNames[binding.Vhost] {
			logger.Error("❌Binding refers to non-existent vhost: " + binding.Vhost)
		}
	}
}

func ValidateExchanges(logger log.Logger) {
	vhostNames := MakeNameMap(cfg.Vhosts, func(vhost Vhost) string { return vhost.Name })

	for _, exchange := range cfg.Exchanges {
		if !vhostNames[exchange.Vhost] {
			logger.Error("❌Exchanges refers to non-existent vhost: " + exchange.Vhost)
		}
	}
}

func ValidatePermisions() {
	userNames := MakeNameMap(cfg.Users, func(user User) string { return user.Name })
	vhostNames := MakeNameMap(cfg.Vhosts, func(vhost Vhost) string { return vhost.Name })
	for _, permission := range cfg.Permissions {
		if !userNames[permission.User] {
			logger.Error("❌Permissions refers to non-existent user: " + permission.User)
		}
		if !vhostNames[permission.Vhost] {
			logger.Error("❌Permissions refers to non-existent vhost: " + permission.Vhost)
		}
	}
}

func ValidatePolicies() {
	vhostNames := MakeNameMap(cfg.Vhosts, func(vhost Vhost) string { return vhost.Name })

	for _, policy := range cfg.Policies {
		if !vhostNames[policy.Vhost] {
			logger.Error("❌Policies refers to non-existent vhost: " + policy.Vhost)
		}
	}
}

func ValidateQueues() {
	vhostNames := MakeNameMap(cfg.Vhosts, func(vhost Vhost) string { return vhost.Name })

	for _, queue := range cfg.Queues {
		if !vhostNames[queue.Vhost] {
			logger.Error("❌Queue refers to non-existent vhost: " + queue.Vhost)
		}
	}
}

func Validate(pathToFile string, logger *slog.Logger) {

	schema, err := jsonschema.Compile("schema.json")
	if err != nil {
		logger.Error("❌failed to compile schema: %v", err)
		panic(err)
	}

	data, err := os.ReadFile(pathToFile)
	if err != nil {
		logger.Error("❌failed to read config: %v", err)
	}

	var jsonMap interface{}
	if err := json.Unmarshal(data, &jsonMap); err != nil {
		logger.Error("❌invalid JSON: %v", err)
	}

	if err := schema.Validate(jsonMap); err != nil {
		logger.Error("❌invalid config: %v", err)
		os.Exit(1)
	}

	//queueNames := map[string]bool{}
	//vhostNames := map[string]bool{}
	//userNames := map[string]bool{}
	//exchangeNames := map[string]bool{}

	if err := json.Unmarshal(data, &cfg); err != nil {
		panic(fmt.Errorf("failed to parse config: %v", err))
	}

	//for _, queue := range cfg.Queues {
	//	queueNames[queue.Name] = true
	//}
	//
	//for _, vhost := range cfg.Vhosts {
	//	vhostNames[vhost.Name] = true
	//}
	//
	//for _, user := range cfg.Users {
	//	userNames[user.Name] = true
	//}
	//
	//for _, exchange := range cfg.Exchanges {
	//	exchangeNames[exchange.Name] = true
	//}

	//for _, binding := range cfg.Bindings {
	//	if binding.DestinationType == "queue" {
	//		if !queueNames[binding.Destination] {
	//			logger.Error("❌Binding refers to non-existent queue: " + binding.Destination)
	//		}
	//	}
	//	if binding.DestinationType == "exchange" {
	//		if !exchangeNames[binding.Destination] {
	//			logger.Error("❌Binding refers to non-existent exchange: " + binding.Destination)
	//		}
	//	}
	//	if !vhostNames[binding.Vhost] {
	//		logger.Error("❌Binding refers to non-existent vhost: " + binding.Vhost)
	//	}
	//}

	//for _, exchange := range cfg.Exchanges {
	//	if !vhostNames[exchange.Vhost] {
	//		logger.Error("❌Exchanges refers to non-existent vhost: " + exchange.Vhost)
	//	}
	//}

	//for _, permission := range cfg.Permissions {
	//	if !userNames[permission.User] {
	//		logger.Error("❌Permissions refers to non-existent user: " + permission.User)
	//	}
	//	if !vhostNames[permission.Vhost] {
	//		logger.Error("❌Permissions refers to non-existent vhost: " + permission.Vhost)
	//	}
	//}

	//for _, policy := range cfg.Policies {
	//	if !vhostNames[policy.Vhost] {
	//		logger.Error("❌Policies refers to non-existent vhost: " + policy.Vhost)
	//	}
	//}

	//for _, queue := range cfg.Queues {
	//	if !vhostNames[queue.Vhost] {
	//		logger.Error("❌Queue refers to non-existent vhost: " + queue.Vhost)
	//	}
	//}

	logger.Info("✅Config JSON is valid")
}
