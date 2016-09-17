package mock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* MockData */
type Product struct {
	ID          int
	Name        string
	Slug        string
	Description string
}

/* We will create our catalog of VR experiences and store them in a slice. */
var products = []Product{
	Product{ID: 1, Name: "Hover Shooters", Slug: "hover-shooters", Description: "Shoot your way to the top on 14 different hoverboards"},
	Product{ID: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description: "Explore the depths of the sea in this one of a kind underwater experience"},
	Product{ID: 3, Name: "Dinosaur Park", Slug: "dinosaur-park", Description: "Go back 65 million years in the past and rIDe a T-Rex"},
	Product{ID: 4, Name: "Cars VR", Slug: "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
	Product{ID: 5, Name: "Robin Hood", Slug: "robin-hood", Description: "Pick up the bow and arrow and master the art of archery"},
	Product{ID: 6, Name: "Real World VR", Slug: "real-world-vr", Description: "Explore the seven wonders of the world in VR"},
}

// MockHandler shows a list of dummy products
func MockHandler(c *gin.Context) {
	var product Product
	slug := c.Param("slug")

	for _, p := range products {
		if p.Slug == slug {
			product = p
		}
	}

	if product.Slug != "" {
		content := gin.H{"content": product}
		c.JSON(http.StatusOK, content)
	} else {
		content := gin.H{"content": "Product Not Found"}
		c.JSON(http.StatusOK, content)
	}
}

func MockProductHandler(c *gin.Context) {
	// Here we are converting the slice of products to json
	content := gin.H{"content": products}
	c.JSON(http.StatusOK, content)
}
