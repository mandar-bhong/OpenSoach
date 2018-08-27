package com.opensoach.vst.Model.DB;

import com.opensoach.vst.Constants.DBTableConstants;
import com.opensoach.vst.DAL.DBTableSchema;

import java.util.Date;

/**
 * Created by Mandar on 07-08-2018.
 */
@DBTableSchema(TableName = DBTableConstants.TABLE_TASK_DATA)
public class DBTaskDataTableRowModel {

    private int  locationId;
    private int  serInID;
    private String title;
    private Date slotStartTime;
    private Date slotEndTime;
    private Date entrytime;
    private boolean isSynced;
    private String authCode;
    private String  value;
    private String  comment;


    public int getLocationId() {
        return locationId;
    }

    public void setLocationId(int locationID) {
        this.locationId = locationID;
    }

    public int getSerInID() {
        return serInID;
    }

    public void setSerInID(int serInID) {
        this.serInID = serInID;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public Date getTaskSlotStartTime() {
        return slotStartTime;
    }

    public void setTaskSlotStartTime(Date taskSlotStartTime) {
        this.slotStartTime = taskSlotStartTime;
    }

    public Date getTaskSlotEndTime() {
        return slotEndTime;
    }

    public void setTaskSlotEndTime(Date taskSlotEndTime) {
        this.slotEndTime = taskSlotEndTime;
    }

    public Date getTaskTime() {
        return entrytime;
    }

    public void setTaskTime(Date taskTime) {
        this.entrytime = taskTime;
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

    public String getValue() {
        return value;
    }

    public void setValue(String value) {
        this.value = value;
    }

    public String getComment() {
        return comment;
    }

    public void setComment(String comment) {
        this.comment = comment;
    }
}
