package router

import (
	"net/http"

	"github.com/ReneVallecillo/office.go/domain"
	"github.com/gin-gonic/gin"
)

//ProfileService references the domain struct
type ProfileService struct {
	Profile Profile
}

type Profile interface {
	Load(id string) (*domain.User, error)
}

//ProfileHandler deals with a Profile Request
func (p *ProfileService) ProfileHandler(c *gin.Context) {
	id := c.Param("id")
	user, err := p.Profile.Load(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "Profile Not Found")
	}
	c.JSON(http.StatusOK, user)

}
