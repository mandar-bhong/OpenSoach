package com.opensoach.hpft.Scheduler;

import android.app.AlarmManager;
import android.app.PendingIntent;
import android.content.Context;
import android.content.Intent;
import android.util.Log;

import java.util.Calendar;

import static android.content.Context.ALARM_SERVICE;

/**
 * Created by samir.s.bukkawar on 3/5/2017.
 */

public class ScheduleManager {
    //private AlarmManager mAlarmManager;
    //private PendingIntent mPendingIntent;
    private final int UNIQUE_REQUEST_CODE = 123456;

    //singleton
    public void init(Context context) {

    }

    public void startScheduler(Context context, int startTime, int endTime, int slotInterval) {
/*
        int startTimeHr = (int) Math.floor(startTime / 60);
        int startTimeMin = startTime % 60;
        int slotIntervalMillsec = 1000 * 60 * slotInterval;
        Intent alarmIntent = new Intent(context, AlarmReveiver.class);
        PendingIntent mPendingIntent = PendingIntent.getBroadcast(context, UNIQUE_REQUEST_CODE,
                alarmIntent, PendingIntent.FLAG_CANCEL_CURRENT);

        // Set the alarm to start at @startTime
        Calendar calendar = Calendar.getInstance();
        // calendar.setTimeInMillis(SystemClock.elapsedRealtime());
        calendar.setTimeInMillis(Calendar.getInstance().getTimeInMillis());

        //DateFormat df = new SimpleDateFormat("EEE, d MMM yyyy, HH:mm:ss");
        //Log.i("AlarmReveiver", "start schedular date : " + df.format(calendar.getTimeInMillis()));

        calendar.set(Calendar.HOUR_OF_DAY, startTimeHr);
        calendar.set(Calendar.MINUTE, startTimeMin);
        calendar.set(Calendar.SECOND, 0);
        calendar.set(Calendar.MILLISECOND, 0);*/

        Calendar calendar = Calendar.getInstance();
        calendar.set(Calendar.HOUR_OF_DAY, 0);
        calendar.set(Calendar.MINUTE, 21);
        calendar.set(Calendar.SECOND, 0);
        calendar.set(Calendar.MILLISECOND, 0);

        Intent alarmIntent = new Intent(context, AlarmReveiver.class);
        PendingIntent mPendingIntent = PendingIntent.getBroadcast(context, UNIQUE_REQUEST_CODE,
                alarmIntent, PendingIntent.FLAG_CANCEL_CURRENT);

        //Log.i("AlarmReveiver", "start schedular date : " + df.format(calendar.getTimeInMillis()));

        //Log.i("AlarmManager", "start schedular slotInterval : " + slotIntervalMillsec);

        // Schedule the alarm!
        AlarmManager mAlarmManager =
                (AlarmManager) context.getSystemService(ALARM_SERVICE);
        mAlarmManager.setRepeating(AlarmManager.RTC, 0,
                1000 * 1 * 1, mPendingIntent);
    }

    public void stopSchedular(Context context) {
        Intent alarmIntent = new Intent(context, AlarmReveiver.class);
        PendingIntent mPendingIntent = PendingIntent.getBroadcast(context, UNIQUE_REQUEST_CODE,
                alarmIntent, PendingIntent.FLAG_CANCEL_CURRENT);
        AlarmManager mAlarmManager =
                (AlarmManager) context.getSystemService(ALARM_SERVICE);

        if (mAlarmManager != null && mPendingIntent != null) {

            Log.i("AlarmManager", "stopSchedular if");
            mAlarmManager.cancel(mPendingIntent);
            mPendingIntent.cancel();
        } else
            Log.i("AlarmManager", "stopSchedular else");
    }

}
