package com.opensoach.hpft.PacketGenerator;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;
import com.opensoach.hpft.Constants.ApplicationConstants;
import com.opensoach.hpft.Constants.CommandConstants;
import com.opensoach.hpft.DAL.DatabaseManager;
import com.opensoach.hpft.Helper.PacketHelper;
import com.opensoach.hpft.Manager.RequestManager;
import com.opensoach.hpft.Model.Communication.CommandRequest;
import com.opensoach.hpft.Model.Communication.PacketModel;
import com.opensoach.hpft.Model.Communication.PacketServiceInstanceTxnModel;
import com.opensoach.hpft.Model.Communication.PacketTaskCompletedDataModel;
import com.opensoach.hpft.Model.DB.DBServiceTaskDataTableQueryModel;
import com.opensoach.hpft.Model.DB.DBServiceTaskDataTableRowModel;
import com.opensoach.hpft.Model.DB.DBTaskDataTableQueryModel;
import com.opensoach.hpft.Model.DB.DBTaskDataTableRowModel;
import com.opensoach.hpft.Model.View.TaskItemDataModel;
import com.opensoach.hpft.Processor.AckChartDataProcessor;
import com.opensoach.hpft.Processor.AckTaskDataProcessor;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Calendar;
import java.util.List;
import java.util.TimeZone;

/**
 * Created by Mandar on 07-08-2018.
 */

public class TaskDataPacketGenerator  implements IPacketGenerator<ArrayList<DBTaskDataTableRowModel>>{

    @Override
    public CommandRequest GenerateRequest(int locationID, ArrayList<DBTaskDataTableRowModel> data) {

        return  GetTaskDataPacket(locationID,data);
    }

    @Override
    public CommandRequest GenerateUnsyncRequest(int locationID) {
        DBTaskDataTableRowModel dbTaskDataTableRowModel = new DBTaskDataTableRowModel();
        dbTaskDataTableRowModel.setLocationId(locationID);
        dbTaskDataTableRowModel.setSynced(false);

        List<DBTaskDataTableRowModel> unSyncChartData = DatabaseManager.SelectByFilter(new DBTaskDataTableQueryModel(), dbTaskDataTableRowModel, DBTaskDataTableQueryModel.FILTER_BY_UNSYNC_DATA);

        if (unSyncChartData.size() == 0) {
            return null;
        }

        return GetTaskDataPacket(locationID, unSyncChartData);
    }

    public CommandRequest GetTaskDataPacket(int locationID, List<DBTaskDataTableRowModel> chartRecords) {

        SimpleDateFormat UTCEntryTimeFormat = new SimpleDateFormat(ApplicationConstants.PACKET_DATE_FORMAT);
        UTCEntryTimeFormat.setTimeZone(TimeZone.getTimeZone("GMT"));

        ArrayList<PacketServiceInstanceTxnModel> txns = new ArrayList<>();

        for (DBTaskDataTableRowModel model : chartRecords) {

            PacketServiceInstanceTxnModel txnModel = new PacketServiceInstanceTxnModel();
            PacketTaskCompletedDataModel packetTaskCompletedDataModel = new PacketTaskCompletedDataModel();
            txnModel.servinid = model.getSerInID();
            txnModel.txndate = UTCEntryTimeFormat.format(model.getTaskTime());
            txnModel.fopcode = model.getAuthCode();

            if (model.getTaskTime().getTime() < model.getTaskSlotEndTime().getTime()) {
                txnModel.status = 1;//On time
            } else {
                txnModel.status = 2; // Delay
            }

            packetTaskCompletedDataModel.taskName = model.getTitle();
            packetTaskCompletedDataModel.comment = model.getComment();
            packetTaskCompletedDataModel.value = model.getValue();

            Calendar calChartStart = Calendar.getInstance();
            calChartStart.setTime(model.getTaskSlotStartTime());
            packetTaskCompletedDataModel.slotStartTime = calChartStart.get(Calendar.HOUR_OF_DAY) * 60 + calChartStart.get(Calendar.MINUTE);

            Calendar calChartEnd = Calendar.getInstance();
            calChartEnd.setTime(model.getTaskSlotEndTime());
            packetTaskCompletedDataModel.slotEndTime = calChartEnd.get(Calendar.HOUR_OF_DAY) * 60 + calChartEnd.get(Calendar.MINUTE);

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
        commandRequest.AckProcessor = new AckTaskDataProcessor();

        RequestManager.Instance().AddRequest(seqid, commandRequest);
        return commandRequest;
    }
}
