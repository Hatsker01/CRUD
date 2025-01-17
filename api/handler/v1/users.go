package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/CRUD/pkg/logger"
	"github.com/CRUD/pkg/models"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	newUser, err := h.storage.UserService().Create(ctx, &user)
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	updatedUser, err := h.storage.UserService().Update(ctx, &updateUser)
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	singleUser, err := h.storage.UserService().Get(ctx, userID)
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	users, err := h.storage.UserService().GetAll(ctx)
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := h.storage.UserService().DeleteUser(ctx, userID)
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	singleUser, err := h.storage.UserService().Get(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting user", logger.Error(err))
		return
	}

	userCountry, err := h.storage.CountryService().GetUserWithCountry(ctx, userID)
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
