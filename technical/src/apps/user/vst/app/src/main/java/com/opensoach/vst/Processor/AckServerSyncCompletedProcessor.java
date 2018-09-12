package com.opensoach.vst.Processor;

import com.opensoach.vst.Constants.CommandConstants;
import com.opensoach.vst.Constants.Constants;
import com.opensoach.vst.Helper.CommonHelper;
import com.opensoach.vst.Helper.PacketHelper;
import com.opensoach.vst.Model.Communication.PacketHeaderModel;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.PacketDecodeResultModel;
import com.opensoach.vst.Model.PacketProcessResultModel;

public class AckServerSyncCompletedProcessor implements IProcessor {

    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {
        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            packetProcessResultModel.CanSendServerCommand = true;

           PacketHeaderModel packetHeaderModel = PacketHelper.CreatePacketHeader(CommandConstants.CMD_CAT_CONFIG,CommandConstants.CMD_CONFIG_GET_TOKEN_LIST,1,1);
            PacketModel<String> packetModel = new PacketModel<>();
            packetModel.Header = packetHeaderModel;


            packetProcessResultModel.ServerCommandPacket = CommonHelper.GetPacketJSON(packetModel);

            packetProcessResultModel.IsSuccess = true;

        } catch (Exception ex) {

        }

        return packetProcessResultModel;
    }
}
