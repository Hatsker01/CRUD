package v1

import (
	"net/http"

	"github.com/CRUD/pkg/logger"
	"github.com/CRUD/pkg/models"
	"github.com/CRUD/storage/postgres"
	"github.com/gin-gonic/gin"
)

// Create Country ....
// @Summary Create Country
// @Description This API for creating Country
// @Tags country
// @Accept json
// @Produce json
// @Param country body models.CreateCountry true "Country"
// @Success 201 {object} models.CountryResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/country [POST]
func (h *handlerV1) CreateCountry(c *gin.Context) {
	var (
		country models.CreateCountry
	)

	if err := c.ShouldBindJSON(&country); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while binding json", logger.Error(err))
		return
	}

	newCountry, err := postgres.NewCountryRepasitory(h.db).Create(&country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating country", logger.Error(err))
		return
	}

	c.JSON(http.StatusCreated, newCountry)
}

// Update Country ....
// @Summary Update Country
// @Description This API for updating Country
// @Tags country
// @Accept json
// @Produce json
// @Param country body models.UpdateCountry true "Update_Country"
// @Success 200 {object} models.CountryResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/country [PUT]
func (h *handlerV1) UpdateCountry(c *gin.Context) {
	var updateCountry models.UpdateCountry

	if err := c.ShouldBindJSON(&updateCountry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while binding json", logger.Error(err))
		return
	}

	updatedCountry, err := postgres.NewCountryRepasitory(h.db).Update(&updateCountry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating country", logger.Error(err))
		return
	}

	c.JSON(http.StatusAccepted, updatedCountry)
}

// Get Single Country ...
// @Summary Get Single Country
// @Description This API for getting single country by id
// @Tags country
// @Accept json
// @Produce json
// @Param id path string true "country_id"
// @Success 200 {object} models.CountryResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/country/{id} [GET]
func (h *handlerV1) GetSingleCountry(c *gin.Context) {
	countryID := c.Param("id")

	singleCountry, err := postgres.NewCountryRepasitory(h.db).Get(countryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting single country", logger.Error(err))
		return
	}

	c.JSON(http.StatusAccepted, singleCountry)
}

// Delete Country ...
// @Summary Delete Country
// @Description This API for deleting country by id
// @Tags country
// @Accept json
// @Produce json
// @Param id path string true "country_id"
// @Success 200 {object} models.Empty
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/country/{id} [DELETE]
func (h *handlerV1) DeleteCountry(c *gin.Context) {

	countryId := c.Param("id")

	if err := postgres.NewCountryRepasitory(h.db).Delete(countryId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting country", logger.Error(err))
		return
	}
}

// Get User Country ...
// @Summary Get User Countries
// @Description This API for getting user countries
// @Tags country
// @Accept json
// @Produce json
// @Param id path string true "user_id"
// @Success 200 {object} models.Countries
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/country/user/{id} [GET]
func (h *handlerV1) GetUserCountries(c *gin.Context) {

	userID := c.Param("id")

	countries, err := postgres.NewCountryRepasitory(h.db).GetUserCountry(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting user countries", logger.Error(err))
		return
	}

	c.JSON(http.StatusAccepted, models.Countries{Countries: *countries})
}
