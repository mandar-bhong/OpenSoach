package com.opensoach.vst.PacketGenerator;

import com.opensoach.vst.Constants.CommandConstants;
import com.opensoach.vst.Helper.PacketHelper;
import com.opensoach.vst.Manager.RequestManager;
import com.opensoach.vst.Model.Communication.CommandRequest;
import com.opensoach.vst.Model.Communication.PacketHeaderModel;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.Communication.PacketTokenClaimDataModel;
import com.opensoach.vst.Model.Communication.PacketTokenCreateDataModel;
import com.opensoach.vst.ViewModels.CreateTokenViewModel;
import com.opensoach.vst.ViewModels.TokenItemViewModel;
import com.opensoach.vst.ViewModels.TokenSelectionViewModel;

public class TokenClaimPacketGenerator implements IPacketGenerator<TokenItemViewModel>  {

    @Override
    public CommandRequest GenerateRequest(int locationID, TokenItemViewModel data) {

        Integer requestID  = RequestManager.Instance().GenerateRequestID();

        CommandRequest<PacketTokenClaimDataModel> request = new CommandRequest<>();

        request.Packet = new PacketModel<PacketTokenClaimDataModel>();
        PacketTokenClaimDataModel packetTokenClaimDataModel = new PacketTokenClaimDataModel();
        packetTokenClaimDataModel.TokenID = data.getDbTokenTableRowModel().getId();
        packetTokenClaimDataModel.OperatorCode = "";

        request.Packet.Payload = packetTokenClaimDataModel;

        request.Packet.Header = new PacketHeaderModel();
        request.Packet.Header = PacketHelper.CreatePacketHeader(CommandConstants.CMD_CAT_DATA,
                CommandConstants.CMD_DATA_CLAIM_TOKEN, requestID, locationID);


        return request;
    }

    @Override
    public CommandRequest GenerateUnsyncRequest(int locationID) {
        return null;
    }

}
