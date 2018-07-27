package com.opensoach.hospital.DAL;


/**
 * Created by Mandar on 3/11/2017.
 */

public class DBConstants {

    public static final String DATABASE_NAME = "SME_APP_DATABASE.db";
    public static final int DATABASE_VERSION = 1;


    public static final int TABLE_JOB_TABLE_ID = 7;

    public static final int TABLE_ID_OPERATOR_CODE =1;

    public static final String DB_FIELD_TINY_INTEGER = " TINYINT";
    public static final String DB_FIELD_INTEGER = " INTEGER";
    public static final String DB_FIELD_VARCHAR = " VARCHAR";
    public static final String DB_FIELD_TINY_INTEGER_AND = " TINYINT,";
    public static final String DB_FIELD_INTEGER_AND = " INTEGER,";
    public static final String DB_FIELD_VARCHAR_AND = " VARCHAR,";


    public static final String TABLE_TASKS = "TABLE_TASKS";
    public static final String TABLE_TASKS_TASK_ID = "taskId";
    public static final String TABLE_TASKS_TASK_NAME = "taskName";
    public static final String TABLE_TASKS_TASK_ORDER = "taskOrder";


    public static final String TABLE_CONFIG = "TABLE_CONFIG";
    public static final String TABLE_CONFIG_CONFIG_KEY = "CONFIG_KEY";
    public static final String TABLE_CONFIG_CONFIG_VALUE = "CONFIG_VALUE";

    public static final String TABLE_TABLE_HASH = "TABLE_TABLE_HASH";
    public static final String TABLE_TABLE_HASH_ID = "tableId";
    public static final String TABLE_TABLE_HASH_CODE = "hashCode";

    public static final String TABLE_LOCATION = "TABLE_LOCATION";
    public static final String TABLE_LOCATION_ID = "locationId";
    public static final String TABLE_LOCATION_NAME = "locationName";

    public static final String TABLE_OPERATOR_CODE = "TABLE_OPERATOR_CODE";
    public static final String TABLE_OPERATOR_CODE_ID = "operatorId";
    public static final String TABLE_OPERATOR_CODE_CODE = "opCode";
    public static final String TABLE_OPERATOR_CODE_JSON = "jsonData";

    public static final String TABLE_AUTH_CODE = "TABLE_AUTH_CODE";
    public static final String TABLE_AUTH_CODE_AUTHCODE = "authCodeJSON";


    public static final String TABLE_JOB_CARD = "TABLE_JOB_CARD";
    public static final String TABLE_JOB_CARD_LOCATION_ID = "locationId";
    public static final String TABLE_JOB_CARD_JOBID = "jobCardId";
    public static final String TABLE_JOB_CARD_CUSTOMER = "customer";
    public static final String TABLE_JOB_CARD_PART_ID = "partId";
    public static final String TABLE_JOB_CARD_PART_COUNT = "partCount";
    public static final String TABLE_JOB_CARD_CODE = "jobCode";
    public static final String TABLE_JOB_CARD_PART_START_DATE = "startDate";
    public static final String TABLE_JOB_CARD_PART_END_DATE = "endDate";
    public static final String TABLE_JOB_CARD_PART_ACTUAL_START_DATE = "actualStartDate";
    public static final String TABLE_JOB_CARD_PART_ACTUAL_END_DATE = "actualEndDate";
    public static final String TABLE_JOB_CARD_PART_STATE = "state";
    public static final String TABLE_JOB_CARD_PART_COMMENTS = "comments";
    public static final String TABLE_JOB_CARD_PART_COMPLETED_COUNT = "completedCount";
    public static final String TABLE_JOB_CARD_JOB_CONFIG = "jobConfig";



    public static final String TABLE_ENGG_PART = "TABLE_ENGG_PART";
    public static final String TABLE_ENGG_PART_PART_ID = "partId";
    public static final String TABLE_ENGG_PART_PART_NO = "partNo";
    public static final String TABLE_ENGG_PART_PART_REVISION = "partRevision";
    public static final String TABLE_ENGG_PART_INTERNAL_PART_NO = "internalPartNo";
    public static final String TABLE_ENGG_PART_PROCESS = "process";
    public static final String TABLE_ENGG_PART_TOOL_JSON = "toolJSON";


    public static final String TABLE_PART_DRAWING = "TABLE_PART_DRAWING";
    public static final String TABLE_PART_DRAWING_ID = "drawingId";
    public static final String TABLE_PART_DRAWING_PART_ID = "partId";
    public static final String TABLE_PART_DRAWING_PATH = "path";


    public static final String TABLE_TOOL = "TABLE_TOOL";
    public static final String TABLE_TOOL_ID = "toolId";
    public static final String TABLE_TOOL_NAME = "toolName";
    public static final String TABLE_TOOL_SHORT_DESCRIPTION = "desc";


