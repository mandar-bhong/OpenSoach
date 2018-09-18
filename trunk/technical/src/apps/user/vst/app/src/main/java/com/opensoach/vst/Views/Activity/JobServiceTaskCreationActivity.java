package com.opensoach.vst.Views.Activity;

import android.databinding.DataBindingUtil;
import android.net.Uri;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.widget.EditText;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.R;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.Views.ClickHandler.TaskCreateCompleteClickHandler;
import com.opensoach.vst.Views.Fragment.HeaderFragment;
import com.opensoach.vst.Views.Fragment.TokenItemFragment;
import com.opensoach.vst.databinding.ActivityJobServiceTaskCreationBinding;
import com.opensoach.vst.databinding.ActivityJobServiceTaskCreationBinding;

public class JobServiceTaskCreationActivity extends AppCompatActivity
        implements TokenItemFragment.OnFragmentInteractionListener,
        HeaderFragment.OnFragmentInteractionListener{

     private EditText etTaskname;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_job_service_task_creation);

        MainViewModel.getInstance().ContextActivity = this;

        //TODO: This step is importent for adding fragment into activity
        android.support.v4.app.FragmentManager fm = getSupportFragmentManager();
        fm.beginTransaction().replace(R.id.headerPlace, HeaderFragment.newInstance("","")).commit();

        setBinding();
    }



    void setBinding(){
        ActivityJobServiceTaskCreationBinding binding = DataBindingUtil.setContentView(this, R.layout.activity_job_service_task_creation);

        JobServiceItemViewModel jobServiceItemViewModel = new JobServiceItemViewModel();
        jobServiceItemViewModel.Parent = AppRepo.getInstance().getJobServiceViewModel().getJobServiceListViewModel();
        jobServiceItemViewModel.ContextActivity = AppRepo.getInstance().getJobServiceViewModel().ContextActivity;

        binding.setVM(jobServiceItemViewModel);
        binding.setClickHandler(new TaskCreateCompleteClickHandler());
    }

    @Override
    public void onFragmentInteraction(Uri uri) {

    }

    @Override
    public void onResume() {
        super.onResume();

        AppRepo.getInstance().getJobServiceViewModel().getJobServiceListViewModel().getJobServiceDataAdapter().notifyDataSetChanged();

        if ((boolean)AppRepo.getInstance().getStore().get(ApplicationConstants.APP_STORE_JOB_SUBMITTED)){
            this.finish();
        }
    }

}
