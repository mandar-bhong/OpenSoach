package com.opensoach.hospital.Model.DB;

import android.content.ContentValues;
import android.database.Cursor;

import com.opensoach.hospital.DAL.DBConstants;
import com.opensoach.hospital.DAL.DBTableSchema;
import com.opensoach.hospital.DAL.IDBRowMapper;

import java.util.Date;

/**
 * Created by Mandar on 9/5/2017.
 */

@DBTableSchema(TableName = DBConstants.TABLE_JOB_CARD)
public class DBJobCardTableQueryModel implements IDBRowMapper<DBJobCardTableRowModel> {

    public static final String SELECT_ALL_FILTER = "SELECT_ALL_FILTER";
    public static final String SELECT_ID_FILTER = "SELECT_ID_FILTER";
    public static final String SELECT_ID_AND_LOCATION_ID_FILTER = "SELECT_ID_AND_LOCATION_ID_FILTER";
    public static final String SELECT_LOCATION_ID_FILTER = "SELECT_LOCATION_ID_FILTER";
    public static final String UPDATE_STATE_BY_ID_FILTER = "UPDATE_STATE_BY_ID_FILTER";
    public static final String UPDATE_STATE_AND_ACTUAL_START_TIME_BY_ID_FILTER = "UPDATE_STATE_AND_ACTUAL_START_TIME_BY_ID_FILTER";
    public static final String UPDATE_QUANTITY_BY_ID_FILTER = "UPDATE_QUANTITY_BY_ID_FILTER ";
    public static final String UPDATE_QUANTITY_BY_ID_AND_LOCATION_ID_FILTER = "UPDATE_QUANTITY_BY_ID_AND_LOCATION_ID_FILTER";
    public static final String UPDATE_JCODE_PARTCOUNT_STARTTIME_ENDTIME_BY_ID_AND_LOCATION_ID_FILTER = "UPDATE_JCODE_PARTCOUNT_STARTTIME_ENDTIME_BY_ID_AND_LOCATION_ID_FILTER";
    public static final String DELETE_BY_ID_FILTER = "DELETE_BY_ID_FILTER";

    @Override
    public DBJobCardTableRowModel Clone() {
        return new DBJobCardTableRowModel();
    }

    @Override
    public void PrepareModel(Cursor cursor, DBJobCardTableRowModel dataModel) {
        dataModel.setJobCardId(cursor.getInt(0));
        dataModel.setPartId(cursor.getInt(1));
        dataModel.setPartCount(cursor.getInt(2));
        dataModel.setJobCode(cursor.getString(3));
        dataModel.setStartDate(new Date(cursor.getLong(4)));
        dataModel.setEndDate(new Date(cursor.getLong(5)));
        dataModel.setActualStartDate(new Date(cursor.getLong(6)));
        dataModel.setActualEndDate(new Date(cursor.getLong(7)));
        dataModel.setState(cursor.getInt(8));
        dataModel.setComments(cursor.getString(9));
        dataModel.setCustomer(cursor.getString(10));
        dataModel.setLocationId(cursor.getInt(11));
        dataModel.setCompletedCount(cursor.getInt(12));
        dataModel.setJobConfig(cursor.getString(13));
    }

    @Override
    public String[] SelectColumn() {
        return new String[]{DBConstants.TABLE_JOB_CARD_JOBID,
                DBConstants.TABLE_JOB_CARD_PART_ID,
                DBConstants.TABLE_JOB_CARD_PART_COUNT,
                DBConstants.TABLE_JOB_CARD_CODE,
                DBConstants.TABLE_JOB_CARD_PART_START_DATE,
                DBConstants.TABLE_JOB_CARD_PART_END_DATE,
                DBConstants.TABLE_JOB_CARD_PART_ACTUAL_START_DATE,
                DBConstants.TABLE_JOB_CARD_PART_ACTUAL_END_DATE,
                DBConstants.TABLE_JOB_CARD_PART_STATE,
                DBConstants.TABLE_JOB_CARD_PART_COMMENTS,
                DBConstants.TABLE_JOB_CARD_CUSTOMER,
                DBConstants.TABLE_JOB_CARD_LOCATION_ID,
                DBConstants.TABLE_JOB_CARD_PART_COMPLETED_COUNT,
                DBConstants.TABLE_JOB_CARD_JOB_CONFIG
        };
    }

