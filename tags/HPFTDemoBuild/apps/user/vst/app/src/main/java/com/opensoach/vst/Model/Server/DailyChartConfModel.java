package com.opensoach.vst.Model.Server;

import com.google.gson.annotations.SerializedName;

public class DailyChartConfModel {
    @SerializedName("taskconf")
    public DailyChartTaskConfModel TaskConf;
    @SerializedName("timeconf")
    public DailyChartTimeConfModel TimeConf;
}
