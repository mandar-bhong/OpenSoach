package com.opensoach.vst.Model.Communication;

import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;

public class PacketServiceJobDetailsDataModel {

    @SerializedName("tokenid")
    public Integer TokenID;

    @SerializedName("serviceconfig")
    public ArrayList<PacketServiceInstanceModel> ServiceConfig;

    @SerializedName("serviceexeconfig")
    public ArrayList<PacketServiceInstanceModel> ServiceExeConfig;

}
