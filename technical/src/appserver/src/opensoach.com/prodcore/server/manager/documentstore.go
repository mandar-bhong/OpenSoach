package manager

import (
	ghelper "opensoach.com/core/helper"
	pcconstants "opensoach.com/prodcore/constants"
	pcmodels "opensoach.com/prodcore/models"
)

type DocumentStoreFileSystem struct {
	Data *pcmodels.DocumentStoreDataModel
}

func DocumentStoreSave(data *pcmodels.DocumentStoreDataModel) (error, string) {

	var document pcmodels.IDocumentStore

	data.DocumentID = ghelper.GenerateUUID()

	switch data.StorageType {
	case pcconstants.DB_DOCUMENT_STORAGE_TYPE_FILE_SYSTEM: //File system
		document = DocumentStoreFileSystem{
			Data: data,
		}
		break
	}

	saveErr := document.Save()
	if saveErr != nil {
		return saveErr, ""
	}

	return nil, data.DocumentID

}

func DocumentStoreGet(data *pcmodels.DocumentStoreDataModel) (error, pcmodels.DocumentData) {

	var document pcmodels.IDocumentStore

	switch data.StorageType {
	case pcconstants.DB_DOCUMENT_STORAGE_TYPE_FILE_SYSTEM: //File system
		document = DocumentStoreFileSystem{
			Data: data,
		}
		break
	}

	getErr, documentData := document.Get()
	if getErr != nil {
		return getErr, pcmodels.DocumentData{}
	}

	return nil, documentData

}

func DocumentStoreUpdate(data *pcmodels.DocumentStoreDataModel) (error, string) {

	var document pcmodels.IDocumentStore

	switch data.StorageType {
	case pcconstants.DB_DOCUMENT_STORAGE_TYPE_FILE_SYSTEM: //File system
		document = DocumentStoreFileSystem{
			Data: data,
		}
		break
	}

	saveErr := document.Update()
	if saveErr != nil {
		return saveErr, ""
	}

	return nil, data.DocumentID

}
