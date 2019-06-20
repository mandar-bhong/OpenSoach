package com.opensoach.vst.Model.DB;

import android.content.ContentValues;
import android.database.Cursor;

import com.opensoach.vst.Constants.DBTableConstants;
import com.opensoach.vst.DAL.DBTableSchema;
import com.opensoach.vst.DAL.IDBRowMapper;

import java.util.Date;

/**
 * Created by Mandar on 08-08-2018.
 */

@DBTableSchema(TableName = DBTableConstants.TABLE_SERVICE_TASK_DATA)
public class DBServiceTaskDataTableQueryModel implements IDBRowMapper<DBServiceTaskDataTableRowModel> {

    public static final String SELECT_ID_FILTER = "SELECT_ID_FILTER";
    public static final String UPDATE_JSON_DATA = "UPDATE_JSON_DATA";
    public static final String SELECT_LOCATION_TIME_FILTER = "SELECT_LOCATION_TIME_FILTER";

    @Override
    public DBServiceTaskDataTableRowModel Clone() {
        return new DBServiceTaskDataTableRowModel();
    }

    @Override
    public void PrepareModel(Cursor cursor, DBServiceTaskDataTableRowModel dataModel) {
        dataModel.setServConfID(cursor.getInt(0));
        dataModel.setSerInID(cursor.getInt(1));
        dataModel.setLocationId(cursor.getInt(2));
        dataModel.setEntryTime(new Date(cursor.getLong(3)));
        dataModel.setSlotStartTime(new Date(cursor.getLong(4)));
        dataModel.setSlotEndTime(new Date(cursor.getLong(5)));
        dataModel.setData(cursor.getString(6));
        dataModel.setData(cursor.getString(7));
    }

    @Override
    public String[] SelectColumn() {
        return new String[]{
                DBTableConstants.TABLE_SERVICE_TASK_DATA_ServConfID,
                DBTableConstants.TABLE_SERVICE_TASK_DATA_SerInID,
                DBTableConstants.TABLE_SERVICE_TASK_DATA_LOCATION_ID,
                DBTableConstants.TABLE_SERVICE_TASK_DATA_TIME,
                DBTableConstants.TABLE_SERVICE_TASK_DATA_SLOT_START_TIME,
                DBTableConstants.TABLE_SERVICE_TASK_DATA_SLOT_END_TIME,
                DBTableConstants.TABLE_SERVICE_TASK_DATA_JSON,
                DBTableConstants.TABLE_SERVICE_TASK_DATA_AUTH_CODE};
    }

    @Override
    public String WhereFilter(String filterName) {
        switch (filterName) {
            case SELECT_ID_FILTER:
                return DBTableConstants.TABLE_SERVICE_TASK_DATA_LOCATION_ID + "=?";
            case SELECT_LOCATION_TIME_FILTER:
                return DBTableConstants.TABLE_SERVICE_TASK_DATA_ServConfID + "=? and " +
                        DBTableConstants.TABLE_SERVICE_TASK_DATA_SerInID + "=? and " +
                        DBTableConstants.TABLE_SERVICE_TASK_DATA_LOCATION_ID + "=? and " +
                        DBTableConstants.TABLE_SERVICE_TASK_DATA_TIME+"=?";
        }
        return "";
    }

    @Override
    public String[] FilterArgs(DBServiceTaskDataTableRowModel dataModel, String filterName) {
        switch (filterName) {
            case SELECT_ID_FILTER:
                return new String[]{String.valueOf(dataModel.getLocationId())};
            case SELECT_LOCATION_TIME_FILTER:
                return new String[]{String.valueOf(dataModel.getServConfID()),
                        String.valueOf(dataModel.getSerInID()),
                        String.valueOf(dataModel.getLocationId()),
                        String.valueOf(dataModel.getEntryTime().getTime())};
        }

        return new String[]{};
    }

    @Override
    public ContentValues UpdateFieldSet(DBServiceTaskDataTableRowModel dataModel, String filterName) {
        ContentValues values = new ContentValues();

        switch (filterName) {
            case UPDATE_JSON_DATA:
                values.put(DBTableConstants.TABLE_SERVICE_TASK_DATA_JSON, dataModel.getData());
                return values;
            case SELECT_LOCATION_TIME_FILTER:
                values.put(DBTableConstants.TABLE_SERVICE_TASK_DATA_JSON, dataModel.getData());
                values.put(DBTableConstants.TABLE_SERVICE_TASK_DATA_SLOT_START_TIME,String.valueOf( dataModel.getSlotStartTime().getTime()));
                values.put(DBTableConstants.TABLE_SERVICE_TASK_DATA_SLOT_END_TIME, String.valueOf(dataModel.getSlotEndTime().getTime()));
                return values;
        }
        return values;
    }
}
