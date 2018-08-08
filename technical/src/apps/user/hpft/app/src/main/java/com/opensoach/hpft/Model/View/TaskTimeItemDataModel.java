package com.opensoach.hpft.Model.View;

import java.util.Date;

/**
 * Created by Mandar on 07-08-2018.
 */

public class TaskTimeItemDataModel {

    private int index;
    private Date startTime;
    private Date endTime;

    public int getIndex() {
        return index;
    }

    public void setIndex(int index) {
        this.index = index;
    }

    public Date getStartTime() {
        return startTime;
    }

    public void setStartTime(Date startTime) {
        this.startTime = startTime;
    }

    public Date getEndTime() {
        return endTime;
    }

    public void setEndTime(Date endTime) {
        this.endTime = endTime;
    }
}
