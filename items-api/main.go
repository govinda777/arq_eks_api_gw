package main

import (
	"github.com/marcosouzatech/items-api/api/src/database"
	"github.com/marcosouzatech/items-api/api/src/config"
	"github.com/marcosouzatech/items-api/api/src/router"
	"fmt"
	"log"
	"net/http"
	"os"

	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

func initNewRelic() {
	_, err := newrelic.NewApplication(
		newrelic.ConfigAppName(os.Getenv("APP_NAME")),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigDebugLogger(os.Stdout),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
func main() {
	if err := config.Load(); err != nil {
		log.Printf("Err ao carregar configurações: %v", err)
		// Em ambiente de container/k8s, às vezes as variáveis já estão injetadas sem .env
	}

	initNewRelic()

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Err ao conectar ao database de dados: %v", err)
	}
	db.Close()

	fmt.Printf("Listening on port %d\n", config.Port)
	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
