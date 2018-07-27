package com.opensoach.hospital.Model.Communication;

import com.google.gson.annotations.SerializedName;

import java.util.Date;

/**
 * Created by Mandar on 25-11-2017.
 */

public class PacketJobStartDataModel extends  PacketPayloadModel{
    @SerializedName("jobid")
    public Integer JobId;
    @SerializedName("operatorcode")
    public String OperatorCode;
    @SerializedName("starttime")
    public Date StartTime;
}
