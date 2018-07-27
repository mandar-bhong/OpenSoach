package com.opensoach.hospital.Utility;

import org.apache.log4j.Logger;

/**
 * Created by samir.s.bukkawar on 5/9/2017.
 */

public class SPLLogger {
    private static SPLLogger mInstance;
    private Logger mLog;

    private SPLLogger() {
        // create logger
        mLog = Logger.getLogger("SPLAPP");
    }

    public static SPLLogger getInstance() {
        if (mInstance == null) {
            mInstance = new SPLLogger();
        }
        return mInstance;
    }

    public void DEBUG(String msg) {
        mLog.debug(msg);
    }

    public void INFO(String msg) {
        mLog.info(msg);
    }

    public void WARN(String msg) {
        mLog.warn(msg);
    }

    public void ERROR(String msg) {
        mLog.error(msg);
    }
}