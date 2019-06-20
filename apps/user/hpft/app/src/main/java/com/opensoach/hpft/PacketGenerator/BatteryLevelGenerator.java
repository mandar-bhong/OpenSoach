package com.opensoach.hpft.PacketGenerator;

import com.opensoach.hpft.Constants.CommandConstants;
import com.opensoach.hpft.Helper.PacketHelper;
import com.opensoach.hpft.Model.Communication.CommandRequest;
import com.opensoach.hpft.Model.Communication.PacketBatteryLevelModel;
import com.opensoach.hpft.Model.Communication.PacketModel;

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
