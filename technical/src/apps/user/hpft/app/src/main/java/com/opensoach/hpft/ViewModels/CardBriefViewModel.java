package com.opensoach.hpft.ViewModels;

import com.opensoach.hpft.Views.Fragment.MedicalDetailsFragment;

/**
 * Created by Mandar on 30-07-2018.
 */

public class CardBriefViewModel extends BaseViewModel {

    public int SerInID;
    public int ServConfID;
    public int locationID;

    private PatientDetailsViewModel patientDetails;
    private MedicalDetailsViewModel medicalDetails;
    private TaskDetailsViewModel taskDetails;


    public PatientDetailsViewModel getPatientDetails() {
        return patientDetails;
    }

    public void setPatientDetails(PatientDetailsViewModel patientDetails) {
        this.patientDetails = patientDetails;
        this.patientDetails.ContextActivity = ContextActivity;
    }

    public MedicalDetailsViewModel getMedicalDetails() {
        return medicalDetails;
    }

    public void setMedicalDetails(MedicalDetailsViewModel medicalDetails) {
        this.medicalDetails = medicalDetails;
        this.medicalDetails.ContextActivity = ContextActivity;
    }

    public TaskDetailsViewModel getTaskDetails() {
        return taskDetails;
    }

    public void setTaskDetails(TaskDetailsViewModel taskDetails) {
        this.taskDetails = taskDetails;
        this.taskDetails.ContextActivity = ContextActivity;
    }

    public int getSerInID() {
        return SerInID;
    }

    public void setSerInID(int serInID) {
        SerInID = serInID;
    }

    public int getServConfID() {
        return ServConfID;
    }

    public void setServConfID(int servConfID) {
        ServConfID = servConfID;
    }

    public int getLocationID() {
        return locationID;
    }

    public void setLocationID(int locationID) {
        this.locationID = locationID;
    }
}
