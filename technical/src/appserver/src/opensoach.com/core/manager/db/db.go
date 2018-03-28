package db

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var spnameparams map[string][]string
var procQueryMap map[string]string

func init() {
	spnameparams = make(map[string][]string, 0)
	procQueryMap = make(map[string]string, 0)
}

//InsertProcContext will be used to insert record with Insert method
//InsertedID will be filled if table have autoincriment column

type QueryType int

const (
	AutoQuery       QueryType = 1
	Query                     = 2
	StoredProcedure           = 3
)

type baseContext struct {
	TableName string
	Type      QueryType
	Query     string
}

type context struct {
	baseContext
	Engine *sqlx.DB
}

type contextTx struct {
	baseContext
	Tx *sqlx.Tx
}

type InsertContext struct {
	context
	SPArgs   interface{}
	InsertID int64
}

type UpdateDeleteContext struct {
	context
	SPArgs       interface{}
	AffectedRows int64
}

type SelectContext struct {
	context
	Dest interface{}
}

type InsertTxContext struct {
	contextTx
	SPArgs   interface{}
	InsertID int64
}

type UpdateDeleteTxContext struct {
	contextTx
	SPArgs       interface{}
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

func (spc *InsertContext) Insert() error {

	switch spc.Type {
	case AutoQuery: //For text based query
		//return spc.Engine.Exec(spc.Query, spc.SPArgs)

		query := GetInsertDynamicQuery(spc.TableName, spc.SPArgs)
		id, err := spc.Engine.NamedExec(query, spc.SPArgs)
		spc.InsertID, _ = id.LastInsertId()
		return err

	case Query:
		id, err := spc.Engine.NamedExec(spc.Query, spc.SPArgs)
		spc.InsertID, _ = id.LastInsertId()
		return err

	case StoredProcedure: //For stored Procedure
		var lastinsertid int64

		tx, txerr := spc.Engine.Beginx()

		if txerr != nil {
			return txerr
		}

		procGetErr, spQuery := getProcQuery(tx, spc.Query)
		fmt.Println("SPQuery")
		fmt.Println(spQuery)

		if procGetErr != nil {
			return procGetErr
		}

		_, execErr := tx.NamedExec(spQuery, spc.SPArgs)

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

func (spc *UpdateDeleteContext) Update() error {

	switch spc.Type {
	case AutoQuery:
		query := GetUpdateDynamicQuery(spc.TableName, spc.SPArgs)
		id, err := spc.Engine.NamedExec(query, spc.SPArgs)
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case Query:
		id, err := spc.Engine.NamedExec(spc.Query, spc.SPArgs)
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case StoredProcedure:
		tx, txBeginErr := spc.Engine.Beginx()

		if txBeginErr != nil {
			return txBeginErr
		}

		procGetErr, spQuery := getProcQuery(tx, spc.Query)

		if procGetErr != nil {
			return procGetErr
		}

		execResult, execErr := tx.NamedExec(spQuery, spc.SPArgs)

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

	switch spc.Type {
	case AutoQuery:
		query := GetDeleteDynamicQuery(spc.TableName, spc.SPArgs)
		id, err := spc.Engine.NamedExec(query, spc.SPArgs)
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case Query:
		id, err := spc.Engine.NamedExec(spc.Query, spc.SPArgs)
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case StoredProcedure:

		tx, txBeginErr := spc.Engine.Beginx()

		if txBeginErr != nil {
			return txBeginErr
		}

		procGetErr, spQuery := getProcQuery(tx, spc.Query)

		if procGetErr != nil {
			return procGetErr
		}

		execResult, execErr := tx.NamedExec(spQuery, spc.SPArgs)

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
	switch spc.Type {
	case AutoQuery:
		query := GetSelectAllDynamicQuery(spc.TableName, spc.Dest)
		err := spc.Engine.Select(spc.Dest, query)
		return err
	case Query:
		err := spc.Engine.Select(spc.Dest, spc.Query)
		return err

	case StoredProcedure:

	}
	return nil
}

func (spc *SelectContext) SelectById(arg int) error {
	switch spc.Type {
	case AutoQuery:
		query := GetSelectByIdDynamicQuery(spc.TableName, spc.Dest)
		err := spc.Engine.Select(spc.Dest, query, arg)
		return err
	case Query:
		err := spc.Engine.Select(spc.Dest, spc.Query, arg)
		return err

	case StoredProcedure:
		break

	}
	return nil
}

func (spc *SelectContext) SelectByFilter(filter interface{}, args ...string) error {
	switch spc.Type {
	case AutoQuery:
		query := GetSelectByFilterDynamicQuery(spc.TableName, filter, args...)
		values := GetFilterValues(spc.TableName, filter, args...)
		err := spc.Engine.Select(spc.Dest, query, values...)
		return err

	case Query:
		values := GetFilterValues(spc.TableName, filter, args...)
		err := spc.Engine.Select(spc.Dest, spc.Query, values...)
		return err

	case StoredProcedure:
		break
	}
	return nil
}

func (spc *SelectContext) Select(args ...interface{}) error {

	switch spc.Type {
	case AutoQuery:
		return errors.New("AutoQuery is not supported for Select method")
	case Query:
		err := spc.Engine.Select(spc.Dest, spc.Query, args...)
		return err
	case StoredProcedure:
		spQuery := ""

		for i := 0; i < len(args); i++ {
			spQuery = spQuery + fmt.Sprintf("%#v", args[i]) + ","
		}

		spQuery = strings.TrimRight(spQuery, ",")

		spQuery = "call " + spc.Query + "(" + spQuery + ")"

		err := spc.Engine.Select(spc.Dest, spQuery)

		return err

	}

	return nil
}

func (spc *InsertTxContext) Insert() error {

	switch spc.Type {
	case AutoQuery: //For text based query
		//return spc.Engine.Exec(spc.Query, spc.SPArgs)

		query := GetInsertDynamicQuery(spc.TableName, spc.SPArgs)
		id, err := spc.Tx.NamedExec(query, spc.SPArgs)
		spc.InsertID, _ = id.LastInsertId()
		return err

	case Query:
		id, err := spc.Tx.NamedExec(spc.Query, spc.SPArgs)
		spc.InsertID, _ = id.LastInsertId()
		return err

	case StoredProcedure:
		var lastinsertid int64

		procGetErr, spQuery := getProcQuery(spc.Tx, spc.Query)

		if procGetErr != nil {
			return procGetErr
		}

		_, execErr := spc.Tx.NamedExec(spQuery, spc.SPArgs)

		if execErr != nil {
			return execErr
		}

		spc.Tx.Get(&lastinsertid, "SELECT Last_Insert_ID()")

		spc.InsertID = lastinsertid

		return nil
	}

	return nil

}

func (spc *UpdateDeleteTxContext) Update() error {

	switch spc.Type {
	case AutoQuery:
		query := GetUpdateDynamicQuery(spc.TableName, spc.SPArgs)
		id, err := spc.Tx.NamedExec(query, spc.SPArgs)
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case Query:
		id, err := spc.Tx.NamedExec(spc.Query, spc.SPArgs)
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case StoredProcedure:
		procGetErr, spQuery := getProcQuery(spc.Tx, spc.Query)

		if procGetErr != nil {
			return procGetErr
		}

		execResult, execErr := spc.Tx.NamedExec(spQuery, spc.SPArgs)

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

	switch spc.Type {
	case AutoQuery:
		query := GetDeleteDynamicQuery(spc.TableName, spc.SPArgs)
		id, err := spc.Tx.NamedExec(query, spc.SPArgs)
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case Query:
		id, err := spc.Tx.NamedExec(spc.Query, spc.SPArgs)
		spc.AffectedRows, _ = id.RowsAffected()
		return err

	case StoredProcedure:
		procGetErr, spQuery := getProcQuery(spc.Tx, spc.Query)

		if procGetErr != nil {
			return procGetErr
		}

		execResult, execErr := spc.Tx.NamedExec(spQuery, spc.SPArgs)

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

	switch spc.Type {
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
