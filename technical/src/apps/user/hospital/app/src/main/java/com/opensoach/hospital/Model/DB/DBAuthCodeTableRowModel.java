package com.opensoach.hospital.Model.DB;

import com.opensoach.hospital.DAL.DBConstants;
import com.opensoach.hospital.DAL.DBTableSchema;

/**
 * Created by Mandar on 8/26/2017.
 */

@DBTableSchema(TableName = DBConstants.TABLE_AUTH_CODE)
public class DBAuthCodeTableRowModel {

    private String authCodeJSON;

    public DBAuthCodeTableRowModel() {
    }


    public String getAuthCodeJSON() {
        return authCodeJSON;
    }

    public void setAuthCode(String authCode) {
        this.authCodeJSON = authCode;
    }
}
