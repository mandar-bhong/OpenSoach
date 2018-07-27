package com.opensoach.hospital.Helper;

import com.opensoach.hospital.Utility.AppLogger;

/**
 * Created by Mandar on 9/4/2017.
 */

public class ExceptionHelper implements Thread.UncaughtExceptionHandler {
    @Override
    public void uncaughtException(Thread t, Throwable e) {

        e.printStackTrace();
        AppLogger.getInstance().Log(AppLogger.LogLevel.Error, e.getMessage());
    }
}
