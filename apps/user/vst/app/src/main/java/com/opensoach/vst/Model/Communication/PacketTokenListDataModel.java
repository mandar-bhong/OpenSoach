package com.opensoach.vst.Model.Communication;

import com.google.gson.annotations.SerializedName;

import java.util.Date;

public class PacketTokenListDataModel {

    @SerializedName("tokenid")
    public Integer TokenID;

    @SerializedName("token")
    public Integer Token;

    @SerializedName("state")
    public Integer State;

    @SerializedName("generatedon")
    public Date GeneratedOn;


    @SerializedName("vehicleno")
    public String VehicleNo;

    @SerializedName("vhlid")
    public Integer VehicleID;

}
