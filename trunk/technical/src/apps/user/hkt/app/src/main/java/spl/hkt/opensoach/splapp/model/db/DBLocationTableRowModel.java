package spl.hkt.opensoach.splapp.model.db;

import spl.hkt.opensoach.splapp.dal.DBConstants;
import spl.hkt.opensoach.splapp.dal.DBTableSchema;

/**
 * Created by Mandar on 4/8/2017.
 */

@DBTableSchema(TableName = DBConstants.TABLE_LOCATION)
public class DBLocationTableRowModel {

    private int locationId;
    private String locationName;
    private int locationCat;

    public DBLocationTableRowModel(){
        locationName = "";
        locationCat = 0;
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

    public int getLocationCat() {
        return locationCat;
    }

    public void setLocationCat(int locationCat) {
        this.locationCat = locationCat;
    }
}
