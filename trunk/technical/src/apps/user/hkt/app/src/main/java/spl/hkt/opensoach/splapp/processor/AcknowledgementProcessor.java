package spl.hkt.opensoach.splapp.processor;

import android.provider.ContactsContract;
import android.util.Log;

import java.text.SimpleDateFormat;
import java.util.List;
import java.util.TimeZone;

import spl.hkt.opensoach.splapp.dal.DatabaseManager;
import spl.hkt.opensoach.splapp.helper.CommandConstants;
import spl.hkt.opensoach.splapp.helper.PacketHelper;
import spl.hkt.opensoach.splapp.manager.RequestManager;
import spl.hkt.opensoach.splapp.model.ChartDataModel;
import spl.hkt.opensoach.splapp.model.PacketDecodeResultModel;
import spl.hkt.opensoach.splapp.model.PacketProcessResultModel;
import spl.hkt.opensoach.splapp.model.communication.DeviceChartDataModel;
import spl.hkt.opensoach.splapp.model.communication.DeviceChartDataStartupSyncModel;
import spl.hkt.opensoach.splapp.model.communication.DeviceDataBaseModel;
import spl.hkt.opensoach.splapp.model.communication.PacketChartDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketChartTaskDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketModel;
import spl.hkt.opensoach.splapp.model.db.DBChartDataTableQueryModel;
import spl.hkt.opensoach.splapp.model.db.DBChartDataTableRowModel;
import spl.hkt.opensoach.splapp.model.db.DBChartTableQueryModel;

/**
 * Created by Mandar on 2/26/2017.
 */

public class AcknowledgementProcessor implements IProcessor {
    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {

        PacketModel resultModel = packetDecodeResultModel.Packet;

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            switch (resultModel.Header.CommandID) {
                case CommandConstants.CMD_ACK_DEVICE_REG:
                    //TODO: Acquire and send offline data

                    List<DBChartDataTableRowModel> unSyncChartData1 = DatabaseManager.SelectAll(new DBChartDataTableQueryModel(), new DBChartDataTableRowModel());
                    List<DBChartDataTableRowModel> unSyncChartData = DatabaseManager.SelectByFilter(new DBChartDataTableQueryModel(), new DBChartDataTableRowModel(), DBChartDataTableQueryModel.FILTER_BY_UNSYNC_DATA);

                    if (unSyncChartData.size() > 0) {
                        packetProcessResultModel = processUnSyncChartData(unSyncChartData);
                    } else {
                        packetProcessResultModel.CanSendServerCommand = true;
                        packetProcessResultModel.ServerCommandPacket = PacketHelper.GetDeviceSynCompletedPacket();
                        packetProcessResultModel.IsSuccess = true;
                    }
                    break;

                case CommandConstants.CMD_ACK_CHART_DATA: {

                    Object data = RequestManager.Instance().GetRequest(resultModel.Header.SeqID);

                    RequestManager.Instance().CompleteRequest(resultModel.Header.SeqID);

                    if (data != null) {

                        DeviceDataBaseModel baseDataModel = (DeviceDataBaseModel) data;

                        switch (baseDataModel.getCommandType()) {
                            case CommandConstants.DEVICE_DATA_COMMAND_CHART_DATA: {
                                DeviceChartDataModel deviceChartDataModel = (DeviceChartDataModel) data;

                                for (ChartDataModel model : deviceChartDataModel.getChartDataModels()) {

                                    DBChartDataTableRowModel dbChartDataTableRowModel = new DBChartDataTableRowModel();
                                    dbChartDataTableRowModel.setChartId(model.getChartId());
                                    dbChartDataTableRowModel.setEntryTime(model.getEntryDate());
                                    dbChartDataTableRowModel.setSynced(true);

                                    DatabaseManager.UpdateRow(new DBChartDataTableQueryModel(), dbChartDataTableRowModel, DBChartDataTableQueryModel.UPDATE_SYNC_STATE_WITH_CHART_ID_ENTRY_TIME);
                                }
                            }
                            break;
                            case CommandConstants.DEVICE_DATA_COMMAND_CHART_UNSYNC_DATA: {

                                DeviceChartDataStartupSyncModel deviceChartDataStartupSyncModel = (DeviceChartDataStartupSyncModel)data;

                                for (DBChartDataTableRowModel model : deviceChartDataStartupSyncModel.getUnSyncChartData()) {
                                    DBChartDataTableRowModel dbChartDataTableRowModel = new DBChartDataTableRowModel();
                                    dbChartDataTableRowModel.setChartId(model.getChartId());
                                    dbChartDataTableRowModel.setEntryTime(model.getEntryTime());
                                    dbChartDataTableRowModel.setSynced(true);

                                    DatabaseManager.UpdateRow(new DBChartDataTableQueryModel(), dbChartDataTableRowModel, DBChartDataTableQueryModel.UPDATE_SYNC_STATE_WITH_CHART_ID_ENTRY_TIME);
                                }

                                packetProcessResultModel.CanSendServerCommand = true;
                                packetProcessResultModel.ServerCommandPacket = PacketHelper.GetDeviceSynCompletedPacket();
                                packetProcessResultModel.IsSuccess = true;
                            }
                            break;
                        }
                    }
                    break;
                }
            }
        } catch (Exception ex) {
            //TODO: Log exception error
            Log.d("AckProcess",ex.getMessage());
        }

