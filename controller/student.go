// package controller student API.
//
// the purpose of this appliation is to provide an application
// ths is using go code to define an rest API
//
//     Schemes: http
//     Host: localhost:8000
//     
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
// swagger:meta
package controller

import (
	model "RESTAPI/models"
	"RESTAPI/utility"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// swagger:operation POST /student student add student
//
// Add new student
//
// Returns new student
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - name: student
//   in: body
//   description: add student data
//   required: true
//   schema:
//     "$ref": "#/definitions/student"
// responses:
//   '200':
//     description: student response
//     schema:
//       "$ref": "#/definitions/student"
//   '405':
//     description: Method Not Allowed
//   '403':
//     description: Forbidden
// InsertStudent in database
func InsertStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client, ctx := utility.Connection()
	data := client.Database("student").Collection("studentinfo")
	var student model.Student
	json.NewDecoder(r.Body).Decode(&student)
	result, err := data.InsertOne(ctx, student)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(result)
}

// swagger:operation GET /student student get student
//
// Get student
//
// Returns existing student filtered by id
//
// ---
// produces:
// - application/json
// parameters:
//  - name: name
//    type: string
//    in: query
//    required: true
//  - name: city
//    type: string
//    in: query
//    required: true
// responses:
//   '200':
//     description: student data
//     schema:
//      "$ref": "#/definitions/student"
//   '405':
//     description: Method Not Allowed
//   '403':
//     description: Forbidden

// GetAllStudent from database
func GetAllStudent(w http.ResponseWriter, r *http.Request) {

	var students []model.Student
	client, ctx := utility.Connection()
	data := client.Database("student").Collection("studentinfo")
	name := r.URL.Query().Get("name")
	city := r.URL.Query().Get("city")
	params := []primitive.M{}
	filter := primitive.M{}
	if name != "" {
		params = append(params, primitive.M{"name": name})
	}
	if city != "" {
		params = append(params, primitive.M{"city": city})
	}
	if len(params) > 0 {
		filter = primitive.M{"$and": params}
	}
	cur, err := data.Find(ctx, filter)
	if err != nil {
		return
	}
	for cur.Next(ctx) {
		var student model.Student
		err := cur.Decode(&student)
		if err != nil {
			return
		}
		students = append(students, student)
	}
	json.NewEncoder(w).Encode(students)

}

// swagger:operation GET /student/{id} student getstudent
//
// Get student
//
// Returns existing student filtered by id
//
// ---
// produces:
// - application/json
// parameters:
//  - name: id
//    type: string
//    in: path
//    required: true
// responses:
//   '200':
//     description: student data
//     schema:
//      "$ref": "#/definitions/student"
//   '405':
//     description: Method Not Allowed, likely url is not correct
//   '403':
//     description: Forbidden, you are not allowed to undertake this operation
// GetStudentById
func GetStudentById(w http.ResponseWriter, r *http.Request) {
	client, ctx := utility.Connection()
	data := client.Database("student").Collection("studentinfo")
	id := mux.Vars(r)["id"]
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var result model.Student
	filter := primitive.M{}

	if id != "" {
		filter = primitive.M{"_id": ID}
	}
	err = data.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

// swagger:operation DELETE /student/{id} delete student by id
//
// delete student by id
// return deleted student
// ---
//produces:
// - application/json
// parameters;
// - name: id
//   type: string
//	 in: path
//	 description: hex id
//   required: true
// responses:
//	'200':
//		description: student respose
//		schema:
//			$ref: "#/definitions/student"
//	'405':
//		description: method not allowed

// DeleteStudentById delete a student by id
func DeleteStudentById(w http.ResponseWriter, r *http.Request) {
	client, ctx := utility.Connection()
	data := client.Database("student").Collection("studentinfo")
	id := mux.Vars(r)["id"]
	ID, _ := primitive.ObjectIDFromHex(id)
	filter := primitive.M{"_id": ID}
	result, err := data.DeleteOne(ctx, filter)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(result)

}

// swagger:operation PUT /student/{id} update student
//
// Update existing student
//
// Update existing student filtered by its id
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - name: id
//   type: string
//   in: path
//   required: true
// - name: student
//   in: body
//   description: add student data
//   required: true
//   schema:
//     "$ref": "#/definitions/student"
// responses:
//   '200':
//     description: student response
//     schema:
//       "$ref": "#/definitions/student"
//   '405':
//     description: Method Not Allowed
//   '403':
//     description: Forbidden
func EditUser(w http.ResponseWriter, r *http.Request) {
	client, ctx := utility.Connection()
	data := client.Database("student").Collection("studentinfo")
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])
	var student model.Student
	json.NewDecoder(r.Body).Decode(&student)
	update := primitive.M{"name": student.Name, "city": student.City, "country": student.Country, "YearOfAdmission": student.YearOfAdmission, "course": student.Course}

	err := data.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": update}).Decode(&student)

	if err != nil {
		return

	}

	json.NewEncoder(w).Encode(student)
}
