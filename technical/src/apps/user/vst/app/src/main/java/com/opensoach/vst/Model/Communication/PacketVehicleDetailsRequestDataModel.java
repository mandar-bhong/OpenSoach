package com.opensoach.vst.Model.Communication;

import com.google.gson.annotations.SerializedName;

public class PacketVehicleDetailsRequestDataModel {

    @SerializedName("vehicleno")
    public String VehicleNumber;

    @SerializedName("vehicledetails")
    public PacketServiceVehicleDetailsDataModel VehicleDetails;
}
