package logic

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sample-todo/config/database"
	"sample-todo/models"
)

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	db := database.GetDB()
	db.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	db := database.GetDB()

	if err := c.BindJSON(&task); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if create := db.Create(&task); create != nil {
		log.Println(create)
	}
	c.JSON(http.StatusOK, &task)
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	db := database.GetDB()

	id := c.Param("id")
	// ID 가 없응 StatusBadRequest
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "id is empty",
		})
		return
	} else if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		if err := c.BindJSON(&task); err != nil {
			return
		}
		db.Save(&task)
	}
}

func DeleteTask(c *gin.Context) {
	var task models.Task
	db := database.GetDB()

	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "id is must not be empty",
		})
	} else if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		db.Delete(&task)
	}
}
