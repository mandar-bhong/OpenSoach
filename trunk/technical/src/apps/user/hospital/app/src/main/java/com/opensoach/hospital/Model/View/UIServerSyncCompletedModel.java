package com.opensoach.hospital.Model.View;

import com.opensoach.hospital.Model.AppNotificationModelBase;
import com.opensoach.hospital.Model.DB.DBLocationTableRowModel;
import com.opensoach.hospital.ViewModels.JobBriefViewModel;

import java.util.List;

/**
 * Created by Mandar on 9/11/2017.
 */

public class UIServerSyncCompletedModel extends AppNotificationModelBase{

    List<JobBriefViewModel> jobBriefViewModels;
    List<DBLocationTableRowModel> locations;

    public List<JobBriefViewModel> getJobBriefViewModels() {
        return jobBriefViewModels;
    }

    public void setJobBriefViewModels(List<JobBriefViewModel> jobBriefViewModels) {
        this.jobBriefViewModels = jobBriefViewModels;
    }

    public List<DBLocationTableRowModel> getLocations() {
        return locations;
    }

    public void setLocations(List<DBLocationTableRowModel> locations) {
        this.locations = locations;
    }
}
