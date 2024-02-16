package market_place

import (
	"api-openlive/blockchain"
	"api-openlive/common"
	"api-openlive/module/market_place/business"
	marketplacemodel "api-openlive/module/market_place/model"
	usermodel "api-openlive/module/user/model"
	"api-openlive/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type MarketPlaceHandler struct {
	Client         *business.MarketPlaceController
	Blc            *blockchain.Client
	OptimisticLock *utils.OptimisticLock
}

// ListBrand List Brand godoc
// @Summary get list brand
// @Tags Market Place
// @Param is_hot query int false "is_hot -> 1 = hot || 0 = not hot || -1 = all"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.ResponsePaginate{data=[]marketplacemodel.Brand}
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/list-brand [get]
func (u MarketPlaceHandler) ListBrand() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var filter marketplacemodel.BrandFilter
		err := ctx.ShouldBind(&filter)
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
		res, err := u.Client.ListBrands(filter, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get list brands successfully", res, paging))
	}
}

// Brand godoc
// @Summary get brand
// @Tags Market Place
// @Param is_hot query int false "is_hot -> 1 = hot || 0 = not hot || -1 = all"
// @Param id path int true "brand id"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.ResponsePaginate{data=[]marketplacemodel.Brand}
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/brand/{id} [get]
func (u MarketPlaceHandler) Brand() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var filter marketplacemodel.BrandFilter
		err := ctx.ShouldBind(&filter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		id := cast.ToInt64(ctx.Param("id"))
		filter.Id = id
		filter.IsHot = -1
		var paging common.Paging
		if errPage := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", errPage.Error()))
			return
		}

		paging.Process()
		res, err := u.Client.ListBrands(filter, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get brand successfully", res, paging))
	}
}

// SellItem Sell Item godoc
// @Summary get detail item
// @Tags Market Place
// @Accept json
// @Produce json
// @Param user body marketplacemodel.CreateMarketPlaceRequest true "login request"
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/sell-item [post]
func (u MarketPlaceHandler) SellItem() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var userReq marketplacemodel.CreateMarketPlaceRequest
		if res := utils.CommonHandle(&userReq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}

		hashKey := utils.GetMD5Hash(userReq.WalletAddress)
		errLock := u.OptimisticLock.Lock(hashKey)
		if errLock != nil {
			if errLock.Error() == utils.StatusLocked {
				ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", "Request is processing"))
				return
			}
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", errLock.Error()))
			return
		}
		defer func(OptimisticLock *utils.OptimisticLock, key string) {
			err := OptimisticLock.Release(key)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
				return
			}
		}(u.OptimisticLock, hashKey)

		err := u.Client.CreateMarketPlace(userReq, userReq.ContractAddress)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Register market user successfully", nil))

	}
}

// BuyItem Sell Item godoc
// @Summary get detail item
// @Tags Market Place
// @Accept json
// @Produce json
// @Param user body marketplacemodel.CreateMarketPlaceSoldRequest true "login request"
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/buy-item [post]
func (u MarketPlaceHandler) BuyItem() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var userReq marketplacemodel.CreateMarketPlaceSoldRequest
		if res := utils.CommonHandle(&userReq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}

		err := u.Client.CreateMarketPlaceSold(userReq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Register market user successfully", nil))

	}
}

// ListItem Sell Item godoc
// @Summary get detail item
// @Tags Market Place
// @Param user_id query int false "user_id"
// @Param sort_by query string false "sort by `field`"
// @Param price query string false "sort by price"
// @Param total_like query string false "sort by total_like"
// @Param is_hot query string false "sort by is_hot"
// @Param sell_type query string false "1 sell || 2 auction"
// @Param created_at query string false "sort by created_at"
// @Param status query int false "sort by status"
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param category_id query int false "filter by category id"
// @Param is_trending query int false "is trending"
// @Param creator query int false "filter by creator"
// @Param price_max query int false "filter by price range"
// @Param price_min query int false "filter by price range"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/list-item [get]
func (u MarketPlaceHandler) ListItem() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.Paging
		_ = ctx.ShouldBind(&paging)
		paging.Process()

		var filter marketplacemodel.MarketPlaceFilter
		_ = ctx.ShouldBind(&filter)

		userId := cast.ToInt64(ctx.Query("user_id"))
		filter.UserId = userId

		filter.Status = common.MARKET_PLACE_STATUS_SELLING
		filter.IsDiscover = 1
		filter.IsHide = common.ITEM_UNHIDE
		data, err := u.Client.ListItemMarketPlace(filter, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get items successfully", data, paging))

	}
}

// ListItem Event Item godoc
// @Summary get list item event
// @Tags Market Place
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param user_id query int false "user_id"
// @Param is_show_all query int false "is_show_all"
// @Param sort_by query string false "sort by `field`"
// @Param category_id query int false "filter by category id"
// @Param search query string false "search by name"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/list-item-event [get]
func (u MarketPlaceHandler) ListItemEvent() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.Paging
		_ = ctx.ShouldBind(&paging)
		paging.Process()

		var filter marketplacemodel.ItemEventFilter
		_ = ctx.ShouldBind(&filter)
		userId := cast.ToInt64(ctx.Query("user_id"))
		filter.CurrentUserId = userId
		filter.IsHide = common.ITEM_UNHIDE
		data, err := u.Client.ListItemEvent(filter, &paging)
		if err != nil {
			ctx.JSON(http.StatusForbidden, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get items successfully", data, paging))

	}
}

// ListMostLikedItem Sell Item godoc
// @Summary get detail item
// @Tags Market Place
// @Param user_id query int false "user_id"
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/most-liked/list [get]
func (u MarketPlaceHandler) ListMostLikedItem() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.Paging
		_ = ctx.ShouldBind(&paging)
		paging.Process()

		var filter marketplacemodel.MarketPlaceFilter
		_ = ctx.ShouldBind(&filter)

		userId := cast.ToInt64(ctx.Query("user_id"))
		filter.UserId = userId

		data, err := u.Client.ListMostLikeItemMarketPlace(filter, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get items successfully", data, paging))

	}
}

