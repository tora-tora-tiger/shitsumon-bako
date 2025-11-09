package database

func Migrate() {
	DB.AutoMigrate(&ImageFile{})
	DB.AutoMigrate(&Question{})
}