package com.opensoach.vst.Processor;

import com.opensoach.vst.Constants.CommandConstants;
import com.opensoach.vst.Constants.Constants;
import com.opensoach.vst.Helper.CommonHelper;
import com.opensoach.vst.Helper.PacketHelper;
import com.opensoach.vst.Manager.RequestManager;
import com.opensoach.vst.Model.Communication.CommandRequest;
import com.opensoach.vst.Model.Communication.PacketHeaderModel;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.Communication.PacketTokenClaimDataModel;
import com.opensoach.vst.Model.PacketDecodeResultModel;
import com.opensoach.vst.Model.PacketProcessResultModel;
import com.opensoach.vst.Utility.AppLogger;

public class AckServerSyncCompletedProcessor implements IProcessor {

    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {
        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            packetProcessResultModel.CanSendServerCommand = true;

            Integer requestID  = RequestManager.Instance().GenerateRequestID();

            CommandRequest<String> request = new CommandRequest<>();
            request.AckProcessor = new TokenListProcessor();

            RequestManager.Instance().AddRequest(requestID,request);

            PacketHeaderModel packetHeaderModel = PacketHelper.CreatePacketHeader(CommandConstants.CMD_CAT_CONFIG,CommandConstants.CMD_CONFIG_GET_TOKEN_LIST,requestID,1);
            PacketModel<String> packetModel = new PacketModel<>();
            packetModel.Header = packetHeaderModel;

            request.Packet = packetModel;

            packetProcessResultModel.ServerCommandPacket = CommonHelper.GetPacketJSON(packetModel);

            packetProcessResultModel.IsSuccess = true;

        } catch (Exception ex) {
            AppLogger.getInstance().Log(ex);
        }

        return packetProcessResultModel;
    }
}
