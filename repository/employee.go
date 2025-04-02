package repository

import (
	"context"

	"github.com/Jkrish1011/projectx-go/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepo struct {
	MongoCollection *mongo.Collection // Collection on which all our operations are based on
}

func (r *EmployeeRepo) InsertEmployee(emp *model.Employee) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), emp)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (r *EmployeeRepo) FindEmployeeByID(empID string) (*model.Employee, error) {
	var emp model.Employee // Data from mongo is from bson format. Needs convertion to json.

	err := r.MongoCollection.FindOne(context.Background(), bson.D{{Key: "employee_id", Value: empID}}).Decode(&emp)

	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func (r *EmployeeRepo) FindAllEmployee() ([]model.Employee, error) {
	ctx := context.Background()
	results, err := r.MongoCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var emps []model.Employee
	err = results.All(ctx, &emps)
	if err != nil {
		return nil, err
	}
	return emps, nil

}

// The return type indicates how many records have been updated
func (r *EmployeeRepo) UpdateEmployeeById(empID string, updateEmp *model.Employee) (int64, error) {
	results, err := r.MongoCollection.UpdateOne(context.Background(),
		bson.D{{Key: "employee_id", Value: empID}},
		bson.D{{Key: "$set", Value: updateEmp}})

	if err != nil {
		return 0, err
	}

	return results.ModifiedCount, nil
}

func (r *EmployeeRepo) DeleteEmployeeByID(empID string) (int64, error) {
	results, err := r.MongoCollection.DeleteOne(context.Background(),
		bson.D{{Key: "employee_id", Value: empID}})

	if err != nil {
		return 0, err
	}
	return results.DeletedCount, nil
}

// To clear the database
func (r *EmployeeRepo) DeleteAllEmployee() (int64, error) {
	results, err := r.MongoCollection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		return 0, err
	}
	return results.DeletedCount, nil
}
