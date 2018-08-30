package com.opensoach.vst.Views.ClickHandler;

import android.app.Activity;
import android.content.Intent;
import android.view.View;

import com.opensoach.vst.BR;
import com.opensoach.vst.Helper.AppAction;
import com.opensoach.vst.Manager.SendPacketManager;
import com.opensoach.vst.ViewModels.CreateTokenViewModel;
import com.opensoach.vst.Views.Activity.CreateTokenActivity;

public class CreateTokenClickHandler {

    public void onClick(View view, CreateTokenViewModel viewModel) {
        viewModel.setGenerateTokenVisible(false);

        SendPacketManager.Instance().send(AppAction.CREATE_TOKEN, viewModel);

    }

    public void onShowCreateToken(View view, CreateTokenViewModel viewModel) {
        viewModel.setGenerateTokenVisible(true);
        viewModel.notifyPropertyChanged(BR._all);
    }

    public void onTokenCreateCompleted(View view, CreateTokenViewModel viewModel) {
        ((Activity)view.getContext()).finish();
    }
}
