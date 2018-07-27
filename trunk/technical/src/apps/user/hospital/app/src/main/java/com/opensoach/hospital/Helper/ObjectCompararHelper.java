package com.opensoach.hospital.Helper;

import com.opensoach.hospital.ViewModels.JobBriefViewModel;

/**
 * Created by Mandar on 9/23/2017.
 */

public class ObjectCompararHelper {

    public static boolean Equal(JobBriefViewModel jobBriefViewModel, int jobID){

        if(jobBriefViewModel.getDbJobCardTableRowModel().getJobCardId() == jobID){
            return true;
        }

        return false;
    }
}
