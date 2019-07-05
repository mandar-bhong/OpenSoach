import { Component, OnInit } from '@angular/core';

export class TestDataRow {
  col1: string;
  col2: string;
  col3: string;
  col4: string;
  col5: string;
  col6: string;
}

export class TestColumn {
  col: string;
  displayname: string;
  breakpoint: string;
  css: string;
}

@Component({
  selector: 'app-default-dashboard',
  templateUrl: './default-dashboard.component.html',
  styleUrls: ['./default-dashboard.component.css']
})
export class DefaultDashboardComponent implements OnInit {
  breakpoint = 'Desktop';
  // displayedColumns = ['col1', 'col2', 'col3', 'col4', 'col5', 'col6'];
  displayedColumns = ['col1', 'col2', 'col3', 'col4', 'col5', 'col6'];
  rows: TestDataRow[] = [];
  // cols: TestColumn[] = [{ col: 'col1', displayname: 'Name', breakpoint: '', css: 'd-block col-md-2 col-sm-4 col-8 col-lg-2' },
  // { col: 'col2', displayname: 'Email', breakpoint: 'xs', css: 'd-none d-sm-block col-sm-4 col-md-2 col-lg-2' },
  // { col: 'col3', displayname: 'Telephone', breakpoint: 'sm', css: 'd-none d-sm-block col-sm-2 col-md-2' },
  // { col: 'col4', displayname: 'Address 1', breakpoint: 'sm', css: 'd-none d-lg-block col-md-2 col-lg-2' },
  // { col: 'col5', displayname: 'Address 2', breakpoint: 'sm', css: 'd-none d-lg-block col-md-2 col-lg-2' },
  // { col: 'col6', displayname: 'Edit', breakpoint: '',  css: 'd-block col-md-2 col-sm-2 col-4 col-lg-2'}];

  cols: TestColumn[] = [{ col: 'col1', displayname: 'Name', breakpoint: '', css: '' },
  { col: 'col2', displayname: 'Email', breakpoint: 'xs', css: '' },
  { col: 'col3', displayname: 'Telephone', breakpoint: 'sm', css: '' },
  { col: 'col4', displayname: 'Address 1', breakpoint: 'sm', css: '' },
  { col: 'col5', displayname: 'Address 2', breakpoint: 'sm', css: '' },
  { col: 'col6', displayname: 'Edit', breakpoint: '',  css: ''}];
  dataSource;
// cols: TestColumn[] = [{ col: 'col1', displayname: 'Name', breakpoint: '' },
// { col: 'col2', displayname: 'Email', breakpoint: 'hidden-xs-down' },
// { col: 'col3', displayname: 'Telephone', breakpoint: 'hidden-sm' },
// { col: 'col4', displayname: 'Address 1', breakpoint: 'hidden-md' },
// { col: 'col5', displayname: 'Address 2', breakpoint: 'hidden-md' },
// { col: 'col6', displayname: 'Edit', breakpoint: '' }];
constructor() { }

ngOnInit() {

  for (let i = 0; i < 10; i++) {
    const row = new TestDataRow();
    row.col1 = 'col1 ' + i;
    row.col2 = 'col2 ' + i;
    row.col3 = 'col3 ' + i;
    row.col4 = 'col4 ' + i;
    row.col5 = 'col5 ' + i;
    row.col6 = 'edit';
    this.rows.push(row);
  }

  this.dataSource = this.rows;
}

}
