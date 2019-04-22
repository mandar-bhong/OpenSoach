package com.opensoach.hpft.Model;

import com.opensoach.hpft.Model.Communication.PacketModel;
import com.opensoach.hpft.Processor.IProcessor;

/**
 * Created by Mandar on 2/26/2017.
 */

public class PacketDecodeResultModel extends ProcessResultModel {
    public String JSONPacket;
    public PacketModel Packet;
    public IProcessor Processor;
}
