package user

import (
	"api-openlive/common"
	bussiness "api-openlive/module/user/business"
	usermodel "api-openlive/module/user/model"
	"api-openlive/utils"
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"github.com/spf13/cast"
	"github.com/twinj/uuid"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type UserHandler struct {
	Client *bussiness.UserController
}

// CheckExistUser Create general godoc
// @Summary check user exist by wallet address
// @Tags User
// @Accept json
// @Produce json
// @Param user body usermodel.CheckExistUserRequest true "login request"
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/check-exist [post]
func (u UserHandler) CheckExistUser() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq usermodel.CheckExistUserRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		res, err := u.Client.CheckExistUser(rq.WalletAddress)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Check user successfully", map[string]interface{}{
			"is_exist": res,
		}))
	}
}

// LoginUser Create general godoc
// @Summary Login
// @Tags User
// @Accept json
// @Produce json
// @Param user body usermodel.UserLoginRequest true "login request"
// @Success 200 {object} common.Response{data=usermodel.Credential}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/login [post]
func (u UserHandler) LoginUser() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq usermodel.UserLoginRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		res, err := u.Client.LoginUser(rq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Login successfully", res))
	}
}

// LoginCMSUser Create general godoc
// @Summary Login CMS
// @Tags CMS
// @Accept json
// @Produce json
// @Param user body usermodel.CMSLoginRequest true "login cms request"
// @Success 200 {object} common.Response{data=usermodel.Credential}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/login [post]
func (u UserHandler) LoginCMSUser() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq usermodel.CMSLoginRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		res, err := u.Client.CMSUserLogin(rq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Login successfully", res))
	}
}

// RegisterCMSUser Create general godoc
// @Summary register CMS user
// @Tags CMS
// @Accept json
// @Produce json
// @Param user body usermodel.CMSUser true "login cms request"
// @Success 200 {object} common.Response{data=usermodel.CMSUser}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/register [post]
func (u UserHandler) RegisterCMSUser() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq usermodel.CMSUser
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		res, err := u.Client.CMSUserRegister(rq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Login successfully", res))
	}
}

// CreateNounce Create general godoc
// @Summary Nonce
// @Tags User
// @Accept json
// @Produce json
// @Param user body usermodel.NonceRequest true "nonce request"
// @Success 200 {object} common.Response{data=usermodel.NounceResponse}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/create-nonce [post]
func (u UserHandler) CreateNounce() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var user usermodel.NonceRequest
		if res := utils.CommonHandle(&user, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}

		res, err := u.Client.CreateNounce(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "create nounce success", res))

	}
}

// SendMailOTP Create general godoc
// @Summary Send Email OTP
// @Tags User
// @Accept json
// @Produce json
// @Param user body usermodel.OTPRequest true "OTP request"
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/send-otp [post]
func (u UserHandler) SendMailOTP() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var otp usermodel.OTPRequest
		if res := utils.CommonHandle(&otp, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}

		err := u.Client.MailOtp(otp)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Send mail successfully", nil))

	}
}

// SendMailOTPReset Create general godoc
// @Summary OTP Reset Password
// @Tags User
// @Accept json
// @Produce json
// @Param user body usermodel.OTPRequestReset true "reset password request"
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/send-otp-reset-password [post]
func (u UserHandler) SendMailOTPReset() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var otp usermodel.OTPRequestReset
		if res := utils.CommonHandle(&otp, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}

		err := u.Client.MailOtpReset(otp)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Send mail successfully", nil))

	}
}

// Register Create general godoc
// @Summary Register
// @Tags User
// @Accept json
// @Produce json
// @Param user body usermodel.CreateUserRequest true "nonce request"
// @Success 200 {object} common.Response{data=usermodel.UserResponse}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/register [post]
func (u UserHandler) Register() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var userReq usermodel.CreateUserRequest
		if res := utils.CommonHandle(&userReq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}

		user, err := u.Client.CreateUser(userReq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Register user successfully", user))

	}
}

// RefreshToken Create general godoc
// @Summary Refresh Token
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{data=usermodel.Credential}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/refresh-token [post]
func (u UserHandler) RefreshToken() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.GetInt("id")
		res, err := u.Client.RefreshToken(int64(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Register user successfully", res))

	}
}

// DetailUser Create User godoc
// @Param Authorization header string false "Insert your access token"
// @Security BearerAuth
// @Summary get detail user
// @Tags User
// @Param user_id query int false "user_id"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=usermodel.User}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/detail [get]
func (u UserHandler) DetailUser() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.GetInt("id")
		var filter = usermodel.UserFilter{CurrentId: int64(id)}
		targetUserId, _ := strconv.Atoi(ctx.Query("user_id"))
		if targetUserId != 0 {
			id = targetUserId
		}
		filter.Id = int64(id)
		res, err := u.Client.DetailUser(filter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// GetUserInfo Create User godoc
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Summary get detail user
// @Tags User
// @Param wallet_address path string true "wallet_address"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=usermodel.User}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/get-user/{wallet_address} [get]
func (u UserHandler) DetailUserTransfer() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var filter usermodel.UserFilter
		filter.WalletAddress = ctx.Param("wallet_address")
		res, err := u.Client.DetailUser(filter)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// DetailOtherUser Create User godoc
// @Summary get detail user
// @Param id path int true "category id"
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=usermodel.User}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/{id} [get]
func (u UserHandler) DetailOtherUser() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := cast.ToInt64(ctx.Param("id"))
		res, err := u.Client.DetailUser(usermodel.UserFilter{Id: id})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// ListFollowers Create User godoc
