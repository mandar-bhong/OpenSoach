package com.opensoach.vst.PacketGenerator;

import com.opensoach.vst.Model.Communication.CommandRequest;

public interface IPacketGenerator<T> {
    CommandRequest GenerateRequest(int locationID,T data);
    CommandRequest GenerateUnsyncRequest(int locationID);
}
