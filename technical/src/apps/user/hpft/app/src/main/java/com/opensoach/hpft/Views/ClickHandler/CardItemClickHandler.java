package com.opensoach.hpft.Views.ClickHandler;

import android.content.Intent;
import android.view.View;

import com.opensoach.hpft.ViewModels.CardBriefViewModel;
import com.opensoach.hpft.ViewModels.CardDetailViewModel;
import com.opensoach.hpft.Views.Activity.CardDetailsActivity;

import java.util.List;

/**
 * Created by Mandar on 8/26/2017.
 */

public class CardItemClickHandler {

    public void onClick(View view, CardBriefViewModel vm) {

        CardBriefViewModel jobBoardViewModel = new CardBriefViewModel();


        Intent i = new Intent(vm.ContextActivity.getBaseContext(), CardDetailsActivity.class);
        vm.ContextActivity.startActivity(i);

    }
}
