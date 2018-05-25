package spl.hkt.opensoach.splapp.helper;

import com.google.gson.Gson;

import java.util.ArrayList;

import spl.hkt.opensoach.splapp.manager.RequestManager;
import spl.hkt.opensoach.splapp.model.communication.APIAuthRequesetModel;
import spl.hkt.opensoach.splapp.model.communication.CommandRequest;
import spl.hkt.opensoach.splapp.model.communication.PacketAuthenticationModel;
import spl.hkt.opensoach.splapp.model.communication.PacketChartDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketFeedbackDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketHeaderModel;
import spl.hkt.opensoach.splapp.model.communication.PacketModel;
import spl.hkt.opensoach.splapp.model.communication.PacketPayloadModel;
import spl.hkt.opensoach.splapp.model.communication.PacketServiceInstanceTxnModel;
import spl.hkt.opensoach.splapp.model.communication.PacketUserComplaintDataModel;
import spl.hkt.opensoach.splapp.processor.AckChartDataProcessor;
import spl.hkt.opensoach.splapp.processor.AckDeviceRegProcessor;

/**
 * Created by Mandar on 2/26/2017.
 */

public class PacketHelper {

    public static PacketHeaderModel CreatePacketHeader(int category, int commandID, int seqID, int locationID) {

        PacketHeaderModel packetHeaderModel = new PacketHeaderModel();
        packetHeaderModel.CommandID = commandID;
        packetHeaderModel.Category = category;
        packetHeaderModel.LocationID = locationID;
        packetHeaderModel.SeqID = seqID;
        packetHeaderModel.CRC = "1";
        return packetHeaderModel;
    }

    public  static  String GetAPIAuthRequestJson(String serialnumber){
        APIAuthRequesetModel apiAuthRequesetModel = new APIAuthRequesetModel();
        apiAuthRequesetModel.SerialNumber = serialnumber;
        return  new Gson().toJson(apiAuthRequesetModel);
    }

    public static String GetStartUpPacket(String authToken) {

        //{"header":{"crc":"12","category":1,"commandid":1,"seqid":3},"payload":{"serialno":"P314140101054022"}}

        PacketModel<PacketAuthenticationModel> packetModel = new PacketModel<>();
        int seqid = RequestManager.Instance().GenerateRequestID();
        packetModel.Header = CreatePacketHeader(CommandConstants.CMD_CAT_DEVICE_REG,
                								CommandConstants.CMD_DEVICE_REGISTRATION, seqid, 0);

        PacketAuthenticationModel packetAuthenticationModel = new PacketAuthenticationModel();
        packetAuthenticationModel.AuthToken = authToken;

        packetModel.Payload = packetAuthenticationModel;

        CommandRequest<PacketAuthenticationModel> commandRequest = new CommandRequest<>();
        commandRequest.Packet= packetModel;
        commandRequest.AckProcessor = new AckDeviceRegProcessor();

        RequestManager.Instance().AddRequest(seqid, commandRequest);

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

    public static String GetChartDataPacket(ArrayList<PacketServiceInstanceTxnModel> model) {
        PacketModel<ArrayList<PacketServiceInstanceTxnModel>> packetModel = new PacketModel<>();
        int seqid = RequestManager.Instance().GenerateRequestID();
        packetModel.Header = CreatePacketHeader(CommandConstants.CMD_CAT_DATA,
                CommandConstants.CMD_DATA_CHART_DATA, seqid, 0);

        packetModel.Payload = model;

        CommandRequest<ArrayList<PacketServiceInstanceTxnModel>> commandRequest = new CommandRequest<>();
        commandRequest.Packet= packetModel;
        commandRequest.AckProcessor = new AckChartDataProcessor();

        RequestManager.Instance().AddRequest(seqid, commandRequest);

        String packetJSON = CommonHelper.GetPacketJSON(packetModel);

        return packetJSON;
    }

    public static String GetComplaintPacket(PacketUserComplaintDataModel model) {

        PacketModel<PacketUserComplaintDataModel> packetModel = new PacketModel<>();

        packetModel.Header = CreatePacketHeader(CommandConstants.CMD_CAT_DATA,CommandConstants.CMD_DATA_COMPLAINT_DATA,0,model.LocationId);
        packetModel.Payload = model;

        String packetJSON = CommonHelper.GetPacketJSON(packetModel);

        return packetJSON;
    }

    public static String GetFeedbackPacket(PacketFeedbackDataModel model) {

        PacketModel<PacketFeedbackDataModel> packetModel = new PacketModel<>();

        packetModel.Header = CreatePacketHeader(CommandConstants.CMD_CAT_DATA,CommandConstants.CMD_DATA_FEEDBACK_DATA,0,model.Rating);
        packetModel.Payload = model;

        String packetJSON = CommonHelper.GetPacketJSON(packetModel);

        return packetJSON;
    }
}
