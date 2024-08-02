package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/koujiman/gobox/api/controller/http/folder"
	"github.com/koujiman/gobox/database"
)

func main() {
	db := database.NewDatabase()

	folderController := folder.InjectFolder(db)

	engine := gin.Default()
	group := engine.Group("/api")

	group.POST("/folders", folderController.CreateFolder)

	if err := engine.Run(); err != nil {
		log.Fatalln(err)
	}
}
