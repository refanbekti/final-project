// package controllers

// import (
// 	"final-project/database"
// 	"final-project/helpers"
// 	"final-project/models"
// 	"net/http"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// )

// func CreatePhoto(c *gin.Context) {
// 	db := database.GetDB()
// 	userData := c.MustGet("userData").(jwt.MapClaims)
// 	contentType := helpers.GetContentType(c)

// 	Product := models.Product{}
// 	userID := uint(userData["id"].(float64))

// 	if contentType == appJSON {
// 		c.ShouldBindJSON(&Product)
// 	} else {
// 		c.ShouldBindJSON(&Product)
// 	}

// 	Product.UserID = userID

// 	err := db.Debug().Create(&Product).Error

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"err":     "Bad Request",
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, Product)
// }

// // func UpdateProduct(c *gin.Context) {
// // 	userData := c.MustGet("userData").(jwt.MapClaims)
// // 	contentType := helpers.GetContentType(c)
// // 	Product := models.Product{}
// // 	productId, err := strconv.Atoi(c.Param("productId"))
// // 	userID := uint(userData["id"].(float64))

// // 	if contentType == appJSON {
// // 		c.ShouldBindJSON(&Product)
// // 	} else {
// // 		c.ShouldBind(&Product)
// // 	}

// // 	Product.UserID = userID
// // 	Product.ID = uint(productId)

// // 	if err != nil {
// // 		c.JSON(http.StatusBadRequest, gin.H{
// // 			"err":     "Bad Request",
// // 			"message": err.Error(),
// // 		})
// // 		return
// // 	}
// // 	c.JSON(http.StatusOK, Product)
// // }
