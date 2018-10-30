package db

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var spnameparams map[string][]string
var procQueryMap map[string]string
var connectionDBEngine map[string]*sqlx.DB

func init() {
	spnameparams = make(map[string][]string, 0)
	procQueryMap = make(map[string]string, 0)
	connectionDBEngine = make(map[string]*sqlx.DB, 0)
}

//InsertProcContext will be used to insert record with Insert method
//InsertedID will be filled if table have autoincriment column

type Type int

const (
	AutoQuery       Type = 1
	Query                = 2
	StoredProcedure      = 3
)

type baseContext struct {
	TableName string
	QueryType Type
	Query     string
}

type context struct {
	baseContext
	DBConnection string
}

type contextTx struct {
	baseContext
	Tx *sqlx.Tx
}

type InsertContext struct {
	context
	Args     interface{}
	InsertID int64
}

type UpdateDeleteContext struct {
	context
	Args         interface{}
	AffectedRows int64
}

type SelectContext struct {
	context
	Dest interface{}
}

type InsertTxContext struct {
	contextTx
	Args     interface{}
	InsertID int64
}

type UpdateDeleteTxContext struct {
	contextTx
	Args         interface{}
	AffectedRows int64
}

type SelectTxContext struct {
	contextTx
	Dest interface{}
}

type SPDiscoveryParam struct {
	SPName    string
	ParamName []string
}

func (spc *contextTx) GetTransaction(dbconn string) (error, *sqlx.Tx) {

	dbEng, err := sqlx.Connect("mysql", dbconn)
	if err != nil {
		return err, nil
	}
	tx, err := dbEng.Beginx()

	return err, tx
}

func (spc *InsertContext) Insert() error {

	switch spc.QueryType {
	case AutoQuery: //For text based query
		//return spc.Engine.Exec(spc.Query, spc.Args)

		query := GetInsertDynamicQuery(spc.TableName, spc.Args)

		dbConnErr, engine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		id, err := engine.NamedExec(query, spc.Args)

		if err != nil {
			return err
		}

		spc.InsertID, _ = id.LastInsertId()
		return err

	case Query:

		dbConnErr, engine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		id, err := engine.NamedExec(spc.Query, spc.Args)
		if err != nil {
			return err
		}
		spc.InsertID, _ = id.LastInsertId()
		return err

	case StoredProcedure: //For stored Procedure
		var lastinsertid int64

		dbConnErr, engine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		tx, txerr := engine.Beginx()

		if txerr != nil {
			return txerr
		}

		procGetErr, spQuery := getProcQuery(tx, spc.Query)
		fmt.Println("SPQuery")
		fmt.Println(spQuery)

		if procGetErr != nil {
			return procGetErr
		}

		_, execErr := tx.NamedExec(spQuery, spc.Args)

		if execErr != nil {
			return execErr
		}

		tx.Get(&lastinsertid, "SELECT Last_Insert_ID()")

		tx.Commit()

		spc.InsertID = lastinsertid

		return nil

	}

	return nil

}

func (spc *UpdateDeleteContext) UpdateByFilter(args ...string) error {

	switch spc.QueryType {
	case AutoQuery:
		query := GetUpdateByFilterDynamicQuery(spc.TableName, spc.Args, args...)

		dbConnErr, dbEngine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		id, err := dbEngine.NamedExec(query, spc.Args)
		if err != nil {
			return err
		}
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case Query:
		dbConnErr, dbEngine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		id, err := dbEngine.NamedExec(spc.Query, spc.Args)
		if err != nil {
			return err
		}
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case StoredProcedure:

		break

	}

	return nil
}

func (spc *UpdateDeleteContext) Update() error {

	switch spc.QueryType {
	case AutoQuery:
		queryErr, query := GetUpdateDynamicQuery(spc.TableName, spc.Args)
		if queryErr != nil {
			return queryErr
		}

		dbConnErr, dbEngine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		id, err := dbEngine.NamedExec(query, spc.Args)
		if err != nil {
			return err
		}
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case Query:
		dbConnErr, dbEngine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		id, err := dbEngine.NamedExec(spc.Query, spc.Args)
		if err != nil {
			return err
		}
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case StoredProcedure:

		dbConnErr, dbEngine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		tx, txBeginErr := dbEngine.Beginx()

		if txBeginErr != nil {
			return txBeginErr
		}

		procGetErr, spQuery := getProcQuery(tx, spc.Query)

		if procGetErr != nil {
			return procGetErr
		}

		execResult, execErr := tx.NamedExec(spQuery, spc.Args)

		if execErr != nil {
			return execErr
		}

		rowAffectedCount, rowAffectedErr := execResult.RowsAffected()

		if rowAffectedErr != nil {
			tx.Rollback()
			return rowAffectedErr
		}

		spc.AffectedRows = rowAffectedCount

		tx.Commit()

		return nil

	}

	return nil
}

