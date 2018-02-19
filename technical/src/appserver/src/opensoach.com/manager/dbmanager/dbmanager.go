package dbmanager

import (
	"fmt"
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

//StoredProcContext will be used to Test
type StoredProcContext struct {
	Engine       *sqlx.DB
	SPName       string
	SPArgs       interface{}
	AffectedRows int64
	InsertID     int64
}

//InsertProcContext will be used to insert record with Insert method
//InsertedID will be filled if table have autoincriment column
type InsertProcContext struct {
	Engine   *sqlx.DB
	SPName   string
	SPArgs   interface{}
	InsertID int64
}

type UpdateDeleteProcContext struct {
	Engine       *sqlx.DB
	SPName       string
	SPArgs       interface{}
	AffectedRows int64
}

type SelectProcContext struct {
	Engine *sqlx.DB
	SPName string
	Dest   interface{}
}

type InsertProcTxContext struct {
	Tx       *sqlx.Tx
	SPName   string
	SPArgs   interface{}
	InsertID int64
}

type UpdateDeleteProcTxContext struct {
	Tx           *sqlx.Tx
	SPName       string
	SPArgs       interface{}
	AffectedRows int64
}

type SelectProcTxContext struct {
	Tx     *sqlx.Tx
	SPName string
	Dest   interface{}
}

type SPDiscoveryParam struct {
	SPName    string
	ParamName []string
}

func (spc *InsertProcContext) Insert() error {

	var lastinsertid int64

	tx, txerr := spc.Engine.Beginx()

	if txerr != nil {
		return txerr
	}

	procGetErr, spQuery := getProcQuery(tx, spc.SPName)

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

func (spc *UpdateDeleteProcContext) Update() error {

	tx, txBeginErr := spc.Engine.Beginx()

	if txBeginErr != nil {
		return txBeginErr
	}

	procGetErr, spQuery := getProcQuery(tx, spc.SPName)

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

func (spc *UpdateDeleteProcContext) Delete() error {

	tx, txBeginErr := spc.Engine.Beginx()

	if txBeginErr != nil {
		return txBeginErr
	}

	procGetErr, spQuery := getProcQuery(tx, spc.SPName)

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

func (spc *SelectProcContext) Select(args ...interface{}) error {

	spQuery := ""

	for i := 0; i < len(args); i++ {
		spQuery = spQuery + fmt.Sprintf("%#v", args[i]) + ","
	}

	spQuery = strings.TrimRight(spQuery, ",")

	spQuery = "call " + spc.SPName + "(" + spQuery + ")"

	err := spc.Engine.Select(spc.Dest, spQuery)

	if err != nil {
		return err
	}

	return nil
}

func (spc *InsertProcTxContext) Insert() error {

	var lastinsertid int64

	procGetErr, spQuery := getProcQuery(spc.Tx, spc.SPName)

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

func (spc *UpdateDeleteProcTxContext) Update() error {

	procGetErr, spQuery := getProcQuery(spc.Tx, spc.SPName)

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

func (spc *UpdateDeleteProcTxContext) Delete() error {
	procGetErr, spQuery := getProcQuery(spc.Tx, spc.SPName)

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

func (spc *SelectProcTxContext) Select(args ...interface{}) error {

	spQuery := ""

	for i := 0; i < len(args); i++ {
		spQuery = spQuery + fmt.Sprintf("%#v", args[i]) + ","
	}

	spQuery = strings.TrimRight(spQuery, ",")

	spQuery = "call " + spc.SPName + "(" + spQuery + ")"

	err := spc.Tx.Select(spc.Dest, spQuery)

	if err != nil {
		return err
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
