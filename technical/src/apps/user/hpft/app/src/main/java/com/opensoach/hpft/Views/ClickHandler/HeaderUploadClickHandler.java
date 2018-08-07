package com.opensoach.hpft.Views.ClickHandler;

import android.content.Intent;
import android.view.View;
import android.widget.Toast;

import com.opensoach.hpft.AppRepo.AppRepo;
import com.opensoach.hpft.R;
import com.opensoach.hpft.ViewModels.CardBriefViewModel;
import com.opensoach.hpft.ViewModels.HeaderViewModel;
import com.opensoach.hpft.ViewModels.MainViewModel;
import com.opensoach.hpft.Views.Activity.CardDetailsActivity;
import com.opensoach.hpft.Views.DialogHelper;

/**
 * Created by Mandar on 06-08-2018.
 */

public class HeaderUploadClickHandler {
    public void onClick(View view, HeaderViewModel vm) {

        DialogHelper.showSingleLineEditTextAlert(
                view.getContext(),
                view.getContext().getResources().getString(R.string.dialog_enter_auth_code),
                new DialogHelper.DialogCallBack() {

                    @Override
                    public boolean onSucess(String authText) {

                        if (AppRepo.getInstance().getAuthCodeList().contains(authText)) {
                            //processChartData(authText);
                            return  true;
                        } else {
                            Toast.makeText(
                                    MainViewModel.getInstance().ContextActivity,
                                    MainViewModel.getInstance().ContextActivity.getResources().getString(R.string.invalid_auth_code),
                                    Toast.LENGTH_LONG).show();

                            return  false;
                        }
                    }

                    @Override
                    public void onSucess(String strData1, String strData2) {

                    }

                    @Override
                    public void onSucess(String strData1, String strData2, String strData3) {

                    }
                });

    }
}
