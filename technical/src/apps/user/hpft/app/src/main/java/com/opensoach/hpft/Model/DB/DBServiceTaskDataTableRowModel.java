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
    private Date entryTime;
    private String data;
    private boolean isSynced;
    private String authCode;
    private Date slotStartTime;
    private Date slotEndTime;

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

    public Date getEntryTime() {
        return entryTime;
    }

    public void setEntryTime(Date entryTime) {
        this.entryTime = entryTime;
    }

    public String getData() {
        return data;
    }

    public void setData(String data) {
        this.data = data;
    }

    public boolean isSynced() {
        return isSynced;
    }

    public void setSynced(boolean synced) {
        isSynced = synced;
    }

    public String getAuthCode() {
        return authCode;
    }

    public void setAuthCode(String authCode) {
        this.authCode = authCode;
    }

    public Date getSlotStartTime() {
        return slotStartTime;
    }

    public void setSlotStartTime(Date slotStartTime) {
        this.slotStartTime = slotStartTime;
    }

    public Date getSlotEndTime() {
        return slotEndTime;
    }

    public void setSlotEndTime(Date slotEndTime) {
        this.slotEndTime = slotEndTime;
    }
}
