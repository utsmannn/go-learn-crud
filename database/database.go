package database

import (
	"database/sql"
	"fmt"
	"learn-crud/data"

	_ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	url := "admin:1234@tcp(localhost:8889)/test_crud"
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("connection success")
	return db
}

func GetAll() ([]data.Student, error) {
	db := connect()
	query, err := db.Query("select * from student order by id")
	if err != nil {
		panic(err.Error())
	}

	student := data.Student{}
	students := []data.Student{}

	for query.Next() {
		var id int
		var name string
		var age int

		err = query.Scan(&id, &name, &age)
		if err != nil {
			panic(err.Error())
		}

		student.Id = id
		student.Name = name
		student.Age = age
		students = append(students, student)
	}

	defer db.Close()
	return students, err
}

func GetById(id string) ([]data.Student, error) {
	db := connect()
	query, err := db.Query("select * from student where id=?", id)
	if err != nil {
		panic(err.Error())
	}

	student := data.Student{}
	students := []data.Student{}
	for query.Next() {
		var id int
		var name string
		var age int

		err = query.Scan(&id, &name, &age)
		if err != nil {
			panic(err.Error())
		}

		student.Id = id
		student.Name = name
		student.Age = age
	}

	if student.Id != 0 {
		students = append(students, student)
	}

	defer db.Close()
	return students, err
}

func Insert(student data.Student) error {
	db := connect()
	query, err := db.Prepare("insert into student(name, age) values(?,?)")
	if err != nil {
		panic(err.Error())
	}

	query.Exec(student.Name, student.Age)
	fmt.Println("insert sucess")

	defer db.Close()
	return err
}

func Update(id string, student data.Student) ([]data.Student, error) {
	db := connect()

	queryString, qType := configQueryUpdate(student)

	query, err := db.Prepare(queryString)
	if err != nil {
		panic(err.Error())
	}

	switch qType {
	case 1:
		query.Exec(student.Age, id)
	case 2:
		query.Exec(student.Name, id)
	case 3:
		query.Exec(student.Name, student.Age, id)
	}

	students, err := GetById(id)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	return students, err
}

func Delete(id string) error {
	db := connect()
	query, err := db.Prepare("delete from student where id=?")
	if err != nil {
		panic(err.Error())
	}
	query.Exec(id)
	defer db.Close()
	return err
}

func configQueryUpdate(student data.Student) (string, int) {
	var query string
	var qType int
	if student.Name == "" {
		query = "update student set age=? where id=?"
		qType = 1
	} else if student.Age == 0 {
		query = "update student set name=? where id=?"
		qType = 2
	} else {
		query = "update student set name=?, age=? where id=?"
		qType = 3
	}

	return query, qType
}
