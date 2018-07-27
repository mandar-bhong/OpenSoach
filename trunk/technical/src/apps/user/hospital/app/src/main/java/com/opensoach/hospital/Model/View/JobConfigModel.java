package com.opensoach.hospital.Model.View;

import com.google.gson.annotations.SerializedName;

/**
 * Created by Mandar on 11-06-2018.
 */

public class JobConfigModel {
    @SerializedName("id")
    public  String workId;
    @SerializedName("work_type")
    public String worktype;
    @SerializedName("work_subtype")
    public String worksubtype;
    @SerializedName("address")
    public  String address;
    @SerializedName("start_dtm")
    public  String startdtm;
    @SerializedName("end_dtm")
    public  String enddtm;

}
