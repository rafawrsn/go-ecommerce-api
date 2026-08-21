package handlers

import (
	"errors"
	"go-ecommerce-api/config"
	"go-ecommerce-api/dto"
	"go-ecommerce-api/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProduct(c *gin.Context){
	var req dto.CreateProductRequest
	var category models.Category

	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"JSON Tidak Valid",
		})
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Name tidak boleh kosong",
		})
		return
	}

	if req.Price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Price tidak boleh kurang dari 0",
		})
		return
	}

	if req.Stock < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Stock tidak boleh kurang dari 0",
		})
		return
	}

	err = config.DB.First(&category, req.CategoryID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			c.JSON(http.StatusNotFound, gin.H{
				"error":"Category tidak ditemukan",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Database error",
		})
		return
	}

	product := models.Product{
		Name: req.Name,
		Price: req.Price,
		Stock: req.Stock,
		CategoryID: req.CategoryID,
	}

	err = config.DB.Create(&product).Error

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Database error",
		})
		return
	}

	response := dto.ProductResponse{
		ID:         product.ID,
    	Name:       product.Name,
    	Price:      product.Price,
    	Stock:      product.Stock,
    	CategoryID: product.CategoryID,
	}

	c.JSON(http.StatusCreated, response)
}

func GetProducts(c *gin.Context){
	var products []models.Product
	var total int64

	pageStr := c.DefaultQuery("page","1")
	limitStr := c.DefaultQuery("limit","10")

	page, err := strconv.Atoi(pageStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Page tidak valid",
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Limit tidak valid",
		})
		return
	}

	if page < 1{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Page tidak boleh kurang dari 1",
		})
		return
	}

	if limit < 1{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Limit tidak boleh kurang dari 1",
		})
		return
	}

	if limit > 100 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Limit tidak boleh lebih dari 100",
		})
		return
	}

	err = config.DB.Model(&models.Product{}).Count(&total).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Database error",
		})
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	offset := (page - 1) * limit

	err = config.DB.Preload("Category").Limit(limit).Offset(offset).Find(&products).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Database error",
		})
		return
	}

	responses := make([]dto.ProductResponse, 0, len(products))

	for _, product := range products {
    response := dto.ProductResponse{
        ID:         product.ID,
        Name:       product.Name,
        Price:      product.Price,
        Stock:      product.Stock,
        CategoryID: product.CategoryID,
        Category: dto.CategoryResponse{
            ID:   product.Category.ID,
            Name: product.Category.Name,
        },
    }

    responses = append(responses, response)
}

	c.JSON(http.StatusOK, gin.H{
		"data": responses,
		"page": page,
		"limit": limit,
		"total": total,
		"total_pages": totalPages,
	})

	}
	

	func GetProductByID(c *gin.Context){
		var product models.Product
		

		id := c.Param("id")

		idInt, err := strconv.Atoi(id)

		
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":"ID Tidak valid",
			})
			return
		}

		err = config.DB.Preload("Category").First(&product, idInt).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound){
				c.JSON(http.StatusNotFound, gin.H{
					"error":"Product tidak ditemukan",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"error":"Database error",
			})
			return
		}

		response := dto.ProductResponse{
		ID:         product.ID,
        Name:       product.Name,
        Price:      product.Price,
        Stock:      product.Stock,
        CategoryID: product.CategoryID,
        Category: dto.CategoryResponse{
            ID:   product.Category.ID,
            Name: product.Category.Name,
        },
		}

		c.JSON(http.StatusOK, response)

	}

func UpdateProduct(c *gin.Context){

	var product models.Product
	var category models.Category
	var req dto.UpdateProductRequest

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"ID Tidak Valid",
		})
		return
	}

	err = config.DB.First(&product, idInt).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			c.JSON(http.StatusNotFound, gin.H{
				"error":"Product tidak ditemukan",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Database error",
		})
		return
	}

	err = c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"JSON tidak valid",
		})
		return
	}

	if req.Name == ""{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Name tidak boleh kosong",
		})
		return
	}

	if req.Price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Price tidak boleh kurang dari 0",
		})
		return
	}

	if req.Stock < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Stock tidak boleh kurang dari 0",
		})
		return
	}

	err = config.DB.First(&category, req.CategoryID).Error

	if err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Category tidak ditemukan",
        })
        return
    }

    	c.JSON(http.StatusInternalServerError, gin.H{
        	"error": "Database error",
    	})
    	return
	}

	product.Name = req.Name
	product.Price = req.Price
	product.Stock = req.Stock
	product.CategoryID = req.CategoryID

	err = config.DB.Save(&product).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Database error",
		})
		return
	}

	c.JSON(http.StatusOK, product)

}

func DeleteProduct(c *gin.Context){
	var product models.Product

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"ID Tidak Valid",
		})
		return
	}
	
	err = config.DB.First(&product, idInt).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			c.JSON(http.StatusNotFound, gin.H{
				"error":"Product tidak ditemukan",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Database error",
		})
		return
	}

	err = config.DB.Delete(&product).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Database error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"Product berhasil di hapus",
	})
}