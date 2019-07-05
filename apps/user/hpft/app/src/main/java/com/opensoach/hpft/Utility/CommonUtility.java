package com.opensoach.hpft.Utility;

import java.text.DateFormat;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Calendar;
import java.util.Date;
import java.util.TimeZone;

/**
 * Created by samir.s.bukkawar on 2/20/2017.
 */

public class CommonUtility {

    private final String TAG = "CommonUtility";
    private static final String SERVER_DATE_FORMAT = "yyyy-MM-dd'T'HH:mm:ss.000'Z'";

    public static void getFirstAlarmTime(Date date, int interval) {

    }

    public static Date getDateTimeStamp(int intTime) {
        int timeHour = (int) Math.floor(intTime / 60);
        int timeMinute = intTime % 60;

        Calendar calendar = Calendar.getInstance();
        // calendar.setTimeInMillis(SystemClock.elapsedRealtime());
        calendar.setTimeInMillis(Calendar.getInstance().getTimeInMillis());

        //DateFormat df = new SimpleDateFormat("EEE, d MMM yyyy, HH:mm:ss");
        calendar.set(Calendar.HOUR_OF_DAY, timeHour);
        calendar.set(Calendar.MINUTE, timeMinute);
        calendar.set(Calendar.SECOND, 0);
        calendar.set(Calendar.MILLISECOND, 0);
        return calendar.getTime();
    }

    public static String getTimeHHMM(int intTime) {
        int timeHour = (int) Math.floor(intTime / 60);
        int timeMinute = intTime % 60;

        Calendar calendar = Calendar.getInstance();
        // calendar.setTimeInMillis(SystemClock.elapsedRealtime());
        calendar.setTimeInMillis(Calendar.getInstance().getTimeInMillis());

        calendar.set(Calendar.HOUR_OF_DAY, timeHour);
        calendar.set(Calendar.MINUTE, timeMinute);
        calendar.set(Calendar.SECOND, 0);
        calendar.set(Calendar.MILLISECOND, 0);
        DateFormat df = new SimpleDateFormat("hh:mm a");
        //Log.i("getTimeHHMM", "intTime : " + intTime + "  : " + df.format(calendar.getTimeInMillis()));
        return df.format(calendar.getTimeInMillis());
    }

    public static Date getCurrentTime() {

        Calendar calendar = Calendar.getInstance();
        // calendar.setTimeInMillis(SystemClock.elapsedRealtime());
        calendar.setTimeInMillis(Calendar.getInstance().getTimeInMillis());
        return calendar.getTime();
    }

    public static String getStringFromDate(Date date) {
        /*SimpleDateFormat dateFormat = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss z");
        dateFormat.setTimeZone(TimeZone.getTimeZone("UTC"));
        Log.i("DATE", "Str >> " + dateFormat.format(date));
        return dateFormat.format(date);*/

        TimeZone tz = TimeZone.getTimeZone("GMT");
        Calendar cal = Calendar.getInstance(tz);
        SimpleDateFormat sdf = new SimpleDateFormat(SERVER_DATE_FORMAT);
        sdf.setCalendar(cal);
        cal.setTime(date);
        //Log.i("DATE", "1 Str >> " + sdf.format(date));
       /* Log.i("DATE", "2 Str >> " + new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.SSS'z'").format(date));
        Log.i("DATE", "3 Str >> " + new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.000'z'").format(date));
        Log.i("DATE", "4 Str >> " + new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.000'Z'").format(date));
        Log.i("DATE", "5 Str >> " + new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.SSS'Z'").format(date));
        Log.i("DATE", "6 Str >> " + new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.000Z").format(date));*/
        return sdf.format(date);
    }


    public static Date getDateFromString(String strDate) {
        try {
            TimeZone tz = TimeZone.getTimeZone("GMT");
            Calendar cal = Calendar.getInstance(tz);
            SimpleDateFormat sdf = new SimpleDateFormat(SERVER_DATE_FORMAT);
            sdf.setCalendar(cal);
            //"2013-07-17T03:58:00.000Z"
            cal.setTime(sdf.parse(strDate));

            //Log.i("DATE", "IND date>> " + sdf.parse(strDate));
            //Log.i("DATE", "GMT date>> " + sdf.format(cal.getTime()));
            return cal.getTime();
        } catch (ParseException e) {
            return null;
        }
    }
}
