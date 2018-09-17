package com.opensoach.vst.Constants;

/**
 * Created by Mandar on 3/28/2017.
 */

public class ApplicationConstants {

    public static   final int UI_PROCESSING_STATERGY_CHART_DATA =1;
    public static   final int UI_PROCESSING_STATERGY_CHART_DATA_START_UP_DISPLAY =2;
    public static   final int UI_PROCESSING_STATERGY_AUTH_CODE_UPDATE = 3 ;
    public static   final int UI_PROCESSING_STATERGY_CARD_LIST_DATA =4;
    public static   final int UI_PROCESSING_STATERGY_TOKEN_CREATED =5;
    public static   final int UI_PROCESSING_STATERGY_TOKEN_LIST =6;



    public static final  int CHART_STATE_ENABLED = 1;
    public static final  int CHART_STATE_BLOCKED = 2;
    public static final  int CHART_STATE_ON_TIME = 3;
    public static final  int CHART_STATE_DELAYED = 4;


    public static final  int DB_CHART_STATE_ON_TIME = 1;
    public static final  int DB_CHART_STATE_DELAYED = 2;

    public static final String PACKET_DATE_FORMAT =  "yyyy-MM-dd'T'HH:mm:ss.SSS'Z'";

    public static   final String UI_DATE_FORMAT ="MMM dd h:mm a";

    public static int DISPLAY_MODE_JOB_CREATION_EDIT = 0;
    public static int DISPLAY_MODE_JOB_CREATION_SUMMARY = 1;
    public static int DISPLAY_MODE_JOB_EXECUTION = 2;

    public enum AppRunningMode {
        Token,
        JobCreation,
        JobExecution
    }


    public static int TOKEN_CREATED = 0;
    public static int TOKEN_CLAIMED = 1;
    public static int TOKEN_JOB_CREATED = 2;
    public static int TOKEN_JOB_EXECUTION_INPROGRESS = 3;
    public static int TOKEN_JOB_EXECUTED = 4;
    public static int TOKEN_JOB_DELIVERED = 5;



    public static final String APP_STORE_SERVICE_INST_ID = "ServiceInstId";
    public static final String APP_STORE_SERVICE_CONFIG_ID = "ServiceConfigId";
}