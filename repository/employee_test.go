package repository

import (
	"context"
	"log"
	"testing"

	"github.com/Jkrish1011/projectx-go/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	ctx := context.Background()
	mongoTestClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://admin:EU2O84oldskBmZa9@cluster0.8ab3ox7.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	if err != nil {
		log.Fatal("Error connecting to mongodb client")
	}

	log.Println("Successfully connected to MongoDB")

	// Connecting to only Primary Instance since it is a cluster
	err = mongoTestClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Ping Failed, ", err)
	}
	log.Println("Ping to Primary Instance was successful!")

	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	ctx := context.Background()
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(ctx)

	emp1 := uuid.New().String()
	emp2 := uuid.New().String()

	collection := mongoTestClient.Database("companydb").Collection("employee_test")

	empRepo := EmployeeRepo{MongoCollection: collection}

	// Insert Employee 1 data
	t.Run("Insert Employee 1", func(t *testing.T) {
		emp := model.Employee{
			Name:       "jayakrishnan",
			Department: "software",
			EmployeeID: emp1,
		}

		result, err := empRepo.InsertEmployee(&emp)
		if err != nil {
			t.Fatal("Insert 1 operation failed, ", err)
		}

		t.Log("Insert 1 successful, ", result)
	})

	t.Run("Insert Employee 2", func(t *testing.T) {
		emp := model.Employee{
			Name:       "steve",
			Department: "physics",
			EmployeeID: emp2,
		}

		result, err := empRepo.InsertEmployee(&emp)
		if err != nil {
			t.Fatal("Insert 2 operation failed, ", err)
		}

		t.Log("Insert 2 successful, ", result)
	})

	t.Run("Get Employee 1", func(t *testing.T) {
		result, err := empRepo.FindEmployeeByID(emp1)
		if err != nil {
			t.Fatal("get emp1 operation failed, ", err)
		}

		t.Log("found emp 1, ", result)
	})

	t.Run("Get All Employees", func(t *testing.T) {
		results, err := empRepo.FindAllEmployee()
		if err != nil {
			t.Fatal("get all employee operation failed, ", err)
		}

		t.Log("found all, ", results)
	})

	t.Run("Update Employee 1 Name", func(t *testing.T) {
		emp := model.Employee{
			Name:       "Jayakrishnan Ashok",
			Department: "Software",
			EmployeeID: emp1,
		}

		result, err := empRepo.UpdateEmployeeById(emp1, &emp)
		if err != nil {
			t.Fatal("Updation Employee 1 failed, ", err)
		}

		t.Log("Updation successful, ", result)
	})

	t.Run("Get Employee 1 after update", func(t *testing.T) {
		result, err := empRepo.FindEmployeeByID(emp1)
		if err != nil {
			t.Fatal("get emp1 operation failed, ", err)
		}

		t.Log("found emp 1, ", result)
	})

	t.Run("Delete Employee 1 ", func(t *testing.T) {
		result, err := empRepo.DeleteEmployeeByID(emp1)
		if err != nil {
			t.Fatal("Delete emp1 operation failed, ", err)
		}

		t.Log("Deleted emp 1, ", result)
	})

	t.Run("Delete All Employees", func(t *testing.T) {
		result, err := empRepo.DeleteAllEmployee()
		if err != nil {
			t.Fatal("Delete all operation failed, ", err)
		}

		t.Log("Deleted all, ", result)
	})

	t.Run("Get All Employees after delete", func(t *testing.T) {
		results, err := empRepo.FindAllEmployee()
		if err != nil {
			t.Fatal("get all employee operation failed, ", err)
		}

		t.Log("found ", results)
	})

}
