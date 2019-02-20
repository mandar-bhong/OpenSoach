package manager

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	pcconstants "opensoach.com/prodcore/constants"
	pcmodels "opensoach.com/prodcore/models"
	"opensoach.com/prodcore/server/dbaccess"
)

func (r DocumentStoreFileSystem) Get() (error, []byte) {

	dbErr, data := dbaccess.GetDocumentByUuid(r.Data.DBContext.GetNodeDBConnection(), r.Data.DocumentID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting field operator by id.", dbErr)
		return dbErr, nil
	}

	dbRecord := *data

	if len(dbRecord) < 1 {
		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Database record not found")
		return errors.New("Database record not found"), nil
	}

	destDir := ghelper.GetExeFolder()
	destDir = filepath.Join(destDir, dbRecord[0].Location)

	byteData, readError := ioutil.ReadFile(destDir)
	if readError != nil {
		return readError, nil
	}

	return nil, byteData
}

func (r DocumentStoreFileSystem) Save() error {

	saveDocErr := r.SaveDocument()
	if saveDocErr != nil {
		return saveDocErr
	}

	saveDBErr := r.SaveDBRecord()
	if saveDBErr != nil {
		return saveDBErr
	}

	return nil

}

func (r DocumentStoreFileSystem) SaveDBRecord() error {

	file := r.Data.FileData.File["file"]

	path := ""

	for _, each := range r.Data.NestedPath {
		path = filepath.Join(each)
	}

	filedata, fileErr := file[0].Open()
	if fileErr != nil {
		return fileErr
	}

	fileBytes, readErr := ioutil.ReadAll(filedata)
	if readErr != nil {
		return readErr
	}

	mimetype := http.DetectContentType(fileBytes)

	fileUrl := "/resources/documents/"
	fileUrl = filepath.Join(fileUrl, path)

	documentStoreInsertRowModel := &pcmodels.DocumentStoreInsertRowModel{}

	documentStoreInsertRowModel.Uuid = r.Data.DocumentID

	fileUrl = filepath.Join(fileUrl, r.Data.DocumentID)

	documentStoreInsertRowModel.Name = file[0].Filename
	documentStoreInsertRowModel.DocType = mimetype
	documentStoreInsertRowModel.StorageType = r.Data.StorageType
	documentStoreInsertRowModel.URI = fileUrl
	documentStoreInsertRowModel.Persisted = pcconstants.DB_PERSISTANT
	documentStoreInsertRowModel.CpmId = r.Data.DBContext.GetCPMID()
	// documentStoreInsertRowModel.Updated_by = 1

	dbErr, _ := dbaccess.InsertDocumentData(r.Data.DBContext.GetNodeDBConnection(), documentStoreInsertRowModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding document details.", dbErr)
		return dbErr
	}

	filedata.Close()

	return nil
}

func (r DocumentStoreFileSystem) SaveDocument() error {

	file := r.Data.FileData.File["file"]

	path := ""

	for _, each := range r.Data.NestedPath {
		path = filepath.Join(each)
	}

	filedata, fileErr := file[0].Open()
	if fileErr != nil {
		return fileErr
	}

	currDir := ghelper.GetExeFolder()

	destDir := filepath.Join(currDir, filepath.Join("resources", filepath.Join("documents", path)))

	dirErr := os.MkdirAll(destDir, os.ModePerm)
	if dirErr != nil {
		return dirErr
	}

	outFile := filepath.Join(destDir, r.Data.DocumentID)

	dst, fileCreateErr := os.Create(outFile)
	if fileCreateErr != nil {
		return fileCreateErr
	}

	_, fileCopyErr := io.Copy(dst, filedata)
	if fileCopyErr != nil {
		return fileCopyErr
	}

	dst.Close()
	filedata.Close()

	return nil
}

func (r DocumentStoreFileSystem) Update() error {

	saveDocErr := r.SaveDocument()
	if saveDocErr != nil {
		return saveDocErr
	}

	saveDBErr := r.UpdateDBRecord()
	if saveDBErr != nil {
		return saveDBErr
	}

	return nil

}

func (r DocumentStoreFileSystem) UpdateDBRecord() error {

	file := r.Data.FileData.File["file"]

	path := ""

	for _, each := range r.Data.NestedPath {
		path = filepath.Join(each)
	}

	filedata, fileErr := file[0].Open()
	if fileErr != nil {
		return fileErr
	}

	fileBytes, readErr := ioutil.ReadAll(filedata)
	if readErr != nil {
		return readErr
	}

	mimetype := http.DetectContentType(fileBytes)

	fileUrl := "/resources/documents/"
	fileUrl = filepath.Join(fileUrl, path)

	documentStoreUpdateRowModel := &pcmodels.DocumentStoreUpdateRowModel{}

	fileUrl = filepath.Join(fileUrl, r.Data.DocumentID)

	documentStoreUpdateRowModel.CpmId = r.Data.DBContext.GetCPMID()
	documentStoreUpdateRowModel.Name = file[0].Filename
	documentStoreUpdateRowModel.DocType = mimetype
	documentStoreUpdateRowModel.Uuid = r.Data.DocumentID
	documentStoreUpdateRowModel.Location = fileUrl
	documentStoreUpdateRowModel.LocationType = r.Data.StorageType
	documentStoreUpdateRowModel.Persisted = pcconstants.DB_PERSISTANT

	dbErr, _ := dbaccess.UpdateDocumentData(r.Data.DBContext.GetNodeDBConnection(), documentStoreUpdateRowModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating document data.", dbErr)
		return dbErr
	}

	filedata.Close()

	return nil
}
