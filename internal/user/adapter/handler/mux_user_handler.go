package handler

import (
	"encoding/json"
	"go-research/internal/pkg/util"
	"go-research/internal/user/port"
	"go-research/internal/user/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func NewUserHandler(userService service.UserService) port.MuxUserHandler {
	return &UserHandler{
		service: userService,
	}
}

type UserHandler struct {
	service service.UserService
}

func (u *UserHandler) GetUsersByName(w http.ResponseWriter, r *http.Request) {
	pageIdxParam := r.URL.Query().Get("pageIdx")
	pageIdx, err := strconv.Atoi(pageIdxParam)
	if err != nil {
		JSON(w, http.StatusBadRequest, util.Reponse{
			Err: util.ErrorWrongTypePageIdx.Error(),
		})
		return
	}
	pageSizeParam := r.URL.Query().Get("pageSize")
	pageSize, err := strconv.Atoi(pageSizeParam)
	if err != nil {
		JSON(w, http.StatusBadRequest, util.Reponse{
			Err: util.ErrorWrongTypePageSize.Error(),
		})
		return
	}
	name := mux.Vars(r)["name"]
	res, err := u.service.GetUsersByName(r.Context(), name, pageSize, pageIdx)
	if err != nil {
		util.Logger.Errorf("internal server error: %v", err)
		JSON(w, http.StatusInternalServerError, util.Reponse{
			Data: res,
			Err:  err.Error(),
		})
	} else {
		JSON(w, http.StatusOK, util.Reponse{
			Data: res,
		})
	}

}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}
