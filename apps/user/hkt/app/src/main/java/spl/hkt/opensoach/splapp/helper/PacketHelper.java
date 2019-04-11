package spl.hkt.opensoach.splapp.helper;

import com.google.gson.Gson;

import spl.hkt.opensoach.splapp.model.communication.APIAuthRequesetModel;
import spl.hkt.opensoach.splapp.model.communication.PacketHeaderModel;

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

    public static String GetAPIAuthRequestJson(String serialnumber) {
        APIAuthRequesetModel apiAuthRequesetModel = new APIAuthRequesetModel();
        apiAuthRequesetModel.SerialNumber = serialnumber;
        apiAuthRequesetModel.ProdCode = "SPL_HKT";
        return new Gson().toJson(apiAuthRequesetModel);
    }
}
