package com.opensoach.vst.Processor;

import com.opensoach.vst.Model.PacketDecodeResultModel;
import com.opensoach.vst.Model.PacketProcessResultModel;

/**
 * Created by Mandar on 2/26/2017.
 */

public interface IProcessor {

    PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel);
}
