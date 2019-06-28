package spl.hkt.opensoach.splapp.model.db;

import android.content.ContentValues;
import android.database.Cursor;

import java.util.Date;

import spl.hkt.opensoach.splapp.dal.DBConstants;
import spl.hkt.opensoach.splapp.dal.DBTableSchema;
import spl.hkt.opensoach.splapp.dal.IDBRowMapper;
import spl.hkt.opensoach.splapp.util.CommonUtility;

/**
 * Created by Mandar on 3/12/2017.
 */

@DBTableSchema(TableName = DBConstants.TABLE_CHART_DATA)
public class DBChartDataTableQueryModel implements IDBRowMapper<DBChartDataTableRowModel> {

    public static final String UPDATE_SYNC_STATE_WITH_CHART_ID = "UPDATE_SYNC_STATE_WITH_CHART_ID";
    public static final String UPDATE_SYNC_STATE_WITH_TASK_ID = "UPDATE_SYNC_STATE_WITH_TASK_ID";
    public static final String UPDATE_SYNC_STATE_WITH_SLOT_ID = "UPDATE_SYNC_STATE_WITH_SLOT_ID";
    public static final String UPDATE_SYNC_STATE_WITH_CHART_TASK_SLOT_ID = "UPDATE_SYNC_STATE_WITH_CHART_TASK_SLOT_ID";
    public static final String FILTER_BY_DAY = "FILTER_BY_DAY";
    public static final String UPDATE_SYNC_STATE_WITH_CHART_ID_ENTRY_TIME = "UPDATE_SYNC_STATE_WITH_CHART_ID_ENTRY_TIME";
    public static final String FILTER_BY_CHARTID_TODAY = "FILTER_BY_CHARTID_TODAY";
    public static final String FILTER_BY_UNSYNC_DATA = "FILTER_BY_UNSYNC_DATA";

    @Override
    public DBChartDataTableRowModel Clone() {
        return new DBChartDataTableRowModel();
    }


    @Override
    public void PrepareModel(Cursor cursor, DBChartDataTableRowModel dataModel) {
        dataModel.setChartId(cursor.getInt(0));
        dataModel.setTaskName(cursor.getString(1));
        dataModel.setSlotId(cursor.getInt(2));
        dataModel.setEntryTime(new Date(cursor.getLong(3)));
        dataModel.setSlotStartTime(new Date(cursor.getLong(4)));
        dataModel.setSlotEndTime(new Date(cursor.getLong(5)));
        dataModel.setCellState(cursor.getInt(6));
        dataModel.setChartDay(new Date(cursor.getLong(7)));
        dataModel.setAuthCode(cursor.getString(8));
        dataModel.setSynced(cursor.getInt(9) == 1);
    }

    @Override
    public String[] SelectColumn() {
        return new String[]{
                DBConstants.TABLE_CHART_DATA_CHART_ID,
                DBConstants.TABLE_CHART_DATA_TASK_NAME,
                DBConstants.TABLE_CHART_DATA_SLOT_ID,
                DBConstants.TABLE_CHART_DATA_ENTRY_TIME,
                DBConstants.TABLE_CHART_DATA_SLOT_START_TIME,
                DBConstants.TABLE_CHART_DATA_SLOT_END_TIME,
                DBConstants.TABLE_CHART_DATA_STATE,
                DBConstants.TABLE_CHART_DATA_DAY,
                DBConstants.TABLE_CHART_DATA_AUTH_CODE,
                DBConstants.TABLE_CHART_DATA_SERVER_SYNC
        };
    }

    @Override
    public String WhereFilter(String filterName) {

        switch (filterName) {
            case UPDATE_SYNC_STATE_WITH_CHART_ID:
                return DBConstants.TABLE_CHART_DATA_CHART_ID + "=?";
            case UPDATE_SYNC_STATE_WITH_TASK_ID:
                return DBConstants.TABLE_CHART_DATA_TASK_NAME + "=?";
            case UPDATE_SYNC_STATE_WITH_SLOT_ID:
                return DBConstants.TABLE_CHART_DATA_SLOT_ID + "=?";
            case FILTER_BY_DAY:
                return DBConstants.TABLE_CHART_DATA_DAY + "=?";
            case UPDATE_SYNC_STATE_WITH_CHART_TASK_SLOT_ID:
                return DBConstants.TABLE_CHART_DATA_CHART_ID + "=?, " + DBConstants.TABLE_CHART_DATA_DAY + "=?";
            case UPDATE_SYNC_STATE_WITH_CHART_ID_ENTRY_TIME:
                return DBConstants.TABLE_CHART_DATA_CHART_ID + "=? and " + DBConstants.TABLE_CHART_DATA_ENTRY_TIME + "=?";
            case FILTER_BY_CHARTID_TODAY:
                return DBConstants.TABLE_CHART_DATA_CHART_ID + "=? and " + DBConstants.TABLE_CHART_DATA_DAY + " = ?";
            case FILTER_BY_UNSYNC_DATA:
                return DBConstants.TABLE_CHART_DATA_SERVER_SYNC + " = ?";
        }
        return "";
    }

    @Override
    public String[] FilterArgs(DBChartDataTableRowModel dataModel, String filterName) {
        switch (filterName) {
            case UPDATE_SYNC_STATE_WITH_CHART_ID:
                return new String[]{String.valueOf(dataModel.getChartId())};
            case UPDATE_SYNC_STATE_WITH_TASK_ID:
                return new String[]{String.valueOf(dataModel.getTaskName())};
            case UPDATE_SYNC_STATE_WITH_SLOT_ID:
                return new String[]{String.valueOf(dataModel.getSlotId())};
            case FILTER_BY_DAY:
                return new String[]{String.valueOf(dataModel.getChartDay())};
            case UPDATE_SYNC_STATE_WITH_CHART_TASK_SLOT_ID:
                return new String[]{String.valueOf(dataModel.getChartId()), String.valueOf(dataModel.getTaskName()), String.valueOf(dataModel.getSlotId())};
            case UPDATE_SYNC_STATE_WITH_CHART_ID_ENTRY_TIME:
                return new String[]{String.valueOf(dataModel.getChartId()), String.valueOf(dataModel.getEntryTime().getTime())};
            case FILTER_BY_CHARTID_TODAY:
                return new String[]{String.valueOf(dataModel.getChartId()), String.valueOf(dataModel.getChartDay().getTime())};
            case FILTER_BY_UNSYNC_DATA:
                return new String[]{String.valueOf((dataModel.isSynced()) ? 1 : 0)};

        }

        return new String[]{};
    }

    @Override
    public ContentValues UpdateFieldSet(DBChartDataTableRowModel dataModel, String filterName) {
        ContentValues values = new ContentValues();

        switch (filterName) {
            case UPDATE_SYNC_STATE_WITH_CHART_TASK_SLOT_ID:
                values.put(DBConstants.TABLE_CHART_DATA_SERVER_SYNC, dataModel.isSynced());
                return values;
            case UPDATE_SYNC_STATE_WITH_CHART_ID_ENTRY_TIME:
                values.put(DBConstants.TABLE_CHART_DATA_SERVER_SYNC, dataModel.isSynced());
                return values;
            case UPDATE_SYNC_STATE_WITH_CHART_ID:
                values.put(DBConstants.TABLE_CHART_DATA_SERVER_SYNC, dataModel.isSynced());
                return values;
        }
        return values;
    }
}
