package com.opensoach.vst.Model.Communication;

import com.google.gson.annotations.SerializedName;

public class PacketServiceOwnerVehicleDetailsDataModel {

    @SerializedName("details")
    public PacketServiceCustomerDetailsDataModel CustomerDetails;

    @SerializedName("vehicleno")
    public String VehicleNo;
}

