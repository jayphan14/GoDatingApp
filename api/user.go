package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/jayphan14/GoDatingApp/sqlc"
	"github.com/jayphan14/GoDatingApp/util"
)

type createUserRequest struct {
	Username    string   `json:"username" binding:"required"`
	Email       string   `json:"email" binding:"required"`
	Password    string   `json:"password" binding:"required"`
	Gender      string   `json:"gender" binding:"required,oneof=M F"`
	University  string   `json:"university" binding:"required"`
	Picture     []byte   `json:"picture" binding:"required"`
	Bio         string   `json:"bio" binding:"required"`
	BioPictures []string `json:"bio_pictures" binding:"required"`
}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req createUserRequest
	// check if the Request has all the needed params
	if errBinding := ctx.ShouldBindJSON(&req); errBinding != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(errBinding))
		return
	}

	arg := db.CreateUserParams{
		Username:    req.Username,
		Email:       req.Email,
		Password:    req.Password,
		Gender:      req.Gender,
		University:  req.University,
		Picture:     req.Picture,
		Bio:         req.Bio,
		BioPictures: req.BioPictures,
	}

	// Check if user exist:
	existingUser, errGettingUserWithEmail := server.store.GetUserByEmail(ctx, req.Email)
	if errGettingUserWithEmail != nil {
		if existingUser.Email != "" {
			ctx.JSON(http.StatusNotFound, util.ErrorResponseString("User already exist"))
			return
		}
	} else {
		ctx.JSON(http.StatusNotFound, util.ErrorResponseString("User already exist"))
		return
	}

	newUser, errCreatingUser := server.store.CreateUser(ctx, arg)

	if errCreatingUser != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(errCreatingUser))
		return
	}

	ctx.JSON(http.StatusOK, newUser)
}

type getUserByEmailRequest struct {
	Email string `uri:"email" binding:"required"`
}

func (server *Server) GetUserByEmail(ctx *gin.Context) {
	var req getUserByEmailRequest

	// check if the Request has all the needed params
	if errBinding := ctx.ShouldBindUri(&req); errBinding != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(errBinding))
		return
	}

	existingUser, errGettingUserWithEmail := server.store.GetUserByEmail(ctx, req.Email)
	if errGettingUserWithEmail != nil {
		// No user exist
		if errGettingUserWithEmail == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, util.ErrorResponse(errGettingUserWithEmail))
			return
		}
		ctx.JSON(http.StatusNotFound, util.ErrorResponse(errGettingUserWithEmail))
	}
	ctx.JSON(http.StatusOK, existingUser)
}

type getUserByIdRequest struct {
	ID pgtype.UUID `uri:"id" binding:"required"`
}

func (server *Server) GetUserById(ctx *gin.Context) {
	var req getUserByIdRequest

	// check if the Request has all the needed params
	if errBinding := ctx.ShouldBindUri(&req); errBinding != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(errBinding))
		return
	}

	existingUser, errGettingUserWithID := server.store.GetUser(ctx, req.ID)
	if errGettingUserWithID != nil {
		// No user exist
		if errGettingUserWithID == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, util.ErrorResponse(errGettingUserWithID))
			return
		}
		ctx.JSON(http.StatusNotFound, util.ErrorResponse(errGettingUserWithID))
	}
	ctx.JSON(http.StatusOK, existingUser)
}

type updateUserRequest struct {
	ID          pgtype.UUID `json:"id" binding:"required"`
	Username    string      `json:"username" binding:"required"`
	Email       string      `json:"email" binding:"required"`
	Password    string      `json:"password" binding:"required"`
	Gender      string      `json:"gender" binding:"required,oneof=M F"`
	University  string      `json:"university" binding:"required"`
	Picture     []byte      `json:"picture" binding:"required"`
	Bio         string      `json:"bio" binding:"required"`
	BioPictures []string    `json:"bio_pictures" binding:"required"`
}

func (server *Server) UpdateUser(ctx *gin.Context) {
	var req updateUserRequest
	// check if the Request has all the needed params
	if errBinding := ctx.ShouldBindJSON(&req); errBinding != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(errBinding))
		return
	}

	arg := db.UpdateUserParams{
		ID:          req.ID,
		Username:    req.Username,
		Email:       req.Email,
		Password:    req.Password,
		Gender:      req.Gender,
		University:  req.University,
		Picture:     req.Picture,
		Bio:         req.Bio,
		BioPictures: req.BioPictures,
	}

	updatedUser, errUpdatingUser := server.store.UpdateUser(ctx, arg)
	if errUpdatingUser != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(errUpdatingUser))
	}
	ctx.JSON(http.StatusOK, updatedUser)
}

type deleteUserRequest struct {
	ID pgtype.UUID `uri:"id" binding:"required"`
}

func (server *Server) DeleteUser(ctx *gin.Context) {
	var req deleteUserRequest

	// check if the Request has all the needed params
	if errBinding := ctx.ShouldBindUri(&req); errBinding != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(errBinding))
		return
	}

	errorDeletingUser := server.store.DeleteUser(ctx, req.ID)
	if errorDeletingUser != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(errorDeletingUser))
	}
	ctx.JSON(http.StatusOK, nil)
}

type listUserRequest struct {
	Limit  int64 `json:"limit" binding:"required"`
	Offset int64 `json:"offset"`
}

func (server *Server) ListUsers(ctx *gin.Context) {
	var req listUserRequest

	// check if the Request has all the needed params
	if errBinding := ctx.ShouldBindJSON(&req); errBinding != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(errBinding))
		return
	}
	arg := db.ListUsersParams{
		Limit:  int32(req.Limit),
		Offset: int32(req.Offset),
	}

	users, errorGettingUsers := server.store.ListUsers(ctx, arg)
	if errorGettingUsers != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(errorGettingUsers))
	}
	ctx.JSON(http.StatusOK, users)
}
