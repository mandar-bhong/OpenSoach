package com.opensoach.vst.Processor;

import com.opensoach.vst.Model.PacketDecodeResultModel;
import com.opensoach.vst.Model.PacketProcessResultModel;

public class ServiceConfigProcessor implements IProcessor  {

    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {
        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {


            packetProcessResultModel.IsSuccess = true;

        }catch (Exception ex){

        }



        return packetProcessResultModel;
    }
}
