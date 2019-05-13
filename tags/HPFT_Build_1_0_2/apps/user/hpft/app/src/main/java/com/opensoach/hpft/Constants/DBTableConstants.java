package com.opensoach.hpft.Constants;

/**
 * Created by Mandar on 08-08-2018.
 */

public class DBTableConstants {

    public static final String TABLE_LOCATION = "TABLE_LOCATION";
    public static final String TABLE_LOCATION_ID = "locationId";
    public static final String TABLE_LOCATION_CAT = "locationCat";
    public static final String TABLE_LOCATION_NAME = "locationName";


    public static final String TABLE_AUTH_LOCATION = "TABLE_AUTH_LOCATION";
    public static final String TABLE_AUTH_LOCATION_LOCATIONID = "locationId";
    public static final String TABLE_AUTH_LOCATION_AUTHCODE = "authCodeJSON";



    public static final String TABLE_SERVICE_CONFIG = "TABLE_SERVICE_CONFIG";
    public static final String TABLE_SERVICE_CONFIG_SerInID= "serInID";
    public static final String TABLE_SERVICE_CONFIG_ConfTypeCode= "confTypeCode";
    public static final String TABLE_SERVICE_CONFIG_ServConfID= "servConfID";
    public static final String TABLE_SERVICE_CONFIG_ConfigName= "configName";
    public static final String TABLE_SERVICE_CONFIG_ServConfJSON= "servConfJSON";
    public static final String TABLE_SERVICE_CONFIG_MedicalDetailsJSON= "medicalDetailsJSON";
    public static final String TABLE_SERVICE_CONFIG_PatientDetailsJSON= "patientDetailsJSON";


    public static final String TABLE_SERVICE_CONFIG_DATA = "TABLE_SERVICE_CONFIG_DATA";

    public static final String TABLE_SERVICE_TASK_DATA = "TABLE_SERVICE_TASK_DATA";
    public static final String TABLE_SERVICE_TASK_DATA_SerInID= "serInID";
    public static final String TABLE_SERVICE_TASK_DATA_ServConfID= "servConfID";
    public static final String TABLE_SERVICE_TASK_DATA_LOCATION_ID = "locationId";
    public static final String TABLE_SERVICE_TASK_DATA_TIME = "entrytime";
    public static final String TABLE_SERVICE_TASK_DATA_SLOT_START_TIME = "slotStartTime";
    public static final String TABLE_SERVICE_TASK_DATA_SLOT_END_TIME = "slotEndTime";
    public static final String TABLE_SERVICE_TASK_DATA_JSON = "data";
    public static final String TABLE_SERVICE_TASK_DATA_SERVER_SYNC = "isSynced";
    public static final String TABLE_SERVICE_TASK_DATA_AUTH_CODE = "authCode";


    public static final String TABLE_TASK_DATA = "TABLE_TASK_DATA";
    public static final String TABLE_TASK_DATA_SerInID= "serInID";
    public static final String TABLE_TASK_DATA_LOCATION_ID = "locationId";
    public static final String TABLE_TASK_DATA_TITLE= "title";
    public static final String TABLE_TASK_DATA_TIME = "entrytime";
    public static final String TABLE_TASK_DATA_VALUE = "value";
    public static final String TABLE_TASK_DATA_COMMENT = "comment";
    public static final String TABLE_TASK_DATA_SLOT_START_TIME = "slotStartTime";
    public static final String TABLE_TASK_DATA_SLOT_END_TIME = "slotEndTime";
    public static final String TABLE_TASK_DATA_SERVER_SYNC = "isSynced";
    public static final String TABLE_TASK_DATA_AUTH_CODE = "authCode";
}
