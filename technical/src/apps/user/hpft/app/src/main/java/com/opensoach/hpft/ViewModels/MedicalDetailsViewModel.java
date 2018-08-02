package com.opensoach.hpft.ViewModels;

import android.databinding.Bindable;

/**
 * Created by Mandar on 01-08-2018.
 */

public class MedicalDetailsViewModel extends BaseViewModel {

    private String treatment;
    private String allergies;
    private String history;


    @Bindable
    public String getTreatment() {
        return treatment;
    }

    @Bindable
    public void setTreatment(String treatment) {
        this.treatment = treatment;

    }

    @Bindable
    public String getAllergies() {
        return allergies;
    }

    @Bindable
    public void setAllergies(String allergies) {
        this.allergies = allergies;
    }

    @Bindable
    public String getHistory() {
        return history;
    }

    @Bindable
    public void setHistory(String history) {
        this.history = history;
    }
}
