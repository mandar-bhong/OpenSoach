package com.opensoach.hpft.Model.Server;

import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;

public class DailyChartTaskConfModel {
    @SerializedName("tasks")
    public ArrayList<DailyChartTaskModel> Tasks;
}
