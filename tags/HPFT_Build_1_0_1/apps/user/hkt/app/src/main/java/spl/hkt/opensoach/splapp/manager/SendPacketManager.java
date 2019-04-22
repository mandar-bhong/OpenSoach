package spl.hkt.opensoach.splapp.manager;

import android.os.AsyncTask;
import android.util.Log;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;

import spl.hkt.opensoach.splapp.apprepo.AppRepo;
import spl.hkt.opensoach.splapp.communication.CommunicationManager;
import spl.hkt.opensoach.splapp.dal.DatabaseManager;
import spl.hkt.opensoach.splapp.helper.AppAction;
import spl.hkt.opensoach.splapp.helper.AppHelper;
import spl.hkt.opensoach.splapp.helper.ApplicationConstants;
import spl.hkt.opensoach.splapp.helper.CommonHelper;
import spl.hkt.opensoach.splapp.helper.SyncState;
import spl.hkt.opensoach.splapp.model.ChartDataModel;
import spl.hkt.opensoach.splapp.model.communication.CommandRequest;
import spl.hkt.opensoach.splapp.model.communication.DeviceChartDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketBatteryLevelModel;
import spl.hkt.opensoach.splapp.model.communication.PacketFeedbackDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketUserComplaintDataModel;
import spl.hkt.opensoach.splapp.model.db.DBChartDataTableRowModel;
import spl.hkt.opensoach.splapp.packetGenerator.AuthDataPacketGenerator;
import spl.hkt.opensoach.splapp.packetGenerator.BatteryLevelGenerator;
import spl.hkt.opensoach.splapp.packetGenerator.ChartDataPacketGenerator;
import spl.hkt.opensoach.splapp.packetGenerator.ComplaintDataPacketGenerator;
import spl.hkt.opensoach.splapp.packetGenerator.DeviceSyncCompletedDataPacketGenerator;
import spl.hkt.opensoach.splapp.packetGenerator.FeedbackDataPacketGenerator;

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
                    }

                    if (request != null) {
                        String packet = CommonHelper.GetPacketJSON(request.Packet);
                        CommunicationManager.getInstance().SendPacket(packet);
                    }
                } catch (Exception ex) {
                    Log.d("ParsingSendChartData", ex.getMessage());
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
                ChartDataPacketGenerator chartDataPacketGenerator = new ChartDataPacketGenerator();
                request = chartDataPacketGenerator.GenerateUnsyncRequest(AppRepo.getInstance().getCurrentLocationId());
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