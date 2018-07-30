package com.opensoach.hpft.Model.Communication;

import com.google.gson.annotations.SerializedName;

public class PacketServiceInstanceTxnModel {

    @SerializedName("servinid")
    public int servinid;

    @SerializedName("txndata")
    public String txndata;

    @SerializedName("txndate")
    public String txndate;

    @SerializedName("status")
    public int status;

    @SerializedName("fopcode")
    public String fopcode;
}
