package com.opensoach.hpft.Model.Communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by Mandar on 07-08-2018.
 */

public class PacketTimeConfigModel {

    @SerializedName("endtime")
    public Integer EndTime;

    @SerializedName("starttime")
    public Integer StartTime;

    @SerializedName("interval")
    public Integer Interval;
}
