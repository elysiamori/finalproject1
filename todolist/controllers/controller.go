package controllers

import (
	"net/http"

	"github.com/elysiamori/finalproject1/kelompok6/config"
	"github.com/elysiamori/finalproject1/kelompok6/models"
	"github.com/gin-gonic/gin"
)

// Get All Todos
// @Summary Get All Todos
// @Description Get All Todos
// @Tags Todos
// @Param request body models.Todos true "Payload Body [RAW]"
// @Accept json
// @Produce json
// @Success 200 {object} models.Todos
// @Failure 400
// @Router /todos [get]
func GetAllTodos(c *gin.Context) {
	db := config.DBConn()
	var todos []models.Todos
	err := db.Find(&todos).Error

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Gagal menampilkan todo",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"todos": todos,
	})

}

// Get Todos By ID
// @Summary Get Todos By ID
// @Description Get Todos By ID
// @Tags Todos
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Todos
// @Failure 404
// @Router /todos/{id} [get]
func GetTodosID(c *gin.Context) {
	db := config.DBConn()
	var todos models.Todos
	id := c.Param("id")
	err := db.First(&todos, id).Error

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "Todo tidak ditemukan",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

// Add Todos
// @Summary Add Todos
// @Description Add Todos
// @Tags Todos
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param completed body bool true "Completed"
// @Success 200 {object} models.Todos
// @Failure 400
// @Router /todos [post]
func AddTodos(c *gin.Context) {
	db := config.DBConn()
	var todos models.Todos
	err := c.ShouldBindJSON(&todos)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Gagal menambahkan todo",
		})
		return
	}

	if err := db.Create(&todos).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Gagal menambahkan todo",
		})
		return
	}
	// inputTodos := models.Todos{Title: todos.Title, Completed: todos.Completed}
	// db.Create(&inputTodos)

	c.IndentedJSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

// Update Todos
// @Summary Update Todos
// @Description Update Todos
// @Tags Todos
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param title body string true "Title"
// @Param completed body bool true "Completed"
// @Success 200 {object} models.Todos
// @Failure 404
// @Failure 400
// @Router /todos/{id} [put]
func UpdateTodos(c *gin.Context) {
	db := config.DBConn()
	var todos models.Todos
	id := c.Param("id")
	err := db.First(&todos, id).Error

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "Todo tidak ditemukan",
		})
		return
	}

	c.ShouldBindJSON(&todos)
	err = db.Save(&todos).Error

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Gagal mengupdate todo",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

// Delete Todos
// @Summary Delete Todos
// @Description Delete Todos
// @Tags Todos
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} string
// @Failure 404
// @Failure 400
// @Router /todos/{id} [delete]
func DeleteTodos(c *gin.Context) {
	db := config.DBConn()
	var todos models.Todos
	id := c.Param("id")
	err := db.First(&todos, id).Error

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "Todo tidak ditemukan",
		})
		return
	}

	err = db.Delete(&todos).Error

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Gagal menghapus todo",
		})
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Berhasil menghapus todo",
	})
}