// @Summary get detail user
// @Param id path int true "category id"
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=usermodel.User}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/{id}/followers [get]
func (u UserHandler) ListFollowers() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := cast.ToInt64(ctx.Param("id"))
		currentId := ctx.GetInt("id")
		res, err := u.Client.ListFollowers(id, int64(currentId))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// ListFollowing Create User godoc
// @Summary get detail user
// @Param id path int true "category id"
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=usermodel.User}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/{id}/following [get]
func (u UserHandler) ListFollowing() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := cast.ToInt64(ctx.Param("id"))
		currentId := ctx.GetInt("id")
		res, err := u.Client.ListFollowing(id, int64(currentId))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// ListFollowersView Create User godoc
// @Summary get detail user
// @Param id path int true "category id"
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=usermodel.User}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/{id}/followers-view [get]
func (u UserHandler) ListFollowersView() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := cast.ToInt64(ctx.Param("id"))
		res, err := u.Client.ListFollowers(id, 0)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// ListFollowingView Create User godoc
// @Summary get detail user
// @Param id path int true "category id"
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=usermodel.User}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/{id}/following-view [get]
func (u UserHandler) ListFollowingView() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := cast.ToInt64(ctx.Param("id"))
		res, err := u.Client.ListFollowing(id, 0)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// ChangeProfile Change profile
// @Param Authorization header string true "Insert your access token"
// @Param user body usermodel.UpdateUserRequest true "nonce request"
// @Security BearerAuth
// @Summary get detail user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=usermodel.User}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/change-profile [patch]
func (u UserHandler) ChangeProfile() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req usermodel.UpdateUserRequest
		if res := utils.CommonHandle(&req, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		req.Id = cast.ToInt64(ctx.GetInt("id"))
		scheme := "https://"
		/*if ctx.Request.TLS != nil {
			scheme = "https://"
		}*/
		options := map[string]interface{}{
			"host": scheme + ctx.Request.Host + "/",
		}
		res, err := u.Client.ChangeProfile(&req, options)
		if err != nil {
			ctx.JSON(http.StatusForbidden, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// SearchHistory Change profile
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Summary get detail user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/search-history [get]
func (u UserHandler) SearchHistory() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := cast.ToInt64(ctx.GetInt("id"))
		res, err := u.Client.SearchHistory(id)
		if err != nil {
			ctx.JSON(http.StatusForbidden, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get history user search successfully", res))

	}
}

// WatchList Change profile
// @Param Authorization header string true "Insert your access token"
// @Param start_date query string false "start_date"
// @Param end_date query string false "end_date"
// @Security BearerAuth
// @Summary get detail user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/watch-list [get]
func (u UserHandler) WatchList() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.PagingWithTime
		_ = ctx.ShouldBind(&paging)
		if !paging.Process() {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid request", nil))
			return
		}
		paging.ProcessLastTime()

		id := cast.ToInt64(ctx.GetInt("id"))
		res, err := u.Client.WatchList(id, &paging)
		if err != nil {
			ctx.JSON(http.StatusForbidden, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get history user search successfully", res))

	}
}

// SearchCollectionCms Change profile
// @Summary searching
// @Tags Common
// @Param value query string false "value to find"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /common/cms/search/collection [get]
func (u UserHandler) SearchCollectionCms() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req usermodel.SearchRequest
		err := ctx.ShouldBind(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		res, err := u.Client.SearchCollection(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// Search Change profile
// @Param user body usermodel.SearchRequest true "nonce request"
// @Summary searching
// @Tags Common
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /common/search [post]
func (u UserHandler) Search() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req usermodel.SearchRequest
		if res := utils.CommonHandle(&req, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		var paging common.Paging
		/*if errPage := ctx.ShouldBind(&paging); errPage != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", errPage.Error()))
			return
		}*/
		paging.Process()
		res, err := u.Client.Search(&req, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error paging", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// ExploreNft explore NFT
// @Summary searching
// @Tags Market Place
// @Param Authorization header string false "Insert your access token"
// @Security BearerAuth
// @Param user_id query int false "user id"
// @Param value query string false "value to find"
// @Param category_ids query string false "many category id"
// @Param price_max query number false "price max"
// @Param price_min query number false "price min"
// @Param status query string false "status"
// @Param is_join_event query int false "is join event ?"
// @Param order_by query int false "order_by"
// @Param collection_id query int false "collection_id"
// @Param sell_type query int false "sell_type"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/explore-nft [get]
func (u UserHandler) ExploreNft() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req usermodel.SearchRequest
		err := ctx.ShouldBind(&req)
		req.IsJustGetNft = 1
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		var paging common.Paging
		if errPage := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", errPage.Error()))
			return
		}
		userId := ctx.GetInt("id")
		if userId > 0 {
			req.UserId = int64(userId)
		}
		paging.Process()
		res, err := u.Client.Search(&req, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get list nft successfully", res, paging))

	}
}

// Upload File
// @Param file formData file true "file upload"
// @Param type path string true "file upload"
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Summary upload file
// @Tags Common
// @Accept mpfd
// @Produce mpfd
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /common/upload/{type} [post]
func (u UserHandler) Upload() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		typeUpload := ctx.Param("type")
		// Single file
		// c.Request.Host/storage/filename
		fileUrl, err := UploadFile(ctx, "file", typeUpload)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid Request", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(
			0,
			"Upload file success",
			map[string]string{
				"file": fileUrl,
			},
		),
		)

	}
}

// UploadImage File
// @Param file formData file true "file upload"
// @Param is_image query int true "file upload"
// @Param cover formData file false "file upload"
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Summary upload file
// @Tags Common
// @Accept mpfd
// @Produce mpfd
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /common/item/upload [post]
func (u UserHandler) UploadImage() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		isImage := cast.ToInt64(ctx.Query("is_image"))
		fileUrl, err := UploadFile(ctx, "file", common.ITEM)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid Request", err.Error()))
			return
		}

		coverUrl := ""
		if isImage != 1 {
			coverUrl, err = UploadFile(ctx, "cover", common.ITEM_IMAGE)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid Request", err.Error()))
				return
			}
		}

		ctx.JSON(http.StatusOK, common.NewResponse(
			0,
			"Upload file success",
			map[string]string{
				"file":  fileUrl,
				"cover": coverUrl,
			},
		),
		)

	}
}

