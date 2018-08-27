package com.opensoach.vst.Views.Activity;




import android.databinding.DataBindingUtil;
import android.net.Uri;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.WindowManager;

import com.opensoach.vst.Model.AppNotificationModelBase;
import com.opensoach.vst.R;
import com.opensoach.vst.ViewModels.CardBriefViewModel;
import com.opensoach.vst.ViewModels.CardGridViewModel;
import com.opensoach.vst.ViewModels.CardListViewModel;
import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.Views.Fragment.HeaderFragment;
import com.opensoach.vst.Views.Interfaces.IFragment;
import com.opensoach.vst.Views.Interfaces.IUIUpdateEvent;
import com.opensoach.vst.databinding.ActivityCardListBinding;

import java.util.ArrayList;


public class CardListActivity extends AppCompatActivity
        implements IFragment<CardListViewModel>,IUIUpdateEvent,HeaderFragment.OnFragmentInteractionListener  {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        MainViewModel.getInstance().ContextActivity = this;

        setContentView(R.layout.activity_card_list);
        setDataContext(MainViewModel.getInstance().getCardListViewModel());

        //TODO: This step is importent for adding fragment into activity
        android.support.v4.app.FragmentManager fm = getSupportFragmentManager();
        fm.beginTransaction().replace(R.id.headerPlace, HeaderFragment.newInstance("","")).commit();

        hideSoftKeyboard();

    }

    public void hideSoftKeyboard() {
        getWindow().setSoftInputMode(WindowManager.LayoutParams.SOFT_INPUT_STATE_HIDDEN);
    }


    @Override
    public  void setDataContext(CardListViewModel viewModel){
        ActivityCardListBinding activityMainBinding = DataBindingUtil.setContentView(this, R.layout.activity_card_list);

        //viewModel.AppContext = this.getBaseContext();
        viewModel.ContextActivity = this;
        viewModel.getCardGridViewModel().ContextActivity = this;

        CardGridViewModel cardGridViewModel = viewModel.getCardGridViewModel();

        ArrayList list = new ArrayList<CardBriefViewModel>();

        cardGridViewModel.setItemsSource(list);

        activityMainBinding.setDataContext(cardGridViewModel);
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


    @Override
    public void onFragmentInteraction(Uri uri){

    }

}
