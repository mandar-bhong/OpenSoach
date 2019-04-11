package com.opensoach.vst.Communication;

import android.os.Handler;
import android.os.Looper;

/**
 * Created by samir.s.bukkawar on 2/18/2017.
 */

public class WebSocketThread extends Thread {

    public Handler handler;
    public static String responseString;
    private final String TAG = "WebSocketThread";

    @Override
    public void run() {
        Looper.prepare();
        handler = new WebSocketHandler();
        Looper.loop();
    }
}
