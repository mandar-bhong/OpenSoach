package spl.hkt.opensoach.splapp.processor;

import java.util.ArrayList;

import spl.hkt.opensoach.splapp.manager.RequestManager;
import spl.hkt.opensoach.splapp.model.PacketDecodeResultModel;
import spl.hkt.opensoach.splapp.model.PacketProcessResultModel;
import spl.hkt.opensoach.splapp.model.communication.CommandRequest;
import spl.hkt.opensoach.splapp.model.communication.PacketFeedbackDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketServiceInstanceTxnModel;
import spl.hkt.opensoach.splapp.model.communication.PacketSimpleAckModel;

public class AckFeedbackProcessor implements IProcessor {
    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {
        CommandRequest<ArrayList<PacketFeedbackDataModel>> request = (CommandRequest) RequestManager.Instance().GetRequest(packetDecodeResultModel.Packet.Header.SeqID);
        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();
        if (request == null) {
            packetProcessResultModel.IsSuccess = true;
            return packetProcessResultModel;
        }

        RequestManager.Instance().CompleteRequest(packetDecodeResultModel.Packet.Header.SeqID);

        PacketSimpleAckModel ack = (PacketSimpleAckModel) packetDecodeResultModel.Packet.Payload;

        if (!ack.Ack) {
            packetProcessResultModel.IsSuccess = false;
            // TODO: save in local storage
            return packetProcessResultModel;
        }

        // TODO: save in local storage

        packetProcessResultModel.IsSuccess = true;

        return packetProcessResultModel;
    }
}

