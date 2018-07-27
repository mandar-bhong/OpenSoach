package com.opensoach.hospital.Processor;

import com.opensoach.hospital.Model.PacketDecodeResultModel;
import com.opensoach.hospital.Model.PacketProcessResultModel;
import com.opensoach.hospital.Processor.Strategy.LocationNoCompairStrategy;

/**
 * Created by Mandar on 8/27/2017.
 */

public class LocationDataProcessor implements IProcessor  {

    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel resultModel) {

        IProcessor processor = new LocationNoCompairStrategy();

        return processor.Process(resultModel);

    }
}
