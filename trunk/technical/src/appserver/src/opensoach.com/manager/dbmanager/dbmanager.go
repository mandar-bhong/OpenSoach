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

type StoredProcContext struct {
	Engine       *sqlx.DB
	SPName       string
	SPArgs       interface{}
	AffectedRows int64
	InsertID     int64
}

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

type SPDiscoveryParam struct {
	SPName    string
	ParamName []string
}

func (spc *InsertProcContext) Insert() error {

	var lastinsertid int64
	procGetErr, spQuery := getProcQuery(spc.Engine, spc.SPName)

	if procGetErr != nil {
		return procGetErr
	}

	tx, txerr := spc.Engine.Beginx()

	if txerr != nil {
		return txerr
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

	procGetErr, spQuery := getProcQuery(spc.Engine, spc.SPName)

	if procGetErr != nil {
		return procGetErr
	}

	tx, txBeginErr := spc.Engine.Beginx()

	if txBeginErr != nil {
		return txBeginErr
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
	procGetErr, spQuery := getProcQuery(spc.Engine, spc.SPName)

	if procGetErr != nil {
		return procGetErr
	}

	tx, txBeginErr := spc.Engine.Beginx()

	if txBeginErr != nil {
		return txBeginErr
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

func getProcQuery(engine *sqlx.DB, procname string) (error, string) {

	procQuery := ""
	spparam := []string{}

	if val, ok := spnameparams[procname]; ok {
		spparam = val
	} else {
		err := engine.Select(&spparam, "select PARAMETER_NAME from information_schema.parameters where SPECIFIC_NAME = ?;", procname)

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
