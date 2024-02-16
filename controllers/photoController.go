package controllers

import (
	"final_project/helpers"
	"final_project/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhotoRepo struct {
	DB *gorm.DB
}

func (p *PhotoRepo) GetAllPhoto(c *gin.Context) {
	Photos := []models.Photo{}

	if err := p.DB.Debug().Preload("Comments").Preload("User").Find(&Photos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "no photo found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successfully retrieve all photos",
		"data":    Photos,
	})
}

func (p *PhotoRepo) GetPhotoByID(c *gin.Context) {
	GetId, _ := strconv.Atoi(c.Param("photoId"))
	Photo := []models.Photo{}

	if err := p.DB.Debug().Preload("Comments").Preload("User").Find(&Photo, GetId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "photo not found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successfully retrieve the photo",
		"data":    Photo,
	})
}

func (p *PhotoRepo) CreatePhoto(c *gin.Context) {
	Photo := models.Photo{}
	contextType := helpers.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	if contextType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = uint(userId)
	Photo.CreatedAt = time.Now()
	Photo.UpdatedAt = time.Now()

	if err := p.DB.Debug().Create(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "fail to upload photo",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "successfully upload your photo",
		"data":    Photo,
	})
}

func (p *PhotoRepo) UpdatePhoto(c *gin.Context) {
	GetId, _ := strconv.Atoi(c.Param("photoId"))

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"]

	contextType := helpers.GetContentType(c)
	Photo := models.Photo{}
	OldPhoto := models.Photo{}

	if contextType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UpdatedAt = time.Now()
	Photo.UserID = uint(userId.(float64))

	if err := p.DB.Debug().First(&OldPhoto, GetId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "photo not found",
			"message": err.Error(),
		})
		return
	}

	if err := p.DB.Debug().Model(&OldPhoto).Updates(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "fail to update photo",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successfully update your photo",
		"data":    OldPhoto,
	})
}

func (p *PhotoRepo) DeletePhoto(c *gin.Context) {
	GetId, _ := strconv.Atoi(c.Param("photoId"))
	Photo := models.Photo{}

	if err := p.DB.Debug().First(&Photo, GetId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "photo not found",
			"message": err.Error(),
		})
		return
	}

	if err := p.DB.Debug().Delete(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "failed to delete photo",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "succes",
		"message": "succesfully delete your photo",
	})
}
