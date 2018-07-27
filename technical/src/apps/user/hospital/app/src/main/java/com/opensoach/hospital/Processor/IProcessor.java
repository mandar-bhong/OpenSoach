package com.opensoach.hospital.Processor;

import com.opensoach.hospital.Model.PacketDecodeResultModel;
import com.opensoach.hospital.Model.PacketProcessResultModel;

/**
 * Created by Mandar on 8/27/2017.
 */

public interface IProcessor {
    PacketProcessResultModel Process(PacketDecodeResultModel resultModel);
}
