package com.opensoach.hpft.Model.Communication;

import com.google.gson.annotations.SerializedName;


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

    @SerializedName("medicaldetails")
    public String MedicalDetails;

    @SerializedName("patientdetails")
    public String PatientDetails;
}
