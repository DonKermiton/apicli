package collections

import (
	"apicli/request"
)

type Collection struct {
	Name     string
	Location Location
	Requests []request.Request
}

func CreateNewCollection(name string, location Location) {
	location.Create()
}

func (c Collection) SaveCollection() {
	c.Location.Save()
}

func (c Collection) GetCollection() {
	c.Location.Get()
}