// ListVotedItem Sell Item godoc
// @Summary get detail item
// @Tags Market Place
// @Param sort_by query string false "sort by `field`"
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/list-item-voted [get]
func (u MarketPlaceHandler) ListVotedItem() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.Paging
		_ = ctx.ShouldBind(&paging)
		paging.Process()

		var filter marketplacemodel.MarketPlaceFilter
		_ = ctx.ShouldBind(&filter)

		data, err := u.Client.ListVotedItemMarketPlace(filter, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get items successfully", data, paging))

	}
}

// DetailSellItem detail Sell Item godoc
// @Summary get detail item
// @Tags Market Place
// @Param nft path string true "user id"
// @Param contract_address path string true "smart contract"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/item/{contract_address}/{nft} [get]
func (u MarketPlaceHandler) DetailSellItem() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		nft := ctx.Param("nft")
		contractAddress := ctx.Param("contract_address")
		userId := cast.ToInt64(ctx.Query("user_id"))

		res, err := u.Client.DetailSellItem(contractAddress, nft, userId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// ChangeStatus ListItem Sell Item godoc
// @Summary get detail item
// @Tags Market Place
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param id path int true "marketplace id"
// @Param id body marketplacemodel.ChangeMPStatusRequest true "marketplace vote request"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/{id}/change-status [put]
func (u MarketPlaceHandler) ChangeStatus() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req marketplacemodel.ChangeMPStatusRequest
		_ = ctx.ShouldBindJSON(&req)

		id := cast.ToInt64(ctx.Param("id"))
		if id == 0 {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid request", common.ErrNotFoundMarketPlaceId))
			return
		}
		req.Id = id

		data, err := u.Client.ChangeItemStatus(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Change status successfully", data))

	}
}

// Vote Vote item
// @Summary vote item
// @Tags Market Place
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param id body marketplacemodel.VoteItemRequest true "marketplace vote request"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/vote [post]
func (u MarketPlaceHandler) Vote() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req marketplacemodel.VoteItemRequest
		_ = ctx.ShouldBindJSON(&req)

		userId := ctx.GetInt("id")
		req.UserId = cast.ToInt64(userId)

		data, err := u.Client.VoteItem(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Vote item success", data))

	}
}

// Like item
// @Summary vote item
// @Tags Market Place
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param id body marketplacemodel.LikeItemRequest true "marketplace vote request"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/like [post]
func (u MarketPlaceHandler) Like() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req marketplacemodel.LikeItemRequest
		_ = ctx.ShouldBindJSON(&req)

		userId := ctx.GetInt("id")
		req.UserId = cast.ToInt64(userId)

		err := u.Client.LikeItem(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "success", nil))

	}
}

// DisLike item
// @Summary vote item
// @Tags Market Place
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param id body marketplacemodel.LikeItemRequest true "marketplace vote request"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/dislike [post]
func (u MarketPlaceHandler) DisLike() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req marketplacemodel.LikeItemRequest
		_ = ctx.ShouldBindJSON(&req)

		userId := ctx.GetInt("id")
		req.UserId = cast.ToInt64(userId)

		err := u.Client.DisLikeItem(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "success", nil))

	}
}

// StakingItem item
// @Summary staking item
// @Tags Market Place
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param id body marketplacemodel.StakingRequest true "marketplace staking request"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/staking-item [post]
func (u MarketPlaceHandler) StakingItem() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req marketplacemodel.StakingRequest
		if res := utils.CommonHandle(&req, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}

		data, err := u.Client.StakingItem(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Staking item success", data))

	}
}

// UpdateStakingStatus StakingItem item
// @Summary staking item
// @Tags Market Place
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param id body marketplacemodel.UpdateStakingStatusRequest true "marketplace update staking request"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/staking-complete [post]
func (u MarketPlaceHandler) UpdateStakingStatus() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req marketplacemodel.UpdateStakingStatusRequest
		if res := utils.CommonHandle(&req, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}

		data, err := u.Client.UpdateStatusStaking(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Update item success", data))

	}
}

// DetailItem Detail Item godoc
// @Param Authorization header string false "Insert your access token"
// @Security BearerAuth
// @Param nft path string true "user id"
// @Param user_id query int false "user id"
// @Param smart_contract path string true "smart contract"
// @Summary get detail item
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{data=marketplacemodel.Item}
// @Failure 400,401,500 {object} common.Response
// @Router /item/{smart_contract}/{nft} [get]
func (u MarketPlaceHandler) DetailItem() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		nft := ctx.Param("nft")
		smartContract := ctx.Param("smart_contract")
		userId := cast.ToInt64(ctx.Query("user_id"))
		res, err := u.Client.DetailItem(marketplacemodel.ItemFilter{
			Nft:             nft,
			ContractAddress: smartContract,
			UserId:          userId,
			IsTrending:      -1,
		})
		if userId > 0 {
			res.UserId = cast.ToInt64(userId)
		}
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// DetailItemEvent Detail Item Status godoc
// @Param nft path string true "user id"
// @Param smart_contract path string true "smart contract"
// @Param user_id query int false "user_id"
// @Summary get detail item
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{data=marketplacemodel.Item}
// @Failure 400,401,500 {object} common.Response
// @Router /item/event/{smart_contract}/{nft} [get]
func (u MarketPlaceHandler) DetailItemEvent() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		nft := ctx.Param("nft")
		smartContract := ctx.Param("smart_contract")
		userId := cast.ToInt64(ctx.Query("user_id"))
		res, err := u.Client.DetailItem(marketplacemodel.ItemFilter{
			Nft:             nft,
			ContractAddress: smartContract,
			IsJoinEvent:     common.ITEM_ON_JOIN_EVENT,
			UserId:          userId,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// ItemEventStatistic Detail Item Status godoc
// @Summary get detail item
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{data=marketplacemodel.Item}
// @Failure 400,401,500 {object} common.Response
// @Router /item/event/statistic [get]
func (u MarketPlaceHandler) ItemEventStatistic() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		res, err := u.Client.NFTEventStatistic()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))
	}
}

