package com.opensoach.vst.Model.Communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by samir.s.bukkawar on 5/6/2017.
 */

public class PacketUserComplaintDataModel {

    @SerializedName("complainttitle")
    public String ComplaintTitle;

    @SerializedName("description")
    public String Description;

    @SerializedName("complaintby")
    public String ComplaintBy;

    @SerializedName("mobileno")
    public String MobileNo;

    @SerializedName("emailid")
    public String EmailId;

    @SerializedName("employeeid")
    public String EmployeeID;

    @SerializedName("raisedon")
    public String RaisedOn;
}
