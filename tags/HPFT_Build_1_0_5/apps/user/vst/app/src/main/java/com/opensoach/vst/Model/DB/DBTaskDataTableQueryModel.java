package com.opensoach.vst.Model.DB;

import android.content.ContentValues;
import android.database.Cursor;

import com.opensoach.vst.Constants.DBTableConstants;
import com.opensoach.vst.DAL.DBTableSchema;
import com.opensoach.vst.DAL.IDBRowMapper;

import java.util.Date;

@DBTableSchema(TableName = DBTableConstants.TABLE_TASK_DATA)
public class DBTaskDataTableQueryModel implements IDBRowMapper<DBTaskDataTableRowModel> {

    public static final String SELECT_ID_FILTER = "SELECT_ID_FILTER";
    public static final String SELECT_TITLE_FILTER = "SELECT_TITLE_FILTER";
    public static final String UPDATE_SYNC_BY_ID_FILTER = "UPDATE_SYNC_BY_ID_FILTER";
    public static final String FILTER_BY_UNSYNC_DATA ="FILTER_BY_UNSYNC_DATA";

    @Override
    public DBTaskDataTableRowModel Clone() {
        return new DBTaskDataTableRowModel();
    }

    @Override
    public void PrepareModel(Cursor cursor, DBTaskDataTableRowModel dataModel) {

        dataModel.setSerInID(cursor.getInt(0));
        dataModel.setLocationId(cursor.getInt(1));
        dataModel.setTitle(cursor.getString(2));
        dataModel.setTaskTime(new Date(cursor.getLong(3)));
        dataModel.setTaskSlotStartTime(new Date(cursor.getLong(4)));
        dataModel.setTaskSlotEndTime(new Date(cursor.getLong(5)));
        dataModel.setSynced(cursor.getInt(6) == 1);
        dataModel.setAuthCode(cursor.getString(7));
        dataModel.setValue(cursor.getString(8));
        dataModel.setComment(cursor.getString(9));
    }

    @Override
    public String[] SelectColumn() {
        return new String[]{
                DBTableConstants.TABLE_TASK_DATA_SerInID,
                DBTableConstants.TABLE_TASK_DATA_LOCATION_ID,
                DBTableConstants.TABLE_TASK_DATA_TITLE,
                DBTableConstants.TABLE_TASK_DATA_TIME,
                DBTableConstants.TABLE_TASK_DATA_SLOT_START_TIME,
                DBTableConstants.TABLE_TASK_DATA_SLOT_END_TIME,
                DBTableConstants.TABLE_TASK_DATA_SERVER_SYNC,
                DBTableConstants.TABLE_TASK_DATA_AUTH_CODE,
                DBTableConstants.TABLE_TASK_DATA_VALUE,
                DBTableConstants.TABLE_TASK_DATA_COMMENT};
    }

    @Override
    public String WhereFilter(String filterName) {
        switch (filterName) {
            case UPDATE_SYNC_BY_ID_FILTER:
            case SELECT_ID_FILTER:
                return DBTableConstants.TABLE_TASK_DATA_LOCATION_ID + "=? and "+
                        DBTableConstants.TABLE_TASK_DATA_SerInID +"=? and " +
                        DBTableConstants.TABLE_TASK_DATA_SLOT_START_TIME +"=?";

            case SELECT_TITLE_FILTER:
                return DBTableConstants.TABLE_TASK_DATA_LOCATION_ID + "=? and "+
                        DBTableConstants.TABLE_TASK_DATA_SerInID +"=? and "+
                        DBTableConstants.TABLE_TASK_DATA_SLOT_START_TIME +"=? and " +
                        DBTableConstants.TABLE_TASK_DATA_TITLE +"=?";

            case FILTER_BY_UNSYNC_DATA:
                return DBTableConstants.TABLE_TASK_DATA_LOCATION_ID + "=? and "+
                        DBTableConstants.TABLE_TASK_DATA_SERVER_SYNC +"=?";
        }
        return "";
    }

    @Override
    public String[] FilterArgs(DBTaskDataTableRowModel dataModel, String filterName) {
        switch (filterName) {
            case UPDATE_SYNC_BY_ID_FILTER:
            case SELECT_ID_FILTER:
                return new String[]{String.valueOf(dataModel.getLocationId()),
                        String.valueOf(dataModel.getSerInID()),
                        String.valueOf(dataModel.getTaskSlotStartTime().getTime())};
            case SELECT_TITLE_FILTER:
                return new String[]{String.valueOf(dataModel.getLocationId()),
                        String.valueOf(dataModel.getSerInID()),
                        String.valueOf(dataModel.getTaskSlotStartTime().getTime()),
                        String.valueOf(dataModel.getTitle())};
            case FILTER_BY_UNSYNC_DATA:
                return new String[]{String.valueOf(dataModel.getLocationId()),
                        String.valueOf(dataModel.isSynced() ? 1 : 0)};
        }

        return new String[]{};
    }

    @Override
    public ContentValues UpdateFieldSet(DBTaskDataTableRowModel dataModel, String filterName) {
        ContentValues values = new ContentValues();

        switch (filterName) {
            case UPDATE_SYNC_BY_ID_FILTER:
                values.put(DBTableConstants.TABLE_TASK_DATA_SERVER_SYNC, dataModel.isSynced() ? 1 : 0);
                return values;

        }
        return values;
    }
}
