package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ashishjuyal/banking-lib/logger"
	"github.com/ashishjuyal/banking/domain"
	"github.com/ashishjuyal/banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}

func Start() {

	//sanityCheck()

	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	dbClient := getDbClient()
	CustomerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)

	ProductRepositoryDb := domain.NewProductRepositoryDb(dbClient)

	ch := CustomerHandlers{service.NewCustomerService(CustomerRepositoryDb)}
	ph := ProductHandlers{service.NewProductService(ProductRepositoryDb)}

	// define routes
	router.
		HandleFunc("/premiumcustomers", ch.getAllCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")
	router.
		HandleFunc("/premiumcustomers/{Id:[0-9]+}", ch.getCustomer).
		Methods(http.MethodGet).
		Name("GetCustomer")

	router.
		HandleFunc("/products", ph.getAllProducts).
		Methods(http.MethodGet).
		Name("getAllProducts")
	router.
		HandleFunc("/products/{id:[0-9]+}", ph.getProduct).
		Methods(http.MethodGet).
		Name("GetProduct")

	//am := AuthMiddleware{domain.NewAuthRepository()}
	//router.Use(am.authorizationHandler())
	// starting server
	address := "localhost" //os.Getenv("SERVER_ADDRESS")
	port := "8000"         //os.Getenv("SERVER_PORT")
	logger.Info(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}

func getDbClient() *sqlx.DB {
	dbUser := "root"      //os.Getenv("DB_USER")
	dbPasswd := ""        //os.Getenv("DB_PASSWD")
	dbAddr := "localhost" //os.Getenv("DB_ADDR")
	dbPort := "3307"      //os.Getenv("DB_PORT")
	dbName := "northwind" //os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