    @Override
    public String WhereFilter(String filterName) {
        switch (filterName) {
            case UPDATE_QUANTITY_BY_ID_FILTER:
            case DELETE_BY_ID_FILTER:
            case UPDATE_STATE_BY_ID_FILTER:
            case SELECT_ID_FILTER:
            case UPDATE_STATE_AND_ACTUAL_START_TIME_BY_ID_FILTER:
                return DBConstants.TABLE_JOB_CARD_JOBID + " = ?";
            case SELECT_LOCATION_ID_FILTER:
                return DBConstants.TABLE_JOB_CARD_LOCATION_ID + " = ?";
            case UPDATE_JCODE_PARTCOUNT_STARTTIME_ENDTIME_BY_ID_AND_LOCATION_ID_FILTER:
            case UPDATE_QUANTITY_BY_ID_AND_LOCATION_ID_FILTER:
            case SELECT_ID_AND_LOCATION_ID_FILTER:
                return DBConstants.TABLE_JOB_CARD_JOBID +" = ? and " +DBConstants.TABLE_JOB_CARD_LOCATION_ID + " = ?";
        }
        return "";
    }

    @Override
    public String[] FilterArgs(DBJobCardTableRowModel dataModel, String filterName) {
        switch (filterName) {
            case UPDATE_QUANTITY_BY_ID_FILTER:
            case DELETE_BY_ID_FILTER:
            case UPDATE_STATE_BY_ID_FILTER:
            case SELECT_ID_FILTER:
            case UPDATE_STATE_AND_ACTUAL_START_TIME_BY_ID_FILTER:
                return new String[]{String.valueOf(dataModel.getJobCardId())};
            case SELECT_LOCATION_ID_FILTER:
                return new String[]{String.valueOf(dataModel.getLocationId())};
            case UPDATE_JCODE_PARTCOUNT_STARTTIME_ENDTIME_BY_ID_AND_LOCATION_ID_FILTER:
            case UPDATE_QUANTITY_BY_ID_AND_LOCATION_ID_FILTER:
            case SELECT_ID_AND_LOCATION_ID_FILTER:
                return new String[]{ String.valueOf(dataModel.getJobCardId()), String.valueOf(dataModel.getLocationId())};
        }

        return new String[]{};
    }

    @Override
    public ContentValues UpdateFieldSet(DBJobCardTableRowModel dataModel, String filterName) {
        ContentValues values = new ContentValues();

        switch (filterName) {
            case UPDATE_STATE_BY_ID_FILTER:
                values.put(DBConstants.TABLE_JOB_CARD_PART_STATE, dataModel.getState());
                return values;
            case UPDATE_STATE_AND_ACTUAL_START_TIME_BY_ID_FILTER:
                values.put(DBConstants.TABLE_JOB_CARD_PART_STATE, dataModel.getState());
                values.put(DBConstants.TABLE_JOB_CARD_PART_ACTUAL_START_DATE, dataModel.getActualStartDate().getTime());
                return values;
            case UPDATE_QUANTITY_BY_ID_AND_LOCATION_ID_FILTER:
            case UPDATE_QUANTITY_BY_ID_FILTER:
                values.put(DBConstants.TABLE_JOB_CARD_PART_COMPLETED_COUNT, dataModel.getCompletedCount());
                return values;
            case UPDATE_JCODE_PARTCOUNT_STARTTIME_ENDTIME_BY_ID_AND_LOCATION_ID_FILTER:
                values.put(DBConstants.TABLE_JOB_CARD_CODE, dataModel.getJobCode());
                values.put(DBConstants.TABLE_JOB_CARD_PART_COUNT, dataModel.getPartCount());
                values.put(DBConstants.TABLE_JOB_CARD_PART_START_DATE, dataModel.getStartDate().getTime());
                values.put(DBConstants.TABLE_JOB_CARD_PART_END_DATE,  dataModel.getEndDate().getTime());
                return values;
        }
        return values;
    }
}
