package spl.hkt.opensoach.splapp;

/**
 * Created by Samir on 15-Feb-17.
 */

public class Constants {

   //public static final String WEB_SOCKET_URL = "ws://servicepoint.live:8080/ws";
    public static final String WEB_SOCKET_URL = "ws://192.168.31.245:8080/ws";
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
        WEB_SOCKET_UNAUTHORIZED,
        WEB_SOCKET_DISSCONNECTED,
        WEB_SOCKET_CONNECTED
    }


    public static String tempString1 = "{\"header\":{\"crc\":\"\",\"category\":2,\"commandid\":3,\"seqid\":1,\"locationid\":1},\"payload\":{\"chartid\":1,\"locationid\":0,\"customerid\":0,\"locationcategoryid\":0,\"chartname\":\"Chart1\"," +
            "\"starttime\":1,\"endtime\":30,\"slotinterval\":1,\"tasks\":[{\"taskid\":1,\"taskname\":\"Task1\",\"taskorder\":0},{\"taskid\":2,\"taskname\":\"Task2\",\"taskorder\":0},{\"taskid\":3,\"taskname\":\"Task3\",\"taskorder\":0},{\"taskid\":4,\"taskname\":\"Task4\",\"taskorder\":0},{\"taskid\":5,\"taskname\":\"Task5\",\"taskorder\":0}]}}";
}