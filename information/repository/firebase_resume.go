package repository

import (
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	information "github.com/passakorn-new/resume/information"
	"github.com/passakorn-new/resume/models"
	"google.golang.org/api/iterator"
)

type firebaseInformationRepository struct {
	app *firestore.Client
}

// NewFirebaseInformation will create an object that represent the Information.Repository interface
func NewFirebaseInformation(app *firestore.Client) information.Repository {
	return &firebaseInformationRepository{app}
}

func (f *firebaseInformationRepository) GetProjects(c *gin.Context) ([]*models.Project, error) {
	ProjectsData := []*models.Project{}
	iter := f.app.Collection("projects").Documents(c)
	defer iter.Stop()
	for {
		ProjectData := models.Project{}
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		mapstructure.Decode(doc.Data(), &ProjectData)
		ProjectsData = append(ProjectsData, &ProjectData)
	}
	return ProjectsData, nil
}

func (f *firebaseInformationRepository) GetIntroduction(c *gin.Context) ([]*models.Introduction, error) {
	IntroductionsData := []*models.Introduction{}
	iter := f.app.Collection("introduction").Documents(c)
	defer iter.Stop()
	for {
		IntroductionData := models.Introduction{}
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		mapstructure.Decode(doc.Data(), &IntroductionData)
		IntroductionsData = append(IntroductionsData, &IntroductionData)
	}
	return IntroductionsData, nil
}

func (f *firebaseInformationRepository) GetWriting(c *gin.Context) ([]*models.Writing, error) {
	WritingsData := []*models.Writing{}
	iter := f.app.Collection("writing").Documents(c)
	defer iter.Stop()
	for {
		WritingData := models.Writing{}
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		mapstructure.Decode(doc.Data(), &WritingData)
		WritingsData = append(WritingsData, &WritingData)
	}
	return WritingsData, nil
}
