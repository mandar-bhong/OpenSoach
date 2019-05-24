package com.opensoach.vst.Views.Fragment;

import android.content.Context;
import android.util.AttributeSet;

import com.opensoach.vst.ViewModels.CardBriefViewModel;
import com.opensoach.vst.ViewModels.CardGridViewModel;
import com.opensoach.vst.Views.Interfaces.IFragment;
import com.opensoach.vst.Views.Interfaces.IList;
import com.opensoach.vst.Views.Miscellaneous.CardBriefViewAdaptor;

import java.util.List;

/**
 * Created by Mandar on 31-07-2018.
 */

public class CardGridView extends CustomGridView implements IFragment<CardGridViewModel>,IList<CardBriefViewModel> {

    public CardGridViewModel DataContext;
    private CardBriefViewAdaptor cardBriefViewAdaptor;

    public CardGridView(Context context){
        super(context);
    }

    public CardGridView(Context context, AttributeSet attrs) {
        super(context, attrs);
    }

    public CardGridView(Context context, AttributeSet attrs, int defStyleAttr) {
        super(context, attrs, defStyleAttr);
    }

    @Override
    public void setDataContext(CardGridViewModel viewModel) {
        //TODO: Check if null condition is required
        if(viewModel == null)return;
        DataContext = viewModel;
        cardBriefViewAdaptor = (CardBriefViewAdaptor) viewModel.getDataAdaptor();
        cardBriefViewAdaptor.ContextActivity = viewModel.ContextActivity;
        setAdapter(cardBriefViewAdaptor);
        cardBriefViewAdaptor.GridViewContainer = this;
    }

    @Override
    public CardGridViewModel getDataContext() {
        return DataContext;
    }

    @Override
    public List<CardBriefViewModel> getItemsSource() {
        return cardBriefViewAdaptor.getItemsSource();
    }


    @Override
    public void setItemsSource(List<CardBriefViewModel> source) {
        cardBriefViewAdaptor.setItemsSource(source);
    }

    @Override
    protected void setDataAdapter() {
        //setAdapter(cardBriefViewAdaptor);
    }
}
