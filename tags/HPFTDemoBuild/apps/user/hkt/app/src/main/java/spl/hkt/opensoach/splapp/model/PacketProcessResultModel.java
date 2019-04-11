package spl.hkt.opensoach.splapp.model;

import spl.hkt.opensoach.splapp.model.communication.PacketModel;
import spl.hkt.opensoach.splapp.processor.IProcessor;

/**
 * Created by Mandar on 2/25/2017.
 */

public class PacketProcessResultModel extends ProcessResultModel {
    public boolean CanUpdateUI;
    public boolean CanSendServerCommand;
    public AppNotificationModelBase UINotifierModel;
    public String ServerCommandPacket;

}
