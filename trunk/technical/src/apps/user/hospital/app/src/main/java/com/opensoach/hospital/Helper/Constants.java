package com.opensoach.hospital.Helper;

/**
 * Created by Mandar on 8/26/2017.
 */

public class Constants {

    //public static  String SERVER_HOST_OR_IP = "servicepoint.live";
    public static  String SERVER_HOST_OR_IP = "139.162.75.182";//"192.168.1.65";//"139.162.75.182";
    public static  String WEB_SERVICE_URL = "http://"+SERVER_HOST_OR_IP+":80/";
    public static  String WEB_SOCKET_URL = "ws://"+SERVER_HOST_OR_IP+":8080/ws";
    public static  String API_VALIDATE_AUTH = "validatedevice";


    public static final long SCREEN_IDLE_TIMEOUT = 1 * 60 * 1000; // 5 min = 5 * 60 * 1000 ms

    //SHARED_PREFERENCE key
    public static final String KEY_LOCATION_NAME = "KEY_LOCATION_NAME";
    public static final String KEY_LOCATION_ID = "KEY_LOCATION_ID";

    //Cell State Key
    public static final int CELL_STATE_NOT_AVAILABLE = 0;
    public static final int CELL_STATE_AVAILABLE = 1;
    public static final int CELL_STATE_COMPLETED_ON_TIME = 2;
    public static final int CELL_STATE_COMPLETED_DELAYED = 3;

    public static enum NETWORK_STATE {
        NW_NOT_AVAILABLE,
        WEB_SOCKET_DISSCONNECTED,
        WEB_SOCKET_CONNECTED
    }


    public static final int RESPONSE_ACK_SUCCESS = 0;
    public static final int JOB_STATE_NEW = 0;
    public static final int JOB_STATE_STARTED = 1;
    public static final int JOB_STATE_COMPLETED = 2;
    public static final int JOB_STATE_DROPPED = 3;
    public static final int JOB_STATE_ABORTED = 4;


    public  static  final int JOB_STATUS_PENDING = 0;
    public  static  final int JOB_STATUS_INPROGRESS = 1;
    public  static  final int JOB_STATUS_COMPLETED = 2;
    public  static  final int JOB_STATUS_DROPED = 3;
    public  static  final int JOB_STATUS_ABORTED = 4;
}
