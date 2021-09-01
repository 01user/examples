package dao

import "github.com/go-pg/pg/v10"

func getDb() *pg.DB {
	params := make(map[string]interface{})

	params["search_path"] = "public"
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "tantan_test",
		Addr:     "localhost:5432",
	}).WithParam("search_path", "public")
	return db
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
