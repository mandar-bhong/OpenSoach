package splserver

import (
	gmodels "opensoach.com/models"
	splmgr "opensoach.com/splserver/manager"
)

func Init(dbconfig *gmodels.ConfigDB) bool {

	splmgr.InitilizeModues(dbconfig)

	return true

}