// Upload File
// @Param file formData file true "file upload"
// @Param type path string true "file upload"
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Summary upload file
// @Tags Common
// @Accept mpfd
// @Produce mpfd
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /common/cms/upload/{type} [post]
func (u UserHandler) UploadCms() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		typeUpload := ctx.Param("type")
		// Single file
		// c.Request.Host/storage/filename
		fileUrl, err := UploadFile(ctx, "file", typeUpload)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid Request", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(
			0,
			"Upload file success",
			map[string]string{
				"file": fileUrl,
			},
		),
		)

	}
}

// UploadImage File
// @Param file formData file true "file upload"
// @Param is_image query int true "file upload"
// @Param cover formData file false "file upload"
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Summary upload file
// @Tags Common
// @Accept mpfd
// @Produce mpfd
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /common/cms/item/upload [post]
func (u UserHandler) UploadImageCms() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		isImage := cast.ToInt64(ctx.Query("is_image"))
		fileUrl, err := UploadFile(ctx, "file", common.ITEM)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid Request", err.Error()))
			return
		}

		coverUrl := ""
		if isImage != 1 {
			coverUrl, err = UploadFile(ctx, "cover", common.ITEM_IMAGE)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid Request", err.Error()))
				return
			}
		}

		ctx.JSON(http.StatusOK, common.NewResponse(
			0,
			"Upload file success",
			map[string]string{
				"file":  fileUrl,
				"cover": coverUrl,
			},
		),
		)

	}
}

// ------CMS---------

// ChangeStatus Uơdate profile CMS
// @Param user body usermodel.UpdateUserStatus true "nonce request"
// @Summary change status user (vip/basic)
// @Tags CMS
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=usermodel.User}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/user/update-status [post]
func (u UserHandler) ChangeStatus() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req usermodel.UpdateUserStatus
		if res := utils.CommonHandle(&req, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}

		if req.Status != common.VIP_USER && req.Status != common.BASIC_USER {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Internal Server Error", common.ErrStatusNotMatch))
			return
		}

		res, err := u.Client.ChangeStatus(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Change status successfully", res))

	}
}

// TopSeller Top Seller
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param start_date query string false "start_date"
// @Param end_date query string false "end_date"
// @Summary get top seller
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=[]usermodel.UserDetail}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/top-seller [get]
func (u UserHandler) TopSeller() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.PagingWithTime
		_ = ctx.ShouldBind(&paging)
		if !paging.Process() {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid request", nil))
			return
		}
		paging.ProcessLastTime()

		res, err := u.Client.TopSeller(&paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get data successfully", res, paging.Paging))

	}
}

// TopCollection Top Seller
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param start_date query string false "start_date"
// @Param end_date query string false "end_date"
// @Summary get top collection
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=[]usermodel.UserDetail}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/top-collection [get]
func (u UserHandler) TopCollection() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.PagingWithTime
		_ = ctx.ShouldBind(&paging)
		if !paging.Process() {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid request", nil))
			return
		}
		paging.ProcessLastTime()

		res, err := u.Client.TopCollection(&paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get data successfully", res, paging.Paging))

	}
}

// TopCreator Top creator
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param start_date query int false "start_date"
// @Param end_date query int false "end_date"
// @Summary get top seller
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/top-creator [get]
func (u UserHandler) TopCreator() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.PagingWithTime
		_ = ctx.ShouldBind(&paging)
		if !paging.Process() {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid request", nil))
			return
		}
		paging.ProcessLastTime()

		res, err := u.Client.TopCreator(&paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get data successfully", map[string]interface{}{
			"total":  paging.Total,
			"limit":  paging.Limit,
			"page":   paging.Page,
			"result": res,
		}))
	}
}

