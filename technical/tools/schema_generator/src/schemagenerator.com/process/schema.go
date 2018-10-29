package process

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"schemagenerator.com/ReadConfig"
	"github.com/stoewer/go-strcase"
	//	"opensoach.com/core/manager/db"
)

type TableSchema struct {
	ColumnName   string `db:"COLUMN_NAME" json:"key"`
	DataType     string `db:"DATA_TYPE" json:"key"`
	IsNullable   string `db:"IS_NULLABLE" json:"key"`
	IsPrimaryKey string `db:"COLUMN_KEY" json:"key"`
	IsAutoIncr   string `db:"EXTRA" json:"key"`
}

type TableName struct {
	TableName string `db:"TABLE_NAME" json:"key"`
}

type PreparedData struct {
	TableName          string
	StructName         string
	StructString       string
	SelectQuery        string
	SelectByIdQuery    string
	InsertQuery        string
	UpdateQuery        string
	SelectProc         string
	SelectByIdProc     string
	InsertProc         string
	UpdateProc         string
	SelectAllFunc      string
	SelectByIdFunc     string
	SelectByFilterFunc string
	InsertFunc         string
	UpdateFunc         string
	DeleteFunc         string
}

func CreateSchema(c ReadConfig.Config) {
	tableName := []TableName{}

	dbengine := sqlx.MustConnect(c.Driver, c.ConnectionString)

	err := dbengine.Select(&tableName, "select TABLE_NAME from information_schema.TABLES where TABLE_SCHEMA = ?", c.DBName)
	if err != nil {
		log.Fatal(err)
	}

	data := []PreparedData{}

	for j := 0; j < len(tableName); j++ {
		tblName := tableName[j].TableName
		tableSchema := CreateTableSchema(c.DBName, tblName, dbengine)
		StructName := getPropStructName(tblName)
		StructString := CreateStruct(tblName, tableSchema)
		SelectQuery := getSelectQuery(tblName, tableSchema)
		SelectByIdQuery := getSelectByIdQuery(tblName, tableSchema)
		InsertQuery := getInsertQuery(tblName, tableSchema)
		UpdateQuery := getUpdateQuery(tblName, tableSchema)
		SelectProc := createStoredProcedure(tblName, SelectQuery, "select")
		SelectByIdProc := createStoredProcedure(tblName, SelectByIdQuery, "selectbyid")
		InsertProc := createStoredProcedure(tblName, InsertQuery, "insert")
		UpdateProc := createStoredProcedure(tblName, UpdateQuery, "update")

		data = append(data, PreparedData{
			TableName:       tblName,
			StructName:      StructName,
			StructString:    StructString,
			SelectQuery:     SelectQuery,
			SelectByIdQuery: SelectByIdQuery,
			InsertQuery:     InsertQuery,
			UpdateQuery:     UpdateQuery,
			SelectProc:      SelectProc,
			SelectByIdProc:  SelectByIdProc,
			InsertProc:      InsertProc,
			UpdateProc:      UpdateProc},
		)
	}

	for i := 0; i < len(data); i++ {
		SelectAllFunc := CreateSelectAllFunc(data[i])
		SelectByIdFunc := CreateSelectByIdFunc(data[i])
		SelectByFilterFunc := CreateSelectByFilterFunc(data[i])
		InsertFunc := CreateInsertFunc(data[i])
		UpdateFunc := CreateUpdateFunc(data[i])
		DeleteFunc := CreateDeleteFunc(data[i])
		data[i].SelectAllFunc = SelectAllFunc
		data[i].SelectByIdFunc = SelectByIdFunc
		data[i].SelectByFilterFunc = SelectByFilterFunc
		data[i].InsertFunc = InsertFunc
		data[i].UpdateFunc = UpdateFunc
		data[i].DeleteFunc = DeleteFunc
		WriteAll(data[i])
		CreateDbAccessFile(data[i].TableName)
		CreateGmodelsFile(data[i].TableName)
	}

}

