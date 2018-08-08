package com.opensoach.hpft.Model.DB;


import com.opensoach.hpft.Constants.DBTableConstants;
import com.opensoach.hpft.DAL.DBTableSchema;

import java.util.Date;

/**
 * Created by Mandar on 07-08-2018.
 */

@DBTableSchema(TableName = DBTableConstants.TABLE_SERVICE_TASK_DATA)
public class DBServiceTaskDataTableRowModel {

    private int  servConfID;
    private int  serInID;
    private int locationId;
    private Date time;
    private String data;

    public DBServiceTaskDataTableRowModel(){
        data = "";
    }

    public int getSerInID() {
        return serInID;
    }

    public void setSerInID(int serInID) {
        this.serInID = serInID;
    }

    public int getServConfID() {
        return servConfID;
    }

    public void setServConfID(int servConfID) {
        this.servConfID = servConfID;
    }

    public int getLocationId() {
        return locationId;
    }

    public void setLocationId(int locationId) {
        this.locationId = locationId;
    }

    public Date getTime() {
        return time;
    }

    public void setTime(Date time) {
        this.time = time;
    }

    public String getData() {
        return data;
    }

    public void setData(String data) {
        this.data = data;
    }
}
