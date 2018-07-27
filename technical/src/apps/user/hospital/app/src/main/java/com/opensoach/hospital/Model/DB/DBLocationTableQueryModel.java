package com.opensoach.hospital.Model.DB;

import android.content.ContentValues;
import android.database.Cursor;

import com.opensoach.hospital.DAL.DBConstants;
import com.opensoach.hospital.DAL.DBTableSchema;
import com.opensoach.hospital.DAL.IDBOutputFormatter;
import com.opensoach.hospital.DAL.IDBRowMapper;

/**
 * Created by Mandar on 8/26/2017.
 */

@DBTableSchema(TableName = DBConstants.TABLE_LOCATION)
public class DBLocationTableQueryModel implements IDBRowMapper<DBLocationTableRowModel>, IDBOutputFormatter {

    public static final String SELECT_ID_FILTER = "SELECT_ID_FILTER";
    public static final String SELECT_ALL_FILTER = "SELECT_ALL_FILTER";
    public static final String UPDATE_INFO_FILTER_BY_ID = "UPDATE_INFO_FILTER_BY_ID";
    public static final String SORT_BY_NAME_DESC = "SORT_BY_NAME_DESC";

    @Override
    public DBLocationTableRowModel Clone() {
        return new DBLocationTableRowModel();
    }

    @Override
    public void PrepareModel(Cursor cursor, DBLocationTableRowModel dataModel) {
        dataModel.setLocationId(cursor.getInt(0));
        dataModel.setLocationName(cursor.getString(1));
    }

    @Override
    public String[] SelectColumn() {
        return new String[]{DBConstants.TABLE_LOCATION_ID,
                DBConstants.TABLE_LOCATION_NAME};
    }

    @Override
    public String WhereFilter(String filterName) {
        switch (filterName) {
            case SELECT_ID_FILTER:
                return DBConstants.TABLE_LOCATION_ID + "=?";
        }
        return "";
    }

    @Override
    public String[] FilterArgs(DBLocationTableRowModel dataModel, String filterName) {
        switch (filterName) {
            case SELECT_ID_FILTER:
                return new String[]{String.valueOf(dataModel.getLocationId())};
        }

        return new String[]{};
    }

    @Override
    public ContentValues UpdateFieldSet(DBLocationTableRowModel dataModel, String filterName) {
        return null;
    }

    @Override
    public String OrderBy(String orderFiter) {

        switch (orderFiter) {
            case SORT_BY_NAME_DESC:
                return DBConstants.TABLE_LOCATION_NAME +" ASC" ;
        }

        return null;
    }
}
