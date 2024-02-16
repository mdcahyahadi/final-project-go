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

type CommentRepo struct {
	DB *gorm.DB
}

func (o *CommentRepo) GetAllComment(c *gin.Context) {
	Comments := []models.Comment{}

	if err := o.DB.Debug().Preload("User").Preload("Photo").Find(&Comments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "no comment found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successfully retrieve all comments",
		"data":    Comments,
	})
}

func (o *CommentRepo) GetCommentByID(c *gin.Context) {
	getId, _ := strconv.Atoi(c.Param("commentId"))
	Comment := []models.Comment{}

	if err := o.DB.Debug().Preload("User").Preload("Photo").Find(&Comment, getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "no comment found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successfully retrieve the comment",
		"data":    Comment,
	})
}

func (o *CommentRepo) CreateComment(c *gin.Context) {
	contentType := helpers.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)
	Comment := models.Comment{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	if err := o.DB.Debug().Find(&models.Comment{}, Comment.PhotoID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "photo not found",
			"message": err.Error(),
		})
		return
	}
	Comment.UserID = uint(userId)
	Comment.CreatedAt = time.Now()
	Comment.UpdatedAt = time.Now()

	if err := o.DB.Debug().Create(&Comment).Error; err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"error":   "fail to create comment",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "successfully create your comment",
		"data":    Comment,
	})
}

func (o *CommentRepo) UpdateComment(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}
	OldComment := models.Comment{}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	getId, _ := strconv.Atoi(c.Param("commentId"))

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = uint(userId)
	Comment.UpdatedAt = time.Now()

	if err := o.DB.Debug().First(&OldComment, getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "fail",
			"error":    "comment not found",
			"messsage": err.Error(),
		})
		return
	}

	if err := o.DB.Debug().Model(&OldComment).Updates(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "fail to update comment",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successfully update your comment",
		"data":    OldComment,
	})
}

func (o *CommentRepo) DeleteComment(c *gin.Context) {
	getId, _ := strconv.Atoi(c.Param("commentId"))

	Comment := models.Comment{}

	if err := o.DB.Debug().First(&Comment, getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "comment not found",
			"message": err.Error(),
		})
		return
	}

	if err := o.DB.Debug().Delete(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "fail to delete comment",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "succes",
		"message": "succesfully delete your comment",
	})

}
