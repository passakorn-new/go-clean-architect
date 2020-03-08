package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/models"
	"github.com/gin-gonic/gin"
	information "github.com/passakorn-new/resume/information"
	"github.com/sirupsen/logrus"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ArticleHandler  represent the httphandler for article
type InformationHandler struct {
	IUsecase information.Usecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewInformationHandlerHandler(g *gin.Engine, us information.Usecase) {

	handler := &InformationHandler{
		IUsecase: us,
	}

	g.GET("/projects", handler.GetProjects)
	g.GET("/introduction", handler.GetIntroduction)
	g.GET("/writing", handler.GetWriting)
}

func (i *InformationHandler) GetProjects(c *gin.Context) {
	projects, err := i.IUsecase.GetProjects(c)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, projects)
}

func (i *InformationHandler) GetIntroduction(c *gin.Context) {
	introductions, err := i.IUsecase.GetIntroduction(c)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, introductions)
}

func (i *InformationHandler) GetWriting(c *gin.Context) {
	writings, err := i.IUsecase.GetWriting(c)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, writings)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
