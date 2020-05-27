package rest

import (
	"encoding/json"
	"io/ioutil"
	"learn-crud/data"
	"learn-crud/database"
	"net/http"
)

func message(err error) (string, bool) {
	if err != nil {
		return "Error", false
	} else {
		return "Success", true
	}
}

func messageRequestField(field string, err error) (string, bool) {
	if err != nil {
		return "Failed", false
	} else if field == "" {
		return "Wrong param", false
	} else {
		return "Success", true
	}
}

func getBody(r *http.Request) (data.Student, error) {
	bodyRaw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var student data.Student
	json.Unmarshal(bodyRaw, &student)

	return student, err
}

func SingleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")

	switch r.Method {
	case "POST":
		insert(w, r)
	case "PUT":
		update(w, r)
	case "DELETE":
		delete(w, r)
	default:
		getId(w, r)
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")

	dataRaw, err := database.GetAll()
	message, status := message(err)

	dataResponse := data.Responses{
		Status:  status,
		Message: message,
		Data:    dataRaw,
	}

	dataRes, err := json.Marshal(dataResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(dataRes)
}

func getId(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	dataRaw, err := database.GetById(idParam)
	message, status := messageRequestField(idParam, err)

	dataResponse := data.Responses{
		Status:  status,
		Message: message,
		Data:    dataRaw,
	}

	dataRes, err := json.Marshal(dataResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(dataRes)
}

func insert(w http.ResponseWriter, r *http.Request) {
	student, _ := getBody(r)
	err := database.Insert(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err.Error())
	}

	w.WriteHeader(http.StatusOK)
}

func update(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	student, err := getBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err.Error())
	}
	students, err := database.Update(idParam, student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err.Error())
	}

	message, status := message(err)
	dataResponse := data.Responses{
		Status:  status,
		Message: message,
		Data:    students,
	}

	dataRes, err := json.Marshal(dataResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(dataRes)
}

func delete(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	err := database.Delete(idParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err.Error())
	}

	w.WriteHeader(http.StatusOK)
}