// DetailItem Detail Item godoc
// @Param nft path string true "user id"
// @Param smart_contract path string true "smart contract"
// @Summary get detail item
// @Tags Item
// @Accept json
// @Produce json
// @Router /item-nft/{smart_contract}/{nft} [get]
func (u MarketPlaceHandler) DetailItemShow() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		nft := ctx.Param("nft")
		smartContract := ctx.Param("smart_contract")
		res, err := u.Client.DetailItem(marketplacemodel.ItemFilter{
			Nft:             nft,
			ContractAddress: smartContract,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponseShow(res.Nft, res.Title, res.Description, res.Image))

	}
}

// ListItems List Item godoc
// @Summary get list item
// @Param Authorization header string false "Insert your access token"
// @Security BearerAuth
// @Param limit query int false "limit"
// @Param user_id query int false "user_id"
// @Param owner_id query int false "owner_id"
// @Param page query int false "page"
// @Param status query int false "status"
// @Param is_favourite query int false "is_favourite"
// @Param is_sold query int false "is_sold"
// @Param is_created query int false "is_created"
// @Param is_owned query int false "is_owned"
// @Param is_join_event query int false "is_join_event"
// @Param stts query string false "status arr"
// @Param search query string false "search by name"
// @Param not_collection query int false "==1 list item hasn't collection"
// @Param price query string false "price: desc or asc"
// @Param order_by query string false "order_by"
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object} common.ResponsePaginate{data=[]marketplacemodel.Item}
// @Failure 400,401,500 {object} common.ResponsePaginate
// @Router /item/list [get]
func (u MarketPlaceHandler) ListItems() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.GetInt("id")
		stts := ctx.Query("stts")
		targetUserId, _ := strconv.Atoi(ctx.Query("user_id"))
		if targetUserId == 0 {
			targetUserId = id
		}
		var itemFilter marketplacemodel.ItemFilter
		// parse paging
		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		paging.Process()

		err := ctx.ShouldBind(&itemFilter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		arr := strings.Split(stts, ",")
		arrStatus := make([]int64, 0, 10)
		for _, v := range arr {
			value := cast.ToInt64(v)
			if value > 0 {
				arrStatus = append(arrStatus, value)
			}
		}
		itemFilter.Stts = arrStatus
		itemFilter.IsNotGetMulti = 1
		itemFilter.IsTrending = -1
		itemFilter.UserId = cast.ToInt64(targetUserId)
		itemFilter.PresentUserId = cast.ToInt64(id)
		itemFilter.IsHide = common.ITEM_UNHIDE
		res, err := u.Client.ListItemWithConditions(itemFilter, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get list item successfully", res, paging))

	}
}

// AddBannerCms list godoc
// @Summary add banner
// @Tags CMS
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param update body marketplacemodel.AddBannerCmsRequest true "update request"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /cms/updatebanner [post]
func (u MarketPlaceHandler) AddBannerCms() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq marketplacemodel.AddBannerCmsRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		err := u.Client.AddBannerCms(&rq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Add banner successfully", nil))
	}
}

// ListBrandCms Create User godoc
// @Summary get list brand
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Param is_hot query int false "is_hot -> 1 = hot || 0 = not hot || -1 = all"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.ResponsePaginate{data=[]marketplacemodel.Brand}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/listbrands [get]
func (u MarketPlaceHandler) ListBrandCms() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var filter marketplacemodel.BrandFilter
		err := ctx.ShouldBind(&filter)
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
		res, err := u.Client.ListBrands(filter, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get list brands successfully", res, paging))
	}
}

// DeleteBrandImageCms Create User godoc
// @Summary delete brand image
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Param id path int true "id"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /cms/delete/brandimage/{id} [delete]
func (u MarketPlaceHandler) DeleteBrandImageCms() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := cast.ToInt64(ctx.Param("id"))
		err := u.Client.DeleteBrandImage(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Delete brand image successfully", nil))
	}
}

// UpdateMarketPlaceCMS godoc
// @Summary update marketplace
// @Tags CMS
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param user body marketplacemodel.UpdateMarketPlaceRequest true "update request"
// @Param id path int true "id"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /cms/market-place/update/{id} [patch]
func (u MarketPlaceHandler) UpdateMarketPlaceCMS() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq marketplacemodel.UpdateMarketPlaceRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		id := cast.ToInt64(ctx.Param("id"))
		rq.Id = id
		err := u.Client.UpdateMarketPlaceCms(&rq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Update marketplace successfully", nil))
	}
}

// ListItemCMS List Item godoc
// @Summary get detail item
// @Tags CMS
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param user_id query int false "user_id"
// @Param sort_by query string false "sort by `field`"
// @Param price query string false "sort by price"
// @Param total_like query string false "sort by total_like"
// @Param is_hot query string false "sort by is_hot"
// @Param created_at query string false "sort by created_at"
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param category_id query int false "filter by category id"
// @Param is_trending query int false "is trending"
// @Param creator query int false "filter by creator"
// @Param price_max query int false "filter by price range"
// @Param price_min query int false "filter by price range"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /cms/market-place/list-item [get]
func (u MarketPlaceHandler) ListItemCMS() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.Paging
		_ = ctx.ShouldBind(&paging)
		paging.Process()

		var filter marketplacemodel.MarketPlaceFilter
		_ = ctx.ShouldBind(&filter)

		userId := cast.ToInt64(ctx.Query("user_id"))
		filter.UserId = userId

		filter.Status = common.MARKET_PLACE_STATUS_SELLING
		filter.IsDiscover = 1
		filter.IsHide = common.ITEM_UNHIDE
		data, err := u.Client.ListItemMarketPlace(filter, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get items successfully", data, paging))

	}
}

