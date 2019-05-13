package com.opensoach.hpft.Model.Communication;
import com.google.gson.annotations.SerializedName;

public  class  PacketModel<T>{

    @SerializedName("header")
    public  PacketHeaderModel Header;
    @SerializedName("payload")
    public T Payload;

}



