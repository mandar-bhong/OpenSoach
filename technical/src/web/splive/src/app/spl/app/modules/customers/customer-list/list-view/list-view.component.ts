import { Component, OnInit } from '@angular/core';
import { CustomerDataListingItemResponse } from '../../../../models/api/customer-models';

export class TestColumn {
  col: string;
  displayname: string;
  breakpoint: string;
  css: string;
}
@Component({
  selector: 'app-list-view',
  templateUrl: './list-view.component.html',
  styleUrls: ['./list-view.component.css']
})
export class ListViewComponent implements OnInit {
  breakpoint = 'Desktop';
  displayedColumns = ['custid', 'corpid', 'custname', 'corpname', 'poc1name', 'poc1emailid', 'poc1mobileno', 'createdon', 'updateon'];
  // customerDataListingItemResponse = [
  //   {custid: 1, corpid: 12, custname: 'customer 1', corpname: 'corp 1', poc1name: 'ram',  },
  //   {custid: 2, corpid: 14, custname: 'customer 2', corpname: 'corp 2', poc1name: 'ram'  }
  // ];
  cols: TestColumn[] = [{ col: 'custid', displayname: 'Customer Id', breakpoint: '', css: '' },
  { col: 'corpid', displayname: 'Corprate Id', breakpoint: 'xs', css: '' },
  { col: 'custname', displayname: 'Customer Name', breakpoint: 'sm', css: '' },
  { col: 'corpname', displayname: 'Corp Name', breakpoint: 'sm', css: '' },
  { col: 'poc1name', displayname: 'Name', breakpoint: 'sm', css: '' },
  { col: 'poc1emailid', displayname: 'Email Id', breakpoint: 'sm', css: '' },
  { col: 'poc1mobileno', displayname: 'Mobile Number', breakpoint: 'sm', css: '' },
  { col: 'createdon', displayname: 'Create', breakpoint: 'sm', css: '' },
  { col: 'updateon', displayname: 'update', breakpoint: '', css: '' }];
  dataSource;
  constructor() { }

  ngOnInit() { // for (let i = 0; i < 10; i++) {
    //   const row = new CustomerDataListingModel();
    //   row.col1 = 'col1 ' + i;
    //   row.col2 = 'col2 ' + i;
    //   row.col3 = 'col3 ' + i;
    //   row.col4 = 'col4 ' + i;
    //   row.col5 = 'col5 ' + i;
    //   row.col6 = 'edit';
    //   this.rows.push(row);
    // }
    // this.dataSource = this.rows;
  }

}
