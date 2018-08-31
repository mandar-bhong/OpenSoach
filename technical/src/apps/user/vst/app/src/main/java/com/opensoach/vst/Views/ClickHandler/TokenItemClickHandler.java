package com.opensoach.vst.Views.ClickHandler;

import android.view.View;

import com.opensoach.vst.ViewModels.TaskTimeItemViewModel;
import com.opensoach.vst.ViewModels.TokenItemViewModel;
import com.opensoach.vst.ViewModels.TokenListViewModel;
import com.opensoach.vst.Views.Adapter.TokensDataAdapter;

public class TokenItemClickHandler {

    public void onClick(View view, TokenItemViewModel vm) {
        ((TokenListViewModel)vm.Parent).getTokensDataAdapter().SelectedIndexChange(vm.position);
        ((TokenListViewModel)vm.Parent).setSelectedToken(vm);
    }
}
