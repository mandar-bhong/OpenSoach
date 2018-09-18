package com.opensoach.vst.Views.Activity;

import android.databinding.DataBindingUtil;
import android.net.Uri;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;

import com.opensoach.vst.R;
import com.opensoach.vst.ViewModels.CreateTokenViewModel;
import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.Views.ClickHandler.JobServiceTokenCreationHandler;
import com.opensoach.vst.Views.Fragment.HeaderFragment;
import com.opensoach.vst.databinding.ActivityCreateTokenBinding;

public class CreateTokenActivity extends AppCompatActivity
        implements HeaderFragment.OnFragmentInteractionListener{

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_create_token);

        MainViewModel.getInstance().ContextActivity = this;

        //TODO: This step is importent for adding fragment into activity
        android.support.v4.app.FragmentManager fm = getSupportFragmentManager();
        fm.beginTransaction().replace(R.id.headerPlace, HeaderFragment.newInstance("","")).commit();

        setBinding();
    }


    void setBinding(){
        ActivityCreateTokenBinding binding = DataBindingUtil.setContentView(this, R.layout.activity_create_token);
        binding.setClickHandler(new JobServiceTokenCreationHandler());

        CreateTokenViewModel createTokenViewModel = new CreateTokenViewModel();
        binding.setVM(createTokenViewModel);
        MainViewModel.getInstance().setCreateTokenViewModel(createTokenViewModel);
    }


        @Override
        protected  void onDestroy() {
            super.onDestroy();
            MainViewModel.getInstance().setCreateTokenViewModel(null);
        }


    @Override
    public void onFragmentInteraction(Uri uri) {

    }
}
