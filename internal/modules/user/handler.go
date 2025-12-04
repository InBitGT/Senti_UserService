package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"UserService/internal/common"
	"UserService/internal/middleware"

	"github.com/gorilla/mux"
)

type Handler struct {
	svc Service
}

func NewHandler(s Service) *Handler {
	return &Handler{s}
}

// ------------------------- CREATE -------------------------
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	u, err := h.svc.Create(&req)
	if err != nil {
		msg := err.Error()
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, msg, &msg)
		return
	}

	common.CreatedResponse(w, common.SUCCESS_CREATED, u, common.HTTP_CREATED)
}

// ------------------------- UPDATE -------------------------
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_FORMAT, nil)
		return
	}

	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	u, err := h.svc.Update(uint(id), &req)
	if err != nil {
		msg := err.Error()
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, msg, &msg)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, u, common.HTTP_OK)
}

// ------------------------- LIST BY TENANT -------------------------
func (h *Handler) ListByTenant(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(middleware.UserCtxKey).(*middleware.UserClaims)

	users, err := h.svc.FindByTenant(claims.TenantID)
	if err != nil {
		msg := err.Error()
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, msg, &msg)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, users, common.HTTP_OK)
}

// ------------------------- DELETE -------------------------
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_FORMAT, nil)
		return
	}

	if err := h.svc.Delete(uint(id)); err != nil {
		msg := err.Error()
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, msg, &msg)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, "user deleted", common.HTTP_OK)
}
