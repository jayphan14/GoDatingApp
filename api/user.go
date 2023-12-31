package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
