package spl.hkt.opensoach.splapp.view;

import android.app.Activity;
import android.app.Fragment;
import android.app.FragmentTransaction;
import android.content.Context;
import android.net.Uri;
import android.os.Bundle;
import android.os.Handler;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.RelativeLayout;



import spl.hkt.opensoach.splapp.model.view.ChartConfigModel;
import spl.hkt.opensoach.splapp.model.view.DisplayChartDataModel;
import spl.hkt.opensoach.splapp.viewModels.MainViewModel;

/**
 * A simple {@link Fragment} subclass.
 * Activities that contain this fragment must implement the
 * {@link ChartTableFragment.OnFragmentInteractionListener} interface
 * to handle interaction events.
 * Use the {@link ChartTableFragment#newInstance} factory method to
 * create an instance of this fragment.
 */
public class ChartTableFragment extends Fragment {
    // TODO: Rename parameter arguments, choose names that match
    // the fragment initialization parameters, e.g. ARG_ITEM_NUMBER
    private static final String ARG_PARAM1 = "param1";
    private static final String ARG_PARAM2 = "param2";

    // TODO: Rename and change types of parameters
    private String mParam1;
    private String mParam2;
    private View tableMainLayout;

    private OnFragmentInteractionListener mListener;

    public ChartTableFragment() {
        // Required empty public constructor
        super();
    }

    /**
     * Use this factory method to create a new instance of
     * this fragment using the provided parameters.
     *
     * @param param1 Parameter 1.
     * @param param2 Parameter 2.
     * @return A new instance of fragment ChartTableFragment.
     */
    // TODO: Rename and change types and number of parameters
    public static ChartTableFragment newInstance(String param1, String param2) {
        ChartTableFragment fragment = new ChartTableFragment();
        Bundle args = new Bundle();
        args.putString(ARG_PARAM1, param1);
        args.putString(ARG_PARAM2, param2);
        fragment.setArguments(args);
        return fragment;
    }

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        if (getArguments() != null) {
            mParam1 = getArguments().getString(ARG_PARAM1);
            mParam2 = getArguments().getString(ARG_PARAM2);
        }
    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        // Inflate the layout for this fragment
        //return inflater.inflate(R.layout.fragment_chart_table, container, false);
        tableMainLayout = new TableMainLayout(getActivity(), MainViewModel.getInstance().getChartViewModel());
        return tableMainLayout;
    }

    // TODO: Rename method, update argument and hook method into UI event
    public void onButtonPressed(Uri uri) {
        if (mListener != null) {
            mListener.onFragmentInteraction(uri);
        }
    }

    @Override
    public void onAttach(Context context) {
        super.onAttach(context);
        if (context instanceof OnFragmentInteractionListener) {
            mListener = (OnFragmentInteractionListener) context;
        } else {
            throw new RuntimeException(context.toString()
                    + " must implement OnFragmentInteractionListener");
        }
    }

    @Override
    public void onDetach() {
        super.onDetach();
        mListener = null;
    }

    /**
     * This interface must be implemented by activities that contain this
     * fragment to allow an interaction in this fragment to be communicated
     * to the activity and potentially other fragments contained in that
     * activity.
     * <p>
     * See the Android Training lesson <a href=
     * "http://developer.android.com/training/basics/fragments/communicating.html"
     * >Communicating with Other Fragments</a> for more information.
     */
    public interface OnFragmentInteractionListener {
        // TODO: Update argument type and name
        void onFragmentInteraction(Uri uri);
    }

    public void setChart(Activity executionContext, ChartConfigModel model) {

        Handler hdl = new Handler(executionContext.getMainLooper());
        hdl.post(new Runnable() {
            Activity executionContext;
            ChartConfigModel chartConfigModel;

            public Runnable init(Activity exeContext, ChartConfigModel model) {
                executionContext = exeContext;
                chartConfigModel = model;
                return this;
            }

            @Override
            public void run() {
                ((TableMainLayout)tableMainLayout).setChart( chartConfigModel);
            }
        }.init(executionContext, model));
    }

    public void setChartData(Activity executionContext, DisplayChartDataModel model){

        Handler hdl = new Handler(executionContext.getMainLooper());
        hdl.post(new Runnable() {
            Activity executionContext;
            DisplayChartDataModel displayChartDataModel;

            public Runnable init(Activity exeContext, DisplayChartDataModel model) {
                executionContext = exeContext;
                displayChartDataModel = model;
                return this;
            }

            @Override
            public void run() {
                ((TableMainLayout)tableMainLayout).setChartData( displayChartDataModel);
            }
        }.init(executionContext, model));
    }
}