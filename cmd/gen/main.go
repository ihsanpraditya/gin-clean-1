package main

import (
	"log"
	"gorm.io/driver/postgres" // Or mysql, depending on your setup
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// 1. Connect to your local DB (same credentials as your config/database.go)
	dsn := "host=localhost user=postgres password=secret dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 2. Initialize generator pointing to your project folders
	g := gen.NewGenerator(gen.Config{
		OutPath:      "query",          // Where the query code will live
		ModelPkgPath: "models",         // Where your existing structs live
		Mode:         gen.WithDefaultQuery | gen.WithGeneric,
	})

	g.UseDB(db)

	// 3. Generate code based on your actual DB tables
	g.ApplyBasic(g.GenerateAllTable()...)

	// 4. Run it
	g.Execute()
}
