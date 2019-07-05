package com.opensoach.vst.Views.Activity;

import android.databinding.DataBindingUtil;
import android.graphics.Color;
import android.net.Uri;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.text.Editable;
import android.text.TextUtils;
import android.text.TextWatcher;
import android.view.View;
import android.view.WindowManager;
import android.widget.Button;
import android.widget.EditText;
import android.widget.SeekBar;
import android.widget.TextView;
import android.widget.Toast;


import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.R;

import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.Views.ClickHandler.JobServiceCreationHandler;
import com.opensoach.vst.Views.Fragment.HeaderFragment;
import com.opensoach.vst.Views.Fragment.TokenItemFragment;
import com.opensoach.vst.databinding.ActivityJobServiceDetailsBinding;


public class JobServiceDetailsActivity extends AppCompatActivity
        implements TokenItemFragment.OnFragmentInteractionListener,
        HeaderFragment.OnFragmentInteractionListener {


    private SeekBar seekBar;
    private TextView textView;
    EditText firstName, lastName, mobileNo, kmRuns;


    protected void onCreate(Bundle savedInstanceState) {

        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_job_service_details);
        MainViewModel.getInstance().ContextActivity = this;

        firstName = (EditText) findViewById(R.id.firstName);
        lastName = (EditText) findViewById(R.id.lastName);
        mobileNo = (EditText) findViewById(R.id.mobileNo);
        kmRuns = (EditText) findViewById(R.id.kmRuns);



        //      TODO: This step is importent for adding fragment into activity
        android.support.v4.app.FragmentManager fm = getSupportFragmentManager();
        fm.beginTransaction().replace(R.id.headerPlace, HeaderFragment.newInstance("", "")).commit();

        setBinding();
        getWindow().setSoftInputMode(WindowManager.LayoutParams.SOFT_INPUT_ADJUST_PAN);

        initializeVariables();
        // Initialize the textview with '0'.
        textView.setText("Petrol Level: " + seekBar.getProgress() + "/" + seekBar.getMax());
        seekBar.setOnSeekBarChangeListener(new SeekBar.OnSeekBarChangeListener() {
            int progress = 0;
            @Override
            public void onProgressChanged(SeekBar seekBar, int progresValue, boolean fromUser) {
                progress = progresValue;
                AppRepo.getInstance().getJobServiceViewModel().getJobServiceDetailsViewModel().setPetrolLevel(String.valueOf(progresValue) );
           }
            @Override
            public void onStartTrackingTouch(SeekBar seekBar) {
            }
            @Override
            public void onStopTrackingTouch(SeekBar seekBar) {
                textView.setText("Petrol Level: " + progress + "%" + "/" + seekBar.getMax()+"%");
            }
        });
    }

   // A private method to help us initialize our variables.
    private void initializeVariables() {
        seekBar = (SeekBar) findViewById(R.id.seekBar1);
        textView = (TextView) findViewById(R.id.textView1);
    }

    void setBinding(){
        ActivityJobServiceDetailsBinding binding = DataBindingUtil.setContentView(this, R.layout.activity_job_service_details);
        binding.setVM(AppRepo.getInstance().getJobServiceViewModel().getJobServiceDetailsViewModel());
        binding.setClickHandler(new JobServiceCreationHandler());
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

//    public boolean Validate(){
//
//        boolean hasError = false;
//        if(firstName.getText().toString().length()==0){
//            firstName.setError("First name not entered");
//            firstName.requestFocus();
//            hasError = true;
//        }
//        if(lastName.getText().toString().length() == 0){
//            lastName.setError("Last name not enterd");
//            lastName.requestFocus();
//            hasError = true;
//        }
//
//
//        return !hasError;
//    }

}
