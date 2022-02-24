/*
 * MassBank Tool API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"github.com/uly55e5/MassBankRepo/api-server/openapi"
	"github.com/uly55e5/MassBankRepo/api-server/server"
	"log"
	"net/http"
	"os"
)

func main() {
	logFile, err := os.OpenFile("/var/log/massbank-server.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)

	server.Init()
	defer server.Close()

	DefaultApiService := openapi.NewDefaultApiService()
	DefaultApiController := openapi.NewDefaultApiController(DefaultApiService)

	router := openapi.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