    public static final String CREATE_TABLE_TASK = "CREATE TABLE " +
            TABLE_TASKS + "( " +
            TABLE_TASKS_TASK_ID + " INTEGER," +
            TABLE_TASKS_TASK_NAME + " VARCHAR," +
            TABLE_TASKS_TASK_ORDER + " INTEGER" + ")";

    public static final String CREATE_TABLE_CONFIG = "CREATE TABLE " +
            TABLE_CONFIG + "( " +
            TABLE_CONFIG_CONFIG_KEY + " VARCHAR," +
            TABLE_CONFIG_CONFIG_VALUE + " VARCHAR" + ")";

    public static final String CREATE_TABLE_TABLE_HASH = "CREATE TABLE " +
            TABLE_TABLE_HASH + "( " +
            TABLE_TABLE_HASH_ID + " INTEGER," +
            TABLE_TABLE_HASH_CODE + " VARCHAR" + ")";


    public static final String CREATE_TABLE_LOCATION = "CREATE TABLE " +
            TABLE_LOCATION + "( " +
    TABLE_LOCATION_ID + " INTEGER," +
    TABLE_LOCATION_NAME + " VARCHAR" + ")";

    public static final String CREATE_TABLE_OPERATOR_CODE = "CREATE TABLE " +
            TABLE_OPERATOR_CODE + "( " +
            TABLE_OPERATOR_CODE_ID + " INTEGER," +
            TABLE_OPERATOR_CODE_CODE + " VARCHAR,"+
            TABLE_OPERATOR_CODE_JSON + " VARCHAR"
            + ")";


    public static final String CREATE_TABLE_AUTH_CODE = "CREATE TABLE " +
            TABLE_AUTH_CODE + "( " +
            TABLE_AUTH_CODE_AUTHCODE + " VARCHAR" + ")";


    public static final String CREATE_TABLE_JOB_CARD = "CREATE TABLE " +
            TABLE_JOB_CARD + "( " +
            TABLE_JOB_CARD_JOBID + DB_FIELD_INTEGER_AND +
            TABLE_JOB_CARD_LOCATION_ID + DB_FIELD_INTEGER_AND +
            TABLE_JOB_CARD_CUSTOMER + DB_FIELD_VARCHAR_AND +
            TABLE_JOB_CARD_PART_ID + DB_FIELD_INTEGER_AND +
            TABLE_JOB_CARD_PART_COUNT + DB_FIELD_INTEGER_AND +
            TABLE_JOB_CARD_CODE + DB_FIELD_VARCHAR_AND +
            TABLE_JOB_CARD_PART_START_DATE + DB_FIELD_VARCHAR_AND +
            TABLE_JOB_CARD_PART_END_DATE + DB_FIELD_VARCHAR_AND +
            TABLE_JOB_CARD_PART_ACTUAL_START_DATE + DB_FIELD_VARCHAR_AND +
            TABLE_JOB_CARD_PART_ACTUAL_END_DATE + DB_FIELD_VARCHAR_AND +
            TABLE_JOB_CARD_PART_STATE + DB_FIELD_TINY_INTEGER_AND +
            TABLE_JOB_CARD_PART_COMMENTS + DB_FIELD_VARCHAR_AND +
            TABLE_JOB_CARD_PART_COMPLETED_COUNT + DB_FIELD_INTEGER + " DEFAULT 0 " + ", "+
            TABLE_JOB_CARD_JOB_CONFIG +
            ")";


    public static final String CREATE_TABLE_ENGG_PART = "CREATE TABLE " +
            TABLE_ENGG_PART + "( " +
            TABLE_ENGG_PART_PART_ID + DB_FIELD_INTEGER_AND +
            TABLE_ENGG_PART_PART_NO + DB_FIELD_VARCHAR_AND +
            TABLE_ENGG_PART_PART_REVISION + DB_FIELD_VARCHAR_AND +
            TABLE_ENGG_PART_INTERNAL_PART_NO + DB_FIELD_VARCHAR_AND +
            TABLE_ENGG_PART_PROCESS + DB_FIELD_VARCHAR_AND +
            TABLE_ENGG_PART_TOOL_JSON + DB_FIELD_VARCHAR +
            ")";

    public static final String CREATE_TABLE_PART_DRAWING = "CREATE TABLE " +
            TABLE_PART_DRAWING + "( " +
            TABLE_PART_DRAWING_ID + DB_FIELD_INTEGER_AND +
            TABLE_PART_DRAWING_PART_ID + DB_FIELD_INTEGER_AND +
            TABLE_PART_DRAWING_PATH + DB_FIELD_VARCHAR +
            ")";


}
