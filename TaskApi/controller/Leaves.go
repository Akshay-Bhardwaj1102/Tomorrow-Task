package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"task.com/config"
	"task.com/model"
)

// POST /api/leave
func CreateLeave(c *gin.Context) {

	var empLeave model.EmpLeave
	if err := c.ShouldBind(&empLeave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if empLeave.LeaveType == "sick" {

		// Handle image upload
		formFile, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Image upload is required"})
			return
		}

		// Read image file content
		fileContent, err := formFile.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image file"})
			return
		}
		defer fileContent.Close()

		// Store image content in the File model
		image := model.File{
			ImageName:    formFile.Filename,
			ImageContent: make([]byte, formFile.Size),
		}
		_, err = fileContent.Read(image.ImageContent)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image content"})
			return
		}

		// Save the EmpLeave data to the database using the db variable
		empLeave.Image = image
		config.DB.Create(&empLeave)
	} else {
		config.DB.Create(&empLeave)
	}

	c.JSON(http.StatusCreated, empLeave)
}

func GetLeaves(c *gin.Context) {
	var empLeave model.EmpLeave
	id := c.Param("id")

	// Find the employee's leave data by ID
	if err := config.DB.Where("id = ?", id).Preload("Image").First(&empLeave).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, empLeave)
}

func GetAllData(c *gin.Context) {
	var empLeaves []model.EmpLeave

	// Retrieve all employee leave data from the database
	if err := config.DB.Preload("Image").Find(&empLeaves).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from the database"})
		return
	}

	c.JSON(http.StatusOK, empLeaves)
}
