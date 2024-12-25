package models

import "sort"

var gathered = [][]Schema{
	Challenges,
	Courses,
	Staff,
	Subscriptions,
	Theory,
	Users,
}

var Schemas []Schema

func init() {
	var schemasLen int

	for _, schemas := range gathered {
		schemasLen += len(schemas)
	}

	Schemas = make([]Schema, 0, schemasLen)

	for _, schemas := range gathered {
		Schemas = append(Schemas, schemas...)
	}

	sort.Slice(Schemas, func(i, j int) bool {
		return Schemas[i].Order < Schemas[j].Order
	})
}
