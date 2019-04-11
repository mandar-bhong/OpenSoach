package com.opensoach.vst.Model.Communication;

import com.google.gson.annotations.SerializedName;

import java.util.Date;

public class PacketTokenCreateResponseDataModel {

    @SerializedName("tokenid")
    public Integer TokenID;

    @SerializedName("token")
    public Integer Token;

    @SerializedName("state")
    public Integer State;

    @SerializedName("generatedon")
    public Date GeneratedOn;

}
