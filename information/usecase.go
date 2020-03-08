package resume

import (
	"github.com/gin-gonic/gin"
	"github.com/passakorn-new/resume/models"
)

// Repository represent the article's repository contract
type Usecase interface {
	GetProjects(c *gin.Context) (res []*models.Project, err error)
	GetIntroduction(c *gin.Context) (res []*models.Introduction, err error)
	GetWriting(c *gin.Context) (res []*models.Writing, err error)
}
