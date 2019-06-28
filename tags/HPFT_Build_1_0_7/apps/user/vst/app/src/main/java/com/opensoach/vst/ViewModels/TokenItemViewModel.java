package com.opensoach.vst.ViewModels;

import android.databinding.Bindable;

import com.opensoach.vst.BR;
import com.opensoach.vst.Model.DB.DBTokenTableRowModel;
import com.opensoach.vst.Model.View.TaskTimeItemDataModel;

import java.text.Format;
import java.text.SimpleDateFormat;

public class TokenItemViewModel extends  BaseViewModel {
    private DBTokenTableRowModel dbTokenTableRowModel;
    public boolean isItemSelected;
    public int position;


    public  TokenItemViewModel(DBTokenTableRowModel item){
        dbTokenTableRowModel = item;
        isItemSelected = false;
    }


    public int getPosition() {
        return position;
    }

    public void setPosition(int position) {
        this.position = position;
    }

    public boolean isItemSelected() {
        return isItemSelected;
    }

    @Bindable
    public void setItemSelected(boolean itemSelected) {
        isItemSelected = itemSelected;
        notifyPropertyChanged(BR.itemSelected);
    }



    public DBTokenTableRowModel getDbTokenTableRowModel() {
        return dbTokenTableRowModel;
    }

    public void setDbTokenTableRowModel(DBTokenTableRowModel dbTokenTableRowModel) {
        this.dbTokenTableRowModel = dbTokenTableRowModel;
    }

    public String getTokenNo(){
        return Integer.toString(  dbTokenTableRowModel.getTokenno());
    }

    public String getVehicleNo(){
        return dbTokenTableRowModel.getVehicleno();
    }

    public String getTime(){

        Format formatter = new SimpleDateFormat("hh:mm a");
        String time = formatter.format(dbTokenTableRowModel.getGeneratedon());
        return time;

    }


}
