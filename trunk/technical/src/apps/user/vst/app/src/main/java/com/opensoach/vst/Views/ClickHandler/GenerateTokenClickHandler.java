package com.opensoach.vst.Views.ClickHandler;

import android.content.Intent;
import android.speech.tts.TextToSpeech;
import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.ViewModels.CardBriefViewModel;
import com.opensoach.vst.Views.Activity.CreateTokenActivity;

public class GenerateTokenClickHandler {

    public void onClick(View view) {
        Intent i = new Intent(view.getContext(), CreateTokenActivity.class);
        i.setAction(TextToSpeech.Engine.ACTION_CHECK_TTS_DATA);
        view.getContext().startActivity(i);
    }
}
