package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ussii39/go_rest_api/model"
	"github.com/ussii39/go_rest_api/view"
)

// UserController require *sql.Db to initialize
// This controller hove CRUD methods
type UserController struct {
	Db *sql.DB
}

// GetUsers return All Users
func (tc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	Users, err := model.GetUsers(r.Context(), tc.Db)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("get Users error: %v", err))
		return
	}
	view.RenderUsers(w, Users)
}

// GetUser は path に含まれる uuid に一致する Users テーブルの レコードを返す
func (tc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	UserUUID := params["uuid"]
	exist, err := model.CheckUserExist(r.Context(), tc.Db, UserUUID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("check User exist error: %v", err))
		return
	}
	if !exist {
		view.RenderNotFound(w, "Users", UserUUID)
		return
	}

	User, err := model.GetUser(r.Context(), tc.Db, UserUUID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("get Users error: %v", err))
		return
	}
	view.RenderUser(w, User, http.StatusOK)
}

// CreateUser create new User, and return that User
func (tc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		return
	}
	var User model.User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		view.RenderBadRequest(w, []string{fmt.Sprintf("read post body error: %v", err)})
		return
	}

	err = json.Unmarshal(body, &User)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("json parse error: %v", err))
		return
	}

	insertID, err := model.CreateUser(r.Context(), tc.Db, &User)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("create User error: %v", err))
		return
	}
	createdUser, err := model.GetUserByID(r.Context(), tc.Db, insertID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("get User error: %v", err))
		return
	}
	view.RenderUser(w, createdUser, http.StatusCreated)
}

// PutUser replace specified User, and return that User
func (tc *UserController) PutUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	UserUUID := params["uuid"]

	exist, err := model.CheckUserExist(r.Context(), tc.Db, UserUUID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("check User exist error: %v", err))
		return
	}
	if !exist {
		view.RenderNotFound(w, "Users", UserUUID)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		view.RenderBadRequest(w, []string{fmt.Sprintf("read post body error: %v", err)})
		return
	}

	var User model.User
	err = json.Unmarshal(body, &User)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("json parse error: %v", err))
		return
	}

	err = model.UpdateUser(r.Context(), tc.Db, &User, UserUUID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("create User error: %v", err))
		return
	}
	updatedUser, err := model.GetUser(r.Context(), tc.Db, UserUUID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("get User error: %v", err))
		return
	}
	view.RenderUser(w, updatedUser, http.StatusOK)
}

// DeleteUser delete specified User, and return only status code
func (tc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	params := mux.Vars(r)
	UserUUID := params["uuid"]

	exist, err := model.CheckUserExist(r.Context(), tc.Db, UserUUID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("check User exist error: %v", err))
		return
	}
	if !exist {
		view.RenderNotFound(w, "Users", UserUUID)
		return
	}

	err = model.DeleteUser(r.Context(), tc.Db, UserUUID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("create User error: %v", err))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
