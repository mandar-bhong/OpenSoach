package spl.hkt.opensoach.splapp.dal;


/**
 * Created by Mandar on 3/11/2017.
 */

public class DBConstants {

    public static final String DATABASE_NAME = "SPLAPP_DATABASE.db";
    public static final int DATABASE_VERSION = 1;

    public static final String DB_FIELD_INTEGER = "INTEGER";
    public static final String DB_FIELD_VARCHAR = "VARCHAR";


    public static final String TABLE_CHART = "TABLE_CHART";
    public static final String TABLE_CHART_CHART_ID = "chartId";
    public static final String TABLE_CHART_LOCATION_ID = "locationId";
    public static final String TABLE_CHART_SERVER_CHART_ID = "serverChartId";
    public static final String TABLE_CHART_CHART_PAYLOAD = "chartPayload";
    public static final String TABLE_CHART_CHART_NAME = "chartName";
    public static final String TABLE_CHART_CHART_DISP_START_DATE = "chartDispStartDate";
    public static final String TABLE_CHART_CHART_DISP_END_DATE = "chartDispEndDate";


    public static final String TABLE_TASKS = "TABLE_TASKS";
    public static final String TABLE_TASKS_TASK_ID = "taskId";
    public static final String TABLE_TASKS_TASK_NAME = "taskName";
    public static final String TABLE_TASKS_TASK_ORDER = "taskOrder";

    public static final String TABLE_CHART_DATA = "TABLE_CHART_DATA";
    public static final String TABLE_CHART_DATA_CHART_ID = TABLE_CHART_CHART_ID;
    public static final String TABLE_CHART_DATA_TASK_NAME = TABLE_TASKS_TASK_NAME;
    public static final String TABLE_CHART_DATA_SLOT_ID = "slotId";
    public static final String TABLE_CHART_DATA_ENTRY_TIME = "entryTime";
    public static final String TABLE_CHART_DATA_SLOT_START_TIME = "slotStartTime";
    public static final String TABLE_CHART_DATA_SLOT_END_TIME = "slotEndTime";
    public static final String TABLE_CHART_DATA_STATE = "cellState";
    public static final String TABLE_CHART_DATA_DAY = "chartDay";
    public static final String TABLE_CHART_DATA_SERVER_SYNC = "isSynced";
    public static final String TABLE_CHART_DATA_AUTH_CODE = "authCode";


    public static final String TABLE_CONFIG = "TABLE_CONFIG";
    public static final String TABLE_CONFIG_CONFIG_KEY = "CONFIG_KEY";
    public static final String TABLE_CONFIG_CONFIG_VALUE = "CONFIG_VALUE";

    public static final String TABLE_LOCATION = "TABLE_LOCATION";
    public static final String TABLE_LOCATION_ID = "locationId";
    public static final String TABLE_LOCATION_CAT = "locationCat";
    public static final String TABLE_LOCATION_NAME = "locationName";


    public static final String TABLE_AUTH_LOCATION = "TABLE_AUTH_LOCATION";
    public static final String TABLE_AUTH_LOCATION_LOCATIONID = "locationId";
    public static final String TABLE_AUTH_LOCATION_AUTHCODE = "authCodeJSON";

    public static final String CREATE_TABLE_CHART = "CREATE TABLE " +
            TABLE_CHART + "( " +
            TABLE_CHART_CHART_ID + " INTEGER," +
            TABLE_CHART_CHART_NAME + " VARCHAR, " +
            TABLE_CHART_SERVER_CHART_ID + " INTEGER," +
            TABLE_CHART_LOCATION_ID + " INTEGER," +
            TABLE_CHART_CHART_PAYLOAD + " VARCHAR, " +
            TABLE_CHART_CHART_DISP_START_DATE + " INTEGER," +
            TABLE_CHART_CHART_DISP_END_DATE + " INTEGER )";

    public static final String CREATE_TABLE_TASK = "CREATE TABLE " +
            TABLE_TASKS + "( " +
            TABLE_TASKS_TASK_ID + " INTEGER," +
            TABLE_TASKS_TASK_NAME + " VARCHAR," +
            TABLE_TASKS_TASK_ORDER + " INTEGER" + ")";

    public static final String CREATE_TABLE_CONFIG = "CREATE TABLE " +
            TABLE_CONFIG + "( " +
            TABLE_CONFIG_CONFIG_KEY + " VARCHAR," +
            TABLE_CONFIG_CONFIG_VALUE + " VARCHAR" + ")";


    public static final String CREATE_TABLE_CHART_DATA = "CREATE TABLE " +
            TABLE_CHART_DATA + "( " +
            TABLE_CHART_DATA_CHART_ID + " INTEGER," +
            TABLE_CHART_DATA_TASK_NAME + " VARCHAR," +
            TABLE_CHART_DATA_SLOT_ID + " TINYINT, " +
            TABLE_CHART_DATA_ENTRY_TIME + " VARCHAR, " +
            TABLE_CHART_DATA_SLOT_START_TIME + " VARCHAR, " +
            TABLE_CHART_DATA_SLOT_END_TIME + " VARCHAR, " +
            TABLE_CHART_DATA_STATE + " TINYINT, " +
            TABLE_CHART_DATA_DAY + " VARCHAR, " +
            TABLE_CHART_DATA_AUTH_CODE + " VARCHAR, " +
            TABLE_CHART_DATA_SERVER_SYNC + " TINYINT DEFAULT 0 " +
            ")";


    public static final String CREATE_TABLE_LOCATION = "CREATE TABLE " +
            TABLE_LOCATION + "( " +
            TABLE_LOCATION_ID + " INTEGER," +
            TABLE_LOCATION_NAME + " VARCHAR," +
            TABLE_LOCATION_CAT + " VARCHAR" + ")";

    public static final String CREATE_TABLE_AUTH_LOCATION = "CREATE TABLE " +
            TABLE_AUTH_LOCATION + "( " +
            TABLE_AUTH_LOCATION_LOCATIONID + " INTEGER," +
            TABLE_AUTH_LOCATION_AUTHCODE + " VARCHAR" + ")";
}
