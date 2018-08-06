package com.opensoach.hpft.ViewModels;

import android.databinding.Bindable;

import com.opensoach.hpft.Model.Communication.PacketMedicalDetailsModel;

/**
 * Created by Mandar on 01-08-2018.
 */

public class MedicalDetailsViewModel extends BaseViewModel {

    private PacketMedicalDetailsModel packetMedicalDetailsModel;

    public MedicalDetailsViewModel(PacketMedicalDetailsModel packetMedicalDetailsModel) {
        this.packetMedicalDetailsModel = packetMedicalDetailsModel;
    }

    @Bindable
    public String getTreatment() {
        return packetMedicalDetailsModel.Treatment;
    }

    @Bindable
    public String getAllergies() {
        return packetMedicalDetailsModel.Allergies;
    }

    @Bindable
    public String getHistory() {
        return packetMedicalDetailsModel.MedicalHistory;
    }


}
