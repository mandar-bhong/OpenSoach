package com.opensoach.vst.SharedPreference;

import android.content.Context;
import android.content.SharedPreferences;

/**
 * Created by samir.s.bukkawar on 3/2/2017.
 */

public class SharedPreferencesHelper {

    private static Context mContext;
    private static SharedPreferencesHelper mInstance;

    private static SharedPreferences sharedpreferences;
    public static final String APP_PREFERENCES = "SPL_APP_SHARED_PREFERANCES";

    private SharedPreferencesHelper() {

    }

    public static SharedPreferencesHelper getInstance(Context context) {
        if (mInstance == null) {
            mContext = context;
            return mInstance = new SharedPreferencesHelper();
        }
        return mInstance;
    }

    public void updateSharedPreference(String key, String value) {

        sharedpreferences = mContext.getSharedPreferences(APP_PREFERENCES, Context.MODE_PRIVATE);
        SharedPreferences.Editor editor = sharedpreferences.edit();
        editor.putString(key, value);
        editor.commit();
        editor.apply();
    }

    public String getDataFromSharedPreference(String key) {
        sharedpreferences = mContext.getSharedPreferences(APP_PREFERENCES, Context.MODE_PRIVATE);
        return sharedpreferences.getString(key, null);
    }
}