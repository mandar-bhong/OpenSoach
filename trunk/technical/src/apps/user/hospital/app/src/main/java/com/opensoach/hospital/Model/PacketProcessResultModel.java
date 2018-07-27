package com.opensoach.hospital.Model;

/**
 * Created by Mandar on 8/27/2017.
 */

public class PacketProcessResultModel extends ProcessResultModel  {
    public boolean CanUpdateUI;
    public boolean CanSendServerCommand;
    public AppNotificationModelBase UINotifierModel;
    public String ServerCommandPacket;
}
