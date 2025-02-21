package business

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
    service BusinessService
}

func NewHandler(service BusinessService) *Handler {
    return &Handler{
        service: service,
    }
}

// CreateBusiness godoc
// @Summary Create a new business
// @Description Create a new business with the given details
// @Tags businesses
// @Accept json
// @Produce json
// @Param business body CreateBusinessRequest true "Business details"
// @Success 201 {object} BusinessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /businesses [post]
func (h *Handler) CreateBusiness(c *gin.Context) {
    var req CreateBusinessRequest  // This comes from your dto.go
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    business := &Business{
        ID:          uuid.New(),
        Name:        req.Name,
        Description: req.Description,
        Category:    req.Category,
        Location:    req.Location,
        Rating:      0, // Default rating for new business
        ContactInfo: req.ContactInfo,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }

    if err := h.service.CreateBusiness(c.Request.Context(), business); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, business.ToResponse())
}

