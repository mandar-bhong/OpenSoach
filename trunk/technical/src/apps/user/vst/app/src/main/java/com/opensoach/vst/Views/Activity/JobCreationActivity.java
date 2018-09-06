package com.opensoach.vst.Views.Activity;

import android.databinding.DataBindingUtil;
import android.net.Uri;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.support.v7.widget.DividerItemDecoration;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.R;
import com.opensoach.vst.ViewModels.JobDetailsViewModel;

import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.Views.ClickHandler.GenerateTokenClickHandler;
import com.opensoach.vst.Views.ClickHandler.JobDetailsCompleteClickHandler;
import com.opensoach.vst.Views.Fragment.HeaderFragment;
import com.opensoach.vst.Views.Fragment.TokenItemFragment;
import com.opensoach.vst.databinding.ActivityJobCreationBinding;
import com.opensoach.vst.databinding.ActivityTokenListBinding;

import static android.support.v7.widget.LinearLayoutManager.VERTICAL;

public class JobCreationActivity extends AppCompatActivity
        implements TokenItemFragment.OnFragmentInteractionListener,
        HeaderFragment.OnFragmentInteractionListener{

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_job_creation);


        MainViewModel.getInstance().ContextActivity = this;

//
////        TODO: This step is importent for adding fragment into activity
        android.support.v4.app.FragmentManager fm = getSupportFragmentManager();
        fm.beginTransaction().replace(R.id.headerPlace, HeaderFragment.newInstance("","")).commit();


        setBinding();

    }


    void setBinding(){
        ActivityJobCreationBinding binding = DataBindingUtil.setContentView(this, R.layout.activity_job_creation);
        binding.setVM(new JobDetailsViewModel());
        binding.setClickHandler(new JobDetailsCompleteClickHandler());

//        binding.setVM(AppRepo.getInstance().getJobServiceViewModel());
    }

    @Override
    public void onFragmentInteraction(Uri uri) {

    }
}
