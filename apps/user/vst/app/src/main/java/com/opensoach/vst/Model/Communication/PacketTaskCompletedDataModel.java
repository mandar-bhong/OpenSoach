package com.opensoach.vst.Model.Communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by Mandar on 07-08-2018.
 */

public class PacketTaskCompletedDataModel {

    @SerializedName("taskname")
    public String taskName;

    @SerializedName("slotstarttime")
    public int slotStartTime;

    @SerializedName("slotendtime")
    public int slotEndTime;

    @SerializedName("comment")
    public String comment;

    @SerializedName("value")
    public String value;

}
