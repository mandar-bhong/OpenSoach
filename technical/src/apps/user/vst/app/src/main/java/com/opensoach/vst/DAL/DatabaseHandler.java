package com.opensoach.vst.DAL;

import android.content.Context;
import android.database.sqlite.SQLiteDatabase;
import android.database.sqlite.SQLiteOpenHelper;
import android.util.Log;

import com.opensoach.vst.Constants.DBCreateTableConstants;


public class DatabaseHandler extends SQLiteOpenHelper {

    private final String TAG = "DatabaseHandler";


    private static DatabaseHandler sInstance;

    public DatabaseHandler(Context context) {
        super(context, DBConstants.DATABASE_NAME, null, DBConstants.DATABASE_VERSION);
    }

  /*  public static synchronized DatabaseHandler getInstance(Context context) {

        // Use the application context, which will ensure that you
        // don't accidentally leak an Activity's context.
        if (sInstance == null) {
            sInstance = new DatabaseHandler(context.getApplicationContext());
        }
        return sInstance;
    }*/

    // Creating Tables
    @Override
    public void onCreate(SQLiteDatabase db) {

        //TODO: This tables need to remove
        db.execSQL(DBConstants.CREATE_TABLE_CHART);
        db.execSQL(DBConstants.CREATE_TABLE_TASK);
        db.execSQL(DBConstants.CREATE_TABLE_CHART_DATA);

        db.execSQL(DBCreateTableConstants.CREATE_TABLE_LOCATION);
        db.execSQL(DBCreateTableConstants.CREATE_TABLE_AUTH_LOCATION);
        db.execSQL(DBCreateTableConstants.CREATE_TABLE_SERVICE_CONFIG);
        db.execSQL(DBCreateTableConstants.CREATE_TABLE_SERVICE_TASK_DATA);
        db.execSQL(DBCreateTableConstants.CREATE_TABLE_TASK_DATA);
    }

    // Upgrading database
    @Override
    public void onUpgrade(SQLiteDatabase db, int oldVersion, int newVersion) {
        Log.w(TAG, "Upgrading database from version "
                + oldVersion + " to "
                + newVersion + " which destroys all old data");
        // Drop older table if existed
        // db.execSQL("DROP TABLE IF EXISTS " + DBKey.TABLE_HEADER);

        // Create tables again
        onCreate(db);
    }
}
