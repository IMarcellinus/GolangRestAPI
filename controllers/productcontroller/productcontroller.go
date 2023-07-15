package productcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeypc/go-crud/config"
	"github.com/jeypc/go-crud/models"
	"gorm.io/gorm"
)

type ProductResponse struct {
	ID          int    `json:"id"`
	NamaProduct string `json:"nama_product"`
	Deskripsi   string `json:"deskripsi"`
}

func Index(c *gin.Context) {
	var products []models.Product
	var productResponses []ProductResponse

	result := config.DB.Model(&models.Product{}).Select("id, nama_product, deskripsi").Find(&products)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	for _, p := range products {
		productResponses = append(productResponses, ProductResponse{
			ID:          p.ID,
			NamaProduct: p.NamaProduct,
			Deskripsi:   p.Deskripsi,
		})
	}

	c.JSON(http.StatusOK, gin.H{"products": productResponses})

}
func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}
func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	config.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}
func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if config.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil update"})
}
func Delete(c *gin.Context) {
	var product models.Product

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if config.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat menghapus product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil dihapus"})
}
