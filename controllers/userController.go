package controllers

import (
	"final_project/helpers"
	"final_project/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func (u *UserRepo) GetAllUser(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	_, _ = u.DB, contentType

	AllUser := []models.User{}

	if err := u.DB.Preload("Photos").Preload("Comments").Preload("SocialMedias").Find(&AllUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"error":   "data not found",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successfully retrieve all users",
		"data":    AllUser,
	})
}

func (u *UserRepo) UserRegister(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	_, _ = u.DB, contentType

	User := models.User{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.CreatedAt = time.Now()
	User.UpdatedAt = time.Now()

	if err := u.DB.Debug().Create(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "fail to create user data",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "successfully create your account",
		"data":    User,
	})
}

func (u *UserRepo) UserLogin(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	_, _ = u.DB, contentType

	User := models.User{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	password := User.Password

	if err := u.DB.Debug().Where("email = ?", User.Email).Take(&User).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"error":   "email not found, sign in to proceed",
			"message": err.Error(),
		})
		return
	}
	fmt.Println((User.Password), (password))
	if comparePass := helpers.ComparePass([]byte(User.Password), []byte(password)); !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"error":   "unauthorized",
			"message": "invalid email/password",
		})
		return
	}
	token := helpers.GenerateToken(uint(User.ID), User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (u *UserRepo) UserUpdate(c *gin.Context) {
	GetId, _ := strconv.Atoi(c.Param("userId"))
	UserData := c.MustGet("userData").(jwt.MapClaims)
	UserId := UserData["id"].(float64)

	contextType := helpers.GetContentType(c)
	_, _ = u.DB, contextType

	User := models.User{}
	OldUser := models.User{}

	if contextType == "application/json" {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.UpdatedAt = time.Now()
	User.ID = uint(UserId)

	if err := u.DB.Where("id = ?", GetId).Take(&OldUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"error":   "user not found",
			"message": err.Error(),
		})
		return
	}
	if err := u.DB.Preload("Photos").Preload("Comments").Preload("SocialMedias").Model(&OldUser).Updates(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "fail to update user data",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successfully update your account",
		"data":    OldUser,
	})
}

func (u *UserRepo) UserDelete(c *gin.Context) {
	UserData := c.MustGet("userData").(jwt.MapClaims)
	UserId := int(UserData["id"].(float64))
	User := models.User{}

	if err := u.DB.First(&User, UserId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"error":   "user not found",
			"message": err.Error(),
		})
		return
	}
	if err := u.DB.Preload("Photos").Preload("Comments").Preload("SocialMedias").Model(&User).Delete(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "fail to delete user data",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "succes",
		"message": "succesfully delete your account",
	})
}
