package com.opensoach.hpft.Model.Communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by Mandar on 2/26/2017.
 */

public class PacketAuthenticationModel extends  PacketPayloadModel{
    @SerializedName("token")
    public String AuthToken;
}
