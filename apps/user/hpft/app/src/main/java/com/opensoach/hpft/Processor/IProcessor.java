package com.opensoach.hpft.Processor;

import com.opensoach.hpft.Model.PacketDecodeResultModel;
import com.opensoach.hpft.Model.PacketProcessResultModel;

/**
 * Created by Mandar on 2/26/2017.
 */

public interface IProcessor {

    PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel);
}
