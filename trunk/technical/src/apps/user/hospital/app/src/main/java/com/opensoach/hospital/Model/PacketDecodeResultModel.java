package com.opensoach.hospital.Model;

import com.opensoach.hospital.Model.Communication.PacketModel;
import com.opensoach.hospital.Processor.IProcessor;

/**
 * Created by Mandar on 8/27/2017.
 */

public class PacketDecodeResultModel extends ProcessResultModel  {

    public String JSONPacket;
    public PacketModel Packet;
    public IProcessor Processor;
}
