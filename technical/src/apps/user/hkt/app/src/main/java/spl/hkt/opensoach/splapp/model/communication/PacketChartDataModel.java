package spl.hkt.opensoach.splapp.model.communication;

import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;

/**
 * Created by samir.s.bukkawar on 3/25/2017.
 */

public class PacketChartDataModel  {

    @SerializedName("tasks")
   public ArrayList<PacketChartTaskDataModel> packetChartTaskDataModels;

    @SerializedName("chartid")
   public int chartId;

    public PacketChartDataModel() {
        packetChartTaskDataModels = new ArrayList<PacketChartTaskDataModel>();
    }
}
