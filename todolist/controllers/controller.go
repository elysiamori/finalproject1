package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/elysiamori/finalproject1/kelompok6/config"
	"github.com/elysiamori/finalproject1/kelompok6/models"
	"github.com/gin-gonic/gin"
)

// Get All Todos
// @Summary Get All Todos
// @Description Get All Todos
// @Tags Todos
// @Accept json
// @Produce json
// @Success 200 {object} models.Todos
// @Failure 400
// @Router /todos [get]
func GetAllTodos(c *gin.Context) {
	db, errDB := config.DBConn()
	if errDB != nil {
		fmt.Println(errDB)
	}

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
	db, errDB := config.DBConn()
	if errDB != nil {
		fmt.Println(errDB)
	}
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
// @Param title query string true "Title"
// @Param completed query boolean true "Completed Status"
// @Success 200 {object} models.Todos
// @Failure 400
// @Router /todos [post]
func AddTodos(c *gin.Context) {
	db, errDB := config.DBConn()
	if errDB != nil {
		fmt.Println(errDB)
	}
	var todos models.Todos

	title := c.DefaultQuery("title", "")
	completedStr := c.DefaultQuery("completed", "false")
	completed, _ := strconv.ParseBool(completedStr)

	todos.Title = title
	todos.Completed = completed

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

// UpdateTodos Todos
// @Summary Update Todos
// @Description Update Todos
// @Tags Todos
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Todos
// @Failure 400
// @Failure 404
// @Router /todos/{id} [put]
func UpdateTodos(c *gin.Context) {
	db, errDB := config.DBConn()
	if errDB != nil {
		fmt.Println(errDB)
	}
	var existingTodo models.Todos
	id := c.Param("id")

	// Check if the todo with the given ID exists
	if err := db.First(&existingTodo, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "Todo tidak ditemukan",
		})
		return
	}

	existingTodo.Completed = true

	// Save the updated todo to the database
	if err := db.Save(&existingTodo).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update todo",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"todos": existingTodo,
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
	db, errDB := config.DBConn()
	if errDB != nil {
		fmt.Println(errDB)
	}
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
