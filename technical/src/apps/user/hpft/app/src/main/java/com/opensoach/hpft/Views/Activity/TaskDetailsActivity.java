package com.opensoach.hpft.Views.Activity;

import android.net.Uri;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;

import com.opensoach.hpft.R;
import com.opensoach.hpft.Views.Fragment.HeaderFragment;

public class TaskDetailsActivity extends AppCompatActivity implements HeaderFragment.OnFragmentInteractionListener {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_task_details);

        //TODO: This step is importent for adding fragment into activity
        android.support.v4.app.FragmentManager fm = getSupportFragmentManager();
        fm.beginTransaction().replace(R.id.headerPlace, HeaderFragment.newInstance("","")).commit();
    }

    @Override
    public void onFragmentInteraction(Uri uri){

    }
}
