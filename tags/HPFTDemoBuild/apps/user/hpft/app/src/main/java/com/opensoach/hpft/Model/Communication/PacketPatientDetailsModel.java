package com.opensoach.hpft.Model.Communication;

import com.google.gson.annotations.SerializedName;

import java.util.Date;

/**
 * Created by Mandar on 03-08-2018.
 */

public class PacketPatientDetailsModel {

    @SerializedName("age")
    public String Age;

    @SerializedName("bedno")
    public String BedNo;

    @SerializedName("patientname")
    public String Name;

    @SerializedName("admissiondate")
    public Date AdmissionDate;

    @SerializedName("dischargedate")
    public Date DischargeDate;

    @SerializedName("emergencycontactno")
    public String EmergencyContactNo;

    @SerializedName("patientregistrationno")
    public String RegistrationNo;

    @SerializedName("weight")
    public String Weight;

    @SerializedName("bloodgroup")
    public String BloodGroup;
}
