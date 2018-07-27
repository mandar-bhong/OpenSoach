package com.opensoach.hospital.Helper;

import com.google.gson.Gson;
import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.Manager.RequestManager;
import com.opensoach.hospital.Model.Communication.APIAuthRequesetModel;
import com.opensoach.hospital.Model.Communication.PacketAuthenticationModel;
import com.opensoach.hospital.Model.Communication.PacketHeaderModel;
import com.opensoach.hospital.Model.Communication.PacketJobQuantityUpdateDataModel;
import com.opensoach.hospital.Model.Communication.PacketJobStartDataModel;
import com.opensoach.hospital.Model.Communication.PacketJobStopDataModel;
import com.opensoach.hospital.Model.Communication.PacketModel;
import com.opensoach.hospital.Model.Communication.PacketPayloadModel;

import java.util.Date;

/**
 * Created by Mandar on 8/26/2017.
 */

public class PacketHelper {

    public static void CreatePacket() {

    }

    public static PacketHeaderModel CreatePacketHeader(int category, int commandID, int seqID, int locationID) {

        PacketHeaderModel packetHeaderModel = new PacketHeaderModel();
        packetHeaderModel.CommandID = commandID;
        packetHeaderModel.Category = category;
        packetHeaderModel.LocationID = locationID;
        packetHeaderModel.SeqID = seqID;
        packetHeaderModel.CRC = "1";
        return packetHeaderModel;
    }

    public static String GetStartUpPacket(String serialNumber) {

        //{"header":{"crc":"12","category":1,"commandid":1,"seqid":3},"payload":{"serialno":"P314140101054022"}}

        PacketModel<PacketAuthenticationModel> packetModel = new PacketModel<>();
        int seqid = RequestManager.Instance().GenerateRequestID();
        packetModel.Header = CreatePacketHeader(CommandConstants.CMD_CAT_DEVICE_REG,
                CommandConstants.CMD_DEVICE_REGISTRATION, seqid, 0);

        PacketAuthenticationModel packetAuthenticationModel = new PacketAuthenticationModel();
        packetAuthenticationModel.SerialNumber = "P706X69030105851";//''serialNumber;
        packetAuthenticationModel.UserName = AppRepo.getInstance().getUserName();
        packetAuthenticationModel.Password = AppRepo.getInstance().getPassword();

        packetModel.Payload = packetAuthenticationModel;

        String packetJSON = CommonHelper.GetPacketJSON(packetModel);

        return packetJSON;

    }

    public static String GetDeviceSynCompletedPacket() {

        //{"header":{"crc":"12","category":1,"commandid":1,"seqid":3},"payload":{"serialno":"P314140101054022"}}

        PacketModel<PacketPayloadModel> packetModel = new PacketModel<>();
        int seqid = RequestManager.Instance().GenerateRequestID();
        packetModel.Header = CreatePacketHeader(CommandConstants.CMD_CAT_CONFIG,
                CommandConstants.CMD_CONFIG_DEVICE_SYNC_COMPLETED, seqid, 0);


        packetModel.Payload = new PacketPayloadModel();

        String packetJSON = CommonHelper.GetPacketJSON(packetModel);

        return packetJSON;

    }

    public static PacketModel<PacketPayloadModel> GetJobStartPacket(int locationId, int jobId, String operatorCode, Date startTime) {
        PacketModel<PacketPayloadModel> packetModel = new PacketModel<>();
        int seqid = RequestManager.Instance().GenerateRequestID();
        packetModel.Header = CreatePacketHeader(CommandConstants.CMD_CAT_DATA,
                CommandConstants.CMD_DATA_START_JOB, seqid, locationId);


        PacketJobStartDataModel packetJobStartDataModel = new PacketJobStartDataModel();
        packetJobStartDataModel.JobId = jobId;
        packetJobStartDataModel.OperatorCode = operatorCode;
        packetJobStartDataModel.StartTime = startTime;

        packetModel.Payload = packetJobStartDataModel;

        return packetModel;
    }

