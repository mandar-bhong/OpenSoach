package com.opensoach.hospital.Processor;

import com.google.gson.Gson;
import com.opensoach.hospital.Model.Communication.PacketModel;
import com.opensoach.hospital.Model.ErrorModel;
import com.opensoach.hospital.Model.PacketDecodeResultModel;
import com.opensoach.hospital.Utility.AppLogger;

/**
 * Created by Mandar on 8/27/2017.
 */

public class HeaderDecoder {
    public static PacketDecodeResultModel Decode(String packet) {

        PacketDecodeResultModel result = new PacketDecodeResultModel();
        try {

            result.JSONPacket = packet;

            Gson gson = new Gson();

            PacketModel packetModel = gson.fromJson(packet, PacketModel.class);

            result.Packet = packetModel;
            result.IsSuccess = true;
        } catch (Exception ex) {
            result.Error = new ErrorModel();
            result.IsSuccess = false;
            AppLogger.getInstance().Log(ex,"Error occured in Header Decoder");
        }
        return result;
    }
}
