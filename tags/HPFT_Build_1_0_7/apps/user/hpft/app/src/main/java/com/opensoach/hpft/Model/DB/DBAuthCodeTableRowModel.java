package com.opensoach.hpft.Model.DB;

import com.opensoach.hpft.DAL.DBConstants;
import com.opensoach.hpft.DAL.DBTableSchema;

/**
 * Created by samir.s.bukkawar on 4/12/2017.
 */

@DBTableSchema(TableName = DBConstants.TABLE_AUTH_LOCATION)
public class DBAuthCodeTableRowModel {

    private int locationId;
    private String authCodeJSON;

    public DBAuthCodeTableRowModel() {
    }

    public int getLocationId() {
        return locationId;
    }

    public void setLocationId(int locationId) {
        this.locationId = locationId;
    }

    public String getAuthCodeJSON() {
        return authCodeJSON;
    }

    public void setAuthCode(String authCode) {
        this.authCodeJSON = authCode;
    }
}