// DeleteBrandCms Create User godoc
// @Summary delete brand image
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Param id path int true "id"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /cms/delete/brand/{id} [delete]
func (u MarketPlaceHandler) DeleteBrandCms() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := cast.ToInt64(ctx.Param("id"))
		err := u.Client.DeleteBrand(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Delete brand image successfully", nil))
	}
}

// ChangeBrandCms list godoc
// @Summary update brands
// @Tags CMS
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param user body marketplacemodel.UpdateBrandCmsRequest true "update request"
// @Param id path int true "brand id"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /cms/change/brands/{id} [patch]
func (u MarketPlaceHandler) ChangeBrandCms() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq marketplacemodel.UpdateBrandCmsRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		id := cast.ToInt64(ctx.Param("id"))
		rq.Id = id
		scheme := "https://"
		/*if ctx.Request.TLS != nil {
			scheme = "https://"
		}*/
		options := map[string]interface{}{
			"host": scheme + ctx.Request.Host + "/",
		}
		err := u.Client.UpdateBrandCms(&rq, options)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Update brands successfully", nil))
	}
}

// ChangeBrandImageCms list godoc
// @Summary update brands image
// @Tags CMS
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param user body marketplacemodel.UpdateBrandImageCmsRequest true "update request"
// @Param id path int true "brand id"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /cms/change/brandsimage/{id} [patch]
func (u MarketPlaceHandler) ChangeBrandImageCms() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq marketplacemodel.UpdateBrandImageCmsRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		id := cast.ToInt64(ctx.Param("id"))
		rq.Id = id
		scheme := "https://"
		/*if ctx.Request.TLS != nil {
			scheme = "https://"
		}*/
		options := map[string]interface{}{
			"host": scheme + ctx.Request.Host + "/",
		}
		err := u.Client.UpdateBrandImageCms(&rq, options)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Update brands successfully", nil))
	}
}

// CreateBrandImageCms list godoc
// @Summary create brands image
// @Tags CMS
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param user body marketplacemodel.CreateBrandImageCmsRequest true "update request"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /cms/create/brandsimage [post]
func (u MarketPlaceHandler) CreateBrandImageCms() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq marketplacemodel.CreateBrandImageCmsRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		scheme := "https://"
		/*if ctx.Request.TLS != nil {
			scheme = "https://"
		}*/
		options := map[string]interface{}{
			"host": scheme + ctx.Request.Host + "/",
		}
		res, err := u.Client.CreateBrandImageCms(&rq, options)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Update brands successfully", res))
	}
}

// CreateBrandCms list godoc
// @Summary create brands
// @Tags CMS
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param user body marketplacemodel.CreateBrandCmsRequest true "update request"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /cms/create/brands [post]
func (u MarketPlaceHandler) CreateBrandCms() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq marketplacemodel.CreateBrandCmsRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		scheme := "https://"
		/*if ctx.Request.TLS != nil {
			scheme = "https://"
		}*/
		options := map[string]interface{}{
			"host": scheme + ctx.Request.Host + "/",
		}
		err := u.Client.CreateBrandCms(&rq, options)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Update brands successfully", nil))
	}
}

// ListItemsCMS List Item godoc
// @Summary get list item cms
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param status query int false "status"
// @Param order_by query string false "order by"
// @Tags CMS
// @Accept json
// @Produce json
// @Success 200 {object} common.ResponsePaginate{data=[]marketplacemodel.Item}
// @Failure 400,401,500 {object} common.ResponsePaginate
// @Router /cms/item/list [get]
func (u MarketPlaceHandler) ListItemsCMS() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var itemFilter marketplacemodel.ItemFilter
		// parse paging
		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		paging.Process()

		err := ctx.ShouldBind(&itemFilter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		itemFilter.Status = 0
		itemFilter.Stts = []int64{common.ITEM_ON_SALE, common.ITEM_NEW, common.ITEM_READY_ON_SELL, common.ITEM_PENDING}
		if len(itemFilter.OrderBy) > 0 {
			itemFilter.OrderBy += ",items.display_order DESC"
		} else {
			itemFilter.OrderBy += "items.display_order DESC"
		}
		itemFilter.IsTrending = -1
		res, err := u.Client.ListItemWithConditions(itemFilter, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get list item successfully", res, paging))

	}
}

// ListCollectionItems of collection List Item godoc
// @Summary get ListItems of collection
// @Param collection_id path int true "collection id"
// @Param current_user_id query int false "current_user_id"
// @Param owner_id query int false "owner_id"
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param status query int false "status"
// @Param search query string false "search nft name"
// @Param stts query string false "status arr"
// @Param is_approval query int false "is_approval filter"
// @Param price query string false "price: desc or asc"
// @Param view query string false "view: desc or asc"
// @Param favourite query string false "favourite: desc or asc"
// @Param created query string false "created: desc or asc"
// @Param order_by query string false "order_by"
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object} common.ResponsePaginate{data=[]marketplacemodel.Item}
// @Failure 400,401,500 {object} common.ResponsePaginate
// @Router /item/collection/item/{id} [get]
func (u MarketPlaceHandler) ListCollectionItems() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var itemFilter marketplacemodel.ItemFilter
		// parse paging
		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		paging.Process()

		err := ctx.ShouldBind(&itemFilter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		itemFilter.CollectionId = cast.ToInt64(ctx.Param("id"))
		itemFilter.IsHide = common.ITEM_UNHIDE
		itemFilter.Stts = []int64{common.ITEM_ON_SALE, common.ITEM_NEW, common.ITEM_READY_ON_SELL}
		itemFilter.IsNotGetMulti = 1
		res, err := u.Client.ListItemWithConditions(itemFilter, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get list item successfully", res, paging))

	}
}

