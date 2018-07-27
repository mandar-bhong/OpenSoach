package com.opensoach.hospital.Model.Communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by Mandar on 8/26/2017.
 */

public class PacketAuthenticationModel extends  PacketPayloadModel{
    @SerializedName("serialno")
    public String SerialNumber;

    @SerializedName("username")
    public String UserName;

    @SerializedName("password")
    public String Password;
}
