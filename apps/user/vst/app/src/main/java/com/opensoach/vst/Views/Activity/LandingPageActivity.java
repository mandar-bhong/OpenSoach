package com.opensoach.vst.Views.Activity;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;

import com.opensoach.vst.R;
import com.opensoach.vst.ViewModels.MainViewModel;

public class LandingPageActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_landing_page);

        MainViewModel.getInstance().ContextActivity = this;
    }
}
