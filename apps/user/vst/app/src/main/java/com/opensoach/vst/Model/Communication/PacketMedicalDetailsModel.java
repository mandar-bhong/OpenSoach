package com.opensoach.vst.Model.Communication;

import com.google.gson.annotations.SerializedName;

/**
 * Created by Mandar on 03-08-2018.
 */

public class PacketMedicalDetailsModel {

    @SerializedName("allergies")
    public String Allergies;

    @SerializedName("treatmentdone")
    public String Treatment;

    @SerializedName("reasonadmission")
    public String ReasonAdmission;

    @SerializedName("patientmedicalhistory")
    public String MedicalHistory;
}
