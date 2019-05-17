package com.opensoach.vst.Model.Communication;

import com.opensoach.vst.Processor.IProcessor;

public class CommandRequest<T> {

   public PacketModel<T> Packet;

   public IProcessor AckProcessor;
}
