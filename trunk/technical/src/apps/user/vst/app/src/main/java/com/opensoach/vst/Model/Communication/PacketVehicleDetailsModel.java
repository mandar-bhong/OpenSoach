package com.opensoach.vst.Model.Communication;

import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;

public class PacketVehicleDetailsModel {

    @SerializedName("tokenid")
    public Integer TokenID;

    @SerializedName("customerdetails")
    public PacketServiceCustomerDetailsDataModel CustomerDetails;

    @SerializedName("vehicledetails")
    public PacketServiceVehicleDetailsDataModel VehicleDetails;

    @SerializedName("tasks")
    public ArrayList<PacketServiceTaskItemDataModel> Tasks;

}
