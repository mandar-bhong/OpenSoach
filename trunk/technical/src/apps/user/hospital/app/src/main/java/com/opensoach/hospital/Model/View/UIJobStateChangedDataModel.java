package com.opensoach.hospital.Model.View;

import com.opensoach.hospital.Model.AppNotificationModelBase;
import com.opensoach.hospital.ViewModels.JobBriefViewModel;

import java.util.List;

/**
 * Created by Mandar on 9/24/2017.
 */

public class UIJobStateChangedDataModel extends AppNotificationModelBase {

    List<JobBriefViewModel> jobBriefViewModels;

    public List<JobBriefViewModel> getJobBriefViewModels() {
        return jobBriefViewModels;
    }

    public void setJobBriefViewModels(List<JobBriefViewModel> jobBriefViewModels) {
        this.jobBriefViewModels = jobBriefViewModels;
    }

}
