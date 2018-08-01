package com.opensoach.hpft.ViewModels;

/**
 * Created by Mandar on 01-08-2018.
 */

public class MedicalDetailsViewModel extends BaseViewModel {

    private String treatment;
    private String allergies;
    private String history;


    public String getTreatment() {
        return treatment;
    }

    public void setTreatment(String treatment) {
        this.treatment = treatment;
    }

    public String getAllergies() {
        return allergies;
    }

    public void setAllergies(String allergies) {
        this.allergies = allergies;
    }

    public String getHistory() {
        return history;
    }

    public void setHistory(String history) {
        this.history = history;
    }
}
