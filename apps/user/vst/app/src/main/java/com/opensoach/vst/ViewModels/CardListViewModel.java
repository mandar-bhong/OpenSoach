package com.opensoach.vst.ViewModels;

import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;

/**
 * Created by Mandar on 30-07-2018.
 */

public class CardListViewModel extends BaseViewModel implements PropertyChangeListener {

    private CardGridViewModel cardGridViewModel;


    public CardListViewModel() {
        this.cardGridViewModel = new CardGridViewModel();
    }

    public CardGridViewModel getCardGridViewModel() {
        return cardGridViewModel;
    }

    public void setCardGridViewModel(CardGridViewModel cardGridViewModel) {
        this.cardGridViewModel = cardGridViewModel;
    }

    @Override
    public void propertyChange(PropertyChangeEvent propertyChangeEvent) {

    }
}
