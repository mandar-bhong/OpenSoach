package spl.hkt.opensoach.splapp.processor;

import android.util.Log;

import com.google.gson.Gson;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Calendar;
import java.util.List;
import java.util.TimeZone;

import spl.hkt.opensoach.splapp.apprepo.AppRepo;
import spl.hkt.opensoach.splapp.dal.DatabaseManager;
import spl.hkt.opensoach.splapp.helper.CommandConstants;
import spl.hkt.opensoach.splapp.helper.PacketHelper;
import spl.hkt.opensoach.splapp.helper.SyncState;
import spl.hkt.opensoach.splapp.manager.RequestManager;
import spl.hkt.opensoach.splapp.manager.SendPacketManager;
import spl.hkt.opensoach.splapp.model.PacketDecodeResultModel;
import spl.hkt.opensoach.splapp.model.PacketProcessResultModel;
import spl.hkt.opensoach.splapp.model.communication.CommandRequest;
import spl.hkt.opensoach.splapp.model.communication.DeviceChartDataStartupSyncModel;
import spl.hkt.opensoach.splapp.model.communication.PacketChartDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketServiceInstanceTxnModel;
import spl.hkt.opensoach.splapp.model.communication.PacketSimpleAckModel;
import spl.hkt.opensoach.splapp.model.db.DBChartDataTableQueryModel;
import spl.hkt.opensoach.splapp.model.db.DBChartDataTableRowModel;

public class AckDeviceRegProcessor implements IProcessor {
    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {

        AppRepo.getInstance().setIsDeviceSyncInProgress(true);
        CommandRequest<ArrayList<PacketServiceInstanceTxnModel>> request = (CommandRequest) RequestManager.Instance().GetRequest(packetDecodeResultModel.Packet.Header.SeqID);
        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();
        if (request == null) {
            packetProcessResultModel.IsSuccess = true;
            return packetProcessResultModel;
        }

        RequestManager.Instance().CompleteRequest(packetDecodeResultModel.Packet.Header.SeqID);

        PacketSimpleAckModel ack = (PacketSimpleAckModel) packetDecodeResultModel.Packet.Payload;

        if (!ack.Ack) {
            packetProcessResultModel.IsSuccess = false;
            return packetProcessResultModel;
        }

        SendPacketManager.Instance().sendOnStateChange(SyncState.DEVICE_REGISTRATION_COMPLETED);

        packetProcessResultModel.IsSuccess = true;
        return packetProcessResultModel;
    }
}

