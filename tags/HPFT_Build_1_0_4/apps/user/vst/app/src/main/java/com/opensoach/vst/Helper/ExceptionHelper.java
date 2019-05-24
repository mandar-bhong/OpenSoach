package com.opensoach.vst.Helper;

import android.content.Context;
import android.util.Log;

import com.opensoach.vst.Utility.AppLogger;

public class ExceptionHelper implements Thread.UncaughtExceptionHandler {
    private Thread.UncaughtExceptionHandler androidDefaultUEH;

    private final Context context;
    private final Thread.UncaughtExceptionHandler rootHandler;


    public ExceptionHelper(Context context) {
        this.context = context;
        // we should store the current exception handler -- to invoke it for all not handled exceptions ...
        rootHandler = Thread.getDefaultUncaughtExceptionHandler();
        // we replace the exception handler now with us -- we will properly dispatch the exceptions ...
        Thread.setDefaultUncaughtExceptionHandler(this);
    }

    @Override
    public void uncaughtException(final Thread thread, final Throwable ex) {
        try {
            Exception exception = new Exception(ex);
            AppLogger.getInstance().Log(exception);
        } catch (Exception e) {
            Log.e("Unhandled", "Exception Logger failed!", e);
        }
    }
}