func WriteAll(data PreparedData) {
	dir := "files/" + data.TableName
	CreateDirIfNotExist(dir)
	path := "files/" + data.TableName + "/" + data.TableName
	structPath := "files/StructModels.txt"
	queryPath := "files/Query.go"
	storProcPath := path + "StrProc.txt"
	funcPath := path + "Func.txt"
	WriteToFile(data.StructString, structPath)
	WriteToFile(data.SelectQuery, queryPath)
	WriteToFile(data.SelectByIdQuery, queryPath)
	WriteToFile(data.InsertQuery, queryPath)
	WriteToFile(data.UpdateQuery, queryPath)
	WriteToFile(data.SelectProc, storProcPath)
	WriteToFile(data.SelectByIdProc, storProcPath)
	WriteToFile(data.InsertProc, storProcPath)
	WriteToFile(data.UpdateProc, storProcPath)
	WriteToFile(data.SelectAllFunc, funcPath)
	WriteToFile(data.SelectByIdFunc, funcPath)
	WriteToFile(data.SelectByFilterFunc, funcPath)
	WriteToFile(data.InsertFunc, funcPath)
	WriteToFile(data.UpdateFunc, funcPath)
	WriteToFile(data.DeleteFunc, funcPath)
}

func CreateTableSchema(dbname string, tblName string, dbengine *sqlx.DB) []TableSchema {
	tableSchema := []TableSchema{}
	err := dbengine.Select(&tableSchema, "select COLUMN_NAME,DATA_TYPE,IS_NULLABLE,COLUMN_KEY,EXTRA from information_schema.COLUMNS where TABLE_SCHEMA = ? and table_name = ?", dbname, tblName)
	if err != nil {
		log.Fatal(err)
	}
	return tableSchema
}

func CreateStruct(tblName string, tableSchema []TableSchema) string {
	tablName := tblName
	propStructName := getPropStructName(tablName)

	structString := "type " + propStructName + "  struct { \n"
	for i := 0; i < len(tableSchema); i++ {
		tblCol := tableSchema[i]
		structString = structString + createStructRow(tblCol.ColumnName, tblCol.DataType, tblCol.IsNullable, tblCol.IsPrimaryKey, tblCol.IsAutoIncr) + "\n"
	}
	structString = structString + "}"

	return structString
}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func WriteToFile(str, filePath string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(str + "\n\n")
	if err != nil {
		log.Fatal(err)
	}
}

func createStructRow(dbColumnName string, dataType string, isNullable string, isPrimaryKey string, isAutoIncr string) string {
	structRow := ""

	goDT := getGolangDataType(dataType)
	propName := getPropNameFromDBName(dbColumnName)

	if isNullable == "YES" {
		goDT = "*" + goDT
	}

	jsonname := strings.Replace(dbColumnName, "_", "", -1)
	jsonname = strings.ToLower(jsonname)

	if isPrimaryKey == "PRI" && isAutoIncr == "auto_increment" {
		structRow = propName + "   " + goDT + "   " + fmt.Sprintf("`db:\"%s\" dbattr:\"%s,%s\"  json:\"%s\"", dbColumnName, "pri", "auto", jsonname)
	} else if isPrimaryKey == "PRI" && isAutoIncr == "" {
		structRow = propName + "   " + goDT + "   " + fmt.Sprintf("`db:\"%s\" dbattr:\"%s\"  json:\"%s\"", dbColumnName, "pri", jsonname)
	} else if isAutoIncr == "auto_increment" && isPrimaryKey == "" {
		structRow = propName + "   " + goDT + "   " + fmt.Sprintf("` db:\"%s\" dbattr:\"%s\" json:\"%s\"", dbColumnName, "auto", jsonname)
	} else {
		structRow = propName + "   " + goDT + "   " + fmt.Sprintf("`db:\"%s\" json:\"%s\"", dbColumnName, jsonname)
	}

	structRow = structRow + "`"

	return structRow
}

func getPropStructName(dbStrucName string) string {
	name := "DB"
	dbStrucName = strcase.UpperCamelCase(dbStrucName)
	dbStrucName = strings.Replace(dbStrucName, "Tbl", "Table", -1)
	name = name + dbStrucName + "RowModel"
	return name
}

func getGolangDataType(dbType string) string {
	switch dbType {
	case "varchar":
		return "string"
	case "int":
	case "tinyint":
		return "int"
	case "smallint":
		return "int"
	case "mediumint":
		return "int"
	case "timestamp":
		return "time.Time"
	case "datetime":
		return "time.Time"
	case "date":
		return "time.Time"
	case "text":
		return "string"
	case "decimal":
		return "float64"
	case "year":
		return "time.Time"
	case "char":
		return "string"
	case "json":
		return "string"
	default:
		fmt.Printf("Unable to find golang datatype for : %s", dbType)
		break
	}
	return "int64"
}

func getPropNameFromDBName(dbColumnName string) string {
	return strcase.UpperCamelCase(dbColumnName)
}

func getPropQueryName(queryConst, queryType string) string {
	queryConst = strings.Replace(queryConst, "tbl", "Table", -1)
	queryname := "const QUERY_"
	queryConst = strings.ToUpper(queryConst)
	queryConst = strings.Replace(queryConst, "tbl", "Table", -1)
	queryname = queryname + queryConst + "_" + queryType
	return queryname
}

func getSelectQuery(tablename string, tableSchema []TableSchema) string {
	tblName := tablename
	queryName := getPropQueryName(tblName, "SELECT_All")
	query := "SELECT "
	for i := 0; i < len(tableSchema); i++ {
		tblCol := tableSchema[i]

		query = query + tblCol.ColumnName + ","
	}
	query = strings.TrimRight(query, ",")
	query = query + " FROM " + tablename
	query = queryName + " = " + "\"" + query + "\""
	return query
}

func getSelectByIdQuery(tablename string, tableSchema []TableSchema) string {
	tblName := tablename
	queryName := getPropQueryName(tblName, "SELECT_BY_ID")
	primaryKey := ""

	for i := 0; i < len(tableSchema); i++ {
		if tableSchema[i].IsPrimaryKey == "PRI" {
			primaryKey = tableSchema[i].ColumnName
		}
	}

	query := "SELECT "
	for i := 0; i < len(tableSchema); i++ {
		tblCol := tableSchema[i]
		if tblCol.IsPrimaryKey != "PRI" {
			query = query + tblCol.ColumnName + ","
		}
	}
	query = strings.TrimRight(query, ",")
	query = query + " FROM " + tablename
	query = query + " WHERE " + primaryKey + " = :" + primaryKey
	query = queryName + " = " + "\"" + query + "\""
	return query
}

func getInsertQuery(tablename string, tableSchema []TableSchema) string {
	tblName := tablename
	queryName := getPropQueryName(tblName, "INSERT")
	query := "INSERT INTO " + tablename + " ("
	for i := 0; i < len(tableSchema); i++ {
		tblCol := tableSchema[i]
		query = query + tblCol.ColumnName + ","
	}
	query = strings.TrimRight(query, ",")
	query = query + ") values ( "
	for i := 0; i < len(tableSchema); i++ {
		tblCol := tableSchema[i]
		query = query + ":" + tblCol.ColumnName + ","
	}
	query = strings.TrimRight(query, ",")
	query = query + ")"
	query = queryName + " = " + "\"" + query + "\""
	return query
}

func getUpdateQuery(tablename string, tableSchema []TableSchema) string {
	tblName := tablename
	queryName := getPropQueryName(tblName, "UPDATE")
	primaryKey := ""

	for i := 0; i < len(tableSchema); i++ {
		if tableSchema[i].IsPrimaryKey == "PRI" {
			primaryKey = tableSchema[i].ColumnName
		}
	}

	query := "UPDATE " + tablename + " SET "
	for i := 0; i < len(tableSchema); i++ {
		tblCol := tableSchema[i]
		if tblCol.IsPrimaryKey != "PRI" {
			query = query + tblCol.ColumnName + " = " + ":" + tblCol.ColumnName + ", "
		}
	}
	query = strings.TrimRight(query, ", ")
	query = query + " WHERE " + primaryKey + " = :" + primaryKey
	query = queryName + " = " + "\"" + query + "\""
	return query
}

func createStoredProcedure(tblName, query, name string) string {
	switch name {
	case "select":
		name = "sp_" + tblName + "_select_all"
	case "selectbyid":
		name = "sp_" + tblName + "_select_by_id"
	case "insert":
		name = "sp_" + tblName + "_insert_all"
	case "update":
		name = "sp_" + tblName + "_update_by_id"
	}
	strProc := "CREATE PROCEDURE `" + name + "`()\n"
	strProc = strProc + "COMMENT '' \nBEGIN \n" + query + ";\nEND"
	return strProc
}

func GetPropFuncName(funcName string) string {
	funcName = strcase.UpperCamelCase(funcName)
	funcName = strings.Replace(funcName, "Tbl", "Table", -1)
	return funcName
}