// CreateCategory Create User godoc
// @Summary create category
// @Param user body marketplacemodel.CreateCategoryRequest true "nonce request"
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=marketplacemodel.Category}
// @Failure 400,401,500 {object} common.Response
// @Router /item/category [post]
func (u MarketPlaceHandler) CreateCategory() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq marketplacemodel.CreateCategoryRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		res, err := u.Client.CreateCategory(&rq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// UpdateCategory Create User godoc
// @Summary update category
// @Param user body marketplacemodel.UpdateCategoryRequest true "nonce request"
// @Param id path int true "category id"
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=marketplacemodel.Category}
// @Failure 400,401,500 {object} common.Response
// @Router /item/category/{id} [put]
func (u MarketPlaceHandler) UpdateCategory() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req marketplacemodel.UpdateCategoryRequest
		id := cast.ToInt64(ctx.Param("id"))
		if res := utils.CommonHandle(&req, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		req.Id = id
		res, err := u.Client.UpdateCategory(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// ListCategory Create User godoc
// @Summary get detail user
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=[]marketplacemodel.Category}
// @Failure 400,401,500 {object} common.Response
// @Router /item/list-category [get]
func (u MarketPlaceHandler) ListCategory() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		res, err := u.Client.ListCategory(marketplacemodel.CategoryFilter{})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// ListCollections Create User godoc
// @Summary get detail user
// @Tags Item
// @Param is_hot query int false "is_hot -> 1 = hot || 0 = not hot || -1 = all"
// @Param user_id query int false "filter user has liked collection"
// @Param limit query int false "limit"
// @Param search query string false "search by name"
// @Param page query int false "page"
// @Param status query int false "status"
// @Param order_by query string false "order_by"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=[]marketplacemodel.Collection}
// @Failure 400,401,500 {object} common.Response
// @Router /item/list-collection [get]
func (u MarketPlaceHandler) ListCollections() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var filter marketplacemodel.CollectionFilter
		err := ctx.ShouldBind(&filter)
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
		if len(filter.OrderBy) > 0 {
			filter.OrderBy += ",collections.display_order DESC"
		} else {
			filter.OrderBy += "collections.display_order DESC"
		}
		filter.IsTrending = -1
		res, err := u.Client.ListCollections(filter, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get list collection successfully", res, paging))
	}
}

// ListCollectionsCms Create User godoc
// @Summary get detail user
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Param is_hot query int false "is_hot -> 1 = hot || 0 = not hot || -1 = all"
// @Param user_id query int false "filter user has liked collection"
// @Param limit query int false "limit"
// @Param search query string false "search by name"
// @Param page query int false "page"
// @Param status query int false "status"
// @Param order_by query string false "order_by"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=[]marketplacemodel.Collection}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/collection [get]
func (u MarketPlaceHandler) ListCollectionsCms() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var filter marketplacemodel.CollectionFilter
		err := ctx.ShouldBind(&filter)
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
		filter.IsHot = -1
		filter.NotGetItem = 1
		filter.IsCms = 1
		filter.IsTrending = -1
		res, err := u.Client.ListCollections(filter, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get list collection successfully", res, paging))
	}
}

// ItemCreator Create User godoc
// @Summary create category
// @Param Authorization header string true "Insert your access token"
// @Param user body marketplacemodel.CreateItemRequest true "nonce request"
// @Security BearerAuth
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=marketplacemodel.Item}
// @Failure 400,401,500 {object} common.Response
// @Router /item/create-item [post]
func (u MarketPlaceHandler) ItemCreator() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq marketplacemodel.CreateItemRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		userId := ctx.GetInt("id")
		scheme := "https://"
		/*if ctx.Request.TLS != nil {
			scheme = "https://"
		}*/
		options := map[string]interface{}{
			"host": scheme + ctx.Request.Host + "/",
		}
		res, err := u.Client.ItemCreator(&rq, int64(userId), options)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "success", res))

	}
}

// CreateCollection godoc
// @Summary create category
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param user body marketplacemodel.CreateCollectionRequest true "nonce request"
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=marketplacemodel.Collection}
// @Failure 400,401,500 {object} common.Response
// @Router /item/collection [post]
func (u MarketPlaceHandler) CreateCollection() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq marketplacemodel.CreateCollectionRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		id := ctx.GetInt("id")
		rq.UserId = cast.ToInt64(id)
		scheme := "https://"
		/*if ctx.Request.TLS != nil {
			scheme = "https://"
		}*/
		options := map[string]interface{}{
			"host": scheme + ctx.Request.Host + "/",
		}
		res, err := u.Client.CreateCollections(&rq, options)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Create collection success ", res))
	}
}

// CollectionDetail godoc
// @Summary CollectionDetail
// @Param id path int true "collection id"
// @Param user_id query int false "user id"
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=marketplacemodel.CollectionDetail}
// @Failure 400,401,500 {object} common.Response
// @Router /item/collection/{id} [get]
func (u MarketPlaceHandler) CollectionDetail() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		tmpCollectionId := ctx.Param("id")
		collectionId := cast.ToInt64(tmpCollectionId)
		userId := cast.ToInt64(ctx.Query("user_id"))

		res, err := u.Client.CollectionDetail(collectionId, userId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Collection not found", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Detail collection success ", res))

	}
}

