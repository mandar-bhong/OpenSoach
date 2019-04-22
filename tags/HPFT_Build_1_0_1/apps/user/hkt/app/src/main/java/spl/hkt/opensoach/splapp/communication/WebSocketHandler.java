package spl.hkt.opensoach.splapp.communication;


import android.os.Bundle;
import android.os.Handler;
import android.os.Message;
import android.util.Log;
import spl.hkt.opensoach.splapp.SPLApplication;
import spl.hkt.opensoach.splapp.model.communication.PacketHeaderModel;
import spl.hkt.opensoach.splapp.model.communication.PacketModel;
import spl.hkt.opensoach.splapp.processor.PacketProcessor;

import com.google.gson.Gson;

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
