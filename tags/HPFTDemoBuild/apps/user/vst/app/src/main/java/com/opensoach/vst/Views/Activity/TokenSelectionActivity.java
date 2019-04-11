package com.opensoach.vst.Views.Activity;

import android.databinding.DataBindingUtil;
import android.net.Uri;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.support.v7.widget.DividerItemDecoration;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;

import com.opensoach.vst.Model.DB.DBTokenTableRowModel;
import com.opensoach.vst.R;
import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.ViewModels.TokenItemViewModel;
import com.opensoach.vst.ViewModels.TokenListViewModel;
import com.opensoach.vst.ViewModels.TokenSelectionViewModel;
import com.opensoach.vst.Views.ClickHandler.TokenSelectionHandler;
import com.opensoach.vst.Views.Fragment.HeaderFragment;
import com.opensoach.vst.Views.Fragment.TokenItemFragment;
import com.opensoach.vst.databinding.ActivityTokenSelectionBinding;

import java.util.ArrayList;

import static android.support.v7.widget.LinearLayoutManager.VERTICAL;

public class TokenSelectionActivity extends AppCompatActivity
        implements TokenItemFragment.OnFragmentInteractionListener,
        HeaderFragment.OnFragmentInteractionListener{

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_token_selection);


        MainViewModel.getInstance().ContextActivity = this;

        //TODO: This step is importent for adding fragment into activity
        android.support.v4.app.FragmentManager fm = getSupportFragmentManager();
        fm.beginTransaction().replace(R.id.headerPlace, HeaderFragment.newInstance("","")).commit();

        MainViewModel.getInstance().getHeaderViewModel().setBackButtonVisiable(false);

        setBinding();
    }

    void setBinding(){
        ActivityTokenSelectionBinding binding = DataBindingUtil.setContentView(this, R.layout.activity_token_selection);

        binding.setClickHandler(new TokenSelectionHandler());


        TokenSelectionViewModel tokenSelectionViewModel = new TokenSelectionViewModel();
        TokenListViewModel tokenListViewModel = MainViewModel.getInstance().getTokenListViewModel();
        tokenListViewModel.ContextActivity = this;
        ArrayList<TokenItemViewModel> list = new ArrayList<>();
        tokenSelectionViewModel.setTokenListViewModel(tokenListViewModel);

        binding.setVM(tokenSelectionViewModel);

        RecyclerView recyclerView = findViewById(R.id.recycler_view);
        recyclerView.setLayoutManager(new LinearLayoutManager(recyclerView.getContext()));
        recyclerView.addItemDecoration(new DividerItemDecoration(recyclerView.getContext(), VERTICAL));

    }

    @Override
    public void onFragmentInteraction(Uri uri) {

    }

}