func UploadFile(ctx *gin.Context, typeFile string, typeUpload string) (string, error) {
	file, err := ctx.FormFile(typeFile)
	if err != nil {
		return "", errors.New("no file is received")
	}
	if !utils.InStringSlice(utils.AllowedTypes[typeUpload], file.Header["Content-Type"][0]) {
		return "", errors.New("file not support")
	}
	fileName := strings.Split(file.Filename, ".")
	newFileName := uuid.NewV4().String() + "." + fileName[len(fileName)-1]
	now := time.Now().Format("2006-01-02")
	dst := "tmp/" + now

	// create file if not exist
	if _, errFound := os.Stat(dst); os.IsNotExist(errFound) {
		if err := os.Mkdir(dst, os.ModePerm); err != nil {
			return "", err
		}
	}

	// Upload the file to specific dst.
	err = ctx.SaveUploadedFile(file, dst+"/"+newFileName)
	if err != nil {
		return "", err
	}

	scheme := "https://"
	/*if ctx.Request.TLS != nil {
		scheme = "https://"
	}*/

	return scheme + ctx.Request.Host + "/" + dst + "/" + newFileName, nil

}

// ListNotification list notification
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param status query int false "status"
// @Param type query string false "type"
// @Summary get list notification of user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /notification/list [get]
func (u UserHandler) ListNotification() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req usermodel.NotificationFilter
		_ = ctx.ShouldBind(&req)
		userId := ctx.GetInt("id")
		req.UserId = cast.ToInt64(userId)

		//var paging common.Paging
		//_ = ctx.ShouldBind(&paging)
		//paging.Process()

		res, err := u.Client.ListNotification(&req, nil)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get data successfully", res))

	}
}

// NotificationSetting list notification
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Summary get list notification of user
// @Tags Notification
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /notification/setting [get]
func (u UserHandler) NotificationSetting() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req usermodel.NotificationSetting
		userId := ctx.GetInt("id")
		req.UserId = cast.ToInt64(userId)

		res, err := u.Client.NotificationSetting(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get data successfully", res))

	}
}

// UpdateStatusNotification Top creator
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param user body usermodel.UpdateNotificationStatus true "nonce request"
// @Summary update status notification
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /notification/change-status [post]
func (u UserHandler) UpdateStatusNotification() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req usermodel.UpdateNotificationStatus
		if res := utils.CommonHandle(&req, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		id := cast.ToInt64(ctx.GetInt("id"))
		req.UserId = id

		err := u.Client.ChangeNotificationStatus(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "update status successfully", nil))

	}
}

// ReadAll Top creator
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Summary update status notification
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /notification/read-all [post]
func (u UserHandler) ReadAll() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req usermodel.UpdateNotificationStatus
		id := cast.ToInt64(ctx.GetInt("id"))
		req.UserId = id

		err := u.Client.ReadAll(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "update status successfully", nil))

	}
}

// UpdateSetting Top creator
// @Param Authorization header string true "Insert your access token"
// @Param id body usermodel.NotificationSetting true "setting request"
// @Security BearerAuth
// @Tags Notification
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /notification/update-setting [post]
func (u UserHandler) UpdateSetting() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req usermodel.NotificationSetting
		_ = ctx.ShouldBindJSON(&req)
		id := cast.ToInt64(ctx.GetInt("id"))
		req.UserId = id

		err := u.Client.UpdateNotiSetting(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "update status successfully", nil))

	}
}

// FollowingUser item
// @Summary vote item
// @Tags User
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param id body usermodel.FollowingRequest true "following request"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/follow [post]
func (u UserHandler) FollowingUser() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req usermodel.FollowingRequest
		_ = ctx.ShouldBindJSON(&req)

		userId := ctx.GetInt("id")
		if cast.ToInt64(userId) == req.FollowerId {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", "cant follow yourself"))
			return
		}
		//req.FollowerId = cast.ToInt64(userId)

		err := u.Client.FollowUser(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "success", nil))

	}
}

// UnfollowUser item
// @Summary vote item
// @Tags User
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param id body usermodel.FollowingRequest true "following request"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/unfollow [post]
func (u UserHandler) UnfollowUser() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req usermodel.FollowingRequest
		_ = ctx.ShouldBindJSON(&req)

		userId := ctx.GetInt("id")
		if cast.ToInt64(userId) == req.FollowerId {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", "cant unfollow yourself"))
			return
		}
		req.UserId = cast.ToInt64(userId)

		err := u.Client.UnfollowUser(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "success", nil))

	}
}

// ProfileStatistic Top creator
// @Param Authorization header string false "Insert your access token"
// @Summary get profile statistic
// @Tags User
// @Param user_id query int false "user_id"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/profile-statistic [get]
func (u UserHandler) ProfileStatistic() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		userId := ctx.GetInt("id")
		targetUserId, _ := strconv.Atoi(ctx.Query("user_id"))
		if targetUserId != 0 {
			userId = targetUserId
		}
		res, err := u.Client.ProfileStatistic(int64(userId))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get data successfully", res))
	}
}

