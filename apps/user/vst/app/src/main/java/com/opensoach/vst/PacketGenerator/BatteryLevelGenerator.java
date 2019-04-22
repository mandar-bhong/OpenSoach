package com.opensoach.vst.PacketGenerator;

import com.opensoach.vst.Constants.CommandConstants;
import com.opensoach.vst.Helper.PacketHelper;
import com.opensoach.vst.Model.Communication.CommandRequest;
import com.opensoach.vst.Model.Communication.PacketBatteryLevelModel;
import com.opensoach.vst.Model.Communication.PacketModel;

public class BatteryLevelGenerator implements IPacketGenerator<PacketBatteryLevelModel> {
    @Override
    public CommandRequest GenerateRequest(int locationID, PacketBatteryLevelModel data) {

        PacketModel<PacketBatteryLevelModel> packetModel = new PacketModel<>();

        packetModel.Header = PacketHelper.CreatePacketHeader(CommandConstants.CMD_CAT_DATA,
                CommandConstants.CMD_DATA_BATTERY_LEVEL_DATA, 0, locationID);

        packetModel.Payload = data;

        CommandRequest<PacketBatteryLevelModel> commandRequest = new CommandRequest<>();
        commandRequest.Packet = packetModel;

        return commandRequest;
    }

    @Override
    public CommandRequest GenerateUnsyncRequest(int locationID) {
        return null;
    }
}
