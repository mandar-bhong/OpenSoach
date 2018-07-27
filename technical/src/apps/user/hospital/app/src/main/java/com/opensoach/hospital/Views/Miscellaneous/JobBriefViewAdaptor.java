package com.opensoach.hospital.Views.Miscellaneous;

import android.databinding.DataBindingUtil;
import android.view.View;
import android.widget.LinearLayout;

import com.opensoach.hospital.R;
import com.opensoach.hospital.ViewModels.JobBriefViewModel;
import com.opensoach.hospital.ViewModels.JobGridViewModel;
import com.opensoach.hospital.Views.ClickHandler.JobBriefClickHandler;
import com.opensoach.hospital.Views.Fragment.JobBrief;
import com.opensoach.hospital.databinding.FragmentJobBriefBinding;

/**
 * Created by Mandar on 8/25/2017.
 */

public class JobBriefViewAdaptor extends CustomBaseAdapter<JobGridViewModel,JobBriefViewModel> {

    View GetJobBrief(JobBriefViewModel dataModel,int position){

        LinearLayout ll = new LinearLayout(ContextActivity.getBaseContext());

        JobBrief jb = new JobBrief();

        FragmentJobBriefBinding fragmentJobBriefBinding = DataBindingUtil.inflate(dataModel.ContextActivity.getLayoutInflater(),
                R.layout.fragment_job_brief,ll,true);

        View v = jb.getView();
        fragmentJobBriefBinding.setData(dataModel);
        fragmentJobBriefBinding.setClickHandler(new JobBriefClickHandler());

        ll.setId(position);

        return ll;
    }


    @Override
    protected View getItemView(JobBriefViewModel dataModel, int position) {
        return GetJobBrief(dataModel,position);
    }

}
