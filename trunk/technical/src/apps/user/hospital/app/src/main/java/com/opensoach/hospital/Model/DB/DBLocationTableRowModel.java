package com.opensoach.hospital.Model.DB;

import com.opensoach.hospital.DAL.DBConstants;
import com.opensoach.hospital.DAL.DBTableSchema;

/**
 * Created by Mandar on 8/26/2017.
 */

@DBTableSchema(TableName = DBConstants.TABLE_LOCATION)
public class DBLocationTableRowModel {

    private int locationId;
    private String locationName;


    public DBLocationTableRowModel(){
        locationName = "";
    }

    public int getLocationId() {
        return locationId;
    }

    public void setLocationId(int locationId) {
        this.locationId = locationId;
    }

    public String getLocationName() {
        return locationName;
    }

    public void setLocationName(String locationName) {
        this.locationName = locationName;
    }

}
