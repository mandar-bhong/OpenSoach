package spl.hkt.opensoach.splapp.model.view;

import java.util.Date;

/**
 * Created by Mandar on 3/27/2017.
 */

public class ChartConfigSlotModel {
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
