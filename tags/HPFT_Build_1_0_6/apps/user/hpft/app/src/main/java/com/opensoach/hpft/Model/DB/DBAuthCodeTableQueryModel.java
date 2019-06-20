package com.opensoach.hpft.Model.DB;


import android.content.ContentValues;
import android.database.Cursor;

import com.opensoach.hpft.DAL.DBConstants;
import com.opensoach.hpft.DAL.DBTableSchema;
import com.opensoach.hpft.DAL.IDBRowMapper;

/**
 * Created by samir.s.bukkawar on 4/12/2017.
 */


@DBTableSchema(TableName = DBConstants.TABLE_AUTH_LOCATION)
public class DBAuthCodeTableQueryModel implements IDBRowMapper<DBAuthCodeTableRowModel> {

    public static final String SELECT_BY_LOCATION_FILTER = "SELECT_BY_LOCATION_FILTER";
    public static final String SELECT_BY_LOCATION_AND_AUTHCODE_FILTER = "SELECT_BY_LOCATION_AND_AUTHCODE_FILTER";
    public static final String UDATE_AUTHCODE_BY_LOCATIONID_FILTER = "UDATE_AUTHCODE_BY_LOCATIONID__FILTER";

    @Override
    public DBAuthCodeTableRowModel Clone() {
        return new DBAuthCodeTableRowModel();
    }

    @Override
    public void PrepareModel(Cursor cursor, DBAuthCodeTableRowModel dataModel) {
        dataModel.setLocationId(cursor.getInt(0));
        dataModel.setAuthCode(cursor.getString(1));
    }


    @Override
    public String[] SelectColumn() {
        return new String[]{DBConstants.TABLE_AUTH_LOCATION_LOCATIONID,
                DBConstants.TABLE_AUTH_LOCATION_AUTHCODE};
    }

    @Override
    public String WhereFilter(String filterName) {
        switch (filterName) {
            case UDATE_AUTHCODE_BY_LOCATIONID_FILTER:
                return DBConstants.TABLE_AUTH_LOCATION_LOCATIONID + "=?";
            case SELECT_BY_LOCATION_FILTER:
                return DBConstants.TABLE_AUTH_LOCATION_LOCATIONID + "=?";
            case SELECT_BY_LOCATION_AND_AUTHCODE_FILTER:
                return DBConstants.TABLE_AUTH_LOCATION_LOCATIONID + "=? and "+DBConstants.TABLE_AUTH_LOCATION_AUTHCODE;
        }
        return "";
    }

    @Override
    public String[] FilterArgs(DBAuthCodeTableRowModel dataModel, String filterName) {
        switch (filterName) {
            case UDATE_AUTHCODE_BY_LOCATIONID_FILTER:
                return new String[]{String.valueOf(dataModel.getLocationId())};
            case SELECT_BY_LOCATION_FILTER:
                return new String[]{String.valueOf(dataModel.getLocationId())};
            case SELECT_BY_LOCATION_AND_AUTHCODE_FILTER:
                return new String[]{String.valueOf(dataModel.getLocationId()),String.valueOf(dataModel.getAuthCodeJSON())};
        }

        return new String[]{};
    }

    @Override
    public ContentValues UpdateFieldSet(DBAuthCodeTableRowModel dataModel, String filterName) {

        ContentValues values = new ContentValues();

        switch (filterName) {
            case UDATE_AUTHCODE_BY_LOCATIONID_FILTER:
                values.put(DBConstants.TABLE_AUTH_LOCATION_AUTHCODE, dataModel.getAuthCodeJSON());
                return values;
        }
        return values;
    }
}
