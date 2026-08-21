package handlers

import (
	"net/http"
	"strconv"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-ecommerce-api/config"
	"go-ecommerce-api/models"
)

func CreateCategory(c *gin.Context) {

    var category models.Category

    err := c.BindJSON(&category)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid JSON",
        })
        return
    }

    if category.Name == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Category Name is Required",
        })
        return
    }

    err = config.DB.Create(&category).Error

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Database Error",
        })
        return
    }

    c.JSON(http.StatusCreated, category)
}

func GetCategories (c *gin.Context){
	var categories []models.Category

	err := config.DB.Find(&categories).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Database Error",
		})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func GetCategoryByID (c *gin.Context){
	var category models.Category

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid ID",
		})
		return
	}

	err = config.DB.First(&category, idInt).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Category Not Found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Database Error",
		})
		return
	}

	c.JSON(http.StatusOK, category)

}

func UpdateCategory (c *gin.Context){
	var category models.Category


	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid ID",
		})
		return
	}

	err = config.DB.First(&category, idInt).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Category Not Found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Database Error",
		})
		return
	}

	err = c.BindJSON(&category)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid JSON",
		})
		return
	}

	category.ID = uint(idInt)

	err = config.DB.Save(&category).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Database Error",
		})
		return
	}

	c.JSON(http.StatusOK, category)


}

func DeleteCategory (c *gin.Context){
	var category models.Category

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid ID",
		})
		return
	}

err = config.DB.First(&category, idInt).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Category Not Found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Database Error",
		})
		return
	}

	err = config.DB.Delete(&category).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Database Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"Category Deleted Successfully",
		})


}
