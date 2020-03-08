package usecase

import (
	"github.com/gin-gonic/gin"
	information "github.com/passakorn-new/resume/information"
	"github.com/passakorn-new/resume/models"
)

type informationUsecase struct {
	informationRepo information.Repository
	// contextTimeout time.Duration
}

// NewArticleUsecase will create new an articleUsecase object representation of article.Usecase interface
func NewInformationUsecase(i information.Repository) information.Usecase {
	return &informationUsecase{
		informationRepo: i,
	}
}

func (i *informationUsecase) GetProjects(c *gin.Context) ([]*models.Project, error) {
	res, err := i.informationRepo.GetProjects(c)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (i *informationUsecase) GetIntroduction(c *gin.Context) ([]*models.Introduction, error) {
	res, err := i.informationRepo.GetIntroduction(c)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (i *informationUsecase) GetWriting(c *gin.Context) ([]*models.Writing, error) {
	res, err := i.informationRepo.GetWriting(c)
	if err != nil {
		return nil, err
	}
	return res, nil
}