// UserActivities Create User godoc
// @Summary get detail user
// @Param Authorization header string false "Insert your access token"
// @Security BearerAuth
// @Param type query string false "type activities"
// @Param time query int false "time filter"
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=usermodel.User}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/activities [get]
func (u UserHandler) UserActivities() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		filter := usermodel.ActivitiesFilter{}
		filter.UserId = int64(ctx.GetInt("id"))
		filter.ActivityType = ctx.Query("type")
		filter.Time = cast.ToInt64(ctx.Query("time"))
		res, err := u.Client.Activities(filter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// Activities Create User godoc
// @Summary get detail user
// @Param type query string false "type activities"
// @Param time query int false "time filter"
// @Param user_id query int false "filter user_id"
// @Tags Common
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=usermodel.User}
// @Failure 400,401,500 {object} common.Response
// @Router /common/activities [get]
func (u UserHandler) Activities() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		filter := usermodel.ActivitiesFilter{}
		filter.UserId = cast.ToInt64(ctx.Query("user_id"))
		filter.ActivityType = ctx.Query("type")
		filter.Time = cast.ToInt64(ctx.Query("time"))
		res, err := u.Client.Activities(filter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// Can Add Ref Create User godoc
// @Summary get detail user
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/can-add-ref [get]
func (u UserHandler) CanAddRef() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.GetInt("id")
		var filter = usermodel.UserFilter{CurrentId: int64(id)}
		filter.Id = int64(id)
		res, err := u.Client.CanAddRefProcess(filter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))
	}
}

func GetImage(c *gin.Context, file string) {
	// Check if file exists and/or if we have permission to access it
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	contentType := http.DetectContentType(fileBytes)
	resizeType, errResize := c.GetQuery("resize")
	if (contentType == "image/jpg" || contentType == "image/jpeg" || contentType == "image/png") && errResize {
		// decode jpeg into image.Image
		var img image.Image
		var errDecode error

		switch contentType {
		case "image/png":
			{
				img, errDecode = png.Decode(bytes.NewReader(fileBytes))
				if errDecode != nil {
					log.Println("unable to write image. ", err.Error())
				}
				break
			}
		default:
			{
				img, _, errDecode = image.Decode(bytes.NewReader(fileBytes))
				if errDecode != nil {
					log.Println("unable to write image. ", err.Error())
				}
				break
			}
		}

		// resize to width 1000 using Lanczos resampling
		// and preserve aspect ratio
		var m image.Image
		buf := new(bytes.Buffer)
		switch resizeType {
		case "big":
			{
				m = resize.Resize(1024, 0, img, resize.Lanczos2)
				err = jpeg.Encode(buf, m, nil)
				if err != nil {
					log.Println("unable to write image. ", err.Error())
				}
				fileBytes = buf.Bytes()
				break
			}
		case "small":
			{
				m = resize.Resize(256, 0, img, resize.Lanczos2)
				err = jpeg.Encode(buf, m, nil)
				if err != nil {
					log.Println("unable to write image. ", err.Error())
				}
				fileBytes = buf.Bytes()
				break
			}
		default:
			{
				m = resize.Resize(512, 0, img, resize.Lanczos2)
				err = jpeg.Encode(buf, m, nil)
				if err != nil {
					log.Println("unable to write image. ", err.Error())
				}
				fileBytes = buf.Bytes()
				break
			}
		}
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Set("Content-Type", contentType)
	c.Writer.Header().Set("Content-Length", strconv.Itoa(len(fileBytes)))
	if _, err := c.Writer.Write(fileBytes); err != nil {
		log.Println("unable to write image.")
	}
}

func (u UserHandler) GetImageStatic(c *gin.Context) {
	file := "./storage" + c.Param("filepath")
	GetImage(c, file)
}

func (u UserHandler) GetTmpImageStatic(c *gin.Context) {
	file := "./tmp" + c.Param("filepath")
	GetImage(c, file)
}

// Ref network list godoc
// @Summary get detail user
// @Tags Market Place
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token"
// @Param time query int false "time"
// @Param sort_user_name query int false "sort_user_name"
// @Param sort_event_volume query int false "sort_event_volume"
// @Param sort_market_volume query int false "sort_market_volume"
// @Security BearerAuth
// @Success 200 {object}{object} common.Response{data=[]marketplacemodel.Collection}
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/ref-network [get]
func (u UserHandler) RefNetwork() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		//paging.Process()
		id := cast.ToInt64(ctx.GetInt("id"))
		if id == 0 {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid request", common.ErrNotFoundMarketPlaceId))
			return
		}
		var filter usermodel.RefNetworkFilter
		_ = ctx.ShouldBind(&filter)
		filter.UserId = id
		res, err := u.Client.RefNetworkProcess(filter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get items successfully", res))
	}
}

