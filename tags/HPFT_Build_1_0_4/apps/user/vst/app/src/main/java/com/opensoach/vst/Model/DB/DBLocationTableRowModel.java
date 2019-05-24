package com.opensoach.vst.Model.DB;

import com.opensoach.vst.DAL.DBConstants;
import com.opensoach.vst.DAL.DBTableSchema;

/**
 * Created by Mandar on 4/8/2017.
 */

@DBTableSchema(TableName = DBConstants.TABLE_LOCATION)
public class DBLocationTableRowModel {

    private int locationId;
    private String locationName;
    private String locationCat;

    public DBLocationTableRowModel(){
        locationName = "";
        locationCat = "";
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

    public String getLocationCat() {
        return locationCat;
    }

    public void setLocationCat(String locationCat) {
        this.locationCat = locationCat;
    }
}
