package com.opensoach.vst.Model.Communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by Mandar on 03-08-2018.
 */

public class PacketCardListConfigurationModel {

    @SerializedName("servinid")
    public int SerInID;
    @SerializedName("conftypecode")
    public String ConfTypeCode;

    @SerializedName("servconfid")
    public int ServConfID;

    @SerializedName("servconfname")
    public String ConfigName;

    @SerializedName("servconf")
    public String ServConfJSON;

    @SerializedName("medicaldetails")
    public String MedicalDetailsJSON;

    @SerializedName("patientdetails")
    public String PatientDetailsJSON;

    public  int LocationID;
    public  PacketPatientDetailsModel PatientDetails;
    public  PacketMedicalDetailsModel MedicalDetails;
    public  PacketServiceConfModel ServiceConf;
}
