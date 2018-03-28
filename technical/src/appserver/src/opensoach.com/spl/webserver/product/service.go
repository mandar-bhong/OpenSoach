package product

import (
	gmodels "opensoach.com/models"
	lmodels "opensoach.com/spl/models"
	repo "opensoach.com/spl/repository"
	"opensoach.com/spl/webserver/product/dbaccess"
	//lhelper "opensoach.com/spl/webserver/product/helper"
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

	//lhelper.CacheGetCPMDetails(repo.Instance().Context.Dynamic.Cache, selectProdReq.ProductID)

	//	isCacheSuccess, cpmData := repo.Instance().Context.Dynamic.Cache.Get(gmodels.CACHE_KEY_PREFIX_CPM_ID + string(selectProdReq.ProductID))

	//mb.Publish("ProductSelection", []byte("User selected product"))

	return true, nil
}

func getCMPDetails() {

	//	isGetSuccess, dbProductRowModel := lhelper.CacheGetCPMDetails(repo.Instance().Context.Dynamic.Cache, gmodels.CACHE_KEY_PREFIX_CPM_ID+string(selectProdReq.ProductID))

	//	if isGetSuccess {

	//		//dbaccess.GetUserProducts()

	//	}

}
