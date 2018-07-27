package com.opensoach.hospital.DAL;

/**
 * Created by samir.s.bukkawar on 2/26/2017.
 */

import android.content.Context;
import android.database.sqlite.SQLiteDatabase;
import android.database.sqlite.SQLiteOpenHelper;
import android.util.Log;


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

//        db.execSQL(DBConstants.CREATE_TABLE_TABLE_HASH);

        db.execSQL(DBConstants.CREATE_TABLE_LOCATION);
        db.execSQL(DBConstants.CREATE_TABLE_AUTH_CODE);
        db.execSQL(DBConstants.CREATE_TABLE_CONFIG);
        db.execSQL(DBConstants.CREATE_TABLE_JOB_CARD);
        db.execSQL(DBConstants.CREATE_TABLE_ENGG_PART);
        db.execSQL(DBConstants.CREATE_TABLE_PART_DRAWING);
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
