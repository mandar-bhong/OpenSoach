package com.opensoach.vst.Processor;

import com.google.gson.Gson;

import com.opensoach.vst.Model.ErrorModel;
import com.opensoach.vst.Model.PacketDecodeResultModel;
import com.opensoach.vst.Model.Communication.PacketModel;

/**
 * Created by Mandar on 2/22/2017.
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
        }
        return result;
    }
}