package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/super-link-manager/utils"
	"net/http"
	"os"
)

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	response := HealthCheckResponse{}

	response.IsHealthy = checkDb()

	err := json.NewEncoder(w).Encode(response)
	utils.CheckErr(err)
}

type HealthCheckResponse struct {
	IsHealthy bool `json:"IsHealthy"`
}

func checkDb() bool {
	var host = os.Getenv("POSTGRES_HOST")
	var port = os.Getenv("POSTGRES_PORT")
	var user = os.Getenv("POSTGRES_USERNAME")
	var password = os.Getenv("POSTGRES_PASSWORD")
	var database = os.Getenv("POSTGRES_DB")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password='%s' dbname=%s sslmode=disable", host, port, user, password, database)
	_, err := sql.Open("postgres", connectionString)
	if err != nil {
		return false
	}

	return true
}
