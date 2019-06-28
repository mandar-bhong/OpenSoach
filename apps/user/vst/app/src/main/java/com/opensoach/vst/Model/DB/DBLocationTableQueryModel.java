package com.opensoach.vst.Model.DB;

import android.content.ContentValues;
import android.database.Cursor;

import com.opensoach.vst.DAL.DBConstants;
import com.opensoach.vst.DAL.DBTableSchema;
import com.opensoach.vst.DAL.IDBRowMapper;

/**
 * Created by Mandar on 4/8/2017.
 */

@DBTableSchema(TableName = DBConstants.TABLE_LOCATION)
public class DBLocationTableQueryModel implements IDBRowMapper<DBLocationTableRowModel> {

    public static final String SELECT_ID_FILTER = "SELECT_ID_FILTER";
    public static final String UPDATE_INFO_FILTER_BY_ID = "UPDATE_INFO_FILTER_BY_ID";

    @Override
    public DBLocationTableRowModel Clone() {
        return new DBLocationTableRowModel();
    }

    @Override
    public void PrepareModel(Cursor cursor, DBLocationTableRowModel dataModel) {
        dataModel.setLocationId(cursor.getInt(0));
        dataModel.setLocationName(cursor.getString(1));
        dataModel.setLocationCat(cursor.getString(2));
    }

    @Override
    public String[] SelectColumn() {
        return new String[]{DBConstants.TABLE_LOCATION_ID,
                DBConstants.TABLE_LOCATION_NAME,
                DBConstants.TABLE_LOCATION_CAT};
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
}
