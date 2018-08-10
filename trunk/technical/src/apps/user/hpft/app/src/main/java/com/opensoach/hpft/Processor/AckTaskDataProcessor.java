package com.opensoach.hpft.Processor;

import com.opensoach.hpft.DAL.DatabaseManager;
import com.opensoach.hpft.Helper.SyncState;
import com.opensoach.hpft.Manager.RequestManager;
import com.opensoach.hpft.Manager.SendPacketManager;
import com.opensoach.hpft.Model.Communication.CommandRequest;
import com.opensoach.hpft.Model.Communication.PacketServiceInstanceTxnModel;
import com.opensoach.hpft.Model.Communication.PacketSimpleAckModel;
import com.opensoach.hpft.Model.DB.DBChartDataTableQueryModel;
import com.opensoach.hpft.Model.DB.DBChartDataTableRowModel;
import com.opensoach.hpft.Model.DB.DBServiceTaskDataTableRowModel;
import com.opensoach.hpft.Model.PacketDecodeResultModel;
import com.opensoach.hpft.Model.PacketProcessResultModel;

import java.util.ArrayList;

public class AckTaskDataProcessor implements IProcessor {

    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {
        CommandRequest<ArrayList<PacketServiceInstanceTxnModel>> request = (CommandRequest) RequestManager.Instance().GetRequest(packetDecodeResultModel.Packet.Header.SeqID);
        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();
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

            DBServiceTaskDataTableRowModel dbChartDataTableRowModel = new DBServiceTaskDataTableRowModel();
            dbChartDataTableRowModel.setSerInID(model.servinid);
            //dbChartDataTableRowModel.setEntryTime(model.txndate);
            dbChartDataTableRowModel.setSynced(true);

            DatabaseManager.UpdateRow(new DBChartDataTableQueryModel(), dbChartDataTableRowModel, DBChartDataTableQueryModel.UPDATE_SYNC_STATE_WITH_CHART_ID);
            break;
        }

        packetProcessResultModel.IsSuccess = true;
        SendPacketManager.Instance().sendOnStateChange(SyncState.CHART_DATA_SYNC_COMPLETED);
        return packetProcessResultModel;
    }
}
