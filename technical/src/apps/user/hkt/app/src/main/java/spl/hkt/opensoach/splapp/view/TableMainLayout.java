package spl.hkt.opensoach.splapp.view;

import android.app.Activity;
import android.content.Context;
import android.graphics.Color;
import android.util.Log;
import android.view.Gravity;
import android.view.LayoutInflater;
import android.view.View;
import android.widget.HorizontalScrollView;
import android.widget.RelativeLayout;
import android.widget.ScrollView;
import android.widget.TableLayout;
import android.widget.TableRow;
import android.widget.TextView;

import java.text.Format;
import java.text.SimpleDateFormat;
import java.util.Calendar;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;

import spl.hkt.opensoach.splapp.R;
import spl.hkt.opensoach.splapp.SPLApplication;
import spl.hkt.opensoach.splapp.logger.AppLogger;
import spl.hkt.opensoach.splapp.model.ChartDataModel;
import spl.hkt.opensoach.splapp.model.view.ChartConfigModel;
import spl.hkt.opensoach.splapp.model.view.ChartConfigSlotModel;
import spl.hkt.opensoach.splapp.model.view.ChartConfigTaskModel;
import spl.hkt.opensoach.splapp.model.view.DisplayChartDataModel;
import spl.hkt.opensoach.splapp.model.view.DisplayChartItemDataModel;
import spl.hkt.opensoach.splapp.viewModels.CellViewModel;
import spl.hkt.opensoach.splapp.viewModels.ChartViewModel;
import spl.hkt.opensoach.splapp.viewModels.MainViewModel;
import spl.hkt.opensoach.splapp.viewModels.TaskRowViewModel;

public class TableMainLayout extends RelativeLayout implements TimeChangeListner, ITableClick {

    public final String TAG = "TableMainLayout";

    //region Private Veriabls
    private final int cellHeight = (int) getResources().getDimension(R.dimen.chart_cellview_height);
    private final int cellWidth = (int) getResources().getDimension(R.dimen.chart_cellview_width);

    private final int cellMargin = (int) getResources().getDimension(R.dimen.chart_cellview_margin);

    private int currentActiveSlot = -1;

    private SPLApplication mSPLApplication;
    private ITableClick iTableClick;
    private ChartConfigModel chartConfigModel;

    private TableLayout tableA; // Time Slot Text table
    private TableLayout tableB; // Chart Slot Menu table
    private TableLayout tableC; // Chart Task Menu table
    private TableLayout tableD; // Chart Data Cell Table

    private HashMap<String, CellViewModel> cellRefVMLookup = new HashMap<String, CellViewModel>();

    private HorizontalScrollView horizontalScrollViewB;
    private HorizontalScrollView horizontalScrollViewD;

    private ScrollView scrollViewC;
    private ScrollView scrollViewD;

    private Context context;

    public ChartConfigModel getChartDataModel() {
        return chartConfigModel;
    }

    private ChartViewModel mChartViewModel;
    //endregion


    //region Constructor
    public TableMainLayout(Context context, ChartViewModel chartViewModel) {

        super(context);
        this.context = context;

        mSPLApplication = SPLApplication.getInstance();
        mSPLApplication.registerTimeChangeListner(this, (Activity) context);

        mChartViewModel = new ChartViewModel();
        mChartViewModel.setiTableClick(this);


        // initialize the main components (TableLayouts, HorizontalScrollView, ScrollView)
        this.initComponents();
        this.setComponentsId();
        this.setScrollViewAndHorizontalScrollViewTag();

        // no need to assemble component A, since it is just a table
        this.horizontalScrollViewB.addView(this.tableB);
        this.horizontalScrollViewB.setHorizontalScrollBarEnabled(false);

        this.scrollViewC.addView(this.tableC);
        scrollViewC.setVerticalScrollBarEnabled(false);

        this.scrollViewD.addView(this.horizontalScrollViewD);
        this.horizontalScrollViewD.addView(this.tableD);
        this.horizontalScrollViewD.setHorizontalScrollBarEnabled(false);

        // add the components to be part of the main layout
        this.addComponentToMainLayout();
        this.setBackgroundColor(Color.WHITE);

        //Time Slot Template Cell Item
        TableRow componentATableRow = new TableRow(this.context);
        TextView textView = this.getTitleTextView(getResources().getString(R.string.txt_time_slot));
        textView.setWidth(cellWidth);
        textView.setHeight(cellHeight);
        textView.setTextSize(getResources().getDimension(R.dimen.chart_cell_font_size));

        TableRow.LayoutParams params = new TableRow.LayoutParams(cellWidth, cellHeight);
        params.setMargins(0, 0, cellMargin, cellMargin);

        componentATableRow.addView(textView, params);

        this.tableA.addView(componentATableRow);


        //for Initial Scroll
/*
        horizontalScrollViewB.post(new Runnable() {
            @Override
            public void run() {
                Log.i("####", "onCreate : " + horizontalScrollViewB.getMaxScrollAmount());
                horizontalScrollViewD.smoothScrollTo(horizontalScrollViewD.getMaxScrollAmount(), 0);
            }
        });*/
    }
    //endregion


    //region Private Methods
    // initalized components
    private void initComponents() {

        this.tableA = new TableLayout(this.context);
        this.tableB = new TableLayout(this.context);
        this.tableC = new TableLayout(this.context);
        this.tableD = new TableLayout(this.context);

        this.horizontalScrollViewB = new ChartDataHorizontalScrollView(this.context);
        this.horizontalScrollViewD = new ChartDataHorizontalScrollView(this.context);
        this.scrollViewC = new ChartDataVerticalScrollView(this.context);
        this.scrollViewD = new ChartDataVerticalScrollView(this.context);

        this.tableA.setBackgroundColor(getResources().getColor(R.color.color_chart_bg_for_border));
        this.tableB.setBackgroundColor(getResources().getColor(R.color.color_chart_bg_for_border));
        this.tableC.setBackgroundColor(getResources().getColor(R.color.color_chart_bg_for_border));
        this.tableD.setBackgroundColor(getResources().getColor(R.color.color_chart_bg_for_border));

        this.horizontalScrollViewB.setBackgroundColor(Color.WHITE);
        this.horizontalScrollViewD.setBackgroundColor(Color.WHITE);
        this.scrollViewC.setBackgroundColor(Color.WHITE);
        this.scrollViewD.setBackgroundColor(Color.WHITE);
    }

    // set essential component IDs
    private void setComponentsId() {
        this.tableA.setId(1);
        this.horizontalScrollViewB.setId(2);
        this.scrollViewC.setId(3);
        this.scrollViewD.setId(4);
    }

    // set tags for some horizontal and vertical scroll view
    private void setScrollViewAndHorizontalScrollViewTag() {

        this.horizontalScrollViewB.setTag("horizontal scroll view b");
        this.horizontalScrollViewD.setTag("horizontal scroll view d");

        this.scrollViewC.setTag("scroll view c");
        this.scrollViewD.setTag("scroll view d");
    }

    // we add the components here in our TableMainLayout
    private void addComponentToMainLayout() {

        // RelativeLayout params were very useful here
        // the addRule method is the key to arrange the components properly
        RelativeLayout.LayoutParams componentB_Params = new RelativeLayout.LayoutParams(LayoutParams.WRAP_CONTENT, LayoutParams.WRAP_CONTENT);
        componentB_Params.addRule(RelativeLayout.RIGHT_OF, this.tableA.getId());

        RelativeLayout.LayoutParams componentC_Params = new RelativeLayout.LayoutParams(LayoutParams.WRAP_CONTENT, LayoutParams.WRAP_CONTENT);
        componentC_Params.addRule(RelativeLayout.BELOW, this.tableA.getId());

        RelativeLayout.LayoutParams componentD_Params = new RelativeLayout.LayoutParams(LayoutParams.WRAP_CONTENT, LayoutParams.WRAP_CONTENT);
        componentD_Params.addRule(RelativeLayout.RIGHT_OF, this.scrollViewC.getId());
        componentD_Params.addRule(RelativeLayout.BELOW, this.horizontalScrollViewB.getId());

        // 'this' is a relative layout,
        // we extend this table layout as relative layout as seen during the creation of this class
        this.addView(this.tableA);
        this.addView(this.horizontalScrollViewB, componentB_Params);
        this.addView(this.scrollViewC, componentC_Params);
        this.addView(this.scrollViewD, componentD_Params);

    }

