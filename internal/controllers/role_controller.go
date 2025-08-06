package controllers

import (
	"furious/iam-api/internal/models"
	"furious/iam-api/internal/services"
	"furious/iam-api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	service *services.RoleService
}

func NewRoleController(service *services.RoleService) *RoleController {
	return &RoleController{
		service: service,
	}
}

func (ctr *RoleController) Search(c *gin.Context) {
	var params models.Role

	id, _ := strconv.ParseUint(c.Query("id"), 10, 32)
	params.ID = uint(id)
	params.Name = c.Query("name")

	var pag utils.Pagination
	pag.Page, _ = strconv.Atoi(c.Query("page"))
	pag.Limit, _ = strconv.Atoi(c.Query("limit"))

	roles, err := ctr.service.Search(&params, pag)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to search for users.")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, roles)
}

func (ctr *RoleController) Persist(c *gin.Context) {
	var role models.Role

	if err := c.ShouldBindJSON(&role); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid data.")
		return
	}

	if err := ctr.service.Persist(&role); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, role)
}

func (ctr *RoleController) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := ctr.service.Delete(uint(id)); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, nil)
}
