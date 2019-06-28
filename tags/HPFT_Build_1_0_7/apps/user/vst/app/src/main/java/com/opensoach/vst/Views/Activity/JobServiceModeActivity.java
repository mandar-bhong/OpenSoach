package com.opensoach.vst.Views.Activity;

import android.content.Context;
import android.content.Intent;
import android.databinding.DataBindingUtil;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.widget.ImageButton;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.Constants.Constants;
import com.opensoach.vst.Helper.AppHelper;
import com.opensoach.vst.R;
import com.opensoach.vst.Scheduler.ScheduleManager;
import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.Views.ClickHandler.JobServiceModeClickHandler;
import com.opensoach.vst.databinding.ActivityJobServiceModeBinding;

import java.util.zip.Inflater;

public class JobServiceModeActivity extends AppCompatActivity {


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_job_service_mode);

        setBinding();
    }

    void setBinding(){
        ActivityJobServiceModeBinding binding = DataBindingUtil.setContentView(this, R.layout.activity_job_service_mode);
        binding.setClickHandler(new JobServiceModeClickHandler());
    }


    @Override
    public void onResume() {
        super.onResume();

        AppHelper.DeInit();
    }
}
