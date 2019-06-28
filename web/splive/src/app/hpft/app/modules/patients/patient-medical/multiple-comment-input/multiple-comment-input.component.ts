import { Component, OnInit, Input, EventEmitter, Output } from '@angular/core';

@Component({
  selector: 'app-multiple-comment-input',
  templateUrl: './multiple-comment-input.component.html',
  styleUrls: ['./multiple-comment-input.component.css']
})
export class MultipleCommentInputComponent implements OnInit {

  @Input() itemList: any;
  @Input() placeHolderText: string;
  @Input() headerText: string;
  @Output() onItemAdd = new EventEmitter();
  contextValue: string;
  constructor() {
  }
  ngOnInit() {    
      this.itemList.sort((a, b) => {
        return (new Date(b.date).getTime() - new Date(a.date).getTime())
      });
  }

  // emit event
  itemAdd() {
    if (this.contextValue) {
      this.onItemAdd.emit(this.contextValue);
      this.contextValue = null;
      this.itemList.sort((a, b) => {
        return (new Date(b.date).getTime() - new Date(a.date).getTime())
      });
    }
  }
}