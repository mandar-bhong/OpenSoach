package com.opensoach.hpft.ViewModels;

import android.databinding.BaseObservable;
import android.databinding.Bindable;

import com.opensoach.hpft.BR;
import com.opensoach.hpft.Constants.ApplicationConstants;
import com.opensoach.hpft.Constants.Constants;

import java.text.SimpleDateFormat;
import java.util.Date;

/**
 * Created by Mandar on 01-08-2018.
 */

public class PatientDetailsViewModel extends BaseViewModel {

    private String name;
    private Integer age;
    private  String emergencyContactNo;
    private String regNo;
    private String roomNo;
    private Date admissionDate;


    @Bindable
    public String getName() {
        return name;
    }

    @Bindable
    public void setName(String name) {
        this.name = name;
        notifyPropertyChanged(BR.name);
    }

    @Bindable
    public Integer getAge() {
        return age;
    }

    @Bindable
    public void setAge(Integer age) {
        this.age = age;
        notifyPropertyChanged(BR.age);
    }

    @Bindable
    public String getEmergencyContactNo() {
        return emergencyContactNo;
    }

    @Bindable
    public void setEmergencyContactNo(String emergencyContactNo) {
        this.emergencyContactNo = emergencyContactNo;
        notifyPropertyChanged(BR.emergencyContactNo);
    }

    @Bindable
    public String getRegNo() {
        return regNo;
    }

    @Bindable
    public void setRegNo(String regNo) {
        this.regNo = regNo;
        notifyPropertyChanged(BR.regNo);
    }

    @Bindable
    public String getRoomNo() {
        return roomNo;
    }

    @Bindable
    public void setRoomNo(String roomNo) {
        this.roomNo = roomNo;
        notifyPropertyChanged(BR.roomNo);
    }

    @Bindable
    public String getAdmissionDateFormatted() {

        SimpleDateFormat dateFormatter = new SimpleDateFormat(ApplicationConstants.UI_DATE_FORMAT);

        if (admissionDate !=null) {
            return dateFormatter.format(admissionDate);
        }else{
            return "NA";
        }
    }

    public void setAdmissionDate(Date admissionDate) {
        this.admissionDate = admissionDate;
        notifyPropertyChanged(BR.admissionDateFormatted);
    }


}
