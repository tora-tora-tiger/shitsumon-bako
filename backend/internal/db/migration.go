package db

func Migrate() {
	DB.AutoMigrate(&ImageFile{})
	DB.AutoMigrate(&Question{})
}