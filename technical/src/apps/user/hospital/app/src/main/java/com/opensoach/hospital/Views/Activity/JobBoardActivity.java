package com.opensoach.hospital.Views.Activity;

import android.app.AlertDialog;
import android.app.Dialog;
import android.content.Context;
import android.content.DialogInterface;
import android.databinding.DataBindingUtil;
import android.graphics.Color;
import android.graphics.Rect;
import android.os.Bundle;
import android.os.Handler;
import android.os.Looper;
import android.os.Message;
import android.support.v4.view.ViewPager;
import android.support.v7.app.AppCompatActivity;
import android.text.Html;
import android.text.method.ScrollingMovementMethod;
import android.view.LayoutInflater;
import android.view.MotionEvent;
import android.view.TouchDelegate;
import android.view.View;
import android.view.ViewGroup;
import android.view.WindowManager;
import android.view.inputmethod.InputMethodManager;
import android.widget.Button;
import android.widget.ImageButton;
import android.widget.ImageView;
import android.widget.LinearLayout;
import android.widget.RelativeLayout;
import android.widget.TextView;
import android.widget.Toast;

import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.Helper.ApplicationConstants;
import com.opensoach.hospital.Helper.CommandConstants;
import com.opensoach.hospital.Helper.CommonHelper;
import com.opensoach.hospital.Helper.Constants;
import com.opensoach.hospital.Model.AppNotificationModelBase;
import com.opensoach.hospital.Model.DB.DBJobCardTableRowModel;
import com.opensoach.hospital.Model.DB.DBPartDrawingTableRowModel;
import com.opensoach.hospital.Model.View.UIViewActionRequestModel;
import com.opensoach.hospital.R;
import com.opensoach.hospital.Utility.AppLogger;
import com.opensoach.hospital.ViewModels.JobBoardViewModel;
import com.opensoach.hospital.ViewModels.JobBriefViewModel;
import com.opensoach.hospital.ViewModels.JobQuantityViewModel;
import com.opensoach.hospital.ViewModels.MainViewModel;
import com.opensoach.hospital.ViewModels.OperatorCodeViewModel;
import com.opensoach.hospital.Views.ClickHandler.IViewAction;
import com.opensoach.hospital.Views.ClickHandler.JobAbortClickHandler;
import com.opensoach.hospital.Views.ClickHandler.JobBoardClickHandler;
import com.opensoach.hospital.Views.ClickHandler.JobDropClickHandler;
import com.opensoach.hospital.Views.ClickHandler.JobQuantityUpdateHandler;
import com.opensoach.hospital.Views.ClickHandler.JobStartClickHandler;
import com.opensoach.hospital.Views.ClickHandler.JobStopClickHandler;
import com.opensoach.hospital.Views.ClickHandler.OperatorCodeInputClickHandler;
import com.opensoach.hospital.Views.Interfaces.IUIUpdateEvent;
import com.opensoach.hospital.Views.Miscellaneous.SliderPagerAdapter;
import com.opensoach.hospital.databinding.ActivityJobBoardBinding;
import com.opensoach.hospital.databinding.FragmentJobQuantityUpdateBinding;
import com.opensoach.hospital.databinding.FragmentOperatorCodeBinding;

import java.util.ArrayList;

public class JobBoardActivity extends AppCompatActivity implements IViewAction,IUIUpdateEvent {

    private LinearLayout mRightSideLayout, mLeftSideLayout, jobToolsDetailsLayout;
    private ImageButton minimizeMaximizeButton;
    private TextView jobToolDescriptionText, jobProcessDescriptionText;
    RelativeLayout jobDescriptionTitleLayout;
    private JobBoardViewModel _jobBoardViewModel;

    private ViewPager vp_slider;
    private LinearLayout ll_dots;
    SliderPagerAdapter sliderPagerAdapter;
    ArrayList<String> slider_image_list;
    private TextView[] dots;
    private AlertDialog _dialog;

    ImageView minimizeMaximizeToolButton;

    Button resizeProcessDataBtn;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_job_board);

        ActivityJobBoardBinding activityJobBoardBinding = DataBindingUtil.setContentView(this, R.layout.activity_job_board);

        _jobBoardViewModel = AppRepo.getInstance().getSelectedJobBoard();
        _jobBoardViewModel.ContextActivity = this;
        //_jobBoardViewModel.setJobCardId(1);
        AppRepo.getInstance().addPropertyChangeListener(_jobBoardViewModel);

        activityJobBoardBinding.setVM(_jobBoardViewModel);
        activityJobBoardBinding.setClickHandler(new JobBoardClickHandler());
        activityJobBoardBinding.setJobboardactivity(this);

        mRightSideLayout = (LinearLayout) findViewById(R.id.rightSideLayout);
        mLeftSideLayout = (LinearLayout) findViewById(R.id.leftSideLayout);
        jobToolsDetailsLayout = (LinearLayout) findViewById(R.id.jobToolsDetailsLayout);
        jobDescriptionTitleLayout = (RelativeLayout) findViewById(R.id.jobDescriptionTitleLayout);

        minimizeMaximizeButton = (ImageButton) findViewById(R.id.minimizeMaximizeButton);
        resizeProcessDataBtn = (Button) findViewById(R.id.resizeProcessDataBtn);
        minimizeMaximizeToolButton = (ImageView) findViewById(R.id.minimizeMaximizeToolButton);

        jobToolDescriptionText = (TextView) findViewById(R.id.jobToolDescriptionText);
        jobProcessDescriptionText = (TextView) findViewById(R.id.jobProcessDescriptionText);
        jobToolDescriptionText.setMovementMethod(new ScrollingMovementMethod());
        jobProcessDescriptionText.setMovementMethod(new ScrollingMovementMethod());

        hideSoftKeyboard();

        // method for initialisation
        init();

        setJobData(_jobBoardViewModel);
        setViewData();

        // method for adding indicators
        addBottomDots(0);

        addExtraTouchArea(minimizeMaximizeButton);
        addExtraTouchArea(resizeProcessDataBtn);
        addExtraTouchArea(minimizeMaximizeToolButton);

        minimizeMaximizeButton.setOnTouchListener(new View.OnTouchListener() {
            @Override
            public boolean onTouch(View v, MotionEvent event) {
                if (event.getAction() == MotionEvent.ACTION_DOWN) {
                    Handler uiHandler = new Handler(Looper.getMainLooper()) {
                        @Override
                        public void handleMessage(Message message) {
                            imageResizeClick();
                        }
                    };

                    Message msg = uiHandler.obtainMessage();
                    Bundle b = new Bundle();
                    // Do what you want

                    uiHandler.sendMessage(msg);
                    return true;
                }
                return false;
            }
        });


        resizeProcessDataBtn.setOnTouchListener(new View.OnTouchListener() {
            @Override
            public boolean onTouch(View v, MotionEvent event) {
                if (event.getAction() == MotionEvent.ACTION_DOWN) {
                    processResizeClick(v);
                    // Do what you want
                    return true;
                }
                return false;
            }
        });

        minimizeMaximizeToolButton.setOnTouchListener(new View.OnTouchListener() {
            @Override
            public boolean onTouch(View v, MotionEvent event) {
                if (event.getAction() == MotionEvent.ACTION_DOWN) {
                    toolsResizeClick(v);
                    // Do what you want
                    return true;
                }
                return false;
            }
        });

    }

    public void hideSoftKeyboard() {
        //getWindow().setSoftInputMode(WindowManager.LayoutParams.SOFT_INPUT_STATE_HIDDEN);

        View view = super.getCurrentFocus();
        if (view != null) {
            InputMethodManager inputManager = (InputMethodManager) super.getSystemService(Context.INPUT_METHOD_SERVICE);
            inputManager.hideSoftInputFromWindow(view.getWindowToken(), InputMethodManager.HIDE_NOT_ALWAYS);
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

    @Override
    protected void onResume() {
        super.onResume();
        AppRepo.getInstance().setForegroundActivityName(ApplicationConstants.FOREGROUND_ACTIVITY_JOB_BOARD);
        AppRepo.getInstance().setForegroundActivityHandler(this);
    }

    @Override
    protected void onDestroy(){
        super.onDestroy();
        AppRepo.getInstance().removePropertyChangeListener(_jobBoardViewModel);
    }

    public void imageResizeClick() {
        if (mRightSideLayout.getVisibility() == View.VISIBLE) {
            mRightSideLayout.setVisibility(View.GONE);
            LinearLayout.LayoutParams param = new LinearLayout.LayoutParams(
                    LinearLayout.LayoutParams.MATCH_PARENT,
                    LinearLayout.LayoutParams.MATCH_PARENT
            );
            param.weight = 1.0f;
            mLeftSideLayout.setLayoutParams(param);
            minimizeMaximizeButton.setBackgroundResource(R.drawable.minimize);
        } else {
            mRightSideLayout.setVisibility(View.VISIBLE);

            LinearLayout.LayoutParams param = new LinearLayout.LayoutParams(
                    LinearLayout.LayoutParams.MATCH_PARENT,
                    LinearLayout.LayoutParams.MATCH_PARENT
            );
            param.weight = 0.5f;
            mLeftSideLayout.setLayoutParams(param);
            mRightSideLayout.setLayoutParams(param);

            minimizeMaximizeButton.setBackgroundResource(R.drawable.fullscreen);

        }

        addExtraTouchArea(minimizeMaximizeButton);
    }

    public void processResizeClick(View v) {
        if (jobToolsDetailsLayout.getVisibility() == View.VISIBLE) {
            jobToolsDetailsLayout.setVisibility(View.GONE);

            jobProcessDescriptionText.setLayoutParams(new LinearLayout.LayoutParams(LinearLayout.LayoutParams.MATCH_PARENT, LinearLayout.LayoutParams.MATCH_PARENT, 1f));
            resizeProcessDataBtn.setBackgroundResource(R.drawable.minimize);
        } else {
            jobToolsDetailsLayout.setVisibility(View.VISIBLE);

            LinearLayout.LayoutParams param = new LinearLayout.LayoutParams(
                    LinearLayout.LayoutParams.MATCH_PARENT,
                    LinearLayout.LayoutParams.MATCH_PARENT
            );
            param.weight = 0.5f;
            jobToolsDetailsLayout.setLayoutParams(param);
            jobProcessDescriptionText.setLayoutParams(param);

            resizeProcessDataBtn.setBackgroundResource(R.drawable.fullscreen);
        }

        addExtraTouchArea(resizeProcessDataBtn);
    }

    public void toolsResizeClick(View v) {
        if (jobProcessDescriptionText.getVisibility() == View.VISIBLE) {
            jobProcessDescriptionText.setVisibility(View.GONE);
            jobDescriptionTitleLayout.setVisibility(View.GONE);

            jobToolsDetailsLayout.setLayoutParams(new LinearLayout.LayoutParams(LinearLayout.LayoutParams.MATCH_PARENT, LinearLayout.LayoutParams.MATCH_PARENT, 1f));
            minimizeMaximizeToolButton.setBackgroundResource(R.drawable.minimize);
        } else {
            jobProcessDescriptionText.setVisibility(View.VISIBLE);
            jobDescriptionTitleLayout.setVisibility(View.VISIBLE);

            LinearLayout.LayoutParams param = new LinearLayout.LayoutParams(
                    LinearLayout.LayoutParams.MATCH_PARENT,
                    LinearLayout.LayoutParams.MATCH_PARENT
            );
            param.weight = 0.5f;
            jobToolsDetailsLayout.setLayoutParams(param);
            jobProcessDescriptionText.setLayoutParams(param);

            minimizeMaximizeToolButton.setBackgroundResource(R.drawable.fullscreen);
        }

        addExtraTouchArea(minimizeMaximizeToolButton);

    }


    public void onSendClick(View v) {

        Dialog dialog = new Dialog(getApplicationContext());

        View parent = getLayoutInflater().inflate(R.layout.activity_job_board, null);

        ViewGroup view = (ViewGroup) findViewById(android.R.id.content);

        FragmentJobQuantityUpdateBinding fragmentJobQuantityUpdateBinding = (FragmentJobQuantityUpdateBinding) DataBindingUtil.inflate(LayoutInflater.from(JobBoardActivity.this),
                R.layout.fragment_job_quantity_update, null, false);
        dialog.getWindow().setType(WindowManager.LayoutParams.TYPE_SYSTEM_ALERT);
        dialog.setContentView(fragmentJobQuantityUpdateBinding.getRoot());
        dialog.setOnDismissListener(new DialogInterface.OnDismissListener() {
            @Override
            public void onDismiss(DialogInterface dialog) {
                hideSoftKeyboard();
            }
        });
        dialog.show();
    }


    public void backButtonClicked(View v) {
        this.finish();
    }


    private void init() {

        vp_slider = (ViewPager) findViewById(R.id.vp_slider);
        ll_dots = (LinearLayout) findViewById(R.id.ll_dots);

        slider_image_list = new ArrayList<>();

        sliderPagerAdapter = new SliderPagerAdapter(JobBoardActivity.this, slider_image_list);
        vp_slider.setAdapter(sliderPagerAdapter);

        vp_slider.setOnPageChangeListener(new ViewPager.OnPageChangeListener() {
            @Override
            public void onPageScrolled(int position, float positionOffset, int positionOffsetPixels) {

            }

            @Override
            public void onPageSelected(int position) {
                addBottomDots(position);
            }

            @Override
            public void onPageScrollStateChanged(int state) {

            }
        });
    }


    private void addBottomDots(int currentPage) {
        dots = new TextView[slider_image_list.size()];

        ll_dots.removeAllViews();
        for (int i = 0; i < dots.length; i++) {
            dots[i] = new TextView(this);
            dots[i].setText(Html.fromHtml("&#8226;"));
            dots[i].setTextSize(35);
            dots[i].setTextColor(Color.parseColor("#AFAFAF"));
            ll_dots.addView(dots[i]);
        }

        if (dots.length > 0)
            dots[currentPage].setTextColor(Color.parseColor("#000000"));
    }


    private void setJobData(JobBoardViewModel jobBoardViewModel) {


        for (DBPartDrawingTableRowModel drawingRow : jobBoardViewModel.getDbPartDrawingTableRowModels()) {
            slider_image_list.add( Constants.WEB_SERVICE_URL + "drawing/" + drawingRow.getPath());
        }

        sliderPagerAdapter.notifyDataSetChanged();

    }

    private void setViewData() {

        if (_jobBoardViewModel.getDbJobCardTableRowModel() == null) return;
        if (_jobBoardViewModel.getDbEnggPartTableRowModel() == null) return;

        jobProcessDescriptionText.setText(_jobBoardViewModel.getDbEnggPartTableRowModel().getProcess());

        String toolDisplay = CommonHelper.ConvertToolJSONTOText(_jobBoardViewModel.getDbEnggPartTableRowModel().getToolJSON());

        jobToolDescriptionText.setText(Html.fromHtml(toolDisplay));
    }

    private void addExtraTouchArea(final View view) {

        final View parent = (View) view.getParent();  // button: the view you want to enlarge hit area
        parent.post(new Runnable() {
            public void run() {
                final Rect rect = new Rect();
                view.getHitRect(rect);
                rect.top -= 100;    // increase top hit area
                rect.left -= 100;   // increase left hit area
                rect.bottom += 100; // increase bottom hit area
                rect.right += 100;  // increase right hit area
                parent.setTouchDelegate(new TouchDelegate(rect, view));
            }
        });
    }

    @Override
    public boolean onTouchEvent(MotionEvent ev) {
        try {
            return super.onTouchEvent(ev);
        } catch (IllegalArgumentException ex) {
            AppLogger.getInstance().Log(AppLogger.LogLevel.Error,ex);
        }
        return false;
    }

    @Override
    public void ExectuteViewTask(int taskID, UIViewActionRequestModel requestData) {

        switch (taskID) {
            case CommandConstants.UI_CMD_BACKGROUND_STOP_JOB:
            case CommandConstants.UI_CMD_BACKGROUND_START_JOB: {

                View parent = getLayoutInflater().inflate(R.layout.activity_job_board, null);
                ViewGroup view = (ViewGroup) findViewById(android.R.id.content);
                FragmentOperatorCodeBinding fragmentOperatorCodeBinding = (FragmentOperatorCodeBinding) DataBindingUtil.inflate(LayoutInflater.from(JobBoardActivity.this),
                        R.layout.fragment_operator_code, null, false);


                fragmentOperatorCodeBinding.setClickHandler(new OperatorCodeInputClickHandler());

                OperatorCodeViewModel operatorCodeViewModel = new OperatorCodeViewModel();
                operatorCodeViewModel.setJobID((int) requestData.Data);
                fragmentOperatorCodeBinding.setVM(operatorCodeViewModel);
                fragmentOperatorCodeBinding.setUIViewActionRequest(requestData);

                AlertDialog.Builder builder = new AlertDialog.Builder(this);
                builder.setView(fragmentOperatorCodeBinding.getRoot());
                builder.setTitle(R.string.title_frag_operator_code);
                _dialog = builder.create();
                _dialog.getWindow().setBackgroundDrawableResource(R.color.color_main_bg);
                _dialog.show();
                //_dialog.getWindow().setLayout(R.dimen.frag_op_code_window_width, WindowManager.LayoutParams.WRAP_CONTENT);
                _dialog.getWindow().setLayout(380, WindowManager.LayoutParams.WRAP_CONTENT);
            }
            break;
            case CommandConstants.UI_CMD_BACKGROUND_UPDATE_JOB_UNIT: {

                View parent = getLayoutInflater().inflate(R.layout.activity_job_board, null);
                ViewGroup view = (ViewGroup) findViewById(android.R.id.content);
                FragmentJobQuantityUpdateBinding fragmentJobQuantityUpdateBinding = (FragmentJobQuantityUpdateBinding) DataBindingUtil.inflate(LayoutInflater.from(JobBoardActivity.this),
                        R.layout.fragment_job_quantity_update, null, false);

                fragmentJobQuantityUpdateBinding.setClickHandler(new JobQuantityUpdateHandler());

                JobQuantityViewModel jobQuantityViewModel = new JobQuantityViewModel();
                jobQuantityViewModel.setJobID((int) requestData.Data);
                fragmentJobQuantityUpdateBinding.setVM(jobQuantityViewModel);

                AlertDialog.Builder builder = new AlertDialog.Builder(this);
                builder.setView(fragmentJobQuantityUpdateBinding.getRoot());
                builder.setTitle(R.string.title_frag_job_quantity_update);
                //builder.create().getWindow().setLayout(200,200);

                _dialog = builder.create();

                _dialog.getWindow().setBackgroundDrawableResource(R.color.color_main_bg);
                _dialog.show();
                _dialog.getWindow().setLayout(460, 320);
            }
            break;
            case CommandConstants.UI_CMD_BACKGROUND_CLOSE_DIALOG: {
                if (_dialog != null) {
                    _dialog.dismiss();
                } else {
                    AppLogger.getInstance().Log(AppLogger.LogLevel.Debug, "_dialog is null to dismiss");
                }
            }
            break;
        }
    }

    @Override
    public void OnUIUpdateEvent(final AppNotificationModelBase model) {

        switch (model.DataProcessStatergyID) {
            case ApplicationConstants.UI_PROCESSING_STATERGY_JOB_STARTED_SUCCESS:
            case ApplicationConstants.UI_PROCESSING_STATERGY_JOB_STARTED_FAILURE:
            case ApplicationConstants.UI_PROCESSING_STATERGY_JOB_STOPED_SUCCESS:
            case ApplicationConstants.UI_PROCESSING_STATERGY_JOB_STOPED_FAILURE:
            case ApplicationConstants.UI_PROCESSING_STATERGY_JOB_QUANTITY_UPDATE_SUCCESS:
            case ApplicationConstants.UI_PROCESSING_STATERGY_JOB_QUANTITY_UPDATE_FAILURE:{
                if(model.Data instanceof  DBJobCardTableRowModel){
                    DBJobCardTableRowModel dbJobCardTableRowModel = (DBJobCardTableRowModel )model.Data;
                    if(dbJobCardTableRowModel.getJobCardId() != _jobBoardViewModel.getJobCardId()){
                        return;
                    }
                }else {
                    return;
                }
            }
            break;
            default:
                return;
        }


        switch (model.DataProcessStatergyID) {
            case ApplicationConstants.UI_PROCESSING_STATERGY_JOB_STARTED_SUCCESS: {
                DBJobCardTableRowModel dbJobCardTableRowModel = (DBJobCardTableRowModel) model.Data;
                _jobBoardViewModel.setDbJobCardTableRowModel(dbJobCardTableRowModel);

                for (JobBriefViewModel jobBriefViewModel : MainViewModel.getInstance().GridViewModel.getItemsSource()) {
                    if(jobBriefViewModel.getDbJobCardTableRowModel().getJobCardId() == dbJobCardTableRowModel.getJobCardId()) {
                        jobBriefViewModel.getDbJobCardTableRowModel().setState(dbJobCardTableRowModel.getState());
                        jobBriefViewModel.getDbJobCardTableRowModel().setActualStartDate(dbJobCardTableRowModel.getActualStartDate());
                        break;
                    }
                }

                _jobBoardViewModel.notifyChange();
                MainViewModel.getInstance().GridViewModel.getDataAdaptor().notifyDataSetChanged();
                MainViewModel.getInstance().setJobStatusTextChanged();
            }
            break;
            case ApplicationConstants.UI_PROCESSING_STATERGY_JOB_STARTED_FAILURE: {
                Toast.makeText(getApplicationContext(), R.string.msg_job_quantity_update_failed, Toast.LENGTH_LONG).show();
            }
            break;
            case ApplicationConstants.UI_PROCESSING_STATERGY_JOB_STOPED_SUCCESS: {
                DBJobCardTableRowModel dbJobCardTableRowModel = (DBJobCardTableRowModel) model.Data;
                JobBriefViewModel itemToRemove = null;
                for (JobBriefViewModel jobBriefViewModel : MainViewModel.getInstance().GridViewModel.getItemsSource()) {

                    if(jobBriefViewModel.getDbJobCardTableRowModel().getJobCardId() == dbJobCardTableRowModel.getJobCardId()) {
                        itemToRemove = jobBriefViewModel;
                        break;
                    }
                }

                if(itemToRemove != null) {
                    //MainViewModel.getInstance().GridViewModel.getItemsSource().remove(itemToRemove);
                    MainViewModel.getInstance().GridViewModel.getDataAdaptor().notifyDataSetChanged();
                    MainViewModel.getInstance().setJobStatusTextChanged();
                }else{
                    AppLogger.getInstance().Log(AppLogger.LogLevel.Error,"JobBoardActivity: After deleting job itemToRemove is found null");
                }

                //finish();
            }
            break;
            case ApplicationConstants.UI_PROCESSING_STATERGY_JOB_STOPED_FAILURE: {
                Toast.makeText(getApplicationContext(), R.string.msg_job_quantity_update_failed, Toast.LENGTH_LONG).show();
            }
            break;
            case ApplicationConstants.UI_PROCESSING_STATERGY_JOB_QUANTITY_UPDATE_SUCCESS: {
                DBJobCardTableRowModel dbJobCardTableRowModel = (DBJobCardTableRowModel) model.Data;
                _jobBoardViewModel.getDbJobCardTableRowModel().setCompletedCount(dbJobCardTableRowModel.getCompletedCount());

                for (JobBriefViewModel jobBriefViewModel : MainViewModel.getInstance().GridViewModel.getItemsSource()) {
                    if(jobBriefViewModel.getDbJobCardTableRowModel().getJobCardId() == dbJobCardTableRowModel.getJobCardId()) {
                        jobBriefViewModel.getDbJobCardTableRowModel().setCompletedCount(dbJobCardTableRowModel.getCompletedCount());
                        break;
                    }
                }

                _jobBoardViewModel.notifyChange();
                MainViewModel.getInstance().GridViewModel.getDataAdaptor().notifyDataSetChanged();
            }
            break;
            case ApplicationConstants.UI_PROCESSING_STATERGY_JOB_QUANTITY_UPDATE_FAILURE: {
                Toast.makeText(getApplicationContext(), R.string.msg_job_quantity_update_failed, Toast.LENGTH_LONG).show();
            }
            break;
        }
    }


    public void OnJobStart(View view, JobBoardViewModel vm) {
        AlertDialog alertDialog = new AlertDialog.Builder(this).create();
        alertDialog.setTitle("Start Job");
        alertDialog.setMessage("Do you want to start this job?");
        alertDialog.setIcon(R.drawable.svg_start);
        alertDialog.setButton(Dialog.BUTTON_POSITIVE, "Yes", new JobStartClickHandler(_jobBoardViewModel));
        alertDialog.setButton(Dialog.BUTTON_NEGATIVE,"No",new JobStartClickHandler(_jobBoardViewModel));
        alertDialog.show();
    }

    public void onJobStop(View view, JobBoardViewModel vm) {
        AlertDialog alertDialog = new AlertDialog.Builder(this).create();
        alertDialog.setTitle("Complete Job");
        alertDialog.setMessage("Do you want to complete this job?");
        alertDialog.setIcon(R.drawable.svg_stop);
        alertDialog.setButton(Dialog.BUTTON_POSITIVE, "Yes", new JobStopClickHandler(_jobBoardViewModel));
        alertDialog.setButton(Dialog.BUTTON_NEGATIVE,"No",new JobStopClickHandler(_jobBoardViewModel));
        alertDialog.show();
    }

    public void OnJobAbort(View view, JobBoardViewModel vm){
        AlertDialog alertDialog = new AlertDialog.Builder(this).create();
        alertDialog.setTitle("Abort Job");
        alertDialog.setMessage("Do you want to abort this job?");
        alertDialog.setIcon(R.drawable.svg_abort);
        alertDialog.setButton(Dialog.BUTTON_POSITIVE, "Yes", new JobAbortClickHandler(_jobBoardViewModel));
        alertDialog.setButton(Dialog.BUTTON_NEGATIVE,"No",new JobAbortClickHandler(_jobBoardViewModel));
        alertDialog.show();
    }

    public void onJobDrop(View view, JobBoardViewModel vm){
        AlertDialog alertDialog = new AlertDialog.Builder(this).create();
        alertDialog.setTitle("Drop Job");
        alertDialog.setMessage("Do you want to drop this job?");
        alertDialog.setIcon(R.drawable.svg_drop);
        alertDialog.setButton(Dialog.BUTTON_POSITIVE, "Yes", new JobDropClickHandler(_jobBoardViewModel));
        alertDialog.setButton(Dialog.BUTTON_NEGATIVE,"No",new JobDropClickHandler(_jobBoardViewModel));
        alertDialog.show();
    }


}
