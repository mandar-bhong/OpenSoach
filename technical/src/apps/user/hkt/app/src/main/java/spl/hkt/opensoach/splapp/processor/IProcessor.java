package spl.hkt.opensoach.splapp.processor;

import spl.hkt.opensoach.splapp.model.PacketDecodeResultModel;
import spl.hkt.opensoach.splapp.model.PacketProcessResultModel;
import spl.hkt.opensoach.splapp.model.communication.PacketModel;

/**
 * Created by Mandar on 2/26/2017.
 */

public interface IProcessor {

    PacketProcessResultModel Process(PacketDecodeResultModel resultModel);
}
