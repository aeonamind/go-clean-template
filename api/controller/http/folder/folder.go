package folder

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koujiman/gobox/domain"
)

type FolderController struct {
	createFolderUseCase  domain.CreateFolderUseCase
	getAllFoldersUseCase domain.GetAllFoldersUseCase
}

func ProvideFolderController(
	createFolderUseCase domain.CreateFolderUseCase,
	getAllFoldersUseCase domain.GetAllFoldersUseCase,
) *FolderController {
	return &FolderController{createFolderUseCase: createFolderUseCase, getAllFoldersUseCase: getAllFoldersUseCase}
}

func (co FolderController) CreateFolder(c *gin.Context) {
	var folder domain.Folder
	err := c.Bind(&folder)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	log.Print("running")

	c.JSON(http.StatusOK, "hehe")
}
