package spl.hkt.opensoach.splapp.helper;

/**
 * Created by Mandar on 2/25/2017.
 */

public class CommandConstants {
    //Category
    public static final int CMD_CAT_DEVICE_REG = 1;
    public static final int CMD_DEVICE_REGISTRATION = 1;


    //Category
    public static final int CMD_CAT_CONFIG = 2;
    public static final int CMD_CONFIG_DEVICE_SYNC_COMPLETED = 1;
    public static final int CMD_CONFIG_LOCATION_SYNC = 2;
    public static final int CMD_CONFIG_CHART_CONFIG = 7;
    public static final int CMD_CONFIG_SERVER_SYNC_COMPLETED = 4;
    public static final int CMD_CONFIG_LOCATION_HCODE = 5;
    public static final int CMD_CONFIG_LOCATION_AUTH_CODE_ADDED = 11;
    public static final int CMD_CONFIG_LOCATION_AUTH_CODE_ASSOCIATED = 8;
    public static final int CMD_CONFIG_LOCATION_AUTH_CODE_REMOVED = 9;

    public static final int CMD_CAT_DATA = 3;//Category
    public static final int CMD_DATA_CHART_DATA = 1;
    public static final int CMD_DATA_BATTERY_LEVEL_DATA = 2;
    public static final int CMD_DATA_COMPLAINT_DATA = 256;
    public static final int CMD_DATA_FEEDBACK_DATA = 257;


    //Category
    public static final int CMD_CAT_NOTIFICATION = 4;
    public static final int CMD_DEVICE_NOTIFICATION_DATA_RECEIVE = 1;
    //Category
    public static final int CMD_CAT_ALARM_EVENT = 5;
    //Category
    public static final int CMD_CAT_ACK = 6;
    public static final int CMD_ACK_DEVICE_REG = 1;
    public static final int CMD_ACK_DEVICE_COMPLAINT_REGISTRATION = 2;
    public static final int CMD_ACK_DEVICE_NOTIFICATION = 3;
    public static final int CMD_ACK_CHART_DATA = 4;

    //Category
    public static final int CMD_CAT_DEVICE_STATUS = 7;
    public static final int CMD_DEVICE_STATUS_BATTERY_STAUS = 1;


    public static final int DEVICE_DATA_COMMAND_CHART_DATA = 1;
    public static final int DEVICE_DATA_COMMAND_CHART_UNSYNC_DATA = 2;

    /*public static final int  PACKET_SYNC_WAIT_STATUS_NOT_STARTED = 0;
    public static final int  PACKET_SYNC_WAIT_STATUS_STARTED = 1;
    public static final int  PACKET_SYNC_WAIT_STATUS_SYNC_CONDITONAL = 2;*/

}
