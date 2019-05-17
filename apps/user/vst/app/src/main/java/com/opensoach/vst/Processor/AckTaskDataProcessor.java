package com.opensoach.vst.Processor;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;
import com.opensoach.vst.DAL.DatabaseManager;
import com.opensoach.vst.Helper.SyncState;
import com.opensoach.vst.Manager.RequestManager;
import com.opensoach.vst.Manager.SendPacketManager;
import com.opensoach.vst.Model.Communication.CommandRequest;
import com.opensoach.vst.Model.Communication.PacketServiceInstanceTxnModel;
import com.opensoach.vst.Model.Communication.PacketSimpleAckModel;
import com.opensoach.vst.Model.Communication.PacketTaskCompletedDataModel;
import com.opensoach.vst.Model.DB.DBTaskDataTableQueryModel;
import com.opensoach.vst.Model.DB.DBTaskDataTableRowModel;
import com.opensoach.vst.Model.PacketDecodeResultModel;
import com.opensoach.vst.Model.PacketProcessResultModel;
import com.opensoach.vst.Utility.AppLogger;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Calendar;

import static com.opensoach.vst.Constants.ApplicationConstants.PACKET_DATE_FORMAT;

public class AckTaskDataProcessor implements IProcessor {

    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {
        CommandRequest<ArrayList<PacketServiceInstanceTxnModel>> request = (CommandRequest) RequestManager.Instance().GetRequest(packetDecodeResultModel.Packet.Header.SeqID);

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();
        try {

            if (request == null) {
                packetProcessResultModel.IsSuccess = true;
                SendPacketManager.Instance().sendOnStateChange(SyncState.CHART_DATA_SYNC_COMPLETED);
                return packetProcessResultModel;
            }

            RequestManager.Instance().CompleteRequest(packetDecodeResultModel.Packet.Header.SeqID);

            PacketSimpleAckModel ack = (PacketSimpleAckModel) packetDecodeResultModel.Packet.Payload;

            if (!ack.Ack) {
                packetProcessResultModel.IsSuccess = true;
                SendPacketManager.Instance().sendOnStateChange(SyncState.CHART_DATA_SYNC_COMPLETED);
                return packetProcessResultModel;
            }

            for (PacketServiceInstanceTxnModel model : request.Packet.Payload) {

                TypeToken<PacketTaskCompletedDataModel> typeToken = new TypeToken<PacketTaskCompletedDataModel>() {
                };

                PacketTaskCompletedDataModel packetTaskCompletedDataModel = new Gson().fromJson(model.txndata, typeToken.getType());

                Calendar cal = Calendar.getInstance();
                SimpleDateFormat txnDate = new SimpleDateFormat(PACKET_DATE_FORMAT);
                //txnDate.setTimeZone(TimeZone.getTimeZone("GMT"));
                cal.setTime(txnDate.parse(model.txndate));

                cal.set(Calendar.MILLISECOND,0);
                cal.set(Calendar.SECOND,0);
                cal.set(Calendar.MINUTE,0);
                cal.set(Calendar.HOUR_OF_DAY,0);

                cal.set(Calendar.MINUTE, packetTaskCompletedDataModel.slotStartTime);

                DBTaskDataTableRowModel dbTaskDataTableRowModel = new DBTaskDataTableRowModel();
                dbTaskDataTableRowModel.setLocationId(request.Packet.Header.LocationID);

                dbTaskDataTableRowModel.setSerInID(model.servinid);
                dbTaskDataTableRowModel.setTaskSlotStartTime(cal.getTime());
                dbTaskDataTableRowModel.setSynced(true);

                DatabaseManager.UpdateRow(new DBTaskDataTableQueryModel(), dbTaskDataTableRowModel, DBTaskDataTableQueryModel.UPDATE_SYNC_BY_ID_FILTER);
            }

            packetProcessResultModel.IsSuccess = true;
            SendPacketManager.Instance().sendOnStateChange(SyncState.CHART_DATA_SYNC_COMPLETED);

        } catch (Exception ex) {
            AppLogger.getInstance().Log(ex);
        }
        return packetProcessResultModel;
    }
}
