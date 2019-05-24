package com.opensoach.vst.Views.ClickHandler;

import android.content.Intent;
import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.ViewModels.CardBriefViewModel;
import com.opensoach.vst.Views.Activity.CardDetailsActivity;

/**
 * Created by Mandar on 8/26/2017.
 */

public class CardItemClickHandler {

    public void onClick(View view, CardBriefViewModel vm) {
        AppRepo.getInstance().setActiveCard(vm);

        Intent i = new Intent(vm.ContextActivity.getBaseContext(), CardDetailsActivity.class);
        vm.ContextActivity.startActivity(i);
    }
}
