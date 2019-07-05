package com.opensoach.hpft.Utility;


import android.util.Log;

/**
 * Created by Mandar on 9/2/2017.
 */

public class AppLogger {

    private static AppLogger singleton;

    private LogLevel logLevel;
    private String AppName = "HPFTApplication";

    public enum LogLevel{
        Debug,
        Error
    }

    //region Properties
    public LogLevel getLogLevel() {
        return logLevel;
    }

    public void setLogLevel(LogLevel logLevel) {
        this.logLevel = logLevel;
    }

    //endregion Properties

    //region Constructor

    private AppLogger() {

    }

    //endregion Constructor

    //region Static Methods

    public static AppLogger getInstance() {

        if (singleton == null)
            singleton = new AppLogger();
        return singleton;
    }

    //endregion Static Methods

    //region Public Methods

    public void Log(LogLevel logLevel, String message){
        Log.i(AppName, "LogLevel: "+logLevel+", "+message);
    }

    public void Log(Exception ex) {
        Log(LogLevel.Error, ex);
    }

    public void Log(LogLevel logLevel,Exception ex){
        Log.i(AppName, "LogLevel: "+logLevel+", "+ex.getMessage());
    }

    public void Log(LogLevel logLevel,Exception ex, String moreInfo ){
        Log.i(AppName, "LogLevel: "+logLevel+", "+"MoreInfo: "+ moreInfo +", Exception: " + ex.getMessage());
    }

    public void Log(Exception ex, String moreInfo ){
        Log.i(AppName, "LogLevel: "+logLevel+", "+"MoreInfo: "+ moreInfo +", Exception: " + ex.getMessage());
    }

    //endregion Public Methods

    //region Private Methods
    //endregion Private Methods

}