func (spc *UpdateDeleteContext) Delete() error {

	switch spc.QueryType {
	case AutoQuery:
		queryErr, query := GetDeleteDynamicQuery(spc.TableName, spc.Args)
		if queryErr != nil {
			return queryErr
		}

		dbConnErr, dbEngine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		id, err := dbEngine.NamedExec(query, spc.Args)
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case Query:
		dbConnErr, dbEngine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		id, err := dbEngine.NamedExec(spc.Query, spc.Args)
		if err != nil {
			return err
		}
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case StoredProcedure:

		dbConnErr, dbEngine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		tx, txBeginErr := dbEngine.Beginx()

		if txBeginErr != nil {
			return txBeginErr
		}

		procGetErr, spQuery := getProcQuery(tx, spc.Query)

		if procGetErr != nil {
			return procGetErr
		}

		execResult, execErr := tx.NamedExec(spQuery, spc.Args)

		if execErr != nil {
			return execErr
		}

		rowAffectedCount, rowAffectedErr := execResult.RowsAffected()

		if rowAffectedErr != nil {
			tx.Rollback()
			return rowAffectedErr
		}

		spc.AffectedRows = rowAffectedCount

		tx.Commit()

		return nil

	}

	return nil

}

func (spc *SelectContext) SelectAll() error {
	switch spc.QueryType {
	case AutoQuery:
		query := GetSelectAllDynamicQuery(spc.TableName, spc.Dest)

		dbConnErr, dbEngine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		err := dbEngine.Select(spc.Dest, query)
		return err
	case Query:

		dbConnErr, dbEngine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		err := dbEngine.Select(spc.Dest, spc.Query)
		return err

	case StoredProcedure:

	}
	return nil
}

func (spc *SelectContext) SelectById(arg int64) error {
	switch spc.QueryType {
	case AutoQuery:
		queryErr, query := GetSelectByIdDynamicQuery(spc.TableName, spc.Dest)
		if queryErr != nil {
			return queryErr
		}

		dbConnErr, dbEngine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		err := dbEngine.Select(spc.Dest, query, arg)
		return err
	case Query:

		dbConnErr, dbEngine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		err := dbEngine.Select(spc.Dest, spc.Query, arg)
		return err

	case StoredProcedure:
		break

	}
	return nil
}

func (spc *SelectContext) SelectByFilter(filter interface{}, args ...string) error {
	switch spc.QueryType {
	case AutoQuery:
		query := GetSelectByFilterDynamicQuery(spc.TableName, filter, args...)
		values := GetFilterValues(spc.TableName, filter, args...)

		dbConnErr, dbEngine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		err := dbEngine.Select(spc.Dest, query, values...)
		return err

	case Query:
		values := GetFilterValues(spc.TableName, filter, args...)

		dbConnErr, dbEngine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		err := dbEngine.Select(spc.Dest, spc.Query, values...)
		return err

	case StoredProcedure:
		break
	}
	return nil
}

func (spc *SelectContext) Select(args ...interface{}) error {

	switch spc.QueryType {
	case AutoQuery:
		return errors.New("AutoQuery is not supported for Select method")
	case Query:
		dbConnErr, dbEngine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		err := dbEngine.Select(spc.Dest, spc.Query, args...)
		return err

	case StoredProcedure:
		spQuery := ""

		for i := 0; i < len(args); i++ {
			spQuery = spQuery + fmt.Sprintf("%#v", args[i]) + ","
		}

		spQuery = strings.TrimRight(spQuery, ",")

		spQuery = "call " + spc.Query + "(" + spQuery + ")"

		dbConnErr, dbEngine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		err := dbEngine.Select(spc.Dest, spQuery)

		return err
	}

	return errors.New(fmt.Sprintf("QueryType Not Set. Got %d", spc.QueryType))
}

