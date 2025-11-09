package main

import (
	"backend/internal/database"

	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/database/query",
		WithUnitTest: true,
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(database.Question{}, database.ImageFile{})

	// Generate the code
	g.Execute()
}
