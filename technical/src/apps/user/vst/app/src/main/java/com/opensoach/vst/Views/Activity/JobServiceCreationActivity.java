package com.opensoach.vst.Views.Activity;

import android.databinding.DataBindingUtil;
import android.net.Uri;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.WindowManager;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.R;

import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.Views.ClickHandler.JobDetailsCompleteClickHandler;
import com.opensoach.vst.Views.Fragment.HeaderFragment;
import com.opensoach.vst.Views.Fragment.TokenItemFragment;
import com.opensoach.vst.databinding.ActivityJobServiceCreationBinding;

public class JobServiceCreationActivity extends AppCompatActivity
        implements TokenItemFragment.OnFragmentInteractionListener,
        HeaderFragment.OnFragmentInteractionListener{

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_job_service_creation);


        MainViewModel.getInstance().ContextActivity = this;

//
////        TODO: This step is importent for adding fragment into activity
        android.support.v4.app.FragmentManager fm = getSupportFragmentManager();
        fm.beginTransaction().replace(R.id.headerPlace, HeaderFragment.newInstance("","")).commit();


        setBinding();

        getWindow().setSoftInputMode(WindowManager.LayoutParams.SOFT_INPUT_ADJUST_PAN);
    }


    void setBinding(){
        ActivityJobServiceCreationBinding binding = DataBindingUtil.setContentView(this, R.layout.activity_job_service_creation);

//
//        JobDetailsViewModel jobDetailsViewModel = new JobDetailsViewModel();
//        jobDetailsViewModel.Parent = AppRepo.getInstance().getJobDetailsViewModel().getJobCustomerDetailsViewModel();
//        jobDetailsViewModel.ContextActivity = AppRepo.getInstance().getJobDetailsViewModel().ContextActivity;

        binding.setVM(AppRepo.getInstance().getJobServiceViewModel().getJobServiceDetailsViewModel());
        binding.setClickHandler(new JobDetailsCompleteClickHandler());

    }

    @Override
    public void onFragmentInteraction(Uri uri) {

    }

    @Override
    public void onResume() {
        super.onResume();

        if ((boolean)AppRepo.getInstance().getStore().get(ApplicationConstants.APP_STORE_JOB_SUBMITTED)){
            AppRepo.getInstance().getStore().put(ApplicationConstants.APP_STORE_JOB_SUBMITTED,false);
            this.finish();
        }
    }

}
