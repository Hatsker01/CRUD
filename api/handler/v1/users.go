package v1

import (
	"net/http"

	"github.com/CRUD/pkg/logger"
	"github.com/CRUD/pkg/models"
	"github.com/CRUD/storage/postgres"
	"github.com/gin-gonic/gin"
)

// This API for creating user
// @Summary Create user
// @Description This API for creating user
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.UserCreate true "Create_user"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/user [POST]
func (h *handlerV1) CreateUser(c *gin.Context) {
	var user models.UserCreate

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while binding json", logger.Error(err))
		return
	}

	newUser, err := postgres.NewUsersRepasitory(h.db).Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating user", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, newUser)
}

// This API for updating user
// @Summary Update user
// @Description This API for updating user
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.UpdateUser true "Update_user"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/user [PUT]
func (h *handlerV1) UpdateUser(c *gin.Context) {
	var updateUser models.UpdateUser

	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating user", logger.Error(err))
		return
	}

	updatedUser, err := postgres.NewUsersRepasitory(h.db).Update(&updateUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating new user", logger.Error(err))
		return
	}

	c.JSON(http.StatusAccepted, updatedUser)
}

// This API for getting single user by ID
// @Summary Get Single user
// @Description This API for getting single user by id
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "user_id"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/user/{id} [GET]
func (h *handlerV1) GetSingleUser(c *gin.Context) {
	userID := c.Param("id")

	singleUser, err := postgres.NewUsersRepasitory(h.db).Get(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting user by id", logger.Error(err))
		return
	}

	c.JSON(http.StatusAccepted, singleUser)
}

// Get All Users ...
// @Summary Get all users
// @Description This API for getting all users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} models.Users
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/user [GET]
func (h *handlerV1) GetAllUsers(c *gin.Context) {

	users, err := postgres.NewUsersRepasitory(h.db).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting all users", logger.Error(err))
		return
	}

	c.JSON(http.StatusAccepted, models.Users{Users: *users})
}

// Delete User ...
// @Summary Delete User
// @Description This API for deleting user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "user_id"
// @Success 200 {object} models.Empty
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/user/{id} [DELETE]
func (h *handlerV1) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	err := postgres.NewUsersRepasitory(h.db).DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting user", logger.Error(err))
		return
	}

	c.JSON(http.StatusAccepted, models.Empty{})
}

// Get User With Country ...
// @Summary Get User with Country
// @Description This API for getting user with countries
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "user_id"
// @Success 200 {object} models.UserWithCountry
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/user/country/{id} [GET]
func (h *handlerV1) GetUserWithCountry(c *gin.Context) {
	var (
		userID = c.Param("id")
	)

	singleUser, err := postgres.NewUsersRepasitory(h.db).Get(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting user", logger.Error(err))
		return
	}

	userCountry, err := postgres.NewCountryRepasitory(h.db).GetUserWithCountry(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting country", logger.Error(err))
		return
	}

	c.JSON(http.StatusAccepted, models.UserWithCountry{
		Name:    singleUser.Name,
		Country: *userCountry,
	})
}
