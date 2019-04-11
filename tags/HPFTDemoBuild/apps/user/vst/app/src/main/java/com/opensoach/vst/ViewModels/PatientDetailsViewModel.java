package com.opensoach.vst.ViewModels;

import android.databinding.Bindable;

import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.Model.Communication.PacketPatientDetailsModel;

import java.text.SimpleDateFormat;

/**
 * Created by Mandar on 01-08-2018.
 */

public class PatientDetailsViewModel extends BaseViewModel {

    private PacketPatientDetailsModel packetPatientDetailsModel;

    public PatientDetailsViewModel(PacketPatientDetailsModel packetPatientDetailsModel) {
        this.packetPatientDetailsModel = packetPatientDetailsModel;
    }

    @Bindable
    public String getName() {
        return packetPatientDetailsModel.Name;
    }

    @Bindable
    public String getAge() {
        return packetPatientDetailsModel.Age + " Years";
    }

    @Bindable
    public String getEmergencyContactNo() {
        return packetPatientDetailsModel.EmergencyContactNo;
    }


    @Bindable
    public String getRegNo() {
        return packetPatientDetailsModel.RegistrationNo;
    }

    @Bindable
    public String getRoomNo() {
        return packetPatientDetailsModel.BedNo;
    }

    @Bindable
    public String getAdmissionDateFormatted() {

        SimpleDateFormat dateFormatter = new SimpleDateFormat(ApplicationConstants.UI_DATE_FORMAT);

        if (packetPatientDetailsModel.AdmissionDate !=null) {
            return dateFormatter.format(packetPatientDetailsModel.AdmissionDate);
        }else{
            return "NA";
        }
    }

    @Bindable
    public String getWeight() {
        return packetPatientDetailsModel.Weight + " Kgs";
    }

    @Bindable
    public String getBloodGroup() {
        return packetPatientDetailsModel.BloodGroup;
    }

}
