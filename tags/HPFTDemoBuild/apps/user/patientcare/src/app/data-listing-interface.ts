import { ObservableArray } from "tns-core-modules/data/observable-array";

export interface DataListingInterface<T>
{
    listSource: Array<T>;
    listItems:ObservableArray<T>;

    // Fetch data from service and push in listSource
    getData();

    // Bind data to listItems filtered by search 
    bindList();

    // check if data recieved exist in listSource, if yes update, else add
    // check if search is applied, if yes do not add item to listItems else add
    onDataReceived(items:T[]);


}