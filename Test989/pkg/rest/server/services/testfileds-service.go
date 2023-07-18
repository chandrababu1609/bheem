package services

import (
	"github.com/chandrababu1609/bheem/test989/pkg/rest/server/daos"
	"github.com/chandrababu1609/bheem/test989/pkg/rest/server/models"
)

type TestfiledsService struct {
	testfiledsDao *daos.TestfiledsDao
}

func NewTestfiledsService() (*TestfiledsService, error) {
	testfiledsDao, err := daos.NewTestfiledsDao()
	if err != nil {
		return nil, err
	}
	return &TestfiledsService{
		testfiledsDao: testfiledsDao,
	}, nil
}

func (testfiledsService *TestfiledsService) CreateTestfileds(testfileds *models.Testfileds) (*models.Testfileds, error) {
	return testfiledsService.testfiledsDao.CreateTestfileds(testfileds)
}

func (testfiledsService *TestfiledsService) UpdateTestfileds(id int64, testfileds *models.Testfileds) (*models.Testfileds, error) {
	return testfiledsService.testfiledsDao.UpdateTestfileds(id, testfileds)
}

func (testfiledsService *TestfiledsService) DeleteTestfileds(id int64) error {
	return testfiledsService.testfiledsDao.DeleteTestfileds(id)
}

func (testfiledsService *TestfiledsService) ListTestfileds() ([]*models.Testfileds, error) {
	return testfiledsService.testfiledsDao.ListTestfileds()
}

func (testfiledsService *TestfiledsService) GetTestfileds(id int64) (*models.Testfileds, error) {
	return testfiledsService.testfiledsDao.GetTestfileds(id)
}
