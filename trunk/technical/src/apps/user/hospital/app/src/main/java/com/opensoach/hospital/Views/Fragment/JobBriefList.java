package com.opensoach.hospital.Views.Fragment;

import android.content.Context;
import android.net.Uri;
import android.os.Bundle;
import android.support.v4.app.Fragment;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import com.opensoach.hospital.R;

/**
 * A simple {@link Fragment} subclass.
 * Activities that contain this fragment must implement the
 * {@link JobBriefList.OnFragmentInteractionListener} interface
 * to handle interaction events.
 * Use the {@link JobBriefList#newInstance} factory method to
 * create an instance of this fragment.
 */
public class JobBriefList extends Fragment  {
    // TODO: Rename parameter arguments, choose names that match
//    private JobBriefViewAdaptor jobBriefViewAdaptor;
//    private JobBriefListViewModel dataContext;
//    private Iterable itemsSource;

    private String mParam2;

    private OnFragmentInteractionListener mListener;

    public JobBriefList() {
        // Required empty public constructor
    }



    /**
     * Use this factory method to create a new instance of
     * this fragment using the provided parameters.
     *
     * @param param1 Parameter 1.
     * @param param2 Parameter 2.
     * @return A new instance of fragment JobBriefList.
     */
    // TODO: Rename and change types and number of parameters
    public static JobBriefList newInstance(String param1, String param2) {
        JobBriefList fragment = new JobBriefList();
        Bundle args = new Bundle();
//        args.putString(ARG_PARAM2, param2);
//        fragment.setArguments(args);
        return fragment;
    }

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        //GridView grid = (GridView) this.getActivity().findViewById(R.id.job_grid_view);
        //jobBriefViewAdaptor = new JobBriefViewAdaptor(getActivity(),null);
       // grid.setAdapter(jobBriefViewAdaptor);
    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        // Inflate the layout for this fragment

        JobGridView gd = (JobGridView) getActivity().findViewById(R.id.job_grid_view);

        return inflater.inflate(R.layout.fragment_job_brief_list, container, false);
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

//      @Override
//    public JobBriefListViewModel getDataContext() {
//        return dataContext;
//    }
//
//    @Override
//    public void setDataContext(JobBriefListViewModel viewModel) {
//        dataContext = viewModel;
//    }

//    @Override
//    public Iterable getItemsSource() {
//        return itemsSource;
//    }
//
//    @Override
//    public void setItemsSource(Iterable source) {
//        itemsSource=source;
//    }
}