func (spc *SelectContext) SelectToMap(args ...interface{}) (error, []map[string]interface{}) {

	switch spc.QueryType {
	case AutoQuery:
		return errors.New("AutoQuery is not supported for Select method"), nil
	case Query:
		dbConnErr, dbEngine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr, nil
		}

		stmt, err := dbEngine.Preparex(spc.Query)
		if err != nil {
			return err, nil
		}
		defer stmt.Close()

		err, result := getMapOfColumnsValues(stmt, args...)
		return err, result

	case StoredProcedure:
		spQuery := ""

		for i := 0; i < len(args); i++ {
			spQuery = spQuery + fmt.Sprintf("%#v", args[i]) + ","
		}

		spQuery = strings.TrimRight(spQuery, ",")

		spQuery = "call " + spc.Query + "(" + spQuery + ")"

		dbConnErr, dbEngine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr, nil
		}

		stmt, err := dbEngine.Preparex(spc.Query)
		if err != nil {
			return err, nil
		}
		defer stmt.Close()

		err, result := getMapOfColumnsValues(stmt, args...)
		return err, result

	}

	return errors.New(fmt.Sprintf("QueryType Not Set. Got %d", spc.QueryType)), nil
}

func (spc *SelectContext) Get(args ...interface{}) error {
	switch spc.QueryType {
	case AutoQuery:
		return errors.New("AutoQuery is not supported for Select method")
	case Query:

		dbConnErr, dbEngine := getConnectionEngine(spc.DBConnection)

		if dbConnErr != nil {
			return dbConnErr
		}

		err := dbEngine.Get(spc.Dest, spc.Query, args...)
		return err
	}

	return nil
}

func (spc *InsertTxContext) Insert() error {

	switch spc.QueryType {
	case AutoQuery: //For text based query
		//return spc.Engine.Exec(spc.Query, spc.Args)

		query := GetInsertDynamicQuery(spc.TableName, spc.Args)
		id, err := spc.Tx.NamedExec(query, spc.Args)
		if err != nil {
			return err
		}
		spc.InsertID, _ = id.LastInsertId()
		return err

	case Query:
		id, err := spc.Tx.NamedExec(spc.Query, spc.Args)
		if err != nil {
			return err
		}
		spc.InsertID, _ = id.LastInsertId()
		return err

	case StoredProcedure:
		var lastinsertid int64

		procGetErr, spQuery := getProcQuery(spc.Tx, spc.Query)

		if procGetErr != nil {
			return procGetErr
		}

		_, execErr := spc.Tx.NamedExec(spQuery, spc.Args)

		if execErr != nil {
			return execErr
		}

		spc.Tx.Get(&lastinsertid, "SELECT Last_Insert_ID()")

		spc.InsertID = lastinsertid

		return nil
	}

	return nil

}

func (spc *UpdateDeleteTxContext) UpdateByFilter(args ...string) error {

	switch spc.QueryType {
	case AutoQuery:
		query := GetUpdateByFilterDynamicQuery(spc.TableName, spc.Args, args...)

		id, err := spc.Tx.NamedExec(query, spc.Args)
		if err != nil {
			return err
		}
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case Query:
		id, err := spc.Tx.NamedExec(spc.Query, spc.Args)
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case StoredProcedure:

		break

	}

	return nil
}

func (spc *UpdateDeleteTxContext) Update() error {

	switch spc.QueryType {
	case AutoQuery:
		queryErr, query := GetUpdateDynamicQuery(spc.TableName, spc.Args)
		if queryErr != nil {
			return queryErr
		}
		id, err := spc.Tx.NamedExec(query, spc.Args)
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case Query:
		id, err := spc.Tx.NamedExec(spc.Query, spc.Args)
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case StoredProcedure:
		procGetErr, spQuery := getProcQuery(spc.Tx, spc.Query)

		if procGetErr != nil {
			return procGetErr
		}

		execResult, execErr := spc.Tx.NamedExec(spQuery, spc.Args)

		if execErr != nil {
			return execErr
		}

		rowAffectedCount, rowAffectedErr := execResult.RowsAffected()

		if rowAffectedErr != nil {
			return rowAffectedErr
		}

		spc.AffectedRows = rowAffectedCount

		return nil
	}

	return nil
}

