package spl.hkt.opensoach.splapp.model.server;

import com.google.gson.annotations.SerializedName;

public class DailyChartConfModel {
    @SerializedName("taskconf")
    public DailyChartTaskConfModel TaskConf;
    @SerializedName("timeconf")
    public DailyChartTimeConfModel TimeConf;
}
