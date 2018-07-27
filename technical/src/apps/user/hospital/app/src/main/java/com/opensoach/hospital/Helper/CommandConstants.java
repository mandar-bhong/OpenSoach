package com.opensoach.hospital.Helper;

/**
 * Created by Mandar on 8/26/2017.
 */

public class CommandConstants {
    //Category
    public static final int CMD_CAT_DEVICE_REG = 1;
    public static final int CMD_DEVICE_REGISTRATION = 1;


    //Category
    public static final int CMD_CAT_CONFIG = 2;
    public static final int CMD_CONFIG_DEVICE_SYNC_COMPLETED = 1;
    public static final int CMD_CONFIG_LOCATION_SYNC = 2;
    public static final int CMD_CONFIG_JOB_CARD = 3;
    //public static final int CMD_CONFIG_CHART_CONFIG = 3;
    public static final int CMD_CONFIG_SERVER_SYNC_COMPLETED = 4;
    public static final int CMD_CONFIG_LOCATION_HCODE = 5;
    public static final int CMD_CONFIG_ENGG_PART = 6;
    public static final int CMD_CONFIG_PART_DRAWING = 7;

    public static final int CMD_CAT_DATA = 3;//Category
    public static final int CMD_DATA_JOB_CARD_DATA = 1;
    public static final int CMD_DATA_COMPLAINT_DATA = 2;
    //public static final int CMD_DATA_FEEDBACK_DATA = 3;
    public static final int CMD_DATA_UPDATE_JOB_UNIT = 3;
    public static final int CMD_DATA_START_JOB = 4;
    public static final int CMD_DATA_STOP_JOB = 5;
    public static final int CMD_DATA_DROP_JOB = 6;
    public static final int CMD_DATA_ABORT_JOB = 7;


    //Category
    public static final int CMD_CAT_NOTIFICATION = 4;
    public static final int CMD_DEVICE_NOTIFICATION_DATA_RECEIVE = 1;
    //Category
    public static final int CMD_CAT_ALARM_EVENT = 5;
    //Category
    public static final int CMD_CAT_ACK = 6;
    public static final int CMD_ACK_DEVICE_REG = 1;
    public static final int CMD_ACK_JOB_STATUS_START = 5;
    public static final int CMD_ACK_JOB_STATUS_STOP = 6;
    public static final int CMD_ACK_JOB_STATUS_DROP = 8;
    public static final int CMD_ACK_JOB_STATUS_ABORT = 9;
    public static final int CMD_ACK_JOB_STATUS_QUANTITY_UPDATE = 7;
    public static final int CMD_ACK_CHART_DATA = 5;


    //Category
    public static final int CMD_CAT_DEVICE_STATUS = 7;
    public static final int CMD_DEVICE_STATUS_BATTERY_STAUS = 1;


    public static final int CMD_CAT_SERVER_DATA_PUSHED = 8;
    public static final int CMD_SERVER_DATA_PUSHED_JOB_CARD = 12;
    public static final int CMD_SERVER_DATA_PUSHED_JOB_CARD_UPDATED = 15;
    public static final int CMD_SERVER_DATA_TABLE_ROW_DELETED = 10;
    public static final int CMD_SERVER_DATA_JOB_STATE_CHANGED = 20;

    public static final int DEVICE_DATA_COMMAND_UPDATE_JOB_UNIT = 1;
    public static final int DEVICE_DATA_COMMAND_START_JOB = 2;
    public static final int DEVICE_DATA_COMMAND_STOP_JOB = 3;
    public static final int DEVICE_DATA_COMMAND_ABORTED_JOB = 4;
    public static final int DEVICE_DATA_COMMAND_DROPED_JOB = 5;


    public static final int UI_CMD_BACKGROUND_START_JOB= 1;
    public static final int UI_CMD_BACKGROUND_UPDATE_JOB_UNIT = 2;
    public static final int UI_CMD_BACKGROUND_STOP_JOB= 3;
    public static final int UI_CMD_BACKGROUND_CLOSE_DIALOG= 4;


    public static final int  UI_VIEW_ACTION_CMD_JOB_START = 1;
    public static final int  UI_VIEW_ACTION_CMD_JOB_STOP = 2;
    public static final int  UI_VIEW_ACTION_CMD_JOB_ABORT = 3;
    public static final int  UI_VIEW_ACTION_CMD_JOB_DROP = 4;

}