        return packetProcessResultModel;
    }


    private PacketProcessResultModel processUnSyncChartData(List<DBChartDataTableRowModel> unSyncChartData){
    //TODO: For multiple chart multiple packet need to create according to chartID
        PacketProcessResultModel packetProcessResultModel =new PacketProcessResultModel();
        PacketChartDataModel packetChartDataModel = new PacketChartDataModel();

        SimpleDateFormat UTCEntryTimeFormat = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.SSS'Z'");
        UTCEntryTimeFormat.setTimeZone(TimeZone.getTimeZone("GMT"));

        for (DBChartDataTableRowModel dbChartDataModel:unSyncChartData) {

            PacketChartTaskDataModel packetChartTaskDataModel = new PacketChartTaskDataModel();
            packetChartTaskDataModel.day = UTCEntryTimeFormat.format(dbChartDataModel.getChartDay().getTime());
            packetChartTaskDataModel.endSlotTimeObject = UTCEntryTimeFormat.format(dbChartDataModel.getSlotEndTime());
            packetChartTaskDataModel.entryTime = UTCEntryTimeFormat.format(dbChartDataModel.getEntryTime());
            packetChartTaskDataModel.slot = dbChartDataModel.getSlotId();
            packetChartTaskDataModel.taskId = dbChartDataModel.getTaskId();
            packetChartTaskDataModel.state = dbChartDataModel.getCellState();
            packetChartTaskDataModel.startSlotTime =UTCEntryTimeFormat.format(dbChartDataModel.getSlotStartTime());

            packetChartDataModel.packetChartTaskDataModels.add(packetChartTaskDataModel);
            packetChartDataModel.chartId = dbChartDataModel.getChartId();
        }

        int requestId = RequestManager.Instance().GenerateRequestID();

        DeviceChartDataStartupSyncModel deviceChartDataStartupSyncModel = new DeviceChartDataStartupSyncModel();
        deviceChartDataStartupSyncModel.setCommandType(CommandConstants.DEVICE_DATA_COMMAND_CHART_UNSYNC_DATA);
        deviceChartDataStartupSyncModel.setUnSyncChartData(unSyncChartData);
        RequestManager.Instance().AddRequest(requestId, deviceChartDataStartupSyncModel);

        packetProcessResultModel.CanSendServerCommand = true;
        packetProcessResultModel.ServerCommandPacket = PacketHelper.GetChartDataPacket(requestId, packetChartDataModel);
        Log.d("SyncDataPck",packetProcessResultModel.ServerCommandPacket);
        packetProcessResultModel.IsSuccess = true;
        return  packetProcessResultModel;

    }
}
