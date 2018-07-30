package com.opensoach.hpft.Model.Server;

import com.google.gson.annotations.SerializedName;

public class DailyChartTimeConfModel {
    @SerializedName("starttime")
    public int StartTime;
    @SerializedName("endtime")
    public int EndTime;
    @SerializedName("interval")
    public int Interval;
}
