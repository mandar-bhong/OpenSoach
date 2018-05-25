package spl.hkt.opensoach.splapp.processor;

import android.util.Log;

import com.google.gson.Gson;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Calendar;
import java.util.List;
import java.util.TimeZone;

import spl.hkt.opensoach.splapp.dal.DatabaseManager;
import spl.hkt.opensoach.splapp.helper.CommandConstants;
import spl.hkt.opensoach.splapp.helper.PacketHelper;
import spl.hkt.opensoach.splapp.manager.RequestManager;
import spl.hkt.opensoach.splapp.model.PacketDecodeResultModel;
import spl.hkt.opensoach.splapp.model.PacketProcessResultModel;
import spl.hkt.opensoach.splapp.model.communication.CommandRequest;
import spl.hkt.opensoach.splapp.model.communication.DeviceChartDataStartupSyncModel;
import spl.hkt.opensoach.splapp.model.communication.PacketChartDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketServiceInstanceTxnModel;
import spl.hkt.opensoach.splapp.model.db.DBChartDataTableQueryModel;
import spl.hkt.opensoach.splapp.model.db.DBChartDataTableRowModel;

public class AckDeviceRegProcessor implements IProcessor {
    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {
        CommandRequest<ArrayList<PacketServiceInstanceTxnModel>> request = (CommandRequest) RequestManager.Instance().GetRequest(packetDecodeResultModel.Packet.Header.SeqID);
        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();
        if (request == null) {
            packetProcessResultModel.IsSuccess = true;
            return packetProcessResultModel;
        }

        RequestManager.Instance().CompleteRequest(packetDecodeResultModel.Packet.Header.SeqID);

        List<DBChartDataTableRowModel> unSyncChartData = DatabaseManager.SelectByFilter(new DBChartDataTableQueryModel(), new DBChartDataTableRowModel(), DBChartDataTableQueryModel.FILTER_BY_UNSYNC_DATA);

        if (unSyncChartData.size() > 0) {
            packetProcessResultModel = processUnSyncChartData(unSyncChartData);
        } else {
            packetProcessResultModel.CanSendServerCommand = true;
            packetProcessResultModel.ServerCommandPacket = PacketHelper.GetDeviceSynCompletedPacket();
            packetProcessResultModel.IsSuccess = true;
        }

        return packetProcessResultModel;
    }

    private PacketProcessResultModel processUnSyncChartData(List<DBChartDataTableRowModel> unSyncChartData){
        //TODO: For multiple chart multiple packet need to create according to chartID
        PacketProcessResultModel packetProcessResultModel =new PacketProcessResultModel();

        SimpleDateFormat UTCEntryTimeFormat = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.SSS'Z'");
        UTCEntryTimeFormat.setTimeZone(TimeZone.getTimeZone("GMT"));

        ArrayList<PacketServiceInstanceTxnModel> txns = new ArrayList<>();

        for (DBChartDataTableRowModel model : unSyncChartData) {
            PacketServiceInstanceTxnModel txnModel = new PacketServiceInstanceTxnModel();
            PacketChartDataModel packetChartDataModel = new PacketChartDataModel();
            txnModel.servinid = model.getChartId();
            txnModel.txndate = UTCEntryTimeFormat.format(model.getEntryTime());
            txnModel.fopcode = model.getAuthCode();
            if (model.getEntryTime().getTime() < model.getSlotEndTime().getTime()) {
                txnModel.status = 1;//On time
            } else {
                txnModel.status = 2; // Delay
            }

            packetChartDataModel.taskName = model.getTaskName();

            Calendar calChartStart = Calendar.getInstance();
            calChartStart.setTime(model.getSlotStartTime());
            packetChartDataModel.slotStartTime = calChartStart.get(Calendar.HOUR_OF_DAY) * 60 + calChartStart.get(Calendar.MINUTE);

            Calendar calChartEnd = Calendar.getInstance();
            calChartEnd.setTime(model.getSlotEndTime());
            packetChartDataModel.slotEndTime = calChartEnd.get(Calendar.HOUR_OF_DAY) * 60 + calChartEnd.get(Calendar.MINUTE);

            txnModel.txndata = new Gson().toJson(packetChartDataModel);

            txns.add(txnModel);
        }

        packetProcessResultModel.CanSendServerCommand = true;
        packetProcessResultModel.ServerCommandPacket = PacketHelper.GetChartDataPacket(txns);
        Log.d("SyncDataPck",packetProcessResultModel.ServerCommandPacket);
        packetProcessResultModel.IsSuccess = true;
        return  packetProcessResultModel;

    }
}

