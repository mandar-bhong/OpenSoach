package spl.hkt.opensoach.splapp.packetGenerator;

import spl.hkt.opensoach.splapp.model.communication.CommandRequest;

public interface IPacketGenerator<T> {
    CommandRequest GenerateRequest(int locationID,T data);
    CommandRequest GenerateUnsyncRequest(int locationID);
}
