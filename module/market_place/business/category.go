package business

import (
	"api-openlive/common"
	marketplacemodel "api-openlive/module/market_place/model"
)

func (i MarketPlaceController) CreateCategory(rq *marketplacemodel.CreateCategoryRequest) (*marketplacemodel.Category, error) {
	cate := &marketplacemodel.Category{
		Name:        rq.Name,
		Description: rq.Description,
	}
	err := i.Store.InsertCategory(cate)
	if err != nil {
		return nil, err
	}
	return cate, nil
}

func (i MarketPlaceController) UpdateCategory(rq *marketplacemodel.UpdateCategoryRequest) (*marketplacemodel.Category, error) {
	cate := &marketplacemodel.Category{
		Name:        rq.Name,
		Description: rq.Description,
		Id:          rq.Id,
	}
	err := i.Store.UpdateCategory(cate)
	if err != nil {
		return nil, err
	}

	return cate, nil
}

func (i MarketPlaceController) ListCategory(filter marketplacemodel.CategoryFilter) ([]*marketplacemodel.Category, error) {
	data, err := i.Store.ListCategories(&filter)
	if err != nil {
		return nil, err
	}

	sortResult := make([]*marketplacemodel.Category, 6)
	for _, v := range data {
		switch v.Id {
		case common.ART:
			sortResult[0] = v
		case common.GAME:
			sortResult[1] = v
		case common.MUSIC:
			sortResult[2] = v
		case common.PHOTOGRAPHY:
			sortResult[3] = v
		case common.VIDEO:
			sortResult[5] = v
		case common.SPORT:
			sortResult[4] = v
		}
	}
	return sortResult, nil
}
