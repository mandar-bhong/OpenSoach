package spl.hkt.opensoach.splapp.model.communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by samir.s.bukkawar on 5/6/2017.
 */

public class PacketUserComplaintDataModel {

    @SerializedName("Description")
    public String Description;

    @SerializedName("ComplaintBy")
    public String ComplaintBy;

    @SerializedName("LocationId")
    public Integer LocationId;

    @SerializedName("RaisedOn")
    public String RaisedOn;

    @SerializedName("EmailId")
    public String EmailId;

    @SerializedName("EmployeeID")
    public String EmployeeID;

    @SerializedName("MobileNo")
    public String MobileNo;
}
