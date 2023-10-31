package muridController

import (
	"gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context)  {
	var dataMurid []models.Murid

	models.DB.Find(&dataMurid)

	c.IndentedJSON(http.StatusOK, gin.H{
		"murid": dataMurid,
	})
}

func Show(c *gin.Context)  {
	id := c.Param("id")
	var dataMurid models.Murid

	if err := models.DB.First(&dataMurid, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"murid": dataMurid})
}

func Store(c *gin.Context)  {
	var dataMurid models.Murid

	if err := c.ShouldBindJSON(&dataMurid); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&dataMurid)
	c.IndentedJSON(http.StatusCreated, gin.H{"murid": dataMurid})
}

func Destroy(c *gin.Context)  {
	id := c.Param("id")

	var dataMurid models.Murid

	if err := models.DB.First(&dataMurid, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	if models.DB.Delete(&dataMurid, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data tidak dapat dihapus"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}

func Update (c *gin.Context) {
	id := c.Param("id")
	var dataMurid models.Murid
	if err := c.ShouldBindJSON(&dataMurid); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&dataMurid).Where("id = ?", id).Updates(&dataMurid).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data tidak berhasil diupdate"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"murid": dataMurid,
		"message": "Data berhasil diupdate",
	})
}