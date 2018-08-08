package com.opensoach.hpft.PacketGenerator;

import com.google.gson.Gson;
import com.opensoach.hpft.Constants.CommandConstants;
import com.opensoach.hpft.Helper.PacketHelper;
import com.opensoach.hpft.Manager.RequestManager;
import com.opensoach.hpft.Model.Communication.CommandRequest;
import com.opensoach.hpft.Model.Communication.PacketModel;
import com.opensoach.hpft.Model.Communication.PacketServiceInstanceTxnModel;
import com.opensoach.hpft.Model.Communication.PacketTaskCompletedDataModel;
import com.opensoach.hpft.Model.DB.DBServiceTaskDataTableRowModel;
import com.opensoach.hpft.Model.View.TaskItemDataModel;
import com.opensoach.hpft.Processor.AckChartDataProcessor;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.List;
import java.util.TimeZone;

/**
 * Created by Mandar on 07-08-2018.
 */

public class TaskDataPacketGenerator  implements IPacketGenerator<ArrayList<TaskItemDataModel>>{

    @Override
    public CommandRequest GenerateRequest(int locationID, ArrayList<TaskItemDataModel> data) {
        return null;
    }

    @Override
    public CommandRequest GenerateUnsyncRequest(int locationID) {

//        List<DBChartDataTableRowModel> unSyncChartData = DatabaseManager.SelectByFilter(new DBChartDataTableQueryModel(), new DBChartDataTableRowModel(), DBChartDataTableQueryModel.FILTER_BY_UNSYNC_DATA);
//
//        if (unSyncChartData.size() == 0) {
//            return null;
//        }
//
//        return GetChartDataPacket(locationID, unSyncChartData);
//
        return null;
    }

    public CommandRequest GetTaskDataPacket(int locationID, List<DBServiceTaskDataTableRowModel> chartRecords) {

        SimpleDateFormat UTCEntryTimeFormat = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.SSS'Z'");
        UTCEntryTimeFormat.setTimeZone(TimeZone.getTimeZone("GMT"));

        ArrayList<PacketServiceInstanceTxnModel> txns = new ArrayList<>();

        for (DBServiceTaskDataTableRowModel model : chartRecords) {
            PacketServiceInstanceTxnModel txnModel = new PacketServiceInstanceTxnModel();
            PacketTaskCompletedDataModel packetTaskCompletedDataModel = new PacketTaskCompletedDataModel();
//            txnModel.servinid = model.getChartId();
//            txnModel.txndate = UTCEntryTimeFormat.format(model.getEntryTime());
//            txnModel.fopcode = model.getAuthCode();
//            if (model.getEntryTime().getTime() < model.getSlotEndTime().getTime()) {
//                txnModel.status = 1;//On time
//            } else {
//                txnModel.status = 2; // Delay
//            }
//
//            packetTaskCompletedDataModel.taskName = model.getTaskName();
//
//            Calendar calChartStart = Calendar.getInstance();
//            calChartStart.setTime(model.getSlotStartTime());
//            packetTaskCompletedDataModel.slotStartTime = calChartStart.get(Calendar.HOUR_OF_DAY) * 60 + calChartStart.get(Calendar.MINUTE);
//
//            Calendar calChartEnd = Calendar.getInstance();
//            calChartEnd.setTime(model.getSlotEndTime());
//            packetTaskCompletedDataModel.slotEndTime = calChartEnd.get(Calendar.HOUR_OF_DAY) * 60 + calChartEnd.get(Calendar.MINUTE);

            txnModel.txndata = new Gson().toJson(packetTaskCompletedDataModel);

            txns.add(txnModel);
        }

        PacketModel<ArrayList<PacketServiceInstanceTxnModel>> packetModel = new PacketModel<>();
        int seqid = RequestManager.Instance().GenerateRequestID();
        packetModel.Header = PacketHelper.CreatePacketHeader(CommandConstants.CMD_CAT_DATA,
                CommandConstants.CMD_DATA_CHART_DATA, seqid, locationID);

        packetModel.Payload = txns;

        CommandRequest<ArrayList<PacketServiceInstanceTxnModel>> commandRequest = new CommandRequest<>();
        commandRequest.Packet = packetModel;
        commandRequest.AckProcessor = new AckChartDataProcessor();

        RequestManager.Instance().AddRequest(seqid, commandRequest);
        return commandRequest;
    }
}
