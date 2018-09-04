package com.opensoach.vst.Views.Activity;

import android.net.Uri;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;

import com.opensoach.vst.R;
import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.Views.Fragment.HeaderFragment;
import com.opensoach.vst.Views.Fragment.TokenItemFragment;

public class JobServiceCreationActivity extends AppCompatActivity
        implements TokenItemFragment.OnFragmentInteractionListener,
        HeaderFragment.OnFragmentInteractionListener{

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_job_service_creation);

        MainViewModel.getInstance().ContextActivity = this;

        //TODO: This step is importent for adding fragment into activity
        android.support.v4.app.FragmentManager fm = getSupportFragmentManager();
        fm.beginTransaction().replace(R.id.headerPlace, HeaderFragment.newInstance("","")).commit();


    }

    @Override
    public void onFragmentInteraction(Uri uri) {

    }
}
