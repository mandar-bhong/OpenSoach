package com.opensoach.vst.Scheduler;

import android.content.BroadcastReceiver;
import android.content.Context;
import android.content.Intent;
import android.util.Log;

import java.text.DateFormat;
import java.text.SimpleDateFormat;
import java.util.Calendar;

import com.opensoach.vst.SPLApplication;

/**
 * Created by samir.s.bukkawar on 3/5/2017.
 */

public class AlarmReveiver extends BroadcastReceiver {
    @Override
    public void onReceive(Context context, Intent intent) {

        DateFormat df = new SimpleDateFormat("EEE, d MMM yyyy, HH:mm:ss");
        String date = df.format(Calendar.getInstance().getTime());
        Log.i("onReceive AlarmReveiver", " date : " + date);

        SPLApplication splApplication = SPLApplication.getInstance();
        splApplication.updateTimeChange();

        //UpdateUI
        //Update DB
        //WS Call
    }

}
