package com.opensoach.hpft.Helper;



import android.support.v7.app.AppCompatActivity;

import com.google.gson.Gson;

import java.util.ArrayList;
import java.util.Calendar;
import java.util.Date;
import java.util.HashMap;
import java.util.List;

import com.opensoach.hpft.Model.Communication.PacketCardListConfigurationModel;
import com.opensoach.hpft.Model.Communication.PacketModel;
import com.opensoach.hpft.Model.Communication.PacketTimeConfigModel;
import com.opensoach.hpft.Model.DB.DBChartTableRowModel;
import com.opensoach.hpft.Model.Server.DailyChartConfModel;
import com.opensoach.hpft.Model.Server.DailyChartTaskModel;
import com.opensoach.hpft.Model.View.ChartConfigModel;
import com.opensoach.hpft.Model.View.ChartConfigSlotModel;
import com.opensoach.hpft.Model.View.ChartConfigTaskModel;
import com.opensoach.hpft.Model.View.TaskTimeItemDataModel;
import com.opensoach.hpft.ViewModels.CardBriefViewModel;
import com.opensoach.hpft.ViewModels.MedicalDetailsViewModel;
import com.opensoach.hpft.ViewModels.PatientDetailsViewModel;
import com.opensoach.hpft.ViewModels.TaskDetailsViewModel;

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

    public static List<CardBriefViewModel> CreateCardListViewModel(AppCompatActivity ctx, ArrayList<PacketCardListConfigurationModel> models){

        CardBriefViewModel cardBriefViewModel =null;
        List<CardBriefViewModel> cardBriefViewModels = new ArrayList<>();

        for (PacketCardListConfigurationModel model : models) {

            cardBriefViewModel = new CardBriefViewModel();
            cardBriefViewModel.ContextActivity = ctx;

            cardBriefViewModel.setSerInID(model.SerInID);
            cardBriefViewModel.setServConfID(model.ServConfID);
            cardBriefViewModel.setLocationID(model.LocationID);


            PatientDetailsViewModel patientDetailsViewModel = new PatientDetailsViewModel(model.PatientDetails);
            MedicalDetailsViewModel medicalDetailsViewModel = new MedicalDetailsViewModel(model.MedicalDetails);


            List<TaskTimeItemDataModel> timeSlots = GenerateTimeSeries(model.ServiceConf.TimeConfig);

            TaskDetailsViewModel taskDetailsViewModel = new TaskDetailsViewModel(model.ServiceConf,timeSlots);


            patientDetailsViewModel.ContextActivity = ctx;
            medicalDetailsViewModel.ContextActivity = ctx;

            //taskDetailsViewModel.setTaskTimeDataViewModel(new TaskTimeDataViewModel());
            taskDetailsViewModel.getTaskTimeDataViewModel().setUp();
            taskDetailsViewModel.setTitle("This is test for databind ele");
            taskDetailsViewModel.ContextActivity = ctx;


            cardBriefViewModel.setPatientDetails(patientDetailsViewModel);
            cardBriefViewModel.setMedicalDetails(medicalDetailsViewModel);
            cardBriefViewModel.setTaskDetails(taskDetailsViewModel);

            // TODO: dummy code to simulate attention timer icon
            if(cardBriefViewModels.size()<3)
            {
                cardBriefViewModel.setNeedsAttention(true);
            }
            else {
                cardBriefViewModel.setNeedsAttention(false);
            }

            cardBriefViewModels.add(cardBriefViewModel);

        }

        return cardBriefViewModels;
    }

    public  static List<TaskTimeItemDataModel> GenerateTimeSeries(PacketTimeConfigModel timeConfModel){

        List<TaskTimeItemDataModel> slots = new ArrayList<>();

        Calendar calChartStart = Calendar.getInstance();
        calChartStart.set(Calendar.HOUR_OF_DAY, 0);
        calChartStart.set(Calendar.MINUTE, timeConfModel.StartTime);
        calChartStart.set(Calendar.SECOND, 0);
        calChartStart.set(Calendar.MILLISECOND, 0);
        Date chartStartTime = calChartStart.getTime();

        int chartEndHour = 0;

        Calendar calChartEnd = Calendar.getInstance();
        calChartEnd.set(Calendar.HOUR_OF_DAY, chartEndHour);
        calChartEnd.set(Calendar.MINUTE, timeConfModel.EndTime);
        calChartEnd.set(Calendar.SECOND, 0);
        calChartEnd.set(Calendar.MILLISECOND, 0);

        if (timeConfModel.StartTime > timeConfModel.EndTime){
            calChartEnd.add(Calendar.HOUR,24);
        }

        Date chartEndTime = calChartEnd.getTime();

        int slotIndex = 0;

        while (chartStartTime.getTime() < chartEndTime.getTime()) {
            TaskTimeItemDataModel taskTimeItemDataModel = new TaskTimeItemDataModel();
            taskTimeItemDataModel.setIndex(slotIndex);
            taskTimeItemDataModel.setStartTime(new Date(chartStartTime.getTime()));

            Calendar cal = Calendar.getInstance();
            cal.setTime(taskTimeItemDataModel.getStartTime());
            cal.add(Calendar.MINUTE, timeConfModel.Interval);

            taskTimeItemDataModel.setEndTime(cal.getTime());

            slots.add(taskTimeItemDataModel);

            calChartStart.add(Calendar.MINUTE, timeConfModel.Interval);
            chartStartTime = calChartStart.getTime();
            slotIndex++;
        }

        return  slots;
    }
}
