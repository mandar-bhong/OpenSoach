package spl.hkt.opensoach.splapp.model.communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by samir.s.bukkawar on 3/25/2017.
 */

public class PacketChartTaskDataModel {

    @SerializedName("taskid")
   public int taskId;

    @SerializedName("entrytime")
    public String entryTime;

    @SerializedName("slot")
    public int slot;

    @SerializedName("state")
    public int state;

    @SerializedName("startSlotTime")
    public String startSlotTime;

    @SerializedName("endSlotTimeObject")
    public String endSlotTimeObject;

    @SerializedName("day")
    public String day;
}
