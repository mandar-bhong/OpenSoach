package com.opensoach.vst.Model.Communication;

import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;

public class PacketServiceJobCreatedDataModel {

    @SerializedName("tokenid")
    public Integer TokenId;

    @SerializedName("tokenno")
    public Integer TokenNo;


    @SerializedName("customerdetails")
    public PacketServiceCustomerDetailsDataModel CustomerDetails;

    @SerializedName("vehicledetails")
    public PacketServiceVehicleDetailsDataModel VehicleDetails;

    @SerializedName("tasks")
    public ArrayList<PacketServiceTaskItemDataModel> Tasks;

}

