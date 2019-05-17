package com.opensoach.vst.Model.Communication;

import com.google.gson.annotations.SerializedName;

public class PacketServiceCustomerDetailsDataModel {
    @SerializedName("firstname")
    public String FirstName;

    @SerializedName("lastname")
    public String LastName;

    @SerializedName("mobno")
    public String MobileNo;
}
