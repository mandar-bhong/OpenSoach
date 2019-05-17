package com.opensoach.hpft.Communication;


import android.os.Bundle;
import android.os.Handler;
import android.os.Message;
import android.util.Log;

import com.opensoach.hpft.Processor.PacketProcessor;

/**
 * Created by samir.s.bukkawar on 2/18/2017.
 */

public class WebSocketHandler extends Handler {

    private final String TAG = "WebSocketHandler";

    int i = 0;

    @Override
    public void handleMessage(Message msg) {
        super.handleMessage(msg);

        Bundle b = msg.getData();
        String strResponse = b.getString ("RESPONSE_MESSAGE");

        Log.i("####", "bundle: " + i++ + ">> " + strResponse);
        PacketProcessor processor = new PacketProcessor();
        processor.handleMessage(msg);


        /*SPLApplication splApplication = SPLApplication.getInstance();
        splApplication.updateActivity(strResponse);*/

        /*runOnUiThread(new Runnable() {
            @Override
            public void run() {
               // textInfo.setText("Message 1");
            }
        });*/


    }

}
