package spl.hkt.opensoach.splapp.model.communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by Mandar on 2/26/2017.
 */

public class PacketAuthenticationModel extends  PacketPayloadModel{
    @SerializedName("token")
    public String AuthToken;
}