// Save setting item general godoc
// @Summary save setting item CMS user
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Accept json
// @Produce json
// @Param user body []usermodel.SystemSetting true "save setting request"
// @Success 200 {object} common.Response{data=[]usermodel.SystemSetting}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/setting/item [post]
func (u UserHandler) SaveSettingItem() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq []usermodel.SystemSetting
		if err := ctx.ShouldBindJSON(&rq); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return

		}
		scheme := "https://"
		options := map[string]interface{}{
			"host": scheme + ctx.Request.Host + "/",
		}
		if len(rq) < 1 {
			ctx.JSON(http.StatusOK, common.NewResponse(0, "Nothing to save", nil))
			return
		}
		res, err := u.Client.SaveSettingProcess(rq, options)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Save setting successfully", res))
	}
}

// delete setting item general godoc
// @Summary delete setting item CMS user
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Param id path int true "id"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/setting/item/{id} [delete]
func (u UserHandler) DeleteSettingItem() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := cast.ToInt64(ctx.Param("id"))
		err := u.Client.DeleteSettingProcess(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Delete successfully", nil))
	}
}

// default setting item general godoc
// @Summary default setting item CMS user
// @Tags CMS
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{data=usermodel.DefaultSetting}
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/setting/default [get]
func (u UserHandler) DefaultSetting() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response, err := u.Client.DefaultSystemSettingProcess()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Delete successfully", response))
	}
}

// setting item general godoc
// @Summary setting item CMS user
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{data=usermodel.DefaultSetting}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/setting [get]
func (u UserHandler) GetSetting() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response, err := u.Client.GetSystemSettingProcess()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Delete successfully", response))
	}
}

// payment success general godoc
// @Summary payment success
// @Tags PAYMENT
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{data=usermodel.DefaultSetting}
// @Failure 400,401,500 {object} common.Response
// @Router /payment/success [get]
func (u UserHandler) PaymentSuccess() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Delete successfully", nil))
	}
}

// payment notification general godoc
// @Summary payment notification
// @Tags PAYMENT
// @Accept json
// @Produce json
// @Param user body usermodel.PaymentNotificationRequest true "request"
// @Success 200 {object} common.Response{data=usermodel.DefaultSetting}
// @Failure 400,401,500 {object} common.Response
// @Router /payment/notification [post]
func (u UserHandler) PaymentNotification() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq usermodel.PaymentNotificationRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Delete successfully", nil))
	}
}

// ExploreUser Create User godoc
// @Summary get detail user
// @Tags User
// @Param order_by query string false "order by"
// @Param search query string false "search"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=usermodel.User}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/explore-user [get]
func (u UserHandler) ExploreUser() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var filter usermodel.ExploreUserFilter
		filter.OrderBy = ctx.Query("order_by")
		filter.Search = ctx.Query("search")
		res, err := u.Client.ExploreUser(filter)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", gin.H{
			"total_user": len(res),
			"users":      res,
		}))

	}
}

// Save daily event setting item general godoc
// @Summary save setting item CMS user
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Accept json
// @Produce json
// @Param user body usermodel.DailyEvent true "save daily event request"
// @Success 200 {object} common.Response{data=usermodel.DailyEvent}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/daily-event [post]
func (u UserHandler) SaveDailyEvent() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq usermodel.DailyEvent
		if err := ctx.ShouldBindJSON(&rq); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return

		}
		scheme := "https://"
		options := map[string]interface{}{
			"host": scheme + ctx.Request.Host + "/",
		}
		rq.Id = 0
		res, err := u.Client.SaveDailyEventProcess(rq, options)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Save setting successfully", res))
	}
}

// Update daily event setting item general godoc
// @Summary save setting item CMS user
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Accept json
// @Produce json
// @Param id path int true "daily event id"
// @Param user body usermodel.DailyEvent true "save daily event request"
// @Success 200 {object} common.Response{data=usermodel.DailyEvent}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/daily-event/{id} [put]
func (u UserHandler) UpdateDailyEvent() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq usermodel.DailyEvent
		if err := ctx.ShouldBindJSON(&rq); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return

		}
		scheme := "https://"
		options := map[string]interface{}{
			"host": scheme + ctx.Request.Host + "/",
		}
		rq.Id = cast.ToInt64(ctx.Param("id"))
		res, err := u.Client.SaveDailyEventProcess(rq, options)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Save setting successfully", res))
	}
}

// Delete daily event setting item general godoc
// @Summary save setting item CMS user
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Accept json
// @Produce json
// @Param id path int true "daily event id"
// @Success 200 {object} common.Response{data=usermodel.DailyEvent}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/daily-event/{id} [delete]
func (u UserHandler) DeleteDailyEvent() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := cast.ToInt64(ctx.Param("id"))
		err := u.Client.DeleteDailyEventProcess(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Save setting successfully", nil))
	}
}

// Detail daily event setting item general godoc
// @Summary save setting item CMS user
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Accept json
// @Produce json
// @Param id path int true "daily event id"
// @Success 200 {object} common.Response{data=usermodel.DailyEvent}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/daily-event/{id} [get]
func (u UserHandler) DetailDailyEvent() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := cast.ToInt64(ctx.Param("id"))
		if id == 0 {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid request", common.ErrNotFoundMarketPlaceId))
			return
		}
		var filter usermodel.DailyEventFilter
		_ = ctx.ShouldBind(&filter)
		filter.Id = id
		dailyEvent, err := u.Client.DetailDailyEventProcess(filter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Save setting successfully", dailyEvent))
	}
}

