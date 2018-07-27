package com.opensoach.hospital.Model.View;

import com.opensoach.hospital.Model.AppNotificationModelBase;

import java.util.List;

/**
 * Created by Mandar on 9/23/2017.
 */

public class UIDeletedJobDataModel extends AppNotificationModelBase {

    List<Integer> deletedJobs;

    public List<Integer> getDeletedJobs() {
        return deletedJobs;
    }

    public void setDeletedJobs(List<Integer> deletedJobs) {
        this.deletedJobs = deletedJobs;
    }
}
