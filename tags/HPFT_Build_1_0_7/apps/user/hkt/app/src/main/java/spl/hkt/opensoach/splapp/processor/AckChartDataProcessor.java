package spl.hkt.opensoach.splapp.processor;

import java.util.ArrayList;

import spl.hkt.opensoach.splapp.dal.DatabaseManager;
import spl.hkt.opensoach.splapp.helper.CommandConstants;
import spl.hkt.opensoach.splapp.helper.SyncState;
import spl.hkt.opensoach.splapp.manager.RequestManager;
import spl.hkt.opensoach.splapp.manager.SendPacketManager;
import spl.hkt.opensoach.splapp.model.ChartDataModel;
import spl.hkt.opensoach.splapp.model.PacketDecodeResultModel;
import spl.hkt.opensoach.splapp.model.PacketProcessResultModel;
import spl.hkt.opensoach.splapp.model.communication.CommandRequest;
import spl.hkt.opensoach.splapp.model.communication.DeviceChartDataModel;
import spl.hkt.opensoach.splapp.model.communication.DeviceDataBaseModel;
import spl.hkt.opensoach.splapp.model.communication.PacketModel;
import spl.hkt.opensoach.splapp.model.communication.PacketServiceInstanceTxnModel;
import spl.hkt.opensoach.splapp.model.communication.PacketSimpleAckModel;
import spl.hkt.opensoach.splapp.model.db.DBChartDataTableQueryModel;
import spl.hkt.opensoach.splapp.model.db.DBChartDataTableRowModel;

public class AckChartDataProcessor implements IProcessor {
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

            DBChartDataTableRowModel dbChartDataTableRowModel = new DBChartDataTableRowModel();
            dbChartDataTableRowModel.setChartId(model.servinid);
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