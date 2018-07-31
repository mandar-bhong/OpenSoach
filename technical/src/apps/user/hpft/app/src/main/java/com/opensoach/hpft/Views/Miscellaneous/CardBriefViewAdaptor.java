package com.opensoach.hpft.Views.Miscellaneous;

import android.databinding.DataBindingUtil;
import android.view.View;
import android.widget.LinearLayout;

import com.opensoach.hpft.ViewModels.CardBriefViewModel;
import com.opensoach.hpft.ViewModels.CardGridViewModel;

/**
 * Created by Mandar on 30-07-2018.
 */

public class CardBriefViewAdaptor extends CustomBaseAdapter<CardGridViewModel,CardBriefViewModel>{

    View GetJobBrief(CardBriefViewModel dataModel, int position){

        LinearLayout ll = new LinearLayout(ContextActivity);

//        JobBrief jb = new JobBrief();
//
//        FragmentJobBriefBinding fragmentJobBriefBinding = DataBindingUtil.inflate(dataModel.ContextActivity.getLayoutInflater(),
//                R.layout.fragment_job_brief,ll,true);
//
//        View v = jb.getView();
//        fragmentJobBriefBinding.setData(dataModel);
//        fragmentJobBriefBinding.setClickHandler(new JobBriefClickHandler());

        ll.setId(position);

        return ll;
    }


    @Override
    protected View getItemView(CardBriefViewModel dataModel, int position) {
        return GetJobBrief(dataModel,position);
    }
}
