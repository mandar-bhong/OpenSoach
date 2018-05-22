package spl.hkt.opensoach.splapp.helper;

import android.content.Context;
import android.net.ConnectivityManager;
import android.net.NetworkInfo;

import com.google.gson.Gson;

import java.lang.reflect.Type;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Calendar;
import java.util.Collections;
import java.util.Date;
import java.util.List;

import spl.hkt.opensoach.splapp.model.communication.PacketChartConfigurationModel;
import spl.hkt.opensoach.splapp.model.communication.PacketModel;
import spl.hkt.opensoach.splapp.model.communication.PacketTaskModel;
import spl.hkt.opensoach.splapp.model.view.ChartConfigModel;
import spl.hkt.opensoach.splapp.model.view.ChartConfigSlotModel;
import spl.hkt.opensoach.splapp.model.view.ChartConfigTaskModel;

/**
 * Created by Mandar on 2/25/2017.
 */

public class CommonHelper {


    public  static String GetPacketJSON (PacketModel model){
      return   new Gson().toJson(model);
    }

    public  static ChartConfigModel CreateChartModel(int locationId, int chartId, String chartJSONData){

        PacketChartConfigurationModel packetModel = new Gson().fromJson(chartJSONData, PacketChartConfigurationModel.class);

        ChartConfigModel chartDataModel = new ChartConfigModel();
        chartDataModel.setChartId(packetModel.ChartID);
        chartDataModel.setChartName(packetModel.ChartName);
        chartDataModel.setSlotInterval(packetModel.SlotInterval);
        chartDataModel.setLocationId(locationId);
        chartDataModel.setServerChartId(packetModel.ChartID);


        List<ChartConfigTaskModel> sortChartTaskList = new ArrayList<ChartConfigTaskModel>();

        for (PacketTaskModel model:packetModel.Tasks )
        {
            ChartConfigTaskModel chartTaskModel=new ChartConfigTaskModel();
            chartTaskModel.setTaskId(model.TaskID);
            chartTaskModel.setTaskName(model.TaskName);
            chartTaskModel.setTaskOrder(model.TaskOrder);
            sortChartTaskList.add(chartTaskModel);
        }

        Collections.sort(sortChartTaskList);
        Collections.reverse(sortChartTaskList);//Reversing array as it is getting add in reverse order in next step

        for (ChartConfigTaskModel model:sortChartTaskList) {
            chartDataModel.getTasks().put(model.getTaskId(),model);
        }

        Calendar.getInstance().getTime();
        Date dateStart = new Date();
        Calendar now = Calendar.getInstance();

        Calendar calChartStart = Calendar.getInstance();
        calChartStart.set(Calendar.HOUR_OF_DAY, 0);
        calChartStart.set(Calendar.MINUTE, packetModel.StartTime);
        calChartStart.set(Calendar.SECOND, 0);
        calChartStart.set(Calendar.MILLISECOND, 0);
        Date chartStartTime =  calChartStart.getTime();

        int chartEndHour = 0;

        if(packetModel.EndTime <= packetModel.StartTime )
        {
            chartEndHour= 24;
        }

        Calendar calChartEnd = Calendar.getInstance();
        calChartEnd.set(Calendar.HOUR_OF_DAY, chartEndHour);
        calChartEnd.set(Calendar.MINUTE, packetModel.EndTime);
        calChartEnd.set(Calendar.SECOND, 0);
        calChartEnd.set(Calendar.MILLISECOND, 0);
        Date chartEndTime =  calChartEnd.getTime();

        int slotIndex=0;

        while(chartStartTime.getTime() < chartEndTime.getTime())
        {
            ChartConfigSlotModel chartSlotModel = new ChartConfigSlotModel();
            chartSlotModel.setIndex(slotIndex);
            chartSlotModel.setStartTime(new Date(chartStartTime.getTime()));

            Calendar cal = Calendar.getInstance();
            cal.setTime(chartSlotModel.getStartTime());
            cal.add(Calendar.MINUTE,chartDataModel.getSlotInterval());

            chartSlotModel.setEndTime(cal.getTime());

            chartDataModel.getSlots().put(slotIndex,chartSlotModel);

            calChartStart.add(Calendar.MINUTE,packetModel.SlotInterval);
            chartStartTime=calChartStart.getTime();
            slotIndex++;
        }

        return  chartDataModel;
    }

}
