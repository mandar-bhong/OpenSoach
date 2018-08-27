package com.opensoach.vst.Helper;

import com.google.gson.Gson;

import com.opensoach.vst.Model.Communication.APIAuthRequesetModel;
import com.opensoach.vst.Model.Communication.PacketHeaderModel;

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
        apiAuthRequesetModel.ProdCode = "SPL_HPFT";
        return new Gson().toJson(apiAuthRequesetModel);
    }
}