    // table cell standard TextView
    TextView getTitleTextView(String label) {

        TextView bodyTextView = new TextView(this.context);
        bodyTextView.setText(label);
        bodyTextView.setGravity(Gravity.CENTER);
        bodyTextView.setBackgroundColor(getResources().getColor(R.color.color_table_head_bg));
        bodyTextView.setTextColor(getResources().getColor(R.color.color_table_head_text));

        return bodyTextView;
    }
    //endregion


    //region Chart Methods
    public void setChart(ChartConfigModel model) {

        try {

            clearPreviousChart();

            chartConfigModel = model;

            currentActiveSlot = identifyCurrentActiveSlot(model);

            this.createChartSlotMenu(model);
            this.createChartTaskMenu(model);
            this.generateChartDataCell(currentActiveSlot, model);

            refreshChartTables();


        } catch (Exception ex) {
            Log.d("TableMainLayout", ex.getMessage());
        }
    }

    public void setChartData(DisplayChartDataModel model) {

        try {

            for (DisplayChartItemDataModel dataItem : model.getChartData()) {

                if (cellRefVMLookup.containsKey(dataItem.getTaskName() + "_" + dataItem.getChartId())) {
                    CellViewModel cellVM = cellRefVMLookup.get(dataItem.getTaskName() + "_" + dataItem.getSlotId());
                    cellVM.setState(dataItem.getState());
                }
            }


            refreshChartTables();


        } catch (Exception ex) {
            Log.d("TableMainLayout", ex.getMessage());
        }
    }

    private void generateChartDataCell(int activeSlot, ChartConfigModel model) {

        for (Map.Entry<String, ChartConfigTaskModel> chartTaskKV : model.getTasks().entrySet()) {
            try {
                TaskRowViewModel taskRowViewModel = new TaskRowViewModel();
                taskRowViewModel.setTaskName(model.getTasks().get(chartTaskKV.getKey()).getTaskName());

                taskRowViewModel.setiRowClick(mChartViewModel);
                mChartViewModel.getTaskRowViewModelList().add(taskRowViewModel);

                TableRow tableRow = new TableRow(this.context);
                TableRow.LayoutParams params = new TableRow.LayoutParams(cellWidth, cellHeight);
                params.setMargins(0, 0, cellMargin, cellMargin);

                for (Integer slotIndex = 0; slotIndex < model.getSlots().size(); slotIndex++) {
                    View cellView = CreateChartDataCell(activeSlot, taskRowViewModel, slotIndex);
                    tableRow.addView(cellView, params);
                }

                //TableRow taleRowForTableD = this.taleRowForTableD(taskRowViewModel, taskRowIndex);
                this.tableD.addView(tableRow);

            } catch (Exception ex) {

            }
        }
    }

    private View CreateChartDataCell(int activeSlot, TaskRowViewModel parent, int slotID) {

        View checkBox = new View(this.context);

        if (activeSlot < slotID) {
            checkBox.setBackgroundResource(R.drawable.custom_cell_not_available);
            checkBox.setEnabled(false);
        } else {
            checkBox.setBackgroundResource(R.drawable.custom_cell_available);
        }

        CellViewModel cellViewModel = new CellViewModel();
        cellViewModel.setTaskName(parent.getTaskName());
        cellViewModel.setSlotID(slotID);
        cellViewModel.setiCellClick(parent);
        cellViewModel.setCheckBox(checkBox);

        checkBox.setOnClickListener(cellViewModel);

        cellRefVMLookup.put(parent.getTaskName() + "_" + slotID, cellViewModel);

        return checkBox;
    }

    private void createChartTaskMenu(ChartConfigModel model) {

        for (ChartConfigTaskModel chartTaskModel : model.getTaskList()) {
            View taskTemplateCell = createChartTaskCell(chartTaskModel.getTaskName());
            taskTemplateCell.setMinimumHeight(cellHeight);
            this.tableC.addView(taskTemplateCell);
        }
    }

    private View createChartTaskCell(String taskName) {

        TableRow.LayoutParams params = new TableRow.LayoutParams(cellWidth, cellHeight);
        params.setMargins(0, 0, cellMargin, cellMargin);

        TableRow tableRowForTableC = new TableRow(this.context);
        LayoutInflater mInflater = ((Activity) this.context).getLayoutInflater();

        View convertView = mInflater.inflate(R.layout.custom_cell_view, null);
        ViewHolder holder = new ViewHolder();
        holder.textView = (TextView) convertView.findViewById(R.id.textView);
        holder.textView.setTextSize(getResources().getDimension(R.dimen.chart_cell_font_size));
        holder.textView.setText(taskName);
        holder.imageView = (View) convertView.findViewById(R.id.imageView);
        holder.imageView.setVisibility(View.GONE);

        tableRowForTableC.addView(convertView, params);
        return tableRowForTableC;
    }


    private void createChartSlotMenu(ChartConfigModel model) {

        TableRow componentBTableRow = new TableRow(this.context);
        TableRow.LayoutParams params = new TableRow.LayoutParams(LayoutParams.WRAP_CONTENT, LayoutParams.MATCH_PARENT);
        params.setMargins(0, 0, cellMargin, cellMargin);

        for (int i = 0; i < model.getSlots().size(); i++) {
            Format formatter = new SimpleDateFormat("hh:mm a");
            String slotDisplayText = formatter.format(model.getSlots().get(i).getStartTime());

            TextView textView = this.getTitleTextView(slotDisplayText);
            textView.setHeight(cellHeight);
            textView.setWidth(cellWidth);
            textView.setTextSize(getResources().getDimension(R.dimen.chart_cell_font_size));

            textView.setLayoutParams(params);
            componentBTableRow.addView(textView);
        }

        this.tableB.addView(componentBTableRow);
    }

    private int identifyCurrentActiveSlot(ChartConfigModel model) {

        int activeSlot = -1;
        Date currentDateTime = new Date();

        if (currentDateTime.getTime() < model.getSlots().get(0).getStartTime().getTime()) {
            return activeSlot;
        }

        if (currentDateTime.getTime() > model.getSlots().get(model.getSlots().size() - 1).getEndTime().getTime()) {
            activeSlot = model.getSlots().size() - 1;
            return activeSlot;
        }

        for (int i = 0; i < model.getSlots().size(); i++) {
            ChartConfigSlotModel configSlotModel = model.getSlots().get(i);

            if (currentDateTime.getTime() > configSlotModel.getStartTime().getTime() &&
                    currentDateTime.getTime() < configSlotModel.getEndTime().getTime()) {
                activeSlot = i;
                break;
            }

        }

        return activeSlot;
    }

    private void refreshChartTables() {
        tableB.invalidate();
        tableB.refreshDrawableState();
        tableC.invalidate();
        tableC.refreshDrawableState();
        tableD.invalidate();
        tableD.refreshDrawableState();
    }

    private void clearPreviousChart() {
        this.tableB.removeAllViews();
        this.tableC.removeAllViews();
        this.tableD.removeAllViews();
    }
    //endregion


    //region Overridden Methods
    @Override
    public void notifyTimeChange() {

        if (chartConfigModel == null) {
            return;
        }

        if (chartConfigModel.getSlots().size() <= currentActiveSlot + 1) {
            Date firstTimeSlot = chartConfigModel.getSlots().get(0).getStartTime();

            Calendar ca = Calendar.getInstance();
            ca.setTime(firstTimeSlot);
            ca.add(Calendar.HOUR_OF_DAY, 23);

            if (ca.getTime().getTime() < (new Date()).getTime()) {

                for (Map.Entry<Integer, ChartConfigSlotModel> chartSlotModelKVP : chartConfigModel.getSlots().entrySet()) {
                    Date chartStartTime = chartSlotModelKVP.getValue().getStartTime();
                    Date tomorrowStartTime = new Date(chartStartTime.getTime() + (1000 * 60 * 60 * 24));
                    chartSlotModelKVP.getValue().setStartTime(tomorrowStartTime);

                    Date chartEndTime = chartSlotModelKVP.getValue().getEndTime();
                    Date tomorrowEndTime = new Date(chartEndTime.getTime() + (1000 * 60 * 60 * 24));
                    chartSlotModelKVP.getValue().setStartTime(tomorrowEndTime);

                }

                setChart(chartConfigModel);

            }

            return;
        }

        Integer nextActiveSlot = currentActiveSlot + 1;

        Date nextSlotStartTime = chartConfigModel.getSlots().get(nextActiveSlot).getStartTime();

        if (nextSlotStartTime.getTime() > (new Date()).getTime()) {
            return;
        }


        for (String taskName : chartConfigModel.getTasks().keySet()) {
            if (cellRefVMLookup.containsKey(taskName + "_" + nextActiveSlot)) {
                cellRefVMLookup.get(taskName + "_" + nextActiveSlot).getCheckBox().setBackgroundResource(R.drawable.custom_cell_available);
                cellRefVMLookup.get(taskName + "_" + nextActiveSlot).getCheckBox().setEnabled(true);
            }
        }

        currentActiveSlot = nextActiveSlot;

        Log.i("####", "notifyTimeChange" + new Date());
    }

    @Override
    public ChartViewModel getChartViewModel() {
        return mChartViewModel;
    }

    @Override
    public void onChartTableClick(ChartViewModel chartViewModel) {
        CellViewModel cellViewModel = chartViewModel.getTaskRowViewModel().getmCellViewModel();

        ChartConfigSlotModel chartConfigSlotModel = chartConfigModel.getSlots().get(cellViewModel.getSlotID());
        ChartConfigTaskModel chartConfigTaskModel = chartConfigModel.getTasks().get(cellViewModel.getTaskID());

        ChartDataModel chartDM = new ChartDataModel();
        chartDM.setChartId(chartConfigModel.getChartId());
        chartDM.setTaskName(chartConfigTaskModel.getTaskName());
        chartDM.setSlotId(chartConfigSlotModel.getIndex());
        //chartDM.setAuthCode();//This code will set later
        chartDM.setEntryDate(new Date());
        chartDM.setSlotStartTime(chartConfigSlotModel.getStartTime());
        chartDM.setSlotEndTime(chartConfigSlotModel.getEndTime());

        HashMap<String, ChartDataModel> currenChartDataModelMap = MainViewModel.getInstance().getCurrenChartDataModelMap();

        if (cellViewModel.isChecked()) {
            currenChartDataModelMap.put(cellViewModel.getTaskID() + "_" + cellViewModel.getSlotID(), chartDM);
            AppLogger.getInstance().Log(AppLogger.LogLevel.Debug, "TableMainLayout :Selected Task ID : " + cellViewModel.getTaskID() + " SlotID " + cellViewModel.getSlotID());
        } else {
            AppLogger.getInstance().Log(AppLogger.LogLevel.Debug, "TableMainLayout :Deselected Task ID : " + cellViewModel.getTaskID() + " SlotID " + cellViewModel.getSlotID());
            currenChartDataModelMap.remove(cellViewModel.getTaskID() + "_" + cellViewModel.getSlotID());
        }
        MainViewModel.getInstance().setCurrenChartDataModelMap(currenChartDataModelMap);
    }
    //endregion


    //region Internal Class
    static class ViewHolder {
        private TextView textView;
        private View imageView;
    }

    // horizontal scroll view custom class
    class ChartDataHorizontalScrollView extends HorizontalScrollView {

        public ChartDataHorizontalScrollView(Context context) {
            super(context);
        }

        @Override
        protected void onScrollChanged(int l, int t, int oldl, int oldt) {
            String tag = (String) this.getTag();


            if (tag.equalsIgnoreCase("horizontal scroll view b")) {
                horizontalScrollViewD.scrollTo(l, 0);
            } else {
                horizontalScrollViewB.scrollTo(l, 0);
            }
        }

    }

    // scroll view custom class
    class ChartDataVerticalScrollView extends ScrollView {

        public ChartDataVerticalScrollView(Context context) {
            super(context);
        }

        @Override
        protected void onScrollChanged(int l, int t, int oldl, int oldt) {

            String tag = (String) this.getTag();

            if (tag.equalsIgnoreCase("scroll view c")) {
                scrollViewD.scrollTo(0, t);
            } else {
                scrollViewC.scrollTo(0, t);
            }
        }
    }
    //endregion

}