package main

import "github.com/SevgiF/location-api/pkg/validator"

func main() {
	//validation
	validator.Validate = validator.InitValidator()
}
