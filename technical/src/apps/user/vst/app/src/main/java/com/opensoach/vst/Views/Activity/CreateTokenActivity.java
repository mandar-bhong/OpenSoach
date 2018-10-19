package com.opensoach.vst.Views.Activity;

import android.databinding.DataBindingUtil;
import android.net.Uri;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.text.Editable;
import android.text.InputFilter;
import android.text.TextWatcher;
import android.widget.EditText;

import com.opensoach.vst.R;
import com.opensoach.vst.ViewModels.CreateTokenViewModel;
import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.Views.ClickHandler.JobServiceTokenCreationHandler;
import com.opensoach.vst.Views.Fragment.HeaderFragment;
import com.opensoach.vst.databinding.ActivityCreateTokenBinding;

public class CreateTokenActivity extends AppCompatActivity
        implements HeaderFragment.OnFragmentInteractionListener{

    EditText VhlNo1;
    EditText VhlNo2;
    EditText VhlNo3;
    EditText VhlNo4;


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_create_token);

        MainViewModel.getInstance().ContextActivity = this;

        //TODO: This step is importent for adding fragment into activity
        android.support.v4.app.FragmentManager fm = getSupportFragmentManager();
        fm.beginTransaction().replace(R.id.headerPlace, HeaderFragment.newInstance("","")).commit();

        setBinding();


        //TODO: View event binding should be after databinding
        AttachedViewEvent();
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


    private void AttachedViewEvent(){
        VhlNo1 = (EditText)findViewById(R.id.vehicalNo1);
        VhlNo2 = (EditText)findViewById(R.id.vehicalNo2);
        VhlNo3 = (EditText)findViewById(R.id.vehicalNo3);
        VhlNo4 = (EditText)findViewById(R.id.vehicalNo4);

        VhlNo1.setFilters(new InputFilter[]{new InputFilter.AllCaps()});
        VhlNo3.setFilters(new InputFilter[]{new InputFilter.AllCaps()});


        VhlNo1.addTextChangedListener(new TextWatcher() {
            @Override
            public void beforeTextChanged(CharSequence s, int start,int count, int after) {}
            @Override
            public void onTextChanged(CharSequence s, int start,int before, int count) {

            }

            @Override
            public void afterTextChanged(Editable s) {
                if(VhlNo1.getText().length() == 2){
                    VhlNo2.requestFocus();
                }
            }
        });

        VhlNo2.addTextChangedListener(new TextWatcher() {

            @Override
            public void beforeTextChanged(CharSequence s, int start,int count, int after) {}
            @Override
            public void onTextChanged(CharSequence s, int start,int before, int count) {}

            @Override
            public void afterTextChanged(Editable s) {
                if(VhlNo2.getText().length() == 2){
                    VhlNo3.requestFocus();
                }
            }
        });

        VhlNo3.addTextChangedListener(new TextWatcher() {
            @Override
            public void beforeTextChanged(CharSequence s, int start,int count, int after) {}
            @Override
            public void onTextChanged(CharSequence s, int start,int before, int count) {}

            @Override
            public void afterTextChanged(Editable s) {
                if(VhlNo3.getText().length() == 2){
                    VhlNo4.requestFocus();
                }
            }
        });

    }
}
