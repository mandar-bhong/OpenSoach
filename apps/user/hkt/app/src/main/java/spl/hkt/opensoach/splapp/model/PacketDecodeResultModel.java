package spl.hkt.opensoach.splapp.model;

import spl.hkt.opensoach.splapp.model.communication.PacketModel;
import spl.hkt.opensoach.splapp.processor.IProcessor;

/**
 * Created by Mandar on 2/26/2017.
 */

public class PacketDecodeResultModel extends ProcessResultModel {
    public String JSONPacket;
    public PacketModel Packet;
    public IProcessor Processor;
}
