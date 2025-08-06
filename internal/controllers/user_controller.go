package controllers

import (
	"furious/iam-api/internal/models"
	"furious/iam-api/internal/services"
	"furious/iam-api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (ctr *UserController) Search(c *gin.Context) {
	var params models.User

	id, _ := strconv.ParseUint(c.Query("id"), 10, 32)
	params.ID = uint(id)
	params.Email = c.Query("email")
	params.Username = c.Query("username")

	var pag utils.Pagination

	pag.Limit, _ = strconv.Atoi(c.Query("limit"))
	pag.Page, _ = strconv.Atoi(c.Query("page"))

	users, err := ctr.userService.Search(&params, pag)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to search for Users.")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, users)
}

func (ctr *UserController) Persist(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid data.")
		return
	}

	if err := ctr.userService.Persist(&newUser); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create user. "+err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, newUser)
}

func (ctr *UserController) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := ctr.userService.Delete(uint(id)); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to delete user."+err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, nil)
}