func CreateInsertFunc(tblData PreparedData) string {

	functnName := tblData.TableName
	functnName = GetPropFuncName(functnName)

	funcString := "func " + functnName + "Insert"

	funcString = funcString + "(dbConn string, insrtStruct lmodels." + tblData.StructName

	funcString = funcString + ") (error, int64) {\n"

	funcString = funcString + "insDBCtx:= dbmgr.InsertContext{}\n"

	funcString = funcString + "insDBCtx.DBConnection = dbConn\n"

	funcString = funcString + "insDBCtx.Args = insrtStruct\n"

	funcString = funcString + "insDBCtx.QueryType=dbmgr.AutoQuery\ninsDBCtx.TableName=" + " \"" + tblData.TableName + "\""

	funcString = funcString + "\ninsertErr := insDBCtx.Insert()\n"

	funcString = funcString + "if insertErr != nil {\nreturn insertErr, 0\n}"

	funcString = funcString + "\nreturn nil, insDBCtx.InsertID\n}"

	return funcString
}

func CreateUpdateFunc(tblData PreparedData) string {

	functnName := tblData.TableName
	functnName = GetPropFuncName(functnName)

	funcString := "func " + functnName + "Update"

	funcString = funcString + "(dbconn string, updtStruct lmodels." + tblData.StructName

	funcString = funcString + ") (error, int64) {\n"

	funcString = funcString + "updateCtx:= dbmgr.UpdateDeleteContext{}\n"

	funcString = funcString + "updateCtx.DBConnection = dbConn\n"

	funcString = funcString + "updateCtx.Args = updtStruct\n"

	funcString = funcString + "updateCtx.QueryType=dbmgr.AutoQuery\nupdateCtx.TableName=" + " \"" + tblData.TableName + "\""

	funcString = funcString + "\nupdateErr := updateCtx.Update()\n"

	funcString = funcString + "if updateErr != nil {\nreturn updateErr, 0\n}"

	funcString = funcString + "\nreturn nil, updateCtx.AffectedRows\n}"

	return funcString
}

func CreateDeleteFunc(tblData PreparedData) string {

	functnName := tblData.TableName
	functnName = GetPropFuncName(functnName)

	funcString := "func " + functnName + "Delete"

	funcString = funcString + "(dbconn string, deltStruct lmodels." + tblData.StructName

	funcString = funcString + ") (error, int64) {\n"

	funcString = funcString + "deleteCtx:= dbmgr.UpdateDeleteContext{}\n"

	funcString = funcString + "deleteCtx.DBConnection = dbConn\n"

	funcString = funcString + "deleteCtx.Args = deltStruct\n"

	funcString = funcString + "deleteCtx.QueryType=dbmgr.AutoQuery\ndeleteCtx.TableName=" + " \"" + tblData.TableName + "\""

	funcString = funcString + "\ndeleteErr := deleteCtx.Delete()\n"

	funcString = funcString + "if deleteErr != nil {\nreturn deleteErr, 0\n}"

	funcString = funcString + "\nreturn nil, deleteCtx.AffectedRows\n}"

	return funcString
}

func CreateSelectAllFunc(tblData PreparedData) string {

	functnName := tblData.TableName
	functnName = GetPropFuncName(functnName)

	funcString := "func " + functnName + "SelectAll"

	funcString = funcString + "(conn string, destStruct *[]gmodels." + tblData.StructName

	funcString = funcString + ") (error, interface{}) {\n"

	funcString = funcString + "err, sqlxDB := getDbEngine(conn)\nif err != nil {\nreturn err, 0\n}\n"

	funcString = funcString + "selectCtx:= db.SelectContext{}\n"

	funcString = funcString + "selectCtx.Engine = sqlxDB\nselectCtx.Dest = destStruct\n"

	funcString = funcString + "selectCtx.Type=db.AutoQuery\nselectCtx.TableName=" + " \"" + tblData.TableName + "\""

	funcString = funcString + "\nselectErr := selectCtx.SelectAll()\n"

	funcString = funcString + "if selectErr != nil {\nreturn selectErr, 0\n}"

	funcString = funcString + "\nreturn nil, selectCtx.Dest\n}"

	return funcString
}

func CreateSelectByIdFunc(tblData PreparedData) string {

	functnName := tblData.TableName
	functnName = GetPropFuncName(functnName)

	funcString := "func " + functnName + "SelectById"

	funcString = funcString + "(conn string, destStruct *[]gmodels." + tblData.StructName

	funcString = funcString + ", Id int) (error, interface{}) {\n"

	funcString = funcString + "err, sqlxDB := getDbEngine(conn)\nif err != nil {\nreturn err, 0\n}\n"

	funcString = funcString + "selectCtx:= db.SelectContext{}\n"

	funcString = funcString + "selectCtx.Engine = sqlxDB\nselectCtx.Dest = destStruct\n"

	funcString = funcString + "selectCtx.Type=db.AutoQuery\nselectCtx.TableName=" + " \"" + tblData.TableName + "\""

	funcString = funcString + "\nselectErr := selectCtx.SelectById(Id)\n"

	funcString = funcString + "if selectErr != nil {\nreturn selectErr, 0\n}"

	funcString = funcString + "\nreturn nil, selectCtx.Dest\n}"

	return funcString
}

