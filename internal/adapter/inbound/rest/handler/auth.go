package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imanudd/inventorybooksvc/internal/adapter/inbound/rest/handler/helper"
	"github.com/imanudd/inventorybooksvc/internal/core/domain"
)

// Login handler
// @Summary login user
// @Description login user
// @Tags auth
// @Accept json
// @Produce json
// @Param input body domain.LoginRequest true "login data"
// @Success 200 {object} helper.JSONResponse
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Router /inventorysvc/auth/login [POST]
func (h *Handler) Login(c *gin.Context) {
	var req *domain.LoginRequest

	if err := c.ShouldBind(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, "error bad request")
		return
	}

	resp, err := h.service.GetAuthService().Login(c, req)
	if err != nil {
		helper.InternalError(c, err)
		return
	}

	helper.Success(c, http.StatusOK, resp)
}

// Register handler
// @Summary register user
// @Description register user
// @Tags auth
// @Accept json
// @Produce json
// @Param input body domain.RegisterRequest true "register data"
// @Success 200 {object} helper.JSONResponse
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Router /inventorysvc/auth/register [POST]
func (h *Handler) Register(c *gin.Context) {
	var req *domain.RegisterRequest

	if err := c.ShouldBind(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, "error bad request")
		return
	}

	err := h.service.GetAuthService().Register(c, req)
	if err != nil {
		helper.InternalError(c, err)
		return
	}

	helper.Success(c, http.StatusOK)
}
