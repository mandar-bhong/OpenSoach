package com.opensoach.hospital.Model.Communication;

import com.google.gson.annotations.SerializedName;

import java.util.Date;

/**
 * Created by Mandar on 9/5/2017.
 */

public class PacketJobCardDataModel {

    @SerializedName("jobcardid")
    public Integer JobID;
    @SerializedName("customer")
    public String Customer;
    @SerializedName("partid")
    public Integer PartID;
    @SerializedName("partcount")
    public Integer PartCount;
    @SerializedName("name")
    public String JobCode;
    @SerializedName("starttime")
    public Date StartTime;
    @SerializedName("endtime")
    public Date EndTime;
    @SerializedName("actualstarttime")
    public Date ActualStartTime;
    @SerializedName("actualendtime")
    public Date ActualEndTime;
    @SerializedName("state")
    public Integer State;
    @SerializedName("comments")
    public String Comments;
    @SerializedName("quantitycompleted")
    public int QuantityCompleted;
    @SerializedName("jobconfig")
    public String JobConfig;

}
