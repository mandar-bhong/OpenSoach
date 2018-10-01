package com.opensoach.vst.Manager;

import android.os.AsyncTask;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Communication.CommunicationManager;
import com.opensoach.vst.DAL.DatabaseManager;
import com.opensoach.vst.Helper.AppAction;
import com.opensoach.vst.Helper.AppHelper;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.Helper.CommonHelper;
import com.opensoach.vst.Helper.SyncState;
import com.opensoach.vst.Model.ChartDataModel;
import com.opensoach.vst.Model.Communication.CommandRequest;
import com.opensoach.vst.Model.Communication.DeviceChartDataModel;
import com.opensoach.vst.Model.Communication.PacketBatteryLevelModel;
import com.opensoach.vst.Model.Communication.PacketFeedbackDataModel;
import com.opensoach.vst.Model.Communication.PacketUserComplaintDataModel;
import com.opensoach.vst.Model.DB.DBChartDataTableRowModel;
import com.opensoach.vst.Model.DB.DBTaskDataTableRowModel;
import com.opensoach.vst.PacketGenerator.AuthDataPacketGenerator;
import com.opensoach.vst.PacketGenerator.BatteryLevelGenerator;
import com.opensoach.vst.PacketGenerator.ChartDataPacketGenerator;
import com.opensoach.vst.PacketGenerator.ComplaintDataPacketGenerator;
import com.opensoach.vst.PacketGenerator.DeviceSyncCompletedDataPacketGenerator;
import com.opensoach.vst.PacketGenerator.FeedbackDataPacketGenerator;
import com.opensoach.vst.PacketGenerator.JobCreatedDataPacketGenerator;
import com.opensoach.vst.PacketGenerator.JobServiceExeDetailsPacketGenerator;
import com.opensoach.vst.PacketGenerator.JobServiceOwnerVehicleDetailsPacketGenerator;
import com.opensoach.vst.PacketGenerator.JobServiceTaskCompletedPacketGenerator;
import com.opensoach.vst.PacketGenerator.TaskDataPacketGenerator;
import com.opensoach.vst.PacketGenerator.TokenClaimPacketGenerator;
import com.opensoach.vst.PacketGenerator.TokenGenerateDataPacketGenerator;
import com.opensoach.vst.PacketGenerator.VehicleDetailsPacketGenerator;
import com.opensoach.vst.Utility.AppLogger;
import com.opensoach.vst.ViewModels.CreateTokenViewModel;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.ViewModels.JobServiceViewModel;
import com.opensoach.vst.ViewModels.TokenItemViewModel;

/**
 * Created by samir.s.bukkawar on 3/25/2017.
 */

public class SendPacketManager {

    private static SendPacketManager singleton;

    /* A private Constructor prevents any other
     * class from instantiating.
     */
    private SendPacketManager() {

    }

    public boolean Init() {
        return true;
    }

    public void DeInit() {
        //stop();//TODO: Deint this class
    }

    /* Static 'instance' method */
    public static SendPacketManager Instance() {
        if (singleton == null)
            singleton = new SendPacketManager();
        return singleton;
    }

