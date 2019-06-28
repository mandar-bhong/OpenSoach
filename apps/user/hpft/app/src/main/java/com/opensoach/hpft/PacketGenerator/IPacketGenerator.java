package com.opensoach.hpft.PacketGenerator;

import com.opensoach.hpft.Model.Communication.CommandRequest;

public interface IPacketGenerator<T> {
    CommandRequest GenerateRequest(int locationID,T data);
    CommandRequest GenerateUnsyncRequest(int locationID);
}
