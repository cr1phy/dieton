package router

import (
	"net/http"
	"strconv"

	"github.com/cr1phy/dieton-backend/pkg/models"
	"github.com/gin-gonic/gin"
)

type RouterHandler struct {
	repo models.DBRepo
}

func NewRouterHandler(r models.DBRepo) *RouterHandler {
	return &RouterHandler{r}
}

func (h *RouterHandler) GetProductInfo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if p, err := h.repo.FindProduct(id, c.Request.Context()); err != nil {
		c.JSON(http.StatusOK, p)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Not found."})
}
