package com.opensoach.hospital.Manager;

import android.os.Bundle;
import android.os.Handler;
import android.os.Looper;
import android.os.Message;
import android.os.Parcelable;

import com.opensoach.hospital.Communication.CommunicationManager;
import com.opensoach.hospital.Helper.CommandConstants;
import com.opensoach.hospital.Helper.CommonHelper;
import com.opensoach.hospital.Helper.PacketHelper;
import com.opensoach.hospital.Model.Communication.Command.DeviceDataAbortJobModel;
import com.opensoach.hospital.Model.Communication.Command.DeviceDataDropJobModel;
import com.opensoach.hospital.Model.Communication.Command.DeviceDataJobQuantityUpdateModel;
import com.opensoach.hospital.Model.Communication.Command.DeviceDataStartJobModel;
import com.opensoach.hospital.Model.Communication.Command.DeviceDataStopJobModel;
import com.opensoach.hospital.Model.Communication.DeviceDataBaseModel;
import com.opensoach.hospital.Model.Communication.PacketModel;
import com.opensoach.hospital.Model.Communication.PacketPayloadModel;
import com.opensoach.hospital.Utility.AppLogger;

/**
 * Created by Mandar on 8/26/2017.
 */

public class SendPacketManager extends Thread{

    private static SendPacketManager singleton;

    private Handler sendPacketHandler;


    /* A private Constructor prevents any other
     * class from instantiating.
     */
    private SendPacketManager() {

    }

    public boolean Init() {
        start();//TODO: Handle error
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

    @Override
    public void run() {
        Thread.currentThread().setName("SendPacketManager");

        Looper.prepare();
        sendPacketHandler = new SendPacketHandler();
        Looper.loop();
    }

    public void send(DeviceDataBaseModel deviceDataBaseModel) {
        Message message = new Message();
        Bundle b = new Bundle();
        b.putParcelable("DEVICE_DATA", (Parcelable) deviceDataBaseModel);
        message.setData(b);

        sendPacketHandler.sendMessage(message);
    }
}


class SendPacketHandler extends Handler {
    @Override
    public void handleMessage(Message msg) {
        super.handleMessage(msg);

        Bundle b = msg.getData();
        DeviceDataBaseModel deviceDataBaseModel = (DeviceDataBaseModel) b.get("DEVICE_DATA");

        try {

            switch (deviceDataBaseModel.getCommandType()) {
                case CommandConstants.DEVICE_DATA_COMMAND_UPDATE_JOB_UNIT:
                    ProcessJobQuantityUpdateCommand(deviceDataBaseModel);
                break;
                case CommandConstants.DEVICE_DATA_COMMAND_START_JOB:
                    ProcessJobStartCommand(deviceDataBaseModel);
                break;
                case CommandConstants.DEVICE_DATA_COMMAND_STOP_JOB:
                    ProcessJobStopCommand(deviceDataBaseModel);
                case CommandConstants.DEVICE_DATA_COMMAND_DROPED_JOB:
                    ProcessJobDropCommand(deviceDataBaseModel);
                case CommandConstants.DEVICE_DATA_COMMAND_ABORTED_JOB:
                    ProcessJobAbortCommand(deviceDataBaseModel);
                break;
            }
        }catch (Exception ex){
            AppLogger.getInstance().Log(ex);
        }
    }


    private void ProcessJobStartCommand(DeviceDataBaseModel deviceDataBaseModel){
        DeviceDataStartJobModel deviceDataStartJobModel = (DeviceDataStartJobModel) deviceDataBaseModel;
        PacketModel<PacketPayloadModel> packetModel = PacketHelper.GetJobStartPacket(
                deviceDataStartJobModel.LocationID(),
                deviceDataStartJobModel.JobID(),
                deviceDataStartJobModel.OperatorCode(),
                deviceDataStartJobModel.StartTime());

        String packetJSON = CommonHelper.GetPacketJSON(packetModel);

        RequestManager.Instance().AddRequest(packetModel.Header.SeqID,packetModel.Payload);

        CommunicationManager.getInstance().SendPacket(packetJSON);
    }

    private void ProcessJobQuantityUpdateCommand(DeviceDataBaseModel deviceDataBaseModel){
        DeviceDataJobQuantityUpdateModel deviceJobQtyModel = (DeviceDataJobQuantityUpdateModel) deviceDataBaseModel;
        PacketModel<PacketPayloadModel> packetModel = PacketHelper.GetJobQuantityUpdatePacket(
                deviceJobQtyModel.LocationID(),
                deviceJobQtyModel.JobID(),
                deviceJobQtyModel.OperatorCode(),
                deviceJobQtyModel.FinishedPartCount(),
                deviceJobQtyModel.CompletionTime(),
                deviceJobQtyModel.Comment());

        String packetJSON = CommonHelper.GetPacketJSON(packetModel);

        RequestManager.Instance().AddRequest(packetModel.Header.SeqID,packetModel.Payload);

        CommunicationManager.getInstance().SendPacket(packetJSON);
    }

    private void ProcessJobStopCommand(DeviceDataBaseModel deviceDataBaseModel) {
        DeviceDataStopJobModel deviceDataStopJobModel = (DeviceDataStopJobModel) deviceDataBaseModel;
        PacketModel<PacketPayloadModel> packetModel = PacketHelper.GetJobStopPacket(
                deviceDataStopJobModel.LocationID(),
                deviceDataStopJobModel.JobID(),
                deviceDataStopJobModel.OperatorCode(),
                deviceDataStopJobModel.EndTime());

        String packetJSON = CommonHelper.GetPacketJSON(packetModel);

        RequestManager.Instance().AddRequest(packetModel.Header.SeqID,packetModel.Payload);

        CommunicationManager.getInstance().SendPacket(packetJSON);
    }

    private void ProcessJobDropCommand(DeviceDataBaseModel deviceDataBaseModel) {
        DeviceDataDropJobModel deviceDataDropJobModel = (DeviceDataDropJobModel) deviceDataBaseModel;
        PacketModel<PacketPayloadModel> packetModel = PacketHelper.GetJobDropPacket(
                deviceDataDropJobModel.LocationID(),
                deviceDataDropJobModel.JobID(),
                deviceDataDropJobModel.OperatorCode(),
                deviceDataDropJobModel.EndTime());

        String packetJSON = CommonHelper.GetPacketJSON(packetModel);

        RequestManager.Instance().AddRequest(packetModel.Header.SeqID,packetModel.Payload);

        CommunicationManager.getInstance().SendPacket(packetJSON);
    }

    private void ProcessJobAbortCommand(DeviceDataBaseModel deviceDataBaseModel) {
        DeviceDataAbortJobModel deviceDataAbortJobModel = (DeviceDataAbortJobModel) deviceDataBaseModel;
        PacketModel<PacketPayloadModel> packetModel = PacketHelper.GetJobAbortPacket(
                deviceDataAbortJobModel.LocationID(),
                deviceDataAbortJobModel.JobID(),
                deviceDataAbortJobModel.OperatorCode(),
                deviceDataAbortJobModel.EndTime());

        String packetJSON = CommonHelper.GetPacketJSON(packetModel);

        RequestManager.Instance().AddRequest(packetModel.Header.SeqID,packetModel.Payload);

        CommunicationManager.getInstance().SendPacket(packetJSON);
    }

//
//    private void ProcessSendMessage(DeviceChartDataModel devideChartDataModel) {
//        int requestId = RequestManager.Instance().GenerateRequestID();
//
//
//        ArrayList<ChartDataModel> chartDataList = devideChartDataModel.getChartDataModels();
//
//        //TODO: Save packet to database
//        try {
//
//            List<DBChartDataTableRowModel> dbChartDataItems = updateTableChartData(devideChartDataModel);
//
//            AppHelper.NotifyChartDataStatusUpdate(dbChartDataItems);
//
//        } catch (ParseException e) {
//            e.printStackTrace();
//        }
//
//
//        //List<DBChartDataTableRowModel> db= DatabaseManager.SelectAll(new DBChartDataTableQueryModel(),new DBChartDataTableRowModel());
//
//        if (AppRepo.getInstance().IsServerConnected()) {
//
//            SimpleDateFormat UTCEntryTimeFormat = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.SSS'Z'");
//            UTCEntryTimeFormat.setTimeZone(TimeZone.getTimeZone("GMT"));
//
//            PacketChartDataModel packetChartDataModel = new PacketChartDataModel();
//            for (ChartDataModel model : devideChartDataModel.getChartDataModels()) {
//                packetChartDataModel.chartId = model.getChartId();
//
//                PacketChartTaskDataModel packetChartTaskDataModel = new PacketChartTaskDataModel();
//                packetChartTaskDataModel.taskId = model.getTaskId();
//                packetChartTaskDataModel.slot = model.getSlotId();
//
//                if(model.getEntryDate().getTime() < model.getSlotEndTime().getTime()){
//                    packetChartTaskDataModel.state = 1;//On time
//                }else{
//                    packetChartTaskDataModel.state = 2; // Delay
//                }
//
//                packetChartTaskDataModel.startSlotTime = UTCEntryTimeFormat.format(model.getSlotStartTime());
//                packetChartTaskDataModel.endSlotTimeObject = UTCEntryTimeFormat.format(model.getSlotEndTime());
//                packetChartTaskDataModel.entryTime = UTCEntryTimeFormat.format(model.getEntryDate());
//
//
//                Calendar dayCal = Calendar.getInstance();
//                dayCal.setTime(model.getEntryDate());
//                dayCal.set(Calendar.HOUR,0);
//                dayCal.set(Calendar.MINUTE,0);
//                dayCal.set(Calendar.SECOND,0);
//                dayCal.set(Calendar.MILLISECOND,0);
//
//                packetChartTaskDataModel.day = UTCEntryTimeFormat.format(dayCal.getTime());
//
//                packetChartDataModel.packetChartTaskDataModels.add(packetChartTaskDataModel);
//            }
//
//            RequestManager.Instance().AddRequest(requestId, devideChartDataModel);
//            String packet = PacketHelper.GetChartDataPacket(requestId, packetChartDataModel);
//
//            CommunicationManager.getInstance().SendPacket(packet);
//        }
//    }
//
//    private List<DBChartDataTableRowModel> updateTableChartData(DeviceChartDataModel devideChartDataModel) throws ParseException {
//
//        List<DBChartDataTableRowModel> dbChartDataItems = new ArrayList<>();
//
//        for (ChartDataModel model : devideChartDataModel.getChartDataModels()) {
//            DBChartDataTableRowModel dbChartDataTableRowModel = new DBChartDataTableRowModel();
//
//            dbChartDataTableRowModel.setChartId(model.getChartId());
//            dbChartDataTableRowModel.setTaskId(model.getTaskId());
//            dbChartDataTableRowModel.setSlotId(model.getSlotId());
//            dbChartDataTableRowModel.setEntryTime(model.getEntryDate());
//            dbChartDataTableRowModel.setSlotStartTime(model.getSlotStartTime());
//            dbChartDataTableRowModel.setSlotEndTime(model.getSlotEndTime());
//            dbChartDataTableRowModel.setSynced(false);
//            dbChartDataTableRowModel.setAuthCode(model.getAuthCode());
//
//            SimpleDateFormat chartStartDate = new SimpleDateFormat("yyyy-MM-dd");
//            Date dateWithoutTime = chartStartDate.parse(chartStartDate.format(model.getEntryDate()));
//
//            dbChartDataTableRowModel.setChartDay(dateWithoutTime);
//
//            if (model.getSlotEndTime().getTime() < model.getEntryDate().getTime()) {
//                dbChartDataTableRowModel.setCellState(ApplicationConstants.DB_CHART_STATE_DELAYED);
//            } else {
//                dbChartDataTableRowModel.setCellState(ApplicationConstants.DB_CHART_STATE_ON_TIME);
//            }
//
//            DatabaseManager.InsertRow(dbChartDataTableRowModel);
//
//            dbChartDataItems.add(dbChartDataTableRowModel);
//        }
//
//        return dbChartDataItems;
//    }

}
