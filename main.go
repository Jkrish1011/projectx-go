package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/Jkrish1011/projectx-go/usecase"
)

var mongoClient *mongo.Client

func init() {
	err := godotenv.Load()
	ctx := context.Background()

	if err != nil {
		log.Fatal("Couldn't load the env file. Quitting...")
	}
	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("Error connecting to the mongo client. Quitting...")
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the primary client, ", err)
	}
	log.Println("Connection to primary instance successful!")
}

func main() {
	ctx := context.Background()
	defer mongoClient.Disconnect(ctx)

	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	empService := usecase.EmployeeService{
		MongoCollection: coll,
	}

	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler).Methods(http.MethodGet)

	r.HandleFunc("/employee", empService.CreateEmployee).Methods(http.MethodPost)
	r.HandleFunc("/employee/{id}", empService.GetEmployeeByID).Methods(http.MethodPost)
	r.HandleFunc("/employee/{id}", empService.GetEmployeeByID).Methods(http.MethodGet)
	r.HandleFunc("/employee", empService.GetAllEmployee).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}", empService.UpdateEmployeeByID).Methods(http.MethodPut)
	r.HandleFunc("/employee/{id}", empService.DeleteEmployeeByID).Methods(http.MethodDelete)
	r.HandleFunc("/employee", empService.DeleteAllEmployee).Methods(http.MethodDelete)

	log.Println("Server running at port 4444")
	http.ListenAndServe(":4444", r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("running..."))
}
