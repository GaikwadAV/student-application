// Package controller API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
//     Schemes: http
//     Host: localhost:8000
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package controller

import (
	model "RESTAPI/models"
	"RESTAPI/utility"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// swagger:operation POST /student Student addstudent
//
// Add new Student
//
// Returns new Student
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - name: Student
//   in: body
//   description: add Student data
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   '201':
//     description: Student response
//     schema:
//       "$ref": "#/definitions/Student"
//   '405':
//     description: Method not allowed

//InsertStudent ...
func InsertStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client, ctx := utility.Connection()
	data := client.Database("student").Collection("studentinfo")
	var student model.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var result *mongo.InsertOneResult
	result, err = data.InsertOne(ctx, student)
	//log.Println("********************######", err)
	if err != nil {
		w.WriteHeader(http.StatusCreated)
		return
	}
	log.Println(student)
	json.NewEncoder(w).Encode(result)
}

// swagger:operation GET /students Student GetStudent
//
// Get Student
//
// Returns existing Student
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: Student data
//     schema:
//      "$ref": "#/definitions/Student"
//   '204':
//     description: No content

//GetAllStudent ...
func GetAllStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
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
		w.WriteHeader(http.StatusNoContent)
		return
	}
	for cur.Next(ctx) {
		var student model.Student
		err := cur.Decode(&student)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			return
		}
		students = append(students, student)
	}
	json.NewEncoder(w).Encode(students)

}

// swagger:operation GET /students/{id} Student GetStudentbyid
//
// Get Student
//
// Returns existing Student filtered by id
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
//     description: Student data
//     schema:
//      "$ref": "#/definitions/Student"
//   '400':
//     description: bad request

//GetStudentById ...
func GetStudentById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
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
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(result)
}

// swagger:operation DELETE /student/{id} Student deleteStudent
//
// delete student
//
// ---
// produces:
// - application/json
// parameters:
//  - name: id
//    in: path
//    type: string
//    required: true
// responses:
//	 '400':
//	   description: bad request
//	 '410':
//	   description: status gone

//deleteStudentById ...
func DeleteStudentById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	client, ctx := utility.Connection()
	data := client.Database("student").Collection("studentinfo")
	id := mux.Vars(r)["id"]
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	filter := primitive.M{"_id": ID}
	result, err := data.DeleteOne(ctx, filter)
	if err != nil {
		w.WriteHeader(http.StatusGone)
		return
	}
	w.WriteHeader(http.StatusGone)

	json.NewEncoder(w).Encode(result)

}

// swagger:operation PUT /students/{id} Student Updatestudent
//
// Update Student
//
// Update existing Student filtered by id
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
// - name: name
//   in: body
//   description: add Student data
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   '200':
//     description: Student response
//     schema:
//       "$ref": "#/definitions/Student"
//   '400':
//     description: bad request, invalid id
//   '201':
//     description: create

//EditUser ...
func EditUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	client, ctx := utility.Connection()
	data := client.Database("student").Collection("studentinfo")
	var params = mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	var student model.Student
	json.NewDecoder(r.Body).Decode(&student)
	update := primitive.M{"name": student.Name, "city": student.City, "country": student.Country, "YearOfAdmission": student.YearOfAdmission, "course": student.Course}

	err = data.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&student)

	if err != nil {
		w.WriteHeader(http.StatusCreated)
		return

	}
	json.NewEncoder(w).Encode(student)
}

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "homepage")
}