// List daily event setting item general godoc
// @Summary save setting item CMS user
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Accept json
// @Produce json
// @Param event_name query int false "event_name search like"
// @Param reward_type query int false "reward_type"
// @Param reward_op query string false "reward_op"
// @Param status query int false "status"
// @Param number_of_task query int false "number_of_task"
// @Param start_date_time query string false "start_date_time timestamp"
// @Param end_date_time query string false "end_date_time timestamp"
// @Param order_by query string false "order_by"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Success 200 {object} common.Response{data=[]usermodel.DailyEvent}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/daily-event [get]
func (u UserHandler) ListDailyEvent() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req usermodel.DailyEventFilter
		err := ctx.ShouldBind(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		var paging common.Paging
		if errPage := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", errPage.Error()))
			return
		}
		paging.Process()
		dailyEvent, err := u.Client.ListDailyEventProcess(req, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Save setting successfully", dailyEvent, paging))
	}
}

// Get user daily event godoc
// @Summary get detail user
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param current_time query string false "current_time timestamp"
// @Param is_history query int false "1 -> get history"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Success 200 {object}{object} common.Response{data=[]usermodel.UserDailyEventExInfo}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/daily-event [get]
func (u UserHandler) GetUserDailyEvent() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := cast.ToInt64(ctx.GetInt("id"))
		if id == 0 {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid request", common.ErrNotFoundMarketPlaceId))
			return
		}
		var req usermodel.DailyEventFilter
		err := ctx.ShouldBind(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		req.UserId = id
		var paging common.Paging
		if errPage := ctx.ShouldBind(&paging); errPage != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", errPage.Error()))
			return
		}
		paging.Process()
		var errProcess error
		var res []usermodel.UserDailyEventExInfo
		if req.IsHistory < 1 {
			res, errProcess = u.Client.UserDailyEventProcess(req, &paging)
		} else {
			res, errProcess = u.Client.UserDailyEventHistoryProcess(req, &paging)
		}
		if errProcess != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", errProcess.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get detail user successfully", res, paging))
	}
}

// Get List reward event godoc
// @Summary get detail user
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param is_available query int false "1 -> get reward available"
// @Param is_past_reward query int false "1 -> is_past_reward"
// @Param is_get_claimed query int false "1 -> is get claimed"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Success 200 {object}{object} common.Response{data=[]usermodel.RewardDepartment}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/list_reward [get]
func (u UserHandler) GetListReWard() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := cast.ToInt64(ctx.GetInt("id"))
		if id == 0 {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid request", common.ErrNotFoundMarketPlaceId))
			return
		}
		var req usermodel.RewardFilter
		err := ctx.ShouldBind(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		if req.IsGetClaimed > 0 {
			req.UserId = id
		}
		var paging common.Paging
		if errPage := ctx.ShouldBind(&paging); errPage != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", errPage.Error()))
			return
		}
		paging.Process()
		res, err := u.Client.ListRewardProcess(req, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get detail user successfully", res, paging))
	}
}

// Get Detail reward event godoc
// @Summary get detail user
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param id path int false "id"
// @Param is_get_claimed query int false "is_get_claimed"
// @Success 200 {object}{object} common.Response{data=usermodel.Reward}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/reward/{id} [get]
func (u UserHandler) GetReWard() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := cast.ToInt64(ctx.GetInt("id"))
		if id == 0 {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid request", common.ErrNotFoundMarketPlaceId))
			return
		}
		var req usermodel.RewardFilter
		err := ctx.ShouldBind(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		//req.UserId = id
		req.RewardId = cast.ToInt64(ctx.Param("id"))
		res, err := u.Client.DetailRewardProcess(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))
	}
}

// Save reward department general godoc
// @Summary save setting item CMS user
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Accept json
// @Produce json
// @Param user body usermodel.CreateRewardDepartmentRequest true "save reward department request"
// @Success 200 {object} common.Response{data=usermodel.RewardDepartment}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/reward-department [post]
func (u UserHandler) SaveRewardDepartment() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq usermodel.CreateRewardDepartmentRequest
		if err := ctx.ShouldBindJSON(&rq); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return

		}
		scheme := "https://"
		options := map[string]interface{}{
			"host": scheme + ctx.Request.Host + "/",
		}
		rq.Id = 0
		rewardDepartment := rq.ConvertToMainStruct()
		res, err := u.Client.SaveRewardDepartmentProcess(rewardDepartment, options)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Save setting successfully", res))
	}
}

// Save reward general godoc
// @Summary save reward CMS user
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Accept json
// @Produce json
// @Param user body usermodel.Reward true "save reward request"
// @Success 200 {object} common.Response{data=usermodel.Reward}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/reward [post]
func (u UserHandler) SaveReward() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq usermodel.Reward
		if err := ctx.ShouldBindJSON(&rq); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return

		}
		scheme := "https://"
		options := map[string]interface{}{
			"host": scheme + ctx.Request.Host + "/",
		}
		rq.Id = 0
		res, err := u.Client.SaveRewardProcess(rq, options)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Save setting successfully", res))
	}
}

