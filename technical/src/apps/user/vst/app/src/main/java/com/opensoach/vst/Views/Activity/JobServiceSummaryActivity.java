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
import com.opensoach.vst.Model.DB.DBTokenTableRowModel;
import com.opensoach.vst.R;
import com.opensoach.vst.ViewModels.CreateTokenViewModel;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.ViewModels.JobServiceListViewModel;
import com.opensoach.vst.ViewModels.JobSummaryViewModel;
import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.Views.ClickHandler.JobServiceCreationHandler;
import com.opensoach.vst.Views.ClickHandler.JobServiceSummaryClickHandler;
import com.opensoach.vst.Views.Fragment.HeaderFragment;
import com.opensoach.vst.Views.Fragment.TokenItemFragment;
import com.opensoach.vst.databinding.ActivityJobServiceSummaryBinding;

import java.util.ArrayList;

import static android.support.v7.widget.LinearLayoutManager.VERTICAL;

public class JobServiceSummaryActivity extends AppCompatActivity
        implements TokenItemFragment.OnFragmentInteractionListener,
        HeaderFragment.OnFragmentInteractionListener{

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_job_service_summary);

        MainViewModel.getInstance().ContextActivity = this;

        //TODO: This step is importent for adding fragment into activity
        android.support.v4.app.FragmentManager fm = getSupportFragmentManager();
        fm.beginTransaction().replace(R.id.headerPlace, HeaderFragment.newInstance("","")).commit();

        AppRepo.getInstance().getJobServiceViewModel().getJobServiceListViewModel().setDisplayMode(ApplicationConstants.DISPLAY_MODE_JOB_CREATION_SUMMARY);

        setBinding();
    }


    void setBinding(){

        ActivityJobServiceSummaryBinding binding = DataBindingUtil.setContentView(this, R.layout.activity_job_service_summary);
        binding.setVM(AppRepo.getInstance().getJobServiceViewModel());
        binding.setClickHandler(new JobServiceCreationHandler());


        RecyclerView recyclerView = findViewById(R.id.recycler_view);
        recyclerView.setLayoutManager(new LinearLayoutManager(recyclerView.getContext()));
        recyclerView.addItemDecoration(new DividerItemDecoration(recyclerView.getContext(), VERTICAL));

    }

    @Override
    public void onFragmentInteraction(Uri uri) {

    }



    @Override
    protected  void onDestroy(){
        super.onDestroy();
        AppRepo.getInstance().getJobServiceViewModel().getJobServiceListViewModel().setDisplayMode(ApplicationConstants.DISPLAY_MODE_JOB_CREATION_EDIT);
        AppRepo.getInstance().getJobServiceViewModel().getJobServiceListViewModel().getJobServiceDataAdapter().NotifyDataSetChanged();
    }


    JobSummaryViewModel GenerateData() {
        JobSummaryViewModel jobSummaryViewModel = new JobSummaryViewModel();
        jobSummaryViewModel.ContextActivity = this;

        ArrayList<JobSummaryViewModel> list = new ArrayList<>();

        DBTokenTableRowModel dbTokenTableRowModel = new DBTokenTableRowModel();



            jobSummaryViewModel.ContextActivity = this;
            jobSummaryViewModel.Parent = jobSummaryViewModel;
            jobSummaryViewModel.setFirstName("Rahul");
            jobSummaryViewModel.setLastName("Mogal");
            jobSummaryViewModel.setMobileNo("9595441245");
            jobSummaryViewModel.setKmRuns("45Km");
            jobSummaryViewModel.setPetrolLevel("Full");
//
            list.add(jobSummaryViewModel);
            jobSummaryViewModel.setData(list);

        return jobSummaryViewModel;

    }
}
