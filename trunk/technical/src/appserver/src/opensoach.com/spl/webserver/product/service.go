package product

import (
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	lmodels "opensoach.com/spl/models"
	repo "opensoach.com/spl/repository"
	"opensoach.com/spl/webserver/product/dbaccess"
	lhelper "opensoach.com/spl/webserver/product/helper"
)

// Implement service with empty struct
type ProductService struct {
}

func (ProductService) GetProducts(pExeContext *gmodels.ExecutionContext) (bool, interface{}) {

	err, data := dbaccess.GetUserProducts(repo.Instance().Context.Dynamic.DB, pExeContext.SessionInfo.UserID)

	if err != nil {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	return true, data
}

func (ProductService) SelectProduct(pExeContext *gmodels.ExecutionContext) (bool, interface{}) {

	selectProdReq := pExeContext.Request.(*lmodels.APIProductSelectRequest)

	pExeContext.SelectedCustomerProduct = selectProdReq.ProductID

	isSuccess, dbProductBriefRowModel := getCMPDetails(selectProdReq.ProductID)

	if !isSuccess {
		respErr := gmodels.APIResponseError{}
		respErr.Code = gmodels.MOD_OPER_ERR_SERVER
		return false, respErr
	}

	return true, dbProductBriefRowModel
}

func getCMPDetails(cpmid int64) (bool, *lmodels.DBProductBriefRowModel) {

	isGetSuccess, dbProductRowModel := lhelper.CacheGetCPMDetails(repo.Instance().Context.Dynamic.Cache, gmodels.CACHE_KEY_PREFIX_CPM_ID+string(cpmid))

	if isGetSuccess == false { //Unable to get form Cache, fetching from database
		err, data := dbaccess.GetCustomerProductDetails(repo.Instance().Context.Dynamic.DB, cpmid)

		if err != nil {
			logger.Context().WithField("CPM ID", cpmid).LogError("PRODUCT", "Unable to fetch product details", err)
			return false, nil
		}

		lhelper.CacheSetCPMDetails(gmodels.CACHE_KEY_PREFIX_CPM_ID+string(cpmid), data, repo.Instance().Context.Dynamic.Cache)

		return true, data
	}

	return true, dbProductRowModel
}
