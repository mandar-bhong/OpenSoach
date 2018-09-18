package com.opensoach.vst.Views.Activity;

import android.databinding.DataBindingUtil;
import android.net.Uri;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.support.v7.widget.DividerItemDecoration;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.R;
import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.Views.ClickHandler.JobServiceCreationHandler;
import com.opensoach.vst.Views.Fragment.HeaderFragment;
import com.opensoach.vst.Views.Fragment.TokenItemFragment;
import com.opensoach.vst.databinding.ActivityJobServiceTaskListBinding;

import static android.support.v7.widget.LinearLayoutManager.VERTICAL;

public class JobServiceTaskListActivity extends AppCompatActivity
        implements TokenItemFragment.OnFragmentInteractionListener,
        HeaderFragment.OnFragmentInteractionListener{

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_job_service_task_list);

        MainViewModel.getInstance().ContextActivity = this;

        //TODO: This step is importent for adding fragment into activity
        android.support.v4.app.FragmentManager fm = getSupportFragmentManager();
        fm.beginTransaction().replace(R.id.headerPlace, HeaderFragment.newInstance("","")).commit();

        setBinding();
    }

    void setBinding(){
        ActivityJobServiceTaskListBinding binding = DataBindingUtil.setContentView(this, R.layout.activity_job_service_task_list);
        binding.setVM(AppRepo.getInstance().getJobServiceViewModel().getJobServiceListViewModel());
        binding.setClickHandler(new JobServiceCreationHandler());


        RecyclerView recyclerView = findViewById(R.id.recycler_view);
        recyclerView.setLayoutManager(new LinearLayoutManager(recyclerView.getContext()));
        recyclerView.addItemDecoration(new DividerItemDecoration(recyclerView.getContext(), VERTICAL));

    }

    @Override
    public void onFragmentInteraction(Uri uri) {

    }

    @Override
    public void onResume(){
        super.onResume();

        if ((boolean)AppRepo.getInstance().getStore().get(ApplicationConstants.APP_STORE_JOB_SUBMITTED)){
            this.finish();
        }

        AppRepo.getInstance().getJobServiceViewModel().getJobServiceListViewModel().setDisplayMode(ApplicationConstants.DISPLAY_MODE_JOB_CREATION_EDIT);

    }


}
