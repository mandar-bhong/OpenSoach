package com.opensoach.hospital.Processor;

import com.opensoach.hospital.Helper.CommandConstants;
import com.opensoach.hospital.Model.Communication.PacketModel;
import com.opensoach.hospital.Model.PacketDecodeResultModel;
import com.opensoach.hospital.Model.PacketProcessResultModel;
import com.opensoach.hospital.Processor.AckProcessor.AckJobAbortedProcessor;
import com.opensoach.hospital.Processor.AckProcessor.AckJobCompletedProcessor;
import com.opensoach.hospital.Processor.AckProcessor.AckJobDroppedProcessor;
import com.opensoach.hospital.Processor.AckProcessor.AckJobQuantityUpdatedProcessor;
import com.opensoach.hospital.Processor.AckProcessor.AckJobStartedProcess;
import com.opensoach.hospital.Processor.AckProcessor.DeviceRegistrationProcessor;
import com.opensoach.hospital.Utility.AppLogger;

/**
 * Created by Mandar on 8/27/2017.
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
                {
                    return new DeviceRegistrationProcessor().Process(packetDecodeResultModel);
                }

                case CommandConstants.CMD_ACK_JOB_STATUS_START: {
                    return new AckJobStartedProcess().Process(packetDecodeResultModel);
                }

                case CommandConstants.CMD_ACK_JOB_STATUS_ABORT: {
                    return new AckJobAbortedProcessor().Process(packetDecodeResultModel);
                }
                case CommandConstants.CMD_ACK_JOB_STATUS_DROP: {
                    return new AckJobDroppedProcessor().Process(packetDecodeResultModel);
                }
                case CommandConstants.CMD_ACK_JOB_STATUS_STOP:
                {
                    return new AckJobCompletedProcessor().Process(packetDecodeResultModel);
                }

                case CommandConstants.CMD_ACK_JOB_STATUS_QUANTITY_UPDATE:
                {
                    return new AckJobQuantityUpdatedProcessor().Process(packetDecodeResultModel);
                }

//                case CommandConstants.CMD_ACK_CHART_DATA: {
//
//                    Object data = RequestManager.Instance().GetRequest(resultModel.Header.SeqID);
//
//                    RequestManager.Instance().CompleteRequest(resultModel.Header.SeqID);
//
//                    if (data != null) {
//
//                        DeviceDataBaseModel baseDataModel = (DeviceDataBaseModel) data;
//
//                        switch (baseDataModel.getCommandType()) {
////                            case CommandConstants.DEVICE_DATA_COMMAND_CHART_DATA: {
////                                DeviceChartDataModel deviceChartDataModel = (DeviceChartDataModel) data;
////
////                                for (ChartDataModel model : deviceChartDataModel.getChartDataModels()) {
////
////                                    DBChartDataTableRowModel dbChartDataTableRowModel = new DBChartDataTableRowModel();
////                                    dbChartDataTableRowModel.setChartId(model.getChartId());
////                                    dbChartDataTableRowModel.setEntryTime(model.getEntryDate());
////                                    dbChartDataTableRowModel.setSynced(true);
////
////                                    DatabaseManager.UpdateRow(new DBChartDataTableQueryModel(), dbChartDataTableRowModel, DBChartDataTableQueryModel.UPDATE_SYNC_STATE_WITH_CHART_ID_ENTRY_TIME);
////                                }
////                            }
////                            break;
////                            case CommandConstants.DEVICE_DATA_COMMAND_CHART_UNSYNC_DATA: {
////
////                                DeviceChartDataStartupSyncModel deviceChartDataStartupSyncModel = (DeviceChartDataStartupSyncModel)data;
////
////                                for (DBChartDataTableRowModel model : deviceChartDataStartupSyncModel.getUnSyncChartData()) {
////                                    DBChartDataTableRowModel dbChartDataTableRowModel = new DBChartDataTableRowModel();
////                                    dbChartDataTableRowModel.setChartId(model.getChartId());
////                                    dbChartDataTableRowModel.setEntryTime(model.getEntryTime());
////                                    dbChartDataTableRowModel.setSynced(true);
////
////                                    DatabaseManager.UpdateRow(new DBChartDataTableQueryModel(), dbChartDataTableRowModel, DBChartDataTableQueryModel.UPDATE_SYNC_STATE_WITH_CHART_ID_ENTRY_TIME);
////                                }
////
////                                packetProcessResultModel.CanSendServerCommand = true;
////                                packetProcessResultModel.ServerCommandPacket = PacketHelper.GetDeviceSynCompletedPacket();
////                                packetProcessResultModel.IsSuccess = true;
////                            }
////                            break;
//                        }
//                    }
//                    break;
//                }
            }
        } catch (Exception ex) {
            //TODO: Log exception error
            packetProcessResultModel.IsSuccess = false;
            AppLogger.getInstance().Log(ex, "Error occured in AcknowledgementProcessor");
        }

        return packetProcessResultModel;
    }


//    private PacketProcessResultModel processUnSyncChartData(List<DBChartDataTableRowModel> unSyncChartData){
//        //TODO: For multiple chart multiple packet need to create according to chartID
//        PacketProcessResultModel packetProcessResultModel =new PacketProcessResultModel();
//        PacketChartDataModel packetChartDataModel = new PacketChartDataModel();
//
//        SimpleDateFormat UTCEntryTimeFormat = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.SSS'Z'");
//        UTCEntryTimeFormat.setTimeZone(TimeZone.getTimeZone("GMT"));
//
//        for (DBChartDataTableRowModel dbChartDataModel:unSyncChartData) {
//
//            PacketChartTaskDataModel packetChartTaskDataModel = new PacketChartTaskDataModel();
//            packetChartTaskDataModel.day = UTCEntryTimeFormat.format(dbChartDataModel.getChartDay().getTime());
//            packetChartTaskDataModel.endSlotTimeObject = UTCEntryTimeFormat.format(dbChartDataModel.getSlotEndTime());
//            packetChartTaskDataModel.entryTime = UTCEntryTimeFormat.format(dbChartDataModel.getEntryTime());
//            packetChartTaskDataModel.slot = dbChartDataModel.getSlotId();
//            packetChartTaskDataModel.taskId = dbChartDataModel.getTaskId();
//            packetChartTaskDataModel.state = dbChartDataModel.getCellState();
//            packetChartTaskDataModel.startSlotTime =UTCEntryTimeFormat.format(dbChartDataModel.getSlotStartTime());
//
//            packetChartDataModel.packetChartTaskDataModels.add(packetChartTaskDataModel);
//            packetChartDataModel.chartId = dbChartDataModel.getChartId();
//        }
//
//        int requestId = RequestManager.Instance().GenerateRequestID();
//
//        DeviceChartDataStartupSyncModel deviceChartDataStartupSyncModel = new DeviceChartDataStartupSyncModel();
//        deviceChartDataStartupSyncModel.setCommandType(CommandConstants.DEVICE_DATA_COMMAND_CHART_UNSYNC_DATA);
//        deviceChartDataStartupSyncModel.setUnSyncChartData(unSyncChartData);
//        RequestManager.Instance().AddRequest(requestId, deviceChartDataStartupSyncModel);
//
//        packetProcessResultModel.CanSendServerCommand = true;
//        packetProcessResultModel.ServerCommandPacket = PacketHelper.GetChartDataPacket(requestId, packetChartDataModel);
//        Log.d("SyncDataPck",packetProcessResultModel.ServerCommandPacket);
//        packetProcessResultModel.IsSuccess = true;
//        return  packetProcessResultModel;
//
//    }
}
