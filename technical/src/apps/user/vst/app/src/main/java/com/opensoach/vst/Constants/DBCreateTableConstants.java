package com.opensoach.vst.Constants;

import static com.opensoach.vst.Constants.DBTableConstants.*;

/**
 * Created by Mandar on 08-08-2018.
 */

public class DBCreateTableConstants {

    public static final String DB_FIELD_TINY_INTEGER = " TINYINT";
    public static final String DB_FIELD_INTEGER = " INTEGER";
    public static final String DB_FIELD_VARCHAR = " VARCHAR";
    public static final String DB_FIELD_TINY_INTEGER_AND = " TINYINT,";
    public static final String DB_FIELD_INTEGER_AND = " INTEGER,";
    public static final String DB_FIELD_VARCHAR_AND = " VARCHAR,";

    public static final String CREATE_TABLE_LOCATION = "CREATE TABLE " +
            TABLE_LOCATION + "( " +
            TABLE_LOCATION_ID + DB_FIELD_INTEGER_AND +
            TABLE_LOCATION_NAME + DB_FIELD_VARCHAR_AND +
            TABLE_LOCATION_CAT + DB_FIELD_VARCHAR + ")";

    public static final String CREATE_TABLE_AUTH_LOCATION = "CREATE TABLE " +
            TABLE_AUTH_LOCATION + "( " +
            TABLE_AUTH_LOCATION_LOCATIONID + DB_FIELD_INTEGER_AND +
            TABLE_AUTH_LOCATION_AUTHCODE + DB_FIELD_VARCHAR + ")";

    public static final String CREATE_TABLE_SERVICE_CONFIG = "CREATE TABLE " +
            TABLE_SERVICE_CONFIG + "(" +
            TABLE_SERVICE_CONFIG_ConfTypeCode + DB_FIELD_VARCHAR_AND+
            TABLE_SERVICE_CONFIG_ServConfID + DB_FIELD_INTEGER_AND +
            TABLE_SERVICE_CONFIG_SerInID +DB_FIELD_INTEGER_AND +
            TABLE_SERVICE_CONFIG_ConfigName + DB_FIELD_VARCHAR_AND +
            TABLE_SERVICE_CONFIG_ServConfJSON + DB_FIELD_VARCHAR_AND +
            TABLE_SERVICE_CONFIG_MedicalDetailsJSON + DB_FIELD_VARCHAR_AND+
            TABLE_SERVICE_CONFIG_PatientDetailsJSON + DB_FIELD_VARCHAR +
            ")";

    public static final String CREATE_TABLE_SERVICE_TASK_DATA = "CREATE TABLE " +
            TABLE_SERVICE_TASK_DATA + "(" +
            TABLE_SERVICE_TASK_DATA_ServConfID + DB_FIELD_INTEGER_AND +
            TABLE_SERVICE_TASK_DATA_SerInID + DB_FIELD_INTEGER_AND +
            TABLE_SERVICE_TASK_DATA_LOCATION_ID + DB_FIELD_INTEGER_AND+
            TABLE_SERVICE_TASK_DATA_TIME + DB_FIELD_VARCHAR_AND +
            TABLE_SERVICE_TASK_DATA_SLOT_START_TIME + DB_FIELD_VARCHAR_AND +
            TABLE_SERVICE_TASK_DATA_SLOT_END_TIME + DB_FIELD_VARCHAR_AND +
            TABLE_SERVICE_TASK_DATA_JSON + DB_FIELD_VARCHAR_AND +
            TABLE_SERVICE_TASK_DATA_SERVER_SYNC + DB_FIELD_TINY_INTEGER + "  DEFAULT 0, " +
            TABLE_SERVICE_TASK_DATA_AUTH_CODE + DB_FIELD_VARCHAR +
            ")";

    public static final String CREATE_TABLE_TASK_DATA = "CREATE TABLE " +
            TABLE_TASK_DATA + "("+
            TABLE_TASK_DATA_SerInID + DB_FIELD_INTEGER_AND +
            TABLE_TASK_DATA_LOCATION_ID + DB_FIELD_INTEGER_AND +
            TABLE_TASK_DATA_TITLE + DB_FIELD_VARCHAR_AND +
            TABLE_TASK_DATA_TIME + DB_FIELD_VARCHAR_AND +
            TABLE_TASK_DATA_SLOT_START_TIME + DB_FIELD_VARCHAR_AND +
            TABLE_TASK_DATA_SLOT_END_TIME + DB_FIELD_VARCHAR_AND +
            TABLE_TASK_DATA_VALUE + DB_FIELD_VARCHAR_AND +
            TABLE_TASK_DATA_COMMENT + DB_FIELD_VARCHAR_AND +
            TABLE_TASK_DATA_SERVER_SYNC +  DB_FIELD_TINY_INTEGER + "  DEFAULT 0, " +
            TABLE_TASK_DATA_AUTH_CODE + DB_FIELD_VARCHAR +
            ")";

}
