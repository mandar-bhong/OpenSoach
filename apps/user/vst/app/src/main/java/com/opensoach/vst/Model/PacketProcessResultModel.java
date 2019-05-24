package com.opensoach.vst.Model;

/**
 * Created by Mandar on 2/25/2017.
 */

public class PacketProcessResultModel extends ProcessResultModel {
    public boolean CanUpdateUI;
    public boolean CanSendServerCommand;
    public AppNotificationModelBase UINotifierModel;
    public String ServerCommandPacket;

}
