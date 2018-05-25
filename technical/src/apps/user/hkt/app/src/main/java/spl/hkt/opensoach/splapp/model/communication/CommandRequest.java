package spl.hkt.opensoach.splapp.model.communication;

import spl.hkt.opensoach.splapp.processor.IProcessor;

public class CommandRequest<T> {

   public PacketModel<T> Packet;

   public IProcessor AckProcessor;
}