    public void send(final AppAction action, final Object data) {
        final int locationID = AppRepo.getInstance().getCurrentLocationId();
        Runnable sendTask = new Runnable() {
            @Override
            public void run() {
                CommandRequest request = null;
                try {

                    switch (action) {
                        case CHART_DATA:
                            DeviceChartDataModel deviceChartDataModel = (DeviceChartDataModel) data;
                            List<DBChartDataTableRowModel> dbChartDataItems = updateTableChartData(deviceChartDataModel);
                            AppHelper.NotifyChartDataStatusUpdate(dbChartDataItems);
                            ChartDataPacketGenerator chartDataPacketGenerator = new ChartDataPacketGenerator();
                            request = chartDataPacketGenerator.GenerateRequest(locationID, dbChartDataItems);
                            break;
                        case TASK_DATA:
                            ArrayList<DBTaskDataTableRowModel> serviceDataItems = (ArrayList<DBTaskDataTableRowModel>) data;
                            TaskDataPacketGenerator taskDataPacketGenerator = new TaskDataPacketGenerator();
                            request = taskDataPacketGenerator.GenerateRequest(locationID, serviceDataItems);
                            break;
                        case COMPLAINT_DATA:
                            ComplaintDataPacketGenerator complaintDataPacketGenerator = new ComplaintDataPacketGenerator();
                            request = complaintDataPacketGenerator.GenerateRequest(locationID, (ArrayList<PacketUserComplaintDataModel>) data);
                            break;
                        case ON_CONNECTION:
                            AuthDataPacketGenerator authDataPacketGenerator = new AuthDataPacketGenerator();
                            request = authDataPacketGenerator.GenerateRequest(locationID, (String) data);
                            break;
                        case FEEDBACK_DATA:
                            FeedbackDataPacketGenerator feedbackDataPacketGenerator = new FeedbackDataPacketGenerator();
                            request = feedbackDataPacketGenerator.GenerateRequest(locationID, (ArrayList<PacketFeedbackDataModel>) data);
                            break;
                        case BATTERY_LEVEL:
                            BatteryLevelGenerator batteryLevelGenerator = new BatteryLevelGenerator();
                            request = batteryLevelGenerator.GenerateRequest(0, (PacketBatteryLevelModel) data);
                            break;
                        case CREATE_TOKEN:
                            TokenGenerateDataPacketGenerator tokenGenerator = new TokenGenerateDataPacketGenerator();
                            request = tokenGenerator.GenerateRequest(locationID, (CreateTokenViewModel) data);
                            break;
                        case CLAIM_TOKEN:
                            TokenClaimPacketGenerator tokenClaimPacketGenerator = new TokenClaimPacketGenerator();
                            request = tokenClaimPacketGenerator.GenerateRequest(locationID, (TokenItemViewModel) data);
                            break;
                        case CREATE_JOB_COMFIRM:
                            JobCreatedDataPacketGenerator jobCreatedDataPacketGenerator = new JobCreatedDataPacketGenerator();
                            request = jobCreatedDataPacketGenerator.GenerateRequest(locationID, (JobServiceViewModel) data);
                            break;

                        case GET_VEHICLE_DETAILS:
                            VehicleDetailsPacketGenerator vehicleDetailsPacketGenerator = new VehicleDetailsPacketGenerator();
                            request = vehicleDetailsPacketGenerator.GenerateRequest(locationID, (String) data);
                            break;

                        case UPADATE_VEHICLE_OWNER_DETAILS: {
                            JobServiceOwnerVehicleDetailsPacketGenerator jsvowd = new JobServiceOwnerVehicleDetailsPacketGenerator();
                            request = jsvowd.GenerateRequest(locationID, (JobServiceViewModel) data);
                            break;
                        }

                        case GET_JOB_DETAILS_FOR_EXE: {
                            JobServiceExeDetailsPacketGenerator jsvowd = new JobServiceExeDetailsPacketGenerator();
                            request = jsvowd.GenerateRequest(locationID, (Integer) data);
                            break;
                        }

                        case UPDATE_JOB_TASK_COMPLETED:
                            JobServiceItemViewModel jobServiceItemViewModel = (JobServiceItemViewModel)data;
                            JobServiceTaskCompletedPacketGenerator taskCompletedPacketGenerator = new JobServiceTaskCompletedPacketGenerator();
                            request = taskCompletedPacketGenerator.GenerateRequest(locationID,jobServiceItemViewModel);
                            break;
                    }

                    if (request != null) {
                        String packet = CommonHelper.GetPacketJSON(request.Packet);
                        CommunicationManager.getInstance().SendPacket(packet);
                    }
                } catch (Exception ex) {
                    AppLogger.getInstance().Log(ex);
                }
            }
        };

        AsyncTask.execute(sendTask);
    }




