package com.opensoach.vst.Model;

import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Processor.IProcessor;

/**
 * Created by Mandar on 2/26/2017.
 */

public class PacketDecodeResultModel extends ProcessResultModel {
    public String JSONPacket;
    public PacketModel Packet;
    public IProcessor Processor;
}
