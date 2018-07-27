package com.opensoach.hospital.Model.Communication;

import com.google.gson.annotations.SerializedName;

import java.util.Date;

/**
 * Created by Mandar on 9/24/2017.
 */

public class PacketJobCardStatusChangedDataModel {

    @SerializedName("jobcardid")
    public Integer JobCardID;

    @SerializedName("state")
    public Integer State;

    @SerializedName("statechangetime")
    public Date StateChangedTime;

}
