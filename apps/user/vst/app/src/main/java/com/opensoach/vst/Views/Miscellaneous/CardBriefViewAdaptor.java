package com.opensoach.vst.Views.Miscellaneous;

import android.databinding.DataBindingUtil;
import android.view.View;
import android.widget.LinearLayout;

import com.opensoach.vst.R;
import com.opensoach.vst.ViewModels.CardBriefViewModel;
import com.opensoach.vst.ViewModels.CardGridViewModel;
import com.opensoach.vst.Views.ClickHandler.CardItemClickHandler;
import com.opensoach.vst.Views.Fragment.CardBriefFragment;
import com.opensoach.vst.databinding.FragmentCardBriefBinding;

/**
 * Created by Mandar on 30-07-2018.
 */

public class CardBriefViewAdaptor extends CustomBaseAdapter<CardGridViewModel,CardBriefViewModel>{

    @Override
    protected View getItemView(CardBriefViewModel dataModel, int position) {
        return GetCardBrief(dataModel,position);
    }

    View GetCardBrief(CardBriefViewModel dataModel, int position){

        LinearLayout ll = new LinearLayout(ContextActivity);

        FragmentCardBriefBinding fragmentCardBriefBinding = DataBindingUtil.inflate(dataModel.ContextActivity.getLayoutInflater(),
                R.layout.fragment_card_brief,ll,true);

        CardBriefFragment cardBriefFragment = new CardBriefFragment();

        View v = cardBriefFragment.getView();
        fragmentCardBriefBinding.setData(dataModel);
        fragmentCardBriefBinding.setClickHandler(new CardItemClickHandler());

        ll.setId(position);

        return ll;
    }



}
