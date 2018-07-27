package com.opensoach.hospital.Views.Activity;

import android.os.Handler;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.view.WindowManager;
import android.view.animation.Animation;
import android.view.animation.AnimationUtils;
import android.widget.ImageView;
import android.widget.ProgressBar;

import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.R;
import com.opensoach.hospital.Utility.AppLogger;

import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;

public class SplashScreenActivity extends AppCompatActivity implements Animation.AnimationListener,PropertyChangeListener {

    Animation animFadein;
    ImageView img1;
    private ProgressBar mProgress;
    boolean isSwitchToMainActivity =false;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_splash_screen);
        mProgress = (ProgressBar) findViewById(R.id.splash_screen_progress_bar);

        if(AppRepo.getInstance().getIsStartUpCompleted()){
            new Handler().postDelayed(new Runnable() {
                @Override
                public void run() {
                    finish();
                }
            },1500);
        }else{

            AppRepo.getInstance().addPropertyChangeListener(this);

            animFadein = AnimationUtils.loadAnimation(getApplicationContext(),
                    R.anim.round);
            animFadein.setAnimationListener(this);
            img1=(ImageView)findViewById(R.id.logo);
            img1.startAnimation(animFadein);

            new Thread(new Runnable() {
                public void run() {
                    doWork();
                }
            }).start();
        }
    }

    private void doWork() {
        for (int progress=0; progress<100; progress+=1) {
            try {
                Thread.sleep(50);
                mProgress.setProgress(progress);
            } catch (Exception e) {
                AppLogger.getInstance().Log(e);
            }

            if(isSwitchToMainActivity){
                progress = 100;
            }
        }
    }

    @Override
    public void onAnimationStart(Animation animation) {

    }

    @Override
    public void onAnimationEnd(Animation animation) {

    }

    @Override
    public void onAnimationRepeat(Animation animation) {

    }

    @Override
    public void propertyChange(PropertyChangeEvent evt) {

        AppLogger.getInstance().Log(AppLogger.LogLevel.Debug,"SplashScreen: PropChanged: PropName: "+evt.getPropertyName());

        switch (evt.getPropertyName()) {
            case AppRepo.IsStartUpCompletedPropName: {
                isSwitchToMainActivity = true;
                AppRepo.getInstance().removePropertyChangeListener(this);
                finish();
            }
            break;
        }
    }

    @Override
    public void onWindowFocusChanged(boolean hasFocus) {
        super.onWindowFocusChanged(hasFocus);

        if (hasFocus) {
            getWindow().getDecorView().setSystemUiVisibility(
                    View.SYSTEM_UI_FLAG_LAYOUT_STABLE
                            | View.SYSTEM_UI_FLAG_LAYOUT_HIDE_NAVIGATION
                            | View.SYSTEM_UI_FLAG_LAYOUT_FULLSCREEN
                            | View.SYSTEM_UI_FLAG_HIDE_NAVIGATION
                            | View.SYSTEM_UI_FLAG_FULLSCREEN
                            | View.SYSTEM_UI_FLAG_IMMERSIVE_STICKY);

            getWindow().addFlags(WindowManager.LayoutParams.FLAG_FULLSCREEN);
        }
    }
}
