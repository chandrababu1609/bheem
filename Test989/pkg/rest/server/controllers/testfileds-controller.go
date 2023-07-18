package controllers

import (
	"github.com/chandrababu1609/bheem/test989/pkg/rest/server/models"
	"github.com/chandrababu1609/bheem/test989/pkg/rest/server/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type TestfiledsController struct {
	testfiledsService *services.TestfiledsService
}

func NewTestfiledsController() (*TestfiledsController, error) {
	testfiledsService, err := services.NewTestfiledsService()
	if err != nil {
		return nil, err
	}
	return &TestfiledsController{
		testfiledsService: testfiledsService,
	}, nil
}

func (testfiledsController *TestfiledsController) CreateTestfileds(context *gin.Context) {
	// validate input
	var input models.Testfileds
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger testfileds creation
	if _, err := testfiledsController.testfiledsService.CreateTestfileds(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Testfileds created successfully"})
}

func (testfiledsController *TestfiledsController) UpdateTestfileds(context *gin.Context) {
	// validate input
	var input models.Testfileds
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger testfileds update
	if _, err := testfiledsController.testfiledsService.UpdateTestfileds(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Testfileds updated successfully"})
}

func (testfiledsController *TestfiledsController) FetchTestfileds(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger testfileds fetching
	testfileds, err := testfiledsController.testfiledsService.GetTestfileds(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, testfileds)
}

func (testfiledsController *TestfiledsController) DeleteTestfileds(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger testfileds deletion
	if err := testfiledsController.testfiledsService.DeleteTestfileds(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Testfileds deleted successfully",
	})
}

func (testfiledsController *TestfiledsController) ListTestfileds(context *gin.Context) {
	// trigger all testfileds fetching
	testfileds, err := testfiledsController.testfiledsService.ListTestfileds()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, testfileds)
}

func (*TestfiledsController) PatchTestfileds(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*TestfiledsController) OptionsTestfileds(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*TestfiledsController) HeadTestfileds(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