    private List<DBChartDataTableRowModel> updateTableChartData(DeviceChartDataModel devideChartDataModel) throws ParseException {

        List<DBChartDataTableRowModel> dbChartDataItems = new ArrayList<>();

        for (ChartDataModel model : devideChartDataModel.getChartDataModels()) {
            DBChartDataTableRowModel dbChartDataTableRowModel = new DBChartDataTableRowModel();

            dbChartDataTableRowModel.setChartId(model.getChartId());
            dbChartDataTableRowModel.setTaskName(model.getTaskName());
            dbChartDataTableRowModel.setSlotId(model.getSlotId());
            dbChartDataTableRowModel.setEntryTime(model.getEntryDate());
            dbChartDataTableRowModel.setSlotStartTime(model.getSlotStartTime());
            dbChartDataTableRowModel.setSlotEndTime(model.getSlotEndTime());
            dbChartDataTableRowModel.setSynced(false);
            dbChartDataTableRowModel.setAuthCode(model.getAuthCode());

            SimpleDateFormat chartStartDate = new SimpleDateFormat("yyyy-MM-dd");
            Date dateWithoutTime = chartStartDate.parse(chartStartDate.format(model.getEntryDate()));

            dbChartDataTableRowModel.setChartDay(dateWithoutTime);

            if (model.getSlotEndTime().getTime() < model.getEntryDate().getTime()) {
                dbChartDataTableRowModel.setCellState(ApplicationConstants.DB_CHART_STATE_DELAYED);
            } else {
                dbChartDataTableRowModel.setCellState(ApplicationConstants.DB_CHART_STATE_ON_TIME);
            }

            DatabaseManager.InsertRow(dbChartDataTableRowModel);

            dbChartDataItems.add(dbChartDataTableRowModel);
        }

        return dbChartDataItems;
    }

    public void sendOnStateChange(SyncState state) {
        if (AppRepo.getInstance().getIsDeviceSyncInProgress() == false) {
            return;
        }
        CommandRequest request = null;
        switch (state) {
            case DEVICE_REGISTRATION_COMPLETED:
//                ChartDataPacketGenerator chartDataPacketGenerator = new ChartDataPacketGenerator();
//                request = chartDataPacketGenerator.GenerateUnsyncRequest(AppRepo.getInstance().getCurrentLocationId());

                TaskDataPacketGenerator taskDataPacketGenerator = new TaskDataPacketGenerator();
                request = taskDataPacketGenerator.GenerateUnsyncRequest(AppRepo.getInstance().getCurrentLocationId());

                if (request == null) {
                    // there is no unsync data, mark chart data sync completed.
                    this.sendOnStateChange(SyncState.CHART_DATA_SYNC_COMPLETED);
                    return;
                }
                break;
            case CHART_DATA_SYNC_COMPLETED:

                ComplaintDataPacketGenerator complaintDataPacketGenerator = new ComplaintDataPacketGenerator();
                request = complaintDataPacketGenerator.GenerateUnsyncRequest(AppRepo.getInstance().getCurrentLocationId());
                if (request == null) {
                    // there is no unsync data, mark complaint data sync completed.
                    this.sendOnStateChange(SyncState.COMPLAINT_DATA_SYNC_COMPLETED);
                    return;
                }
                break;

            case COMPLAINT_DATA_SYNC_COMPLETED:
                AppRepo.getInstance().setIsDeviceSyncInProgress(false);
                DeviceSyncCompletedDataPacketGenerator deviceSyncCompletedDataPacketGenerator = new DeviceSyncCompletedDataPacketGenerator();
                request = deviceSyncCompletedDataPacketGenerator.GenerateRequest(0, null);

                break;
        }

        if (request != null) {
            String packet = CommonHelper.GetPacketJSON(request.Packet);
            CommunicationManager.getInstance().SendPacket(packet);
        }
    }
}