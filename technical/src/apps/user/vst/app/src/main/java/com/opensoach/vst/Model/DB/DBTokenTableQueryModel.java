package com.opensoach.vst.Model.DB;

import android.content.ContentValues;
import android.database.Cursor;

import java.util.Date;

import com.opensoach.vst.Constants.DBTableConstants;
import com.opensoach.vst.DAL.DBConstants;
import com.opensoach.vst.DAL.DBTableSchema;
import com.opensoach.vst.DAL.IDBRowMapper;

@DBTableSchema(TableName = DBTableConstants.TABLE_TOKEN_DATA)
public class DBTokenTableQueryModel implements IDBRowMapper<DBTokenTableRowModel> {

    public static final String SELECT_ID_FILTER = "SELECT_ID_FILTER";
    public static final String SELECT_TOKEN_NO_FILTER = "SELECT_TOKEN_NO_FILTER";

    @Override
    public DBTokenTableRowModel Clone() {
        return new DBTokenTableRowModel();
    }

    @Override
    public void PrepareModel(Cursor cursor, DBTokenTableRowModel dataModel) {
        dataModel.setId(cursor.getInt(0));
        dataModel.setTokenno(cursor.getInt(1));
        dataModel.setVehicleno(cursor.getString(2));
        dataModel.setMapping(cursor.getString(3));
        dataModel.setState(cursor.getInt(4));
        dataModel.setGeneratedon(new Date(cursor.getLong(5)));
    }

    @Override
    public String[] SelectColumn() {

        return new String[]{DBTableConstants.TABLE_TOKEN_ID,
                DBTableConstants.TABLE_TOKEN_NO,
                DBTableConstants.TABLE_TOKEN_VEHICLE_NO,
                DBTableConstants.TABLE_TOKEN_MAPPING_DETAILS,
                DBTableConstants.TABLE_TOKEN_STATE,
                DBTableConstants.TABLE_TOKEN_GENERATED_ON
        };

    }

    @Override
    public String WhereFilter(String filterName) {
        switch (filterName) {
            case SELECT_ID_FILTER:
                return DBTableConstants.TABLE_TOKEN_ID + "=?";
            case SELECT_TOKEN_NO_FILTER:
                return DBTableConstants.TABLE_TOKEN_NO+ "=?";
        }
        return "";
    }

    @Override
    public String[] FilterArgs(DBTokenTableRowModel dataModel, String filterName) {
        switch (filterName) {
            case SELECT_ID_FILTER:
                return new String[]{String.valueOf(dataModel.getId())};
            case SELECT_TOKEN_NO_FILTER:
                return new String[]{String.valueOf(dataModel.getTokenno())};
        }

        return new String[]{};
    }

    @Override
    public ContentValues UpdateFieldSet(DBTokenTableRowModel dataModel, String filterName) {
        ContentValues values = new ContentValues ();

        switch (filterName) {
            case SELECT_ID_FILTER:
                values.put(DBTableConstants.TABLE_TOKEN_ID, dataModel.getId());
                return values;
            case SELECT_TOKEN_NO_FILTER:
                values.put(DBTableConstants.TABLE_TOKEN_NO, dataModel.getTokenno());
                return values;
        }

        return values;
    }
}
