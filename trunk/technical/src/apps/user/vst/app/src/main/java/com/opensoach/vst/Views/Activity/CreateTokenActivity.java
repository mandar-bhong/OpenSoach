package com.opensoach.vst.Views.Activity;

import android.databinding.DataBindingUtil;
import android.net.Uri;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.support.v7.widget.DividerItemDecoration;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;

import com.opensoach.vst.R;
import com.opensoach.vst.ViewModels.CreateTokenViewModel;
import com.opensoach.vst.Views.ClickHandler.CreateTokenClickHandler;
import com.opensoach.vst.Views.ClickHandler.GenerateTokenClickHandler;
import com.opensoach.vst.Views.Fragment.HeaderFragment;
import com.opensoach.vst.databinding.ActivityCreateTokenBinding;
import com.opensoach.vst.databinding.ActivityTokenListBinding;

import static android.support.v7.widget.LinearLayoutManager.VERTICAL;

public class CreateTokenActivity extends AppCompatActivity
        implements HeaderFragment.OnFragmentInteractionListener{

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_create_token);

        //TODO: This step is importent for adding fragment into activity
        android.support.v4.app.FragmentManager fm = getSupportFragmentManager();
        fm.beginTransaction().replace(R.id.headerPlace, HeaderFragment.newInstance("","")).commit();


        setBinding();
    }


    void setBinding(){
        ActivityCreateTokenBinding binding = DataBindingUtil.setContentView(this, R.layout.activity_create_token);
        binding.setClickHandler(new CreateTokenClickHandler());
        binding.setVM(new CreateTokenViewModel());
    }



    @Override
    public void onFragmentInteraction(Uri uri) {

    }
}
