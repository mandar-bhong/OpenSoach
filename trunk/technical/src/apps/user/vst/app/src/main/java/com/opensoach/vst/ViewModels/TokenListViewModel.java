package com.opensoach.vst.ViewModels;

import android.databinding.Bindable;

import com.opensoach.vst.Views.Adapter.TaskTimeDataAdapter;
import com.opensoach.vst.Views.Adapter.TokensDataAdapter;

import java.util.List;

public class TokenListViewModel extends  BaseViewModel {

    private TokensDataAdapter tokensDataAdapter;
    private List<TokenItemViewModel> data;


    public  TokenListViewModel(){
        tokensDataAdapter = new TokensDataAdapter();
    }


    @Bindable
    public TokensDataAdapter getTokensDataAdapter() {
        return tokensDataAdapter;
    }

    public List<TokenItemViewModel> getData() {
        return data;
    }

    public void setData(List<TokenItemViewModel> data) {
        this.data = data;
    }
}
