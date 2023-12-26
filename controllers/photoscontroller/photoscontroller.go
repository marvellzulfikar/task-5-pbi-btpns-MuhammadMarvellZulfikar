package photoscontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marvellzulfikar/task-5-pbi-btpns-MuhammadMarvellZulfikar/models"
)

func Index(c *gin.Context) {
	var photos []models.Photos

	models.DB.Find(&photos)
	c.JSON(http.StatusOK, gin.H{"photos": photos})
}

func Create(c *gin.Context) {
	var photos models.Photos

	if err := c.ShouldBindJSON(&photos); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&photos)
	c.JSON(http.StatusOK, gin.H{"photos": photos})

}

func Update(c *gin.Context) {
	var photos models.Photos
	id := c.Param("id")

	if err := c.ShouldBindJSON(&photos); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&photos).Where("id = ?", id).Updates(&photos).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func Delete(c *gin.Context) {
	var photos models.Photos

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&photos, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})

}