    public static PacketModel<PacketPayloadModel> GetJobStopPacket(int locationId, int jobId, String operatorCode, Date endTime) {
        PacketModel<PacketPayloadModel> packetModel = new PacketModel<>();
        int seqid = RequestManager.Instance().GenerateRequestID();
        packetModel.Header = CreatePacketHeader(CommandConstants.CMD_CAT_DATA,
                CommandConstants.CMD_DATA_STOP_JOB, seqid, locationId);

        PacketJobStopDataModel packetJobStopDataModel = new PacketJobStopDataModel();
        packetJobStopDataModel.JobId = jobId;
        packetJobStopDataModel.OperatorCode = operatorCode;
        packetJobStopDataModel.StopTime = endTime;

        packetModel.Payload = packetJobStopDataModel;

        return packetModel;
    }

    public static PacketModel<PacketPayloadModel> GetJobDropPacket(int locationId, int jobId, String operatorCode, Date endTime) {
        PacketModel<PacketPayloadModel> packetModel = new PacketModel<>();
        int seqid = RequestManager.Instance().GenerateRequestID();
        packetModel.Header = CreatePacketHeader(CommandConstants.CMD_CAT_DATA,
                CommandConstants.CMD_DATA_DROP_JOB, seqid, locationId);

        PacketJobStopDataModel packetJobStopDataModel = new PacketJobStopDataModel();
        packetJobStopDataModel.JobId = jobId;
        packetJobStopDataModel.OperatorCode = operatorCode;
        packetJobStopDataModel.StopTime = endTime;

        packetModel.Payload = packetJobStopDataModel;

        return packetModel;
    }

    public static PacketModel<PacketPayloadModel> GetJobAbortPacket(int locationId, int jobId, String operatorCode, Date endTime) {
        PacketModel<PacketPayloadModel> packetModel = new PacketModel<>();
        int seqid = RequestManager.Instance().GenerateRequestID();
        packetModel.Header = CreatePacketHeader(CommandConstants.CMD_CAT_DATA,
                CommandConstants.CMD_DATA_ABORT_JOB, seqid, locationId);

        PacketJobStopDataModel packetJobStopDataModel = new PacketJobStopDataModel();
        packetJobStopDataModel.JobId = jobId;
        packetJobStopDataModel.OperatorCode = operatorCode;
        packetJobStopDataModel.StopTime = endTime;

        packetModel.Payload = packetJobStopDataModel;

        return packetModel;
    }

    public static PacketModel<PacketPayloadModel> GetJobQuantityUpdatePacket(int locationId, int jobId, String operatorCode,int finishedPartCount, Date completionTime,String comment) {

        PacketModel<PacketPayloadModel> packetModel = new PacketModel<>();
        int seqid = RequestManager.Instance().GenerateRequestID();
        packetModel.Header = CreatePacketHeader(CommandConstants.CMD_CAT_DATA,
                CommandConstants.CMD_DATA_UPDATE_JOB_UNIT, seqid, locationId);

        PacketJobQuantityUpdateDataModel packetJobStartDataModel = new PacketJobQuantityUpdateDataModel();
        packetJobStartDataModel.JobId = jobId;
        packetJobStartDataModel.OperatorCode = operatorCode;
        packetJobStartDataModel.FinishedPartCount = finishedPartCount;
        packetJobStartDataModel.CompletionTime = completionTime;
        packetJobStartDataModel.Comment = comment;

        packetModel.Payload = packetJobStartDataModel;

        return packetModel;
    }

    public static String GetAPIAuthRequestJson(String username,String password) {
        APIAuthRequesetModel apiAuthRequesetModel = new APIAuthRequesetModel();
        apiAuthRequesetModel.UserName = username;
        apiAuthRequesetModel.Password = password;
        return new Gson().toJson(apiAuthRequesetModel);
    }

}
