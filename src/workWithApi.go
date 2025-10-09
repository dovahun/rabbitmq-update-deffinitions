package src

import (
	"crypto/tls"
	"log/slog"
	"net/http"
	"os"
)

func WorkWithApi(rmqPathToFile string, rmqUrl string, rmqPass string, rmqUser string, logger *slog.Logger) {
	file, err := os.Open(rmqPathToFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	req, err := http.NewRequest("POST", rmqUrl, file)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	// авторизация Basic
	req.SetBasicAuth(rmqUser, rmqPass)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // <- отключает проверку (ОПАСНО)
	}

	// Транспорт с нашим tlsConfig
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	// Клиент, который использует этот транспорт
	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Status: ", resp.Status)
		panic(err)
	}
	logger.Info("🕒 Post request to RabbitMQ for update: " + rmqUrl)

	switch resp.StatusCode {
	case 200, 204:
		logger.Info("✅ Response OK", "status", resp.Status)
	default:
		logger.Error("❌ Unexpected response status", "status", resp.Status)
	}
	defer resp.Body.Close()
}
