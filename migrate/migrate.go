package migrate

import (
	"github.com/Eswarakash/GoGinGORM/initializers"
	"github.com/Eswarakash/GoGinGORM/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func MigrateAuto() {
	initializers.DB.AutoMigrate(&models.Post{})
}
