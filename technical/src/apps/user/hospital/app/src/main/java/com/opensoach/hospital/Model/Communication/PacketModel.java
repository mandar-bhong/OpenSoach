package com.opensoach.hospital.Model.Communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by Mandar on 8/26/2017.
 */

public class PacketModel<T> {
    @SerializedName("header")
    public  PacketHeaderModel Header;
    @SerializedName("payload")
    public T Payload;
}
