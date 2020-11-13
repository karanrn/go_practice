package main

import (
	"errors"
	"fmt"
	"reflect"
)

// Employee type for employee details
type Employee struct {
	name    string
	age     int
	salary  int
	country string
}

// Project type for project details in the organisation
type Project struct {
	name    string
	client  string
	country string
}

// createQuery returns insert query for the given struct
func createQuery(q interface{}) (string, error) {
	// Support only for struct types
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values (", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s %d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				return "", errors.New("unsupported data type for the field")
			}
		}

		query = fmt.Sprintf("%s)", query)
		return query, nil
	}
	return "", errors.New("unsupported type")
}

func main() {
	p := Project{
		name:    "ABC",
		client:  "Amazon",
		country: "USA",
	}

	e := Employee{
		name:    "John",
		age:     35,
		salary:  200000,
		country: "India",
	}

	pQuery, err := createQuery(p)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Project: %s\n", pQuery)

	eQuery, err := createQuery(e)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("\nEmployee: %s\n", eQuery)
}
