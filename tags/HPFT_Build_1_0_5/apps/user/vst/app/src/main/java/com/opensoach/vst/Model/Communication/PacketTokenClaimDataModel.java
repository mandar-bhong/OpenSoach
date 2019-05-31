package com.opensoach.vst.Model.Communication;

import com.google.gson.annotations.SerializedName;

public class PacketTokenClaimDataModel {

    @SerializedName("tokenid")
    public Integer TokenID;

    @SerializedName("fopcode")
    public String OperatorCode;

}
