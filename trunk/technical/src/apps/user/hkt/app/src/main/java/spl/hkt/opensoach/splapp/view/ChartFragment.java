package spl.hkt.opensoach.splapp.view;

import android.app.Fragment;
import android.os.Bundle;
import android.view.Gravity;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TableLayout;
import android.widget.TableRow;
import android.widget.TextView;
import android.widget.Toast;

import spl.hkt.opensoach.splapp.R;
import spl.hkt.opensoach.splapp.model.db.DBChartTableRowModel;

/**
 * Created by samir.s.bukkawar on 3/5/2017.
 */

public class ChartFragment extends Fragment {

    TableLayout mChartTableLayout;

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container, Bundle savedInstanceState) {

        View view = inflater.inflate(R.layout.chart_fragment, container, false);
        TextView tv = (TextView) view.findViewById(R.id.fragmentTextView);
        mChartTableLayout = (TableLayout) view.findViewById(R.id.chartTableLayout);
        mChartTableLayout.removeAllViews();

        DBChartTableRowModel model = new DBChartTableRowModel();
        int taskCount = 5;
        int slotInterval = 10;
        int rowCount = taskCount;
        int colCount = slotInterval + 1;

        //Add Titles
        mChartTableLayout.addView(getChartTitle(colCount));

        for (int i = 0; i < rowCount; i++) {
            TableRow tr = new TableRow(getActivity());
            tr.setLayoutParams(new TableLayout.LayoutParams(TableLayout.LayoutParams.WRAP_CONTENT, TableLayout.LayoutParams.WRAP_CONTENT));
            //  tr.setLayoutParams(new TableLayout.LayoutParams(TableLayout.LayoutParams.FILL_PARENT, TableLayout.LayoutParams.FILL_PARENT));
            tr.setWeightSum(colCount);
            //TableRow rowView = (TableRow)inflater.inflate(R.layout.my_row, null);
            //model.setTasksListId(i);

            //Added +1 in colCount for first column having task name
            for (int j = 0; j < colCount; j++) {

                //model.setStartTime(j);
                final TextView tv1 = new TextView(getActivity());

                if (j < 1) {
                    //tv1.setBackgroundColor(getResources().getColor(R.color.colorAccent));
                    tv1.setBackgroundResource(R.drawable.custom_cell_col_title);
                    tv1.setText("Task " + (i + 1));
                    tv1.setGravity(Gravity.CENTER);
                    tv1.setTextColor(getResources().getColor(R.color.color_cell_available));
                } else if (j < 4) {
                    //tv1.setBackgroundColor(getResources().getColor(R.color.colorAccent));
                    tv1.setBackgroundResource(R.drawable.custom_cell_available_complete);
                } else if (j < 7) {
                    tv1.setBackgroundResource(R.drawable.custom_cell_available);
                    // tv1.setBackgroundColor(getResources().getColor(R.color.color_table_bg));

                    tv1.setOnClickListener(new View.OnClickListener() {
                        @Override
                        public void onClick(View view) {

                            Toast.makeText(getActivity(), " >> " + tv1.getText(), Toast.LENGTH_SHORT).show();
                        }
                    });
                    //tv1.setBackgroundResource(null);
                } else {
                    tv1.setBackgroundResource(R.drawable.custom_cell_not_available);
                }

                // tv1.setPadding(2, 2, 2, 2);

                tv1.setHeight(30);
                // tv1.setWidth(60);
                // tv1.setMinimumWidth(20);

                tv1.setTag(model);


                tv1.setLayoutParams(new TableRow.LayoutParams(0, TableRow.LayoutParams.WRAP_CONTENT, 1f));
                tr.addView(tv1);
            }
            //mChartTableLayout.addView(tr);
            mChartTableLayout.addView(tr);
        }

        return view;
    }

    private TableRow getChartTitle(int colCount) {

        TableRow tr = new TableRow(getActivity());
        tr.setLayoutParams(new TableLayout.LayoutParams(TableLayout.LayoutParams.WRAP_CONTENT, TableLayout.LayoutParams.WRAP_CONTENT));
        tr.setWeightSum(colCount);
        for (int j = 0; j < colCount; j++) {

            final TextView tv1 = new TextView(getActivity());
            tv1.setHeight(30);
            tv1.setLayoutParams(new TableRow.LayoutParams(0, TableRow.LayoutParams.WRAP_CONTENT, 1f));
            if (j == 0) {
                tv1.setText("");
//                tv1.setBackgroundResource(R.drawable.custom_cell_col_title);
                tv1.setBackgroundColor(getResources().getColor(R.color.color_cell_available));
            } else {
                tv1.setText("" + j);
                tv1.setTextColor(getResources().getColor(R.color.color_cell_available));
                tv1.setBackgroundResource(R.drawable.custom_cell_col_title);
            }
            tv1.setGravity(Gravity.CENTER);
            tr.addView(tv1);

        }
        return tr;
    }

}