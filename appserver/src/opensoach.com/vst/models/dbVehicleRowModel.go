package models

type DBVehicleDataModel struct {
	VehicleNo string `db:"vehicle_no" json:"vehicleno"`
	Details   string `db:"details" json:"details"`
}

type DBVehicleInsertRowModel struct {
	DBVehicleDataModel
	CpmId int64 `db:"cpm_id_fk" json:"cpmid"`
}

type DBVehicleUpdateRowModel struct {
	VehicleId int64  `db:"id" dbattr:"pri,auto"  json:"vehicleid"`
	CpmId     int64  `db:"cpm_id_fk" json:"cpmid"`
	VehicleNo string `db:"vehicle_no" json:"vehicleno"`
	Details   string `db:"details" json:"details"`
}

type DBVehicleDetailsUpdateModel struct {
	VehicleNo string `db:"vehicle_no" json:"vehicleno"`
	Details   string `db:"details" json:"details"`
}
