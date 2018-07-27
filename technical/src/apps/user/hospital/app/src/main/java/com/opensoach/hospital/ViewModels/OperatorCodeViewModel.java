package com.opensoach.hospital.ViewModels;

/**
 * Created by Mandar on 26-11-2017.
 */

public class OperatorCodeViewModel extends BaseViewModel  {

    private int jobID;
    private String operatorCode;

    public int getJobID() {
        return jobID;
    }

    public void setJobID(int jobID) {
        this.jobID = jobID;
    }

    public String getOperatorCode() {
        return operatorCode;
    }

    public void setOperatorCode(String operatorCode) {
        this.operatorCode = operatorCode;
    }
}
