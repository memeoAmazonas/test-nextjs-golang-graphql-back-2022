package controller

import (
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/database"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/graph/model"
	"log"
)

const (
	DB  = "test_graphql"
	COL = "book"
)

type ControllerB interface {
	Save(book *model.Book)
	FindAll() []*model.Book
}

func Save(document *model.Book) interface{} {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer database.Close(client, ctx, cancel)
	cursor, err := database.SaveOne(client, ctx, DB, COL, document)
	if err != nil {
		panic(err)
	}
	return cursor.InsertedID
}
func FindAll() []*model.Book {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer database.Close(client, ctx, cancel)
	cursor, err := database.Query(client, ctx, DB, COL)
	if err != nil {
		panic(err)
	}
	var result []*model.Book
	for cursor.Next(ctx) {
		var v *model.Book
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, v)
	}
	return result
}
