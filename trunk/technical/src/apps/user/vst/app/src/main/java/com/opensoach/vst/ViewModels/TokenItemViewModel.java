package com.opensoach.vst.ViewModels;

import com.opensoach.vst.Model.DB.DBTokenTableRowModel;
import com.opensoach.vst.Model.View.TaskTimeItemDataModel;

public class TokenItemViewModel extends  BaseViewModel {
    private DBTokenTableRowModel dbTokenTableRowModel;


    public  TokenItemViewModel(DBTokenTableRowModel item){
        dbTokenTableRowModel = item;
    }


    public String getTokenNo(){
        return Integer.toString(  dbTokenTableRowModel.getTokenno());
    }

    public String getVehicleNo(){
        return dbTokenTableRowModel.getVehicleno();
    }

    public String getTime(){
        return dbTokenTableRowModel.getGeneratedon().toString();
    }



}