// Claim reward general godoc
// @Summary Claim reward CMS user
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param user body usermodel.ClaimReward true "claim reward request"
// @Success 200 {object} common.Response{data=usermodel.UserReward}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/claim-reward [post]
func (u UserHandler) ClaimReward() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq usermodel.ClaimReward
		id := cast.ToInt64(ctx.GetInt("id"))
		if id == 0 {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid request", common.ErrNotFoundMarketPlaceId))
			return
		}
		if err := ctx.ShouldBindJSON(&rq); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return

		}
		rq.UserId = id
		res, err := u.Client.ClaimRewardProcess(rq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Save setting successfully", res))
	}
}

// Get user claim reward godoc
// @Summary get detail user
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=[]usermodel.UserDailyEventExInfo}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/user_claim_reward [get]
func (u UserHandler) GetUserReward() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := cast.ToInt64(ctx.GetInt("id"))
		if id == 0 {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid request", common.ErrNotFoundMarketPlaceId))
			return
		}
		var req usermodel.RewardFilter
		req.UserRewardUserId = id
		req.UserRewardStatus = common.USER_REWARD_STATUS_NEW
		res, err := u.Client.GetUserRewardProcess(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))
	}
}

// Get user daily event check godoc
// @Summary user daily event check user
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param is_share query int false "1 -> is_share"
// @Success 200 {object}{object} common.Response{}
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/daily-event/check/{is_share} [put]
func (u UserHandler) UserDailyEventCheck() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := cast.ToInt64(ctx.GetInt("id"))
		if id == 0 {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid request", common.ErrNotFoundMarketPlaceId))
			return
		}
		isShare := cast.ToInt64(ctx.Param("is_share"))
		dailyTypeEvent := common.DAILY_EVENT_TYPE_SWAP
		if isShare == 1 {
			dailyTypeEvent = common.DAILY_EVENT_TYPE_SHARE
		}
		u.Client.BuildUserDailyEvent(id, dailyTypeEvent)
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", nil))
	}
}

// GetAllListReWardDepartmentCms reward event godoc
// @Summary get list reward
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Accept json
// @Produce json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Success 200 {object}{object} common.Response{data=map[int64]string}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/all/list_reward_department [get]
func (u UserHandler) GetAllListReWardDepartmentCms() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.Paging
		if errPage := ctx.ShouldBind(&paging); errPage != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", errPage.Error()))
			return
		}
		paging.Process()
		res, err := u.Client.ListAllRewardDepartment(&paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get list reward successfully", res, paging))
	}
}

// GetListReWardDepartmentCms reward event godoc
// @Summary get list reward
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Accept json
// @Produce json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Success 200 {object}{object} common.Response{data=[]usermodel.RewardDepartment}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/list_reward_department [get]
func (u UserHandler) GetListReWardDepartmentCms() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.Paging
		if errPage := ctx.ShouldBind(&paging); errPage != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", errPage.Error()))
			return
		}
		paging.Process()
		res, err := u.Client.ListRewardDepart(&paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get list reward successfully", res, paging))
	}
}

// GetListReWardCms reward event godoc
// @Summary get list reward
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Accept json
// @Produce json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Success 200 {object}{object} common.Response{data=[]usermodel.RewardDepartment}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/list_reward [get]
func (u UserHandler) GetListReWardCms() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.Paging
		if errPage := ctx.ShouldBind(&paging); errPage != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", errPage.Error()))
			return
		}
		paging.Process()
		res, err := u.Client.ListReward(&paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get list reward successfully", res, paging))
	}
}

// SendVerifyEmail Create general godoc
// @Summary Send Verify Email
// @Param Authorization header string false "Insert your access token"
// @Param user body usermodel.SendVerifyEmailRequest true "Verify email request"
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/send-verify-email [post]
func (u UserHandler) SendVerifyEmail() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var res usermodel.SendVerifyEmailRequest
		if res := utils.CommonHandle(&res, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		userId := ctx.GetInt("id")
		if userId > 0 {
			res.UserId = int64(userId)
		}
		err := u.Client.SendVerifyMail(res)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Send verify mail successfully", nil))

	}
}

// VerifyEmail Create general godoc
// @Summary Verify Email
// @Tags User
// @Param Authorization header string false "Insert your access token"
// @Param user body usermodel.VerifyEmailRequest true "Verify email request"
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /v1/user/verify-email [post]
func (u UserHandler) VerifyEmail() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var res usermodel.VerifyEmailRequest
		if errRes := utils.CommonHandle(&res, ctx); errRes != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		userId := ctx.GetInt("id")
		if userId > 0 {
			res.UserId = int64(userId)
		}
		err := u.Client.VerifyMail(res)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Send verify mail successfully", nil))

	}
}

// ExchangeOPV exchange opv to USD
// @Param opv query string true "opv"
// @Tags Common
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /common/exchange [get]
func (u UserHandler) ExchangeOPV() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req usermodel.ExchangeOPV
		req.OPV = cast.ToFloat64(ctx.Query("opv"))
		res, err := u.Client.ExchangeOPV(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Exchange OPV successfully", res))

	}
}
