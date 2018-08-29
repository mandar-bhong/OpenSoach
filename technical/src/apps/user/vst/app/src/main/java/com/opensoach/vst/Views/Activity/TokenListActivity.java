package com.opensoach.vst.Views.Activity;

import android.app.Activity;
import android.databinding.DataBindingUtil;
import android.net.Uri;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.support.v7.widget.DividerItemDecoration;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Model.DB.DBTokenTableRowModel;
import com.opensoach.vst.R;
import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.ViewModels.TokenItemViewModel;
import com.opensoach.vst.ViewModels.TokenListViewModel;
import com.opensoach.vst.Views.ClickHandler.GenerateTokenClickHandler;
import com.opensoach.vst.Views.Fragment.HeaderFragment;
import com.opensoach.vst.Views.Fragment.TokenItemFragment;
import com.opensoach.vst.databinding.ActivityTaskDetailsBinding;
import com.opensoach.vst.databinding.ActivityTokenListBinding;

import java.util.ArrayList;
import java.util.Date;

import static android.support.v7.widget.LinearLayoutManager.VERTICAL;

public class TokenListActivity extends AppCompatActivity
        implements TokenItemFragment.OnFragmentInteractionListener,
        HeaderFragment.OnFragmentInteractionListener{

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_token_list);

        MainViewModel.getInstance().ContextActivity = this;

        //TODO: This step is importent for adding fragment into activity
        android.support.v4.app.FragmentManager fm = getSupportFragmentManager();
        fm.beginTransaction().replace(R.id.headerPlace, HeaderFragment.newInstance("","")).commit();


        setBinding();

    }


    void setBinding(){
        ActivityTokenListBinding binding = DataBindingUtil.setContentView(this, R.layout.activity_token_list);
        binding.setVM(GenerateData());
        binding.setClickHandler(new GenerateTokenClickHandler());

        RecyclerView recyclerView = findViewById(R.id.recycler_view);
        recyclerView.setLayoutManager(new LinearLayoutManager(recyclerView.getContext()));
        recyclerView.addItemDecoration(new DividerItemDecoration(recyclerView.getContext(), VERTICAL));

    }

    @Override
    public void onFragmentInteraction(Uri uri) {

    }

    TokenListViewModel GenerateData() {
        TokenListViewModel tokenListViewModel = new TokenListViewModel();
        tokenListViewModel.ContextActivity = this;

        ArrayList<TokenItemViewModel> list = new ArrayList<>();

        DBTokenTableRowModel dbTokenTableRowModel = new DBTokenTableRowModel();
        dbTokenTableRowModel.setTokenno(5);
        dbTokenTableRowModel.setGeneratedon(new Date());
        dbTokenTableRowModel.setVehicleno("MH 12 DC3422");

        TokenItemViewModel tokenItemViewModel = new TokenItemViewModel(dbTokenTableRowModel);
        tokenItemViewModel.ContextActivity = this;
        tokenItemViewModel.Parent = tokenListViewModel;


        list.add(tokenItemViewModel);
        list.add(tokenItemViewModel);
        list.add(tokenItemViewModel);

        tokenListViewModel.setData(list);

        return tokenListViewModel;

    }
}
