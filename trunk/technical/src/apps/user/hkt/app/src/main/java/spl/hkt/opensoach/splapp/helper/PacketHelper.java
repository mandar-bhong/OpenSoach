package spl.hkt.opensoach.splapp.helper;

import spl.hkt.opensoach.splapp.manager.RequestManager;
import spl.hkt.opensoach.splapp.model.communication.PacketAuthenticationModel;
import spl.hkt.opensoach.splapp.model.communication.PacketChartDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketFeedbackDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketHeaderModel;
import spl.hkt.opensoach.splapp.model.communication.PacketModel;
import spl.hkt.opensoach.splapp.model.communication.PacketPayloadModel;
import spl.hkt.opensoach.splapp.model.communication.PacketUserComplaintDataModel;

/**
 * Created by Mandar on 2/26/2017.
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
        packetAuthenticationModel.SerialNumber = serialNumber;

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

    public static String GetChartDataPacket(int requestId, PacketChartDataModel model) {


        PacketModel<PacketChartDataModel> packetModel = new PacketModel<>();
        int seqid = RequestManager.Instance().GenerateRequestID();
        packetModel.Header = CreatePacketHeader(CommandConstants.CMD_CAT_DATA,
                CommandConstants.CMD_DATA_CHART_DATA, requestId, 0);

        packetModel.Payload = model;

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
