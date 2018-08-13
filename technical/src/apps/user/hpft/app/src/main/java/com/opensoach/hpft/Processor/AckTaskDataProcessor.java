package com.opensoach.hpft.Processor;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;
import com.opensoach.hpft.Constants.ApplicationConstants;
import com.opensoach.hpft.DAL.DatabaseManager;
import com.opensoach.hpft.Helper.SyncState;
import com.opensoach.hpft.Manager.RequestManager;
import com.opensoach.hpft.Manager.SendPacketManager;
import com.opensoach.hpft.Model.Communication.CommandRequest;
import com.opensoach.hpft.Model.Communication.PacketLocationDataModel;
import com.opensoach.hpft.Model.Communication.PacketModel;
import com.opensoach.hpft.Model.Communication.PacketServiceInstanceTxnModel;
import com.opensoach.hpft.Model.Communication.PacketSimpleAckModel;
import com.opensoach.hpft.Model.Communication.PacketTaskCompletedDataModel;
import com.opensoach.hpft.Model.DB.DBChartDataTableQueryModel;
import com.opensoach.hpft.Model.DB.DBChartDataTableRowModel;
import com.opensoach.hpft.Model.DB.DBServiceTaskDataTableRowModel;
import com.opensoach.hpft.Model.DB.DBTaskDataTableQueryModel;
import com.opensoach.hpft.Model.DB.DBTaskDataTableRowModel;
import com.opensoach.hpft.Model.PacketDecodeResultModel;
import com.opensoach.hpft.Model.PacketProcessResultModel;
import com.opensoach.hpft.Utility.AppLogger;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Calendar;
import java.util.Date;
import java.util.TimeZone;

import static com.opensoach.hpft.Constants.ApplicationConstants.PACKET_DATE_FORMAT;

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
                cal.set(Calendar.HOUR,0);

                cal.set(Calendar.HOUR, packetTaskCompletedDataModel.slotStartTime/60);
                cal.set(Calendar.MINUTE, packetTaskCompletedDataModel.slotStartTime%60);

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