// CanCreateCollection godoc
// @Summary create category
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param user body marketplacemodel.CanCreateCollectionRequest true "nonce request"
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=marketplacemodel.Collection}
// @Failure 400,401,500 {object} common.Response
// @Router /item/is-can-create-collection [post]
func (u MarketPlaceHandler) IsCanCreateCollection() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq marketplacemodel.CanCreateCollectionRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		id := ctx.GetInt("id")
		rq.UserId = cast.ToInt64(id)
		res, err := u.Client.CanCreateCollections(&rq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Can Create collection", res))
	}
}

// ListSelfCollections self Collections godoc
// @Summary get detail user
// @Param Authorization header string false "Insert your access token"
// @Security BearerAuth
// @Param user_id query int false "user_id"
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param order_by query string false "order_by"
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=[]marketplacemodel.Collection}
// @Failure 400,401,500 {object} common.Response
// @Router /item/user-collection [get]
func (u MarketPlaceHandler) ListSelfCollections() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.GetInt("id")
		targetUserId, _ := strconv.Atoi(ctx.Query("user_id"))
		if targetUserId != 0 {
			id = targetUserId
		}
		orderBy := ctx.Query("order_by")
		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		paging.Process()

		res, err := u.Client.ListCollections(marketplacemodel.CollectionFilter{UserId: cast.ToInt64(id), IsHot: -1, OrderBy: orderBy, IsTrending: -1}, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get detail user successfully", res, paging))

	}
}

// Transasction list godoc
// @Summary get detail user
// @Tags Market Place
// @Param sort_by query string false "sort by `field`"
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=[]marketplacemodel.Collection}
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/transaction [get]
func (u MarketPlaceHandler) ListTransaction() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.Paging
		_ = ctx.ShouldBind(&paging)
		//paging.Process()
		var filter marketplacemodel.MarketPlaceFilter
		_ = ctx.ShouldBind(&filter)
		res, err := u.Client.ListTransaction(filter, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get items successfully", res, paging))
	}
}

// UpdateItem list godoc
// @Summary update item
// @Tags Item
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param user body marketplacemodel.UpdateItemRequest true "nonce request"
// @Param id path int true "category id"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=marketplacemodel.Item}
// @Failure 400,401,500 {object} common.Response
// @Router /item/{id} [patch]
func (u MarketPlaceHandler) UpdateItem() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq marketplacemodel.UpdateItemRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		id := cast.ToInt64(ctx.Param("id"))
		userid := ctx.GetInt("id")
		rq.Id = id
		rq.UserId = int64(userid)
		if rq.Status != common.ITEM_PENDING && rq.Status != 0 {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", "cant not update status"))
			return
		}
		res, err := u.Client.UpdateItem(&rq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get items successfully", res))
	}
}

// UpdateItemStatusCMS list godoc
// @Summary update item cms
// @Tags CMS
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param user body marketplacemodel.UpdateItemStatusRequest true "nonce request"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=marketplacemodel.Item}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/item/update-status [post]
func (u MarketPlaceHandler) UpdateItemStatusCMS() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq marketplacemodel.UpdateItemStatusRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		err := u.Client.UpdateItemStatusCMS(&rq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Update status items successfully", nil))
	}
}

// UpdateItemStatusCMS list godoc
// @Summary update item cms
// @Tags CMS
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param user body marketplacemodel.UpdateItemCmsRequest true "update request"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=marketplacemodel.Item}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/item/update/{id} [put]
func (u MarketPlaceHandler) UpdateItemCMS() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := cast.ToInt64(ctx.Param("id"))
		var rq marketplacemodel.UpdateItemCmsRequest
		rq.Id = id
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		err := u.Client.UpdateItemCMS(&rq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Update status items successfully", nil))
	}
}

// DetailItemCMS Detail Item godoc
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param id path int true "item id"
// @Summary get detail item
// @Tags CMS
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{data=marketplacemodel.Item}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/item/{id} [get]
func (u MarketPlaceHandler) DetailItemCMS() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := cast.ToInt64(ctx.Param("id"))
		res, err := u.Client.DetailItem(marketplacemodel.ItemFilter{
			Id: id,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// LikeCollection item
// @Summary like collection
// @Tags Item
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param id body marketplacemodel.LikeCollectionRequest true "marketplace vote request"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /item/collection/like [post]
func (u MarketPlaceHandler) LikeCollection() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req marketplacemodel.LikeCollectionRequest
		_ = ctx.ShouldBindJSON(&req)

		if req.Id == 0 {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", "collection not found"))
			return
		}
		userId := ctx.GetInt("id")
		req.UserId = cast.ToInt64(userId)

		err := u.Client.LikeCollection(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "success", nil))

	}
}

// DisLikeCollection item
// @Summary dislike collection
// @Tags Item
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param id body marketplacemodel.LikeCollectionRequest true "marketplace vote request"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /item/collection/dislike [post]
func (u MarketPlaceHandler) DisLikeCollection() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req marketplacemodel.LikeCollectionRequest
		_ = ctx.ShouldBindJSON(&req)

		userId := ctx.GetInt("id")
		req.UserId = cast.ToInt64(userId)

		err := u.Client.DisLikeCollection(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "success", nil))

	}
}

// UpdateCollection list godoc
// @Summary update item
// @Tags Item
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param user body marketplacemodel.UpdateCollectionRequest true "nonce request"
// @Param id path int true "collection id"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=marketplacemodel.Collection}
// @Failure 400,401,500 {object} common.Response
// @Router /item/collection/{id} [patch]
func (u MarketPlaceHandler) UpdateCollection() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq marketplacemodel.UpdateCollectionRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		id := cast.ToInt64(ctx.Param("id"))
		userid := ctx.GetInt("id")
		rq.Id = id
		rq.UserId = int64(userid)
		scheme := "https://"
		options := map[string]interface{}{
			"host": scheme + ctx.Request.Host + "/",
		}

		res, err := u.Client.UpdateCollection(&rq, options)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get items successfully", res))
	}
}

