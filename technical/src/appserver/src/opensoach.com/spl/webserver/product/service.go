package product

import (
	gmodels "opensoach.com/models"
	lmodels "opensoach.com/spl/models"
	repo "opensoach.com/spl/repository"
	"opensoach.com/spl/webserver/product/dbaccess"
)

// Implement service with empty struct
type ProductService struct {
}

func (ProductService) GetProducts(pExeContext *gmodels.ExecutionContext) (bool, interface{}) {

	err, data := dbaccess.GetUserProducts(repo.Instance().MasterDBConnection, pExeContext.SessionInfo.UserID)

	if err != nil {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	return true, data
}

func (ProductService) SelectProduct(pExeContext *gmodels.ExecutionContext) (bool, interface{}) {

	selectProdReq := pExeContext.Request.(*lmodels.APIProductSelectRequest)

	pExeContext.SelectedProduct = selectProdReq.ProductID

	//mb.Publish("ProductSelection", []byte("User selected product"))

	return true, nil
}
