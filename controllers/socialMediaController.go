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

type SocialMediaRepo struct {
	DB *gorm.DB
}

func (s *SocialMediaRepo) GetAllSocialMedia(c *gin.Context) {
	SocialMedias := []models.SocialMedia{}

	if err := s.DB.Debug().Preload("User").Find(&SocialMedias).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "no social media found",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successfully retrieve all social medias",
		"data":    SocialMedias,
	})
}

func (s *SocialMediaRepo) GetSocialMediaByID(c *gin.Context) {
	getId, _ := strconv.Atoi(c.Param("socialMediaId"))
	SocialMedia := []models.SocialMedia{}

	if err := s.DB.Debug().Preload("User").Find(&SocialMedia, getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "social media not found",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successfully retrieve the social media",
		"data":    SocialMedia,
	})
}

func (s *SocialMediaRepo) CreateSocialMedia(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	if contentType == "application/json" {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = uint(userId)
	SocialMedia.CreatedAt = time.Now()
	SocialMedia.UpdatedAt = time.Now()

	if err := s.DB.Debug().Create(&SocialMedia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "fail to create social media",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "successfully create your social media",
		"data":    SocialMedia,
	})
}

func (s *SocialMediaRepo) UpdateSocialMedia(c *gin.Context) {
	contentType := helpers.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	getId, _ := strconv.Atoi(c.Param("socialMediaId"))

	SocialMedia := models.SocialMedia{}
	OldSocialMedia := models.SocialMedia{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UpdatedAt = time.Now()
	SocialMedia.UserID = uint(userId)

	if err := s.DB.Debug().First(&OldSocialMedia, getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "social media not found",
			"message": err.Error(),
		})
		return
	}
	if err := s.DB.Debug().Model(&OldSocialMedia).Updates(&SocialMedia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "fail to update media",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successfully update your social media",
		"data":    OldSocialMedia,
	})

}

func (s *SocialMediaRepo) DeletSocialeMedia(c *gin.Context) {
	getId, _ := strconv.Atoi(c.Param("socialMediaId"))
	SocialMedia := models.SocialMedia{}

	if err := s.DB.Debug().First(&SocialMedia, getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "social media not found",
			"message": err.Error(),
		})
		return
	}
	if err := s.DB.Debug().Delete(&SocialMedia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "fail to delete media",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "succes",
		"message": "succesfully delete your social media",
	})
}
