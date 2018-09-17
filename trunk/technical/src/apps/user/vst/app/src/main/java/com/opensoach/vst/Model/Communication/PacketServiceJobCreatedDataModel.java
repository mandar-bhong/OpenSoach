package com.opensoach.vst.Model.Communication;

import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;
import java.util.Date;

public class PacketServiceJobCreatedDataModel {

    @SerializedName("tokenid")
    public Integer TokenId;

    @SerializedName("tokenno")
    public Integer TokenNo;

    @SerializedName("servinid")
    public Integer ServInId;

    @SerializedName("status")
    public Integer Status;

    @SerializedName("txndata")
    public String TxnData;


    @SerializedName("txndate")
    public Date TxnDate;


}

