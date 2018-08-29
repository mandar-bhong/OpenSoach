package com.opensoach.vst.ViewModels;

import android.databinding.Bindable;

public class CreateTokenViewModel extends BaseViewModel  {

    public String vehicleNuber;
    public String vehicleNo1;
    public String vehicleNo2;
    public String vehicleNo3;


    public String getVehicleNuber() {
        return vehicleNuber;
    }

    public String getVehicleNo1() {
        return vehicleNo1;
    }

    @Bindable
    public void setVehicleNo1(String vehicleNo1) {
        this.vehicleNo1 = vehicleNo1;
    }


    public String getVehicleNo2() {
        return vehicleNo2;
    }

    public void setVehicleNo2(String vehicleNo2) {
        this.vehicleNo2 = vehicleNo2;
    }

    public String getVehicleNo3() {
        return vehicleNo3;
    }

    public void setVehicleNo3(String vehicleNo3) {
        this.vehicleNo3 = vehicleNo3;
    }
}
