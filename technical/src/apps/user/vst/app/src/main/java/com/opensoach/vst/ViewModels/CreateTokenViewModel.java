package com.opensoach.vst.ViewModels;

import android.databinding.Bindable;

import com.opensoach.vst.BR;
import com.opensoach.vst.Model.DB.DBTokenTableRowModel;

import java.text.Format;
import java.text.SimpleDateFormat;

public class CreateTokenViewModel extends BaseViewModel  {


    public String vehicleNo1;
    public boolean generateTokenVisible;

    public DBTokenTableRowModel dbTokenTableRowModel;


    public CreateTokenViewModel(){
        generateTokenVisible = true;
    }


    public DBTokenTableRowModel getDbTokenTableRowModel() {
        return dbTokenTableRowModel;
    }

    public void setDbTokenTableRowModel(DBTokenTableRowModel dbTokenTableRowModel) {
        this.dbTokenTableRowModel = dbTokenTableRowModel;
        notifyPropertyChanged(BR.generatedOn);
        notifyPropertyChanged(BR.generatedToken);
        notifyPropertyChanged(BR.newVehicleNumber);
    }

    public String getVehicleNumber() {
        return  vehicleNo1 ;
    }

    public String getVehicleNo1() {
        return vehicleNo1;
    }

    @Bindable
    public void setVehicleNo1(String vehicleNo1) {
        this.vehicleNo1 = vehicleNo1;
    }





    @Bindable
    public boolean isGenerateTokenVisible() {
        return generateTokenVisible;
    }

    @Bindable
    public void setGenerateTokenVisible(boolean generateTokenVisibility) {
        this.generateTokenVisible = generateTokenVisibility;
        notifyPropertyChanged(BR.generateTokenVisible);
    }

    @Bindable
    public String getGeneratedToken(){

        if (dbTokenTableRowModel == null){
            return "";
        }
        return Integer.toString( dbTokenTableRowModel.getTokenno());
    }

    @Bindable
    public String getGeneratedOn(){

        if (dbTokenTableRowModel == null){
            return "";
        }

        Format formatter = new SimpleDateFormat("hh:mm a");
        String time = formatter.format(dbTokenTableRowModel.getGeneratedon());
        return time;

    }

    @Bindable
    public String getNewVehicleNumber(){
        if (dbTokenTableRowModel == null){
            return "";
        }

        return dbTokenTableRowModel.getVehicleno();
    }

}