func CreateSelectByFilterFunc(tblData PreparedData) string {

	functnName := tblData.TableName
	functnName = GetPropFuncName(functnName)

	funcString := "func " + functnName + "SelectByFilter"

	funcString = funcString + "(conn string, destStruct *[]gmodels." + tblData.StructName

	funcString = funcString + ", filter gmodels." + tblData.StructName + ", args ...string) (error, interface{}) {\n"

	funcString = funcString + "err, sqlxDB := getDbEngine(conn)\nif err != nil {\nreturn err, 0\n}\n"

	funcString = funcString + "selectCtx:= db.SelectContext{}\n"

	funcString = funcString + "selectCtx.Engine = sqlxDB\nselectCtx.Dest = destStruct\n"

	funcString = funcString + "selectCtx.Type=db.AutoQuery\nselectCtx.TableName=" + " \"" + tblData.TableName + "\""

	funcString = funcString + "\nselectErr := selectCtx.SelectByFilter(filter, args...)\n"

	funcString = funcString + "if selectErr != nil {\nreturn selectErr, 0\n}"

	funcString = funcString + "\nreturn nil, selectCtx.Dest\n}"

	return funcString
}

func CreateDbAccessFile(tablename string) {
	dbAccessFileName := tablename
	dir := "files/" + tablename + "/dbaccess"
	dbAccessFileName = strcase.UpperCamelCase(dbAccessFileName)
	dbAccessFileName = strings.Replace(dbAccessFileName, "_", "", -1)
	dbAccessFileName = strings.Replace(dbAccessFileName, "Tbl", "", -1)
	dbAccessFileName = "db" + dbAccessFileName + "Access.go"
	dbaccessfilepath := "files/" + tablename + "/dbaccess/" + dbAccessFileName
	funcFilePath := "files/" + tablename + "/" + tablename + "Func.txt"
	CreateDirIfNotExist(dir)

	dbaccessfile, err := os.OpenFile(dbaccessfilepath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer dbaccessfile.Close()

	dbaccesstmpltFile, err := os.Open("dbaccesstmplt.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dbaccesstmpltFile.Close()

	dbaccesstmpltFileData, err := ioutil.ReadAll(dbaccesstmpltFile)
	if err != nil {
		log.Fatal(err)
	}

	funcFile, err := os.Open(funcFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer funcFile.Close()

	funcFileData, err := ioutil.ReadAll(funcFile)
	if err != nil {
		log.Fatal(err)
	}

	s := strings.Replace(string(dbaccesstmpltFileData), "$$funcs$$", string(funcFileData), -1)

	_, err = dbaccessfile.WriteString(s)
	if err != nil {
		log.Fatal(err)
	}

}

func CreateGmodelsFile(tablename string) {

	dir := "files/gmodels"
	gmodelsfilepath := "files/gmodels/gmodels.go"
	structFilePath := "files/StructModels.txt"
	CreateDirIfNotExist(dir)

	gmodelsfile, err := os.OpenFile(gmodelsfilepath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer gmodelsfile.Close()

	gmodelstmpltFile, err := os.Open("gmodelstmplt.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer gmodelstmpltFile.Close()

	gmodelstmpltFileData, err := ioutil.ReadAll(gmodelstmpltFile)
	if err != nil {
		log.Fatal(err)
	}

	structFile, err := os.Open(structFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer structFile.Close()

	structFileData, err := ioutil.ReadAll(structFile)
	if err != nil {
		log.Fatal(err)
	}

	structFileDataStr := string(structFileData)

	str := ""

	if strings.Contains(structFileDataStr, "time.Time") {
		str = strings.Replace(string(gmodelstmpltFileData), "$$importTime$$", " \"time\"", 1)
	} else {
		str = strings.Replace(string(gmodelstmpltFileData), "$$importTime$$", "", 1)
	}

	str = strings.Replace(str, "$$models$$", structFileDataStr, -1)

	_, err = gmodelsfile.WriteString(str)
	if err != nil {
		log.Fatal(err)
	}
}