// UpdateCollection list godoc
// @Summary update item
// @Tags Item
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param user body marketplacemodel.UpdateCollectionCmsRequest true "update request"
// @Param id path int true "collection id"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=marketplacemodel.Collection}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/collection/{id} [patch]
func (u MarketPlaceHandler) UpdateCollectionCms() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq marketplacemodel.UpdateCollectionCmsRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		id := cast.ToInt64(ctx.Param("id"))
		rq.Id = id

		err := u.Client.UpdateCollectionCms(&rq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get items successfully", nil))
	}
}

// CanJoinEvent Check Item godoc
// @Summary get check item can join event.
// @Tags Item
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param smart_contract path string true "smart_contract"
// @Param nft path string true "nft"
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /item/can-join-event/{smart_contract}/{nft} [get]
func (u MarketPlaceHandler) CanJoinEvent() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		nft := ctx.Param("nft")
		smartContract := ctx.Param("smart_contract")

		var result marketplacemodel.CheckCanJoinEventResponse

		userId := ctx.GetInt("id")

		res, err := u.Client.CheckItemCanJoinEvent(smartContract, nft, userId)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		result.Status = res
		ctx.JSON(http.StatusOK, common.NewResponse(0, "CheckItemCanJoinEvent successfully", result))
	}
}

// Stats Create general godoc
// @Summary stats CMS
// @Tags Item
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param time query int false "time"
// @Param id path int true "collection id"
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /item/collection/{id}/stats [get]
func (u MarketPlaceHandler) Stats() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		tmpCollectionId := ctx.Param("id")
		time := cast.ToInt64(ctx.Query("time"))
		if time == 0 {
			time = 7
		}
		res, err := u.Client.CmsStats(cast.ToInt64(tmpCollectionId), time)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Login successfully", res))
	}
}

// ExploreStatistic list godoc
// @Summary get detail user
// @Tags Market Place
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=[]marketplacemodel.Collection}
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/explore-statistic [get]
func (u MarketPlaceHandler) ExploreStatistic() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		//paging.Process()
		var filter marketplacemodel.MarketPlaceFilter
		_ = ctx.ShouldBind(&filter)
		res, err := u.Client.ExploreStatistic()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get items successfully", res))
	}
}

// History Detail Item godoc
// @Param nft path string true "user id"
// @Param smart_contract path string true "smart contract"
// @Param type query string false "type activities"
// @Summary get detail item
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{data=marketplacemodel.Item}
// @Failure 400,401,500 {object} common.Response
// @Router /item/{smart_contract}/{nft}/history [get]
func (u MarketPlaceHandler) History() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		filter := usermodel.ActivitiesFilter{}
		filter.Nft = ctx.Param("nft")
		filter.ContractAddress = ctx.Param("smart_contract")
		filter.ActivityType = ctx.Query("type")
		res, err := u.Client.ItemHistory(filter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// CollectionHistory Detail Item godoc
// @Param id path string true "collection id"
// @Param type query string false "type activities"
// @Summary get detail item
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{data=marketplacemodel.Item}
// @Failure 400,401,500 {object} common.Response
// @Router /item/collection/{id}/history [get]
func (u MarketPlaceHandler) CollectionHistory() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		filter := usermodel.ActivitiesFilter{}
		filter.CollectionId = cast.ToInt64(ctx.Param("id"))
		filter.ActivityType = ctx.Query("type")
		res, err := u.Client.CollectionHistory(filter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get detail user successfully", res))

	}
}

// CMSHideNft list godoc
// @Summary update item cms
// @Tags Item
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param user body marketplacemodel.HideNftRequest true "update request"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=marketplacemodel.Item}
// @Failure 400,401,500 {object} common.Response
// @Router /item/hide [post]
func (u MarketPlaceHandler) CMSHideNft() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		userId := cast.ToInt64(ctx.GetInt("id"))
		var rq marketplacemodel.HideNftRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		rq.UserId = userId
		rq.IsAdmin = 0
		err := u.Client.NFTHide(rq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Update status items successfully", nil))
	}
}

// CMSUnHideNft list godoc
// @Summary update item cms
// @Tags Item
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param user body marketplacemodel.HideNftRequest true "update request"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=marketplacemodel.Item}
// @Failure 400,401,500 {object} common.Response
// @Router /item/unhide [post]
func (u MarketPlaceHandler) CMSUnHideNft() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		userId := cast.ToInt64(ctx.GetInt("id"))
		var rq marketplacemodel.HideNftRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		rq.UserId = userId
		rq.IsAdmin = 0
		err := u.Client.NFTUnHide(rq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Update status items successfully", nil))
	}
}

// ExploreCollection Detail Item godoc
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param order_by query string false "order_by"
// @Param total_item_min query int false "total_item_min"
// @Param total_item_max query int false "total_item_max"
// @Param floor_price_min query int false "floor_price_min"
// @Param floor_price_max query int false "floor_price_max"
// @Summary get detail item
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{data=marketplacemodel.Item}
// @Failure 400,401,500 {object} common.Response
// @Router /item/collection/explore [get]
func (u MarketPlaceHandler) ExploreCollection() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var filter marketplacemodel.ExploreCollectionFilter
		_ = ctx.ShouldBind(&filter)

		var paging common.Paging
		_ = ctx.ShouldBind(&paging)
		paging.Process()
		res, err := u.Client.ExploreCollection(filter, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get detail user successfully", res, paging))

	}
}

// GetBannerCms list godoc
// @Summary get banner
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=[]marketplacemodel.BrandDetail}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/getbanner [get]
func (u MarketPlaceHandler) GetBannerCms() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		res, err := u.Client.ListBanner()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get list banner successfully", res))
	}
}

