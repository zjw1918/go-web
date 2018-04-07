package main

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	Name string
	Age int
	Address string

	Grade int
}
