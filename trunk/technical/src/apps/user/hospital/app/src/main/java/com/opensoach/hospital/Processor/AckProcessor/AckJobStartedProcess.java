package com.opensoach.hospital.Processor.AckProcessor;

import com.opensoach.hospital.DAL.DatabaseManager;
import com.opensoach.hospital.Helper.ApplicationConstants;
import com.opensoach.hospital.Helper.Constants;
import com.opensoach.hospital.Manager.RequestManager;
import com.opensoach.hospital.Model.AppNotificationModelBase;
import com.opensoach.hospital.Model.Communication.PacketJobStartDataModel;
import com.opensoach.hospital.Model.Communication.PacketModel;
import com.opensoach.hospital.Model.Communication.PacketSimpleAckModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableQueryModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableRowModel;
import com.opensoach.hospital.Model.PacketDecodeResultModel;
import com.opensoach.hospital.Model.PacketProcessResultModel;
import com.opensoach.hospital.Processor.IProcessor;
import com.opensoach.hospital.Utility.AppLogger;

import java.util.List;

/**
 * Created by Mandar on 25-11-2017.
 */

public class AckJobStartedProcess implements IProcessor {

    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel resultModel) {

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {
            PacketModel<PacketSimpleAckModel> ackModel = (PacketModel<PacketSimpleAckModel>) resultModel.Packet.Payload;

            Object dataModel = RequestManager.Instance().GetRequest(ackModel.Header.SeqID);
            RequestManager.Instance().CompleteRequest(ackModel.Header.SeqID);

            if (dataModel == null ) {
                AppLogger.getInstance().Log(AppLogger.LogLevel.Error, "AckJobStarted: With Request id data model is null. Packet:" + resultModel.JSONPacket);
                return packetProcessResultModel;
            }

            if ( !(dataModel instanceof PacketJobStartDataModel)) {
                AppLogger.getInstance().Log(AppLogger.LogLevel.Error, "AckJobStarted: Request is not of PacketJobStartDataModel. Packet:" + resultModel.JSONPacket);
                return packetProcessResultModel;
            }

            PacketJobStartDataModel packetJobStartDataModel = (PacketJobStartDataModel) dataModel;

            DBJobCardTableRowModel dbSelectModel = new DBJobCardTableRowModel();
            dbSelectModel.setJobCardId(packetJobStartDataModel.JobId);

            List<DBJobCardTableRowModel> dbJobCardTableRowModelList = DatabaseManager.SelectByFilter(new DBJobCardTableQueryModel(), dbSelectModel, DBJobCardTableQueryModel.SELECT_ID_FILTER);

            if (dbJobCardTableRowModelList.size() == 0) {
                AppLogger.getInstance().Log(AppLogger.LogLevel.Error, "AckJobStarted: No row found in local db. Packet: " + resultModel.JSONPacket);
                return packetProcessResultModel;
            }

            DBJobCardTableRowModel dbJobCardTableRowModelItem = dbJobCardTableRowModelList.get(0);
            dbJobCardTableRowModelItem.setActualStartDate( packetJobStartDataModel.StartTime);

            if (ackModel.Payload.Ack == Constants.RESPONSE_ACK_SUCCESS) {
                dbJobCardTableRowModelItem.setState(Constants.JOB_STATE_STARTED);

                DatabaseManager.UpdateRow(new DBJobCardTableQueryModel(), dbJobCardTableRowModelItem, DBJobCardTableQueryModel.UPDATE_STATE_BY_ID_FILTER);

                packetProcessResultModel.UINotifierModel = new AppNotificationModelBase();
                packetProcessResultModel.UINotifierModel.Data = dbJobCardTableRowModelItem;
                packetProcessResultModel.UINotifierModel.DataProcessStatergyID = ApplicationConstants.UI_PROCESSING_STATERGY_JOB_STARTED_SUCCESS;

                packetProcessResultModel.CanUpdateUI = true;
                packetProcessResultModel.IsSuccess = true;

            } else {
                packetProcessResultModel.UINotifierModel = new AppNotificationModelBase();
                packetProcessResultModel.UINotifierModel.Data = dbJobCardTableRowModelItem;
                packetProcessResultModel.UINotifierModel.DataProcessStatergyID = ApplicationConstants.UI_PROCESSING_STATERGY_JOB_STARTED_FAILURE;

                packetProcessResultModel.CanUpdateUI = true;
                packetProcessResultModel.IsSuccess = true;
            }
        } catch (Exception ex) {
            packetProcessResultModel.IsSuccess = false;
            AppLogger.getInstance().Log(ex, "Error occured in AckJobStartedProcessor");
        }
        return packetProcessResultModel;
    }
}