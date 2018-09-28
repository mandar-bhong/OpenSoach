package com.opensoach.vst.Model.Communication;

import com.google.gson.annotations.SerializedName;

public class PacketServiceInstanceModel {

    @SerializedName("servinid")
    public int servinid;

    @SerializedName("servintxnid")
    public int servintxnid;

    @SerializedName("status")
    public int status;

    @SerializedName("txndata")
    public String txndata;

    @SerializedName("txndate")
    public String txndate;

}
