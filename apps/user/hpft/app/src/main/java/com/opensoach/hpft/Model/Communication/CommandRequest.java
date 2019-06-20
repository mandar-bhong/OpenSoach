package com.opensoach.hpft.Model.Communication;

import com.opensoach.hpft.Processor.IProcessor;

public class CommandRequest<T> {

   public PacketModel<T> Packet;

   public IProcessor AckProcessor;
}
