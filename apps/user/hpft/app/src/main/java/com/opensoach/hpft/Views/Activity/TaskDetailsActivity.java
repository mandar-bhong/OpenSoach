package com.opensoach.hpft.Views.Activity;

import android.databinding.DataBindingUtil;
import android.net.Uri;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;

import com.opensoach.hpft.AppRepo.AppRepo;
import com.opensoach.hpft.Model.View.TaskItemDataModel;
import com.opensoach.hpft.R;
import com.opensoach.hpft.Utility.AppLogger;
import com.opensoach.hpft.ViewModels.TaskItemViewModel;
import com.opensoach.hpft.Views.Fragment.HeaderFragment;
import com.opensoach.hpft.databinding.ActivityTaskDetailsBinding;

public class TaskDetailsActivity extends AppCompatActivity implements HeaderFragment.OnFragmentInteractionListener {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        try {
            super.onCreate(savedInstanceState);
            setContentView(R.layout.activity_task_details);

            //TODO: This step is importent for adding fragment into activity
            android.support.v4.app.FragmentManager fm = getSupportFragmentManager();
            fm.beginTransaction().replace(R.id.headerPlace, HeaderFragment.newInstance("", "")).commit();

            ActivityTaskDetailsBinding binding = DataBindingUtil.setContentView(this, R.layout.activity_task_details);
            binding.setVM(AppRepo.getInstance().getActiveTaskItem());
            binding.setBriefView(AppRepo.getInstance().getActiveCard());
        }catch (Exception ex){
            AppLogger.getInstance().Log(ex);
        }
    }

    @Override
    public void onFragmentInteraction(Uri uri){

    }

    public void ok(View view) {
        this.finish();
    }
}
