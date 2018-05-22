package spl.hkt.opensoach.splapp.util;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;

import java.lang.reflect.Type;
import java.text.DateFormat;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Calendar;
import java.util.Date;
import java.util.List;
import java.util.TimeZone;

import spl.hkt.opensoach.splapp.Constants;
import spl.hkt.opensoach.splapp.model.communication.PacketChartConfigurationModel;
import spl.hkt.opensoach.splapp.model.communication.PacketModel;
import spl.hkt.opensoach.splapp.model.communication.PacketTaskModel;
import spl.hkt.opensoach.splapp.viewModels.ChartViewModel;
import spl.hkt.opensoach.splapp.viewModels.TaskRowViewModel;

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


    public static ChartViewModel getChartViewModel() {
        ChartViewModel chartViewModel = new ChartViewModel();

        TypeToken<PacketModel<PacketChartConfigurationModel>> typeToken = new TypeToken<PacketModel<PacketChartConfigurationModel>>() {
        };
        Type packetType = typeToken.getType();
        //packetDecodeResultModel.Packet.Payload = new Gson().fromJson(packet, packetType);

        String packet = Constants.tempString1;
        PacketModel<PacketChartConfigurationModel> packetChartConfigurationModel =
                (PacketModel<PacketChartConfigurationModel>) new Gson().fromJson(packet, packetType);

        PacketChartConfigurationModel chartConfigModel = packetChartConfigurationModel.Payload;
        ArrayList<String> tableColumnTitleList = new ArrayList<String>();
        ArrayList<String> tableRowTitleList = new ArrayList<String>();

        chartViewModel.setLocationName(chartConfigModel.ChartName);
        chartViewModel.setSlotInterval(chartConfigModel.SlotInterval);
        chartViewModel.setTaskStartTime(CommonUtility.getDateTimeStamp(chartConfigModel.StartTime));
        chartViewModel.setTaskEndTime(CommonUtility.getDateTimeStamp(chartConfigModel.EndTime));

        ArrayList<TaskRowViewModel> taskRowViewModelList = new ArrayList<TaskRowViewModel>();
        List<PacketTaskModel> tasksList = chartConfigModel.Tasks;
        for (int i = 0; i < tasksList.size(); i++) {

            TaskRowViewModel taskRowViewModel = new TaskRowViewModel();
            taskRowViewModel.setTaskName(tasksList.get(i).TaskName);
            taskRowViewModel.setTaskStartTime(CommonUtility.getDateTimeStamp(chartConfigModel.StartTime));
            taskRowViewModel.setTaskEndTime(CommonUtility.getDateTimeStamp(chartConfigModel.EndTime));
            taskRowViewModel.setTaskOrder(tasksList.get(i).TaskOrder);
            taskRowViewModel.setTaskID(tasksList.get(i).TaskID);

            tableRowTitleList.add(tasksList.get(i).TaskName);

           /* ArrayList<CellViewModel> cellViewModelList = new ArrayList<CellViewModel>();
            int cellCount = (chartConfigModel.EndTime - chartConfigModel.StartTime) / chartConfigModel.SlotInterval;
            for (int j = 0; j < cellCount; j++) {
                CellViewModel cellViewModel = new CellViewModel();

                cellViewModel.setTaskID(tasksList.get(i).TaskID);
                cellViewModel.setTaskName(tasksList.get(i).TaskName);
                cellViewModel.setCellStartTime(CommonUtility.getDateTimeStamp(chartConfigModel.StartTime + chartConfigModel.SlotInterval * (j)));
                cellViewModel.setCellEndTime(CommonUtility.getDateTimeStamp(chartConfigModel.StartTime + chartConfigModel.SlotInterval * (j + 1)));
                cellViewModel.setCellState(Constants.CELL_STATE_NOT_AVAILABLE);
                cellViewModel.setCellColumnID(j);
                //TODO Need to confirm slotId = columnId
                cellViewModel.setSlotID(j);
                cellViewModel.setCellRowID(i);

                cellViewModelList.add(cellViewModel);
            }
            taskRowViewModel.setCellViewModelList(cellViewModelList);
            taskRowViewModelList.add(taskRowViewModel);*/
        }
        chartViewModel.setTableRowTitleList(tableRowTitleList);
        chartViewModel.setTaskRowViewModelList(taskRowViewModelList);

        int cellCount = (chartConfigModel.EndTime - chartConfigModel.StartTime) / chartConfigModel.SlotInterval;
        for (int j = 0; j < cellCount; j++) {
            tableColumnTitleList.add(CommonUtility.getTimeHHMM(chartConfigModel.StartTime + chartConfigModel.SlotInterval * j));
        }
        chartViewModel.setTableColumnTitleList(tableColumnTitleList);

        return chartViewModel;
    }
}
