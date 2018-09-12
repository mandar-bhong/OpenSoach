package com.opensoach.vst.ViewModels;

import android.databinding.Bindable;

import com.opensoach.vst.BR;
import com.opensoach.vst.Views.Adapter.TaskTimeDataAdapter;
import com.opensoach.vst.Views.Adapter.TokensDataAdapter;

import java.util.ArrayList;
import java.util.List;

public class TokenListViewModel extends  BaseViewModel {

    private TokensDataAdapter tokensDataAdapter;
    private List<TokenItemViewModel> data;
    private TokenItemViewModel selectedToken;


    public  TokenListViewModel(){
        tokensDataAdapter = new TokensDataAdapter();
        data = new ArrayList<>() ;
    }

    public TokenItemViewModel getSelectedToken() {
        return selectedToken;
    }

    public void setSelectedToken(TokenItemViewModel selectedToken) {
        this.selectedToken = selectedToken;
        notifyPropertyChanged(BR.tokenSelected);
    }

    @Bindable
    public boolean isTokenSelected(){
        return (this.selectedToken != null) ? true: false;
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
