package main

import (
	"github.com/SevgiF/location-api/pkg/database/mysql"
	"github.com/SevgiF/location-api/pkg/validator"
)

func main() {
	//validation
	validator.Validate = validator.InitValidator()
	//mysql db connection
	db := mysql.NewGormDBManager().DB
}
