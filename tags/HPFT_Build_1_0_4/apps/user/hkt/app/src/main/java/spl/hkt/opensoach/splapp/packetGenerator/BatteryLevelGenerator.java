package spl.hkt.opensoach.splapp.packetGenerator;

import java.util.ArrayList;

import spl.hkt.opensoach.splapp.helper.CommandConstants;
import spl.hkt.opensoach.splapp.helper.PacketHelper;
import spl.hkt.opensoach.splapp.manager.RequestManager;
import spl.hkt.opensoach.splapp.model.communication.CommandRequest;
import spl.hkt.opensoach.splapp.model.communication.PacketBatteryLevelModel;
import spl.hkt.opensoach.splapp.model.communication.PacketFeedbackDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketModel;
import spl.hkt.opensoach.splapp.processor.AckFeedbackProcessor;

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
