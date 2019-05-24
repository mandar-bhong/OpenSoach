package spl.hkt.opensoach.splapp.model.communication;

import com.google.gson.annotations.SerializedName;

import java.util.List;


public class PacketChartConfigurationModel {

    @SerializedName("servinid")
    public int ChartID;
    @SerializedName("conftypecode")
    public String ConfTypeCode;

    @SerializedName("servconfid")
    public int ServConfID;

    @SerializedName("servconfname")
    public String ChartName;

    @SerializedName("servconf")
    public String ServConf;
}
