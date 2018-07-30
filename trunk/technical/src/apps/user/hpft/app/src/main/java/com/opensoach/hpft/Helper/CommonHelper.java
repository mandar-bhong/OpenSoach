package com.opensoach.hpft.Helper;

import com.google.gson.Gson;

import java.util.Calendar;
import java.util.Date;

import com.opensoach.hpft.Model.Communication.PacketModel;
import com.opensoach.hpft.Model.DB.DBChartTableRowModel;
import com.opensoach.hpft.Model.Server.DailyChartConfModel;
import com.opensoach.hpft.Model.Server.DailyChartTaskModel;
import com.opensoach.hpft.Model.View.ChartConfigModel;
import com.opensoach.hpft.Model.View.ChartConfigSlotModel;
import com.opensoach.hpft.Model.View.ChartConfigTaskModel;

/**
 * Created by Mandar on 2/25/2017.
 */

public class CommonHelper {


    public static String GetPacketJSON(PacketModel model) {
        return new Gson().toJson(model);
    }

    public static ChartConfigModel CreateChartModel(DBChartTableRowModel dbChartTableRowModel) {
        DailyChartConfModel dailyChartConfModel = new Gson().fromJson(dbChartTableRowModel.getChartPayload(), DailyChartConfModel.class);
        ChartConfigModel chartDataModel = new ChartConfigModel();
        chartDataModel.setChartId(dbChartTableRowModel.getChartId());
        chartDataModel.setChartName(dbChartTableRowModel.getChartName());
        chartDataModel.setSlotInterval(dailyChartConfModel.TimeConf.Interval);
        chartDataModel.setLocationId(dbChartTableRowModel.getLocationId());
        chartDataModel.setServerChartId(dbChartTableRowModel.getServerChartId());

        for (DailyChartTaskModel model : dailyChartConfModel.TaskConf.Tasks) {
            ChartConfigTaskModel chartTaskModel = new ChartConfigTaskModel();
            chartTaskModel.setTaskName(model.TaskName);
            chartDataModel.getTasks().put(model.TaskName, chartTaskModel);
            chartDataModel.getTaskList().add(chartTaskModel);
        }

        Calendar calChartStart = Calendar.getInstance();
        calChartStart.set(Calendar.HOUR_OF_DAY, 0);
        calChartStart.set(Calendar.MINUTE, dailyChartConfModel.TimeConf.StartTime);
        calChartStart.set(Calendar.SECOND, 0);
        calChartStart.set(Calendar.MILLISECOND, 0);
        Date chartStartTime = calChartStart.getTime();

        int chartEndHour = 0;

        if (dailyChartConfModel.TimeConf.EndTime <= dailyChartConfModel.TimeConf.StartTime) {
            chartEndHour = 24;
        }

        Calendar calChartEnd = Calendar.getInstance();
        calChartEnd.set(Calendar.HOUR_OF_DAY, chartEndHour);
        calChartEnd.set(Calendar.MINUTE, dailyChartConfModel.TimeConf.EndTime);
        calChartEnd.set(Calendar.SECOND, 0);
        calChartEnd.set(Calendar.MILLISECOND, 0);
        Date chartEndTime = calChartEnd.getTime();

        int slotIndex = 0;

        while (chartStartTime.getTime() < chartEndTime.getTime()) {
            ChartConfigSlotModel chartSlotModel = new ChartConfigSlotModel();
            chartSlotModel.setIndex(slotIndex);
            chartSlotModel.setStartTime(new Date(chartStartTime.getTime()));

            Calendar cal = Calendar.getInstance();
            cal.setTime(chartSlotModel.getStartTime());
            cal.add(Calendar.MINUTE, chartDataModel.getSlotInterval());

            chartSlotModel.setEndTime(cal.getTime());

            chartDataModel.getSlots().put(slotIndex, chartSlotModel);

            calChartStart.add(Calendar.MINUTE, dailyChartConfModel.TimeConf.Interval);
            chartStartTime = calChartStart.getTime();
            slotIndex++;
        }

        return chartDataModel;
    }

}
