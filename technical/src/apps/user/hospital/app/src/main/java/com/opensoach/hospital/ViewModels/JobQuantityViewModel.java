package com.opensoach.hospital.ViewModels;

/**
 * Created by Mandar on 29-10-2017.
 */

public class JobQuantityViewModel extends BaseViewModel {

    private int jobID;
    private String operatorCode;
    private String finishedQuantity;
    private String comment;

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

    public String getFinishedQuantity() {
        return finishedQuantity;
    }

    public void setFinishedQuantity(String finishedQuantity) {
        this.finishedQuantity = finishedQuantity;
    }

    public String getComment() {
        if(comment == null)
            return "";
        return comment;
    }

    public void setComment(String comment) {
        this.comment = comment;
    }


    public int getQuantity() {

        if (finishedQuantity == null || finishedQuantity == "")
            return 0;
        else
            return Integer.parseInt(finishedQuantity);
    }
}
