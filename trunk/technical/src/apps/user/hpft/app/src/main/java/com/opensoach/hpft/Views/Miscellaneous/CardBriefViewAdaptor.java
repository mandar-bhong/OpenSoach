package com.opensoach.hpft.Views.Miscellaneous;

import android.databinding.DataBindingUtil;
import android.view.View;
import android.widget.LinearLayout;

import com.opensoach.hpft.R;
import com.opensoach.hpft.ViewModels.CardBriefViewModel;
import com.opensoach.hpft.ViewModels.CardGridViewModel;
import com.opensoach.hpft.Views.Fragment.CardBriefFragment;
import com.opensoach.hpft.databinding.FragmentCardBriefBinding;

/**
 * Created by Mandar on 30-07-2018.
 */

public class CardBriefViewAdaptor extends CustomBaseAdapter<CardGridViewModel,CardBriefViewModel>{

    @Override
    protected View getItemView(CardBriefViewModel dataModel, int position) {
        return GetJobBrief(dataModel,position);
    }

    View GetJobBrief(CardBriefViewModel dataModel, int position){

        LinearLayout ll = new LinearLayout(ContextActivity);

        FragmentCardBriefBinding fragmentCardBriefBinding = DataBindingUtil.inflate(dataModel.ContextActivity.getLayoutInflater(),
                R.layout.fragment_card_brief,ll,true);

        CardBriefFragment jb = new CardBriefFragment();


        View v = jb.getView();
        fragmentCardBriefBinding.setData(dataModel);
        //fragmentCardBriefBinding.setClickHandler(new JobBriefClickHandler());

        ll.setId(position);

        return ll;
    }



}
