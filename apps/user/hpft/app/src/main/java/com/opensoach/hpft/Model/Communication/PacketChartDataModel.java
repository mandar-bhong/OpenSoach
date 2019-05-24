package com.opensoach.hpft.Model.Communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by samir.s.bukkawar on 3/25/2017.
 */

public class PacketChartDataModel {

    @SerializedName("taskname")
    public String taskName;

    @SerializedName("slotstarttime")
    public int slotStartTime;

    @SerializedName("slotendtime")
    public int slotEndTime;

    public PacketChartDataModel() {
    }
}
