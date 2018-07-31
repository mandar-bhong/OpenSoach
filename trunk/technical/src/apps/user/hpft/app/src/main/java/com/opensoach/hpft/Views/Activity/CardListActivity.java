package com.opensoach.hpft.Views.Activity;


import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.view.WindowManager;

import com.opensoach.hpft.Model.AppNotificationModelBase;
import com.opensoach.hpft.R;
import com.opensoach.hpft.ViewModels.CardListViewModel;
import com.opensoach.hpft.ViewModels.MainViewModel;
import com.opensoach.hpft.Views.Interfaces.IFragment;
import com.opensoach.hpft.Views.Interfaces.IUIUpdateEvent;


public class CardListActivity extends AppCompatActivity implements IFragment<CardListViewModel>,IUIUpdateEvent {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_card_list);
        setDataContext(MainViewModel.getInstance().getCardListViewModel());

        hideSoftKeyboard();
    }

    public void hideSoftKeyboard() {
        getWindow().setSoftInputMode(WindowManager.LayoutParams.SOFT_INPUT_STATE_HIDDEN);
    }

    @Override
    public void onWindowFocusChanged(boolean hasFocus) {
        super.onWindowFocusChanged(hasFocus);

        if (hasFocus) {
            getWindow().getDecorView().setSystemUiVisibility(
                    View.SYSTEM_UI_FLAG_LAYOUT_STABLE
                            | View.SYSTEM_UI_FLAG_LAYOUT_HIDE_NAVIGATION
                            | View.SYSTEM_UI_FLAG_LAYOUT_FULLSCREEN
                            | View.SYSTEM_UI_FLAG_HIDE_NAVIGATION
                            | View.SYSTEM_UI_FLAG_FULLSCREEN
                            | View.SYSTEM_UI_FLAG_IMMERSIVE_STICKY);

            getWindow().addFlags(WindowManager.LayoutParams.FLAG_FULLSCREEN);
        }
    }

    @Override
    public  void setDataContext(CardListViewModel vm){

    }

    @Override
    public CardListViewModel getDataContext() {
        return null;
    }

    @Override
    public void OnUIUpdateEvent(final AppNotificationModelBase model) {
        switch (model.DataProcessStatergyID) {

        }
    }


}