// CollectionProperties Create general godoc
// @Summary stats CMS
// @Tags Item
// @Accept json
// @Produce json
// @Param id path int true "collection id"
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /item/collection-properties/{id} [get]
func (u MarketPlaceHandler) CollectionProperties() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		tmpCollectionId := ctx.Param("id")
		res, err := u.Client.CollectionItemProperties(cast.ToInt64(tmpCollectionId))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get properties successfully", res))
	}
}

// ExploreNft Detail Item godoc
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param order_by query string false "order_by"
// @Param price_max query int false "price_max"
// @Param price_min query int false "price_min"
// @Param collection_id query int false "collection_id"
// @Param category_id query int false "category_id"
// @Param properties query string false "properties"
// @Param search query string false "search"
// @Summary get detail item
// @Param Authorization header string false "Insert your access token"
// @Security BearerAuth
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{data=marketplacemodel.Item}
// @Failure 400,401,500 {object} common.Response
// @Router /item/explore [get]
func (u MarketPlaceHandler) ExploreNft() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.GetInt("id")
		var filter marketplacemodel.ExploreNftFilter
		_ = ctx.ShouldBind(&filter)
		filter.UserId = cast.ToInt64(id)

		var paging common.Paging
		_ = ctx.ShouldBind(&paging)
		paging.Process()
		res, err := u.Client.ExploreNft(filter, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get detail user successfully", res, paging))

	}
}

// ItemBidDetail Create general godoc
// @Summary Item Detail Bid
// @Tags Item
// @Accept json
// @Produce json
// @Param id path int true "item id"
// @Success 200 {object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /item/bid/{id} [get]
func (u MarketPlaceHandler) ItemBidDetail() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		tmpItemId := ctx.Param("id")
		res, err := u.Client.ItemDetailBid(cast.ToInt64(tmpItemId))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get properties successfully", res))
	}
}

// ListReportCms List Report godoc
// @Summary get list report
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.ResponsePaginate{data=[]marketplacemodel.ReportTableResponse}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/list/report [get]
func (u MarketPlaceHandler) ListReportCms() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.Paging
		if errPage := ctx.ShouldBind(&paging); errPage != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", errPage.Error()))
			return
		}
		paging.Process()
		res, err := u.Client.ListReports(&paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponseWithPaging(0, "Get list brands successfully", res, paging))
	}
}

// ChangeReportCms change Report godoc
// @Summary change list report
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags CMS
// @Param update body marketplacemodel.ChangeReportRequest true "update request"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /cms/update/report [put]
func (u MarketPlaceHandler) ChangeReportCms() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq marketplacemodel.ChangeReportRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		err := u.Client.ChangeReport(&rq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Update report successfully", nil))
	}
}

// RequestReport change Report godoc
// @Summary change list report
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Tags Market Place
// @Param update body marketplacemodel.RequestReport true "update request"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/request/report [post]
func (u MarketPlaceHandler) RequestReport() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var rq marketplacemodel.RequestReport
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}

		UserId := ctx.GetInt("id")
		if UserId > 0 {
			rq.UserId = cast.ToInt64(UserId)
		}

		err := u.Client.RequestReport(&rq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Update report successfully", nil))
	}
}

// List token godoc
// @Summary list token
// @Tags Market Place
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token"
// @Param ItemGetListIdFilter body marketplacemodel.ItemGetListIdFilter true "ItemGetListIdFilter"
// @Security BearerAuth
// @Success 200 {object}{object} common.Response{data=[]string}
// @Failure 400,401,500 {object} common.Response
// @Router /market-place/get-list-item-token [post]
func (u MarketPlaceHandler) GetListItemToken() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		//paging.Process()
		id := cast.ToInt64(ctx.GetInt("id"))
		if id < 1 {
			ctx.JSON(http.StatusBadRequest, common.NewResponse(1, "Invalid request", common.ErrNotFoundMarketPlaceId))
			return
		}

		var tmpFilter marketplacemodel.ItemGetListIdFilter
		_ = ctx.ShouldBind(&tmpFilter)
		var filter marketplacemodel.ItemFilter
		filter.UserId = id
		filter.IsGetMultiId = tmpFilter.ItemId
		filter.IsGetMulti = 1
		filter.IsBuy = tmpFilter.IsBuy
		if tmpFilter.IsBuy == 1 {
			filter.Stts = append(filter.Stts, common.ITEM_ON_SALE)
			filter.UserId = tmpFilter.SellerId
		} else {
			filter.Stts = append(filter.Stts, common.ITEM_NEW, common.ITEM_READY_ON_SELL)
		}

		var paging common.Paging
		paging.Limit = tmpFilter.Number
		paging.Process()

		res, err := u.Client.ListItemToken(filter, &paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Get items successfully", res))
	}
}

// HideNftCMS list godoc
// @Summary update item cms
// @Tags Item
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param user body marketplacemodel.HideNftRequest true "update request"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=marketplacemodel.Item}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/hide [post]
func (u MarketPlaceHandler) HideNftCMS() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		userId := cast.ToInt64(ctx.GetInt("id"))
		var rq marketplacemodel.HideNftRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		rq.UserId = userId
		rq.IsAdmin = 1
		err := u.Client.NFTHide(rq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Update status items successfully", nil))
	}
}

// UnHideNftCMS list godoc
// @Summary update item cms
// @Tags Item
// @Param Authorization header string true "Insert your access token"
// @Security BearerAuth
// @Param user body marketplacemodel.HideNftRequest true "update request"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=marketplacemodel.Item}
// @Failure 400,401,500 {object} common.Response
// @Router /cms/unhide [post]
func (u MarketPlaceHandler) UnHideNftCMS() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		userId := cast.ToInt64(ctx.GetInt("id"))
		var rq marketplacemodel.HideNftRequest
		if res := utils.CommonHandle(&rq, ctx); res != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		rq.UserId = userId
		rq.IsAdmin = 1
		err := u.Client.NFTUnHide(rq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResponse(0, "Update status items successfully", nil))
	}
}
