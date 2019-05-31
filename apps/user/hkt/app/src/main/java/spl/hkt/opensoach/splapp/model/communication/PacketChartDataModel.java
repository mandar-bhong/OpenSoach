package spl.hkt.opensoach.splapp.model.communication;

import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;

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
