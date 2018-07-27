package com.opensoach.hospital.Views.Fragment;

import android.content.Context;
import android.util.AttributeSet;

import com.opensoach.hospital.ViewModels.JobBriefViewModel;
import com.opensoach.hospital.ViewModels.JobGridViewModel;
import com.opensoach.hospital.Views.Interfaces.IFragment;
import com.opensoach.hospital.Views.Interfaces.IList;
import com.opensoach.hospital.Views.Miscellaneous.JobBriefViewAdaptor;

import java.util.List;

/**
 * Created by Mandar on 8/25/2017.
 */

public class JobGridView extends CustomGridView implements IFragment<JobGridViewModel>,IList<JobBriefViewModel> {

public JobGridViewModel DataContext;
    private JobBriefViewAdaptor jobBriefViewAdaptor;


    public JobGridView(Context context){
        super(context);
    }

    public JobGridView(Context context, AttributeSet attrs) {
        super(context, attrs);
    }

    public JobGridView(Context context, AttributeSet attrs, int defStyleAttr) {
        super(context, attrs, defStyleAttr);
    }

    @Override
    protected void setDataAdapter() {

        //setAdapter(jobBriefViewAdaptor);
    }


    @Override
    public JobGridViewModel getDataContext() {
        return DataContext;
    }

    @Override
    public void setDataContext(JobGridViewModel viewModel) {
		//TODO: Check if null condition is required
        if(viewModel == null)return;
        DataContext = viewModel;
        jobBriefViewAdaptor = (JobBriefViewAdaptor) viewModel.getDataAdaptor();
        jobBriefViewAdaptor.ContextActivity = viewModel.ContextActivity;
        setAdapter(jobBriefViewAdaptor);
        jobBriefViewAdaptor.GridViewContainer = this;
    }

    @Override
    public List<JobBriefViewModel> getItemsSource() {
        return jobBriefViewAdaptor.getItemsSource();
    }

    @Override
    public void setItemsSource(List<JobBriefViewModel> source) {
        jobBriefViewAdaptor.setItemsSource(source);
    }

}