func (spc *UpdateDeleteTxContext) Delete() error {

	switch spc.QueryType {
	case AutoQuery:
		queryErr, query := GetDeleteDynamicQuery(spc.TableName, spc.Args)
		if queryErr != nil {
			return queryErr
		}
		id, err := spc.Tx.NamedExec(query, spc.Args)
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case Query:
		id, err := spc.Tx.NamedExec(spc.Query, spc.Args)
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case StoredProcedure:
		procGetErr, spQuery := getProcQuery(spc.Tx, spc.Query)

		if procGetErr != nil {
			return procGetErr
		}

		execResult, execErr := spc.Tx.NamedExec(spQuery, spc.Args)

		if execErr != nil {
			return execErr
		}

		rowAffectedCount, rowAffectedErr := execResult.RowsAffected()

		if rowAffectedErr != nil {
			return rowAffectedErr
		}

		spc.AffectedRows = rowAffectedCount

		return nil
	}

	return nil
}

func (spc *SelectTxContext) Select(args ...interface{}) error {

	switch spc.QueryType {
	case AutoQuery:
		return errors.New("AutoQuery is not supported for Select method")
	case Query:
		err := spc.Tx.Select(spc.Dest, spc.Query, args...)
		return err
	case StoredProcedure:
		spQuery := ""

		for i := 0; i < len(args); i++ {
			spQuery = spQuery + fmt.Sprintf("%#v", args[i]) + ","
		}

		spQuery = strings.TrimRight(spQuery, ",")

		spQuery = "call " + spc.Query + "(" + spQuery + ")"

		err := spc.Tx.Select(spc.Dest, spQuery)

		if err != nil {
			return err
		}

		return nil
	}
	return nil
}

func (spc *SelectTxContext) SelectToMap(args ...interface{}) (error, []map[string]interface{}) {

	switch spc.QueryType {
	case AutoQuery:
		return errors.New("AutoQuery is not supported for Select method"), nil
	case Query:

		stmt, err := spc.Tx.Preparex(spc.Query)
		if err != nil {
			return err, nil
		}
		defer stmt.Close()

		err, result := getMapOfColumnsValues(stmt, args...)
		return err, result

	case StoredProcedure:
		spQuery := ""

		for i := 0; i < len(args); i++ {
			spQuery = spQuery + fmt.Sprintf("%#v", args[i]) + ","
		}

		spQuery = strings.TrimRight(spQuery, ",")

		spQuery = "call " + spc.Query + "(" + spQuery + ")"

		stmt, err := spc.Tx.Preparex(spc.Query)
		if err != nil {
			return err, nil
		}
		defer stmt.Close()

		err, result := getMapOfColumnsValues(stmt, args...)
		return err, result

	}

	return errors.New(fmt.Sprintf("QueryType Not Set. Got %d", spc.QueryType)), nil
}

func getProcQuery(tx *sqlx.Tx, procname string) (error, string) {

	procQuery := ""
	spparam := []string{}

	if val, ok := spnameparams[procname]; ok {
		spparam = val
	} else {
		err := tx.Select(&spparam, "select PARAMETER_NAME from information_schema.parameters where SPECIFIC_NAME = ?;", procname)

		if err != nil {
			return err, ""
		}

		for i := 0; i < len(spparam); i++ {
			spparam[i] = strings.TrimLeft(spparam[i], "in_")
		}

		for i := 0; i < len(spparam); i++ {
			procQuery = procQuery + ":" + spparam[i] + ","
		}

		procQuery = strings.TrimRight(procQuery, ",")

		procQuery = "call " + procname + "(" + procQuery + ")"

		procQueryMap[procname] = procQuery
	}

	return nil, procQuery
}

func getDBTags(user interface{}) []string {

	dbTagName := "db"

	t := reflect.TypeOf(user)

	var dbTags []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		dbTags = append(dbTags, field.Tag.Get(dbTagName))

		//		if field.Name == propName {
		//			dbTagValue := field.Tag.Get(dbTagName)
		//			return dbTagValue
		//		}
	}

	return []string{""}
}

func getMapOfColumnsValues(stmt *sqlx.Stmt, args ...interface{}) (error, []map[string]interface{}) {

	rows, err := stmt.Query(args...)
	if err != nil {
		return err, nil
	}

	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return err, nil
	}

	colLength := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, colLength)
	valuePtrs := make([]interface{}, colLength)
	for rows.Next() {
		for i := 0; i < colLength; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		m := make(map[string]interface{})
		for i, col := range columns {
			var newVal interface{}
			val := values[i]

			switch val.(type) {
			case []byte:
				newVal = string(val.([]byte))
			case time.Time:
				newVal = val.(time.Time).Format("2006-01-02 15:04:05")
			default:
				newVal = val
			}
			m[col] = newVal
		}
		tableData = append(tableData, m)
	}

	return nil, tableData
}
