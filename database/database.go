package database

import "fmt"

var connection string

func init() { // this function will executed when package database imported by other
	fmt.Println("function init executed")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
