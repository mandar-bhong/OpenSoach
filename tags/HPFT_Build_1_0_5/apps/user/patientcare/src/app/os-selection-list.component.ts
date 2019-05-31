import { Component, EventEmitter, Input, OnInit, Output, ViewChild, ElementRef } from '@angular/core';
import { RadListViewComponent } from 'nativescript-ui-listview/angular/listview-directives';
import { PlatformHelper } from './helpers/platform-helper';
import { screen } from "tns-core-modules/platform/platform";
import { ItemDataViewModal } from './models/ui/os-selection-list-model';
export enum SELECTION_TYPE {
    SINGLE = 0,
    MULTIPLE = 1
}
@Component({
    moduleId: module.id,
    selector: "os-selection-list",
    template: `
    <RadListView  #customListView tkExampleTitle tkToggleNavButton [items]="listItems" style="background-color: #e0e0e0; padding:3px 3px">
      <ng-template tkListItemTemplate let-data="item" let-i="index">
        <GridLayout columns="auto,auto" rows="45" class="list-group-item m-2 active item-background" style="padding: 0 4px">
            <Label row="0" col="0" class="homeicon mdi icon" (tap)="itemTap(data,i)"
                [text]="data.isselected ? '&#xf133;':'&#xf12f;'"
                [style.color]="data.isselected ? '#FF8910' : '#e0e0e0'"></Label>
            <Label  row="0" col="1" [text]="data.displayText" (tap)="itemTap(data,i)" width="{{ItemWidth}}"  class="icon_value"></Label>
        </GridLayout>
      </ng-template>
    <ListViewLinearLayout tkListViewLayout scrollDirection="Horizontal" ></ListViewLinearLayout>
    </RadListView >
    `,
    styles: [
        `       
        .icon {
            vertical-align: center;
            margin-right: 10px;
        }
        
        .icon_value {
            color: black;
            display: inline;
            font-size: 14px;
            vertical-align: center;
            margin-right: 10px;
        }
        
        .item-background {
            background: white;
            padding: 10px;
            widows: auto;
            padding-right: 2px;
            border: 1px solid gray;
        }
        `
    ]

})

export class OsSelectionListComponent implements OnInit {

    @Input() SelctionMode: number;
    @Input() SelectedItems: any[];
    @Output() Checked = new EventEmitter();
    @Output() Unchecked = new EventEmitter();
    @Output() Tap = new EventEmitter();
    @Input('DisplayText') DisplayText: (item: any) => string;

    public Items: any[];
    public ItemWidth: number;

    public listItems: ItemDataViewModal[] = [];

    currentIndex: number;
    scrollToItemIndex = 0;

    @ViewChild('customListView') catListViewComponent: RadListViewComponent;
    constructor() {
        this.SelectedItems = [];
        this.Items = [];
        this.currentIndex = 0;
        this.scrollToItemIndex = 0;
    }


    ngOnInit() {
        this.Items = [];
    }

    public AddItem(element: any) {
        let vm = new ItemDataViewModal();
        vm.id = PlatformHelper.API.getRandomUUID();
        vm.item = element;
        vm.displayText = this.DisplayText(element);
        if (this.SelectedItems) {
            if (this.SelectedItems.length > 0) {
                const isSelected = this.SelectedItems.filter(data => data === element)[0] || null;
                if (isSelected) {
                    vm.isselected = true;
                } else {
                    vm.isselected = false
                }
            } else {
                vm.isselected = false
            }
        }
        this.listItems.push(vm);
        this.setWidth(screen.mainScreen.widthDIPs, this.Items.length);
    }



    public Init() {
        // console.log('selectedItems --->', this.SelectedItems)
        // console.log('items get for control ------->', this.Items);
        // console.log('SELECTION_TYPE ----', this.SelctionMode);
        // console.log('ItemWidth ----', this.ItemWidth);
        this.Items.forEach(element => {
            let vm = new ItemDataViewModal();
            vm.id = PlatformHelper.API.getRandomUUID();
            vm.item = element;
            vm.index = this.currentIndex;
            vm.displayText = this.DisplayText(element);
            if (this.SelectedItems.length > 0) {
                const isSelected = this.SelectedItems.filter(data => data === element)[0] || null;
                if (isSelected) {
                    vm.isselected = true;
                    this.scrollToItemIndex = vm.index;
                } else {
                    vm.isselected = false
                }
            } else {
                vm.isselected = false
            }
            this.listItems.push(vm);
            this.setWidth(screen.mainScreen.widthDIPs, this.Items.length);

            this.currentIndex++;
            // console.log('listItems  ----->', this.listItems);
        });

        setTimeout(() => {
            this.scrollToIndex(this.scrollToItemIndex);
        }, 300);


    }

    itemTap(selectedRec: ItemDataViewModal, index: number) {
        this.scrollToIndex(index);
        switch (this.SelctionMode) {
            case SELECTION_TYPE.SINGLE:
                if (selectedRec.isselected == true) {
                    return;
                }

                this.listItems.forEach((dataItem) => {
                    if (dataItem.id != selectedRec.id) {
                        if (dataItem.isselected == true) {
                            dataItem.isselected = false;
                            this.Unchecked.emit(dataItem.item);
                        }
                        dataItem.isselected = false;
                    }
                });
                selectedRec.isselected = !selectedRec.isselected;
                if (selectedRec.isselected) {
                    this.Checked.emit(selectedRec.item);
                } else {
                    this.Unchecked.emit(selectedRec.item);
                }
                this.SelectedItems = [];
                this.SelectedItems.push(selectedRec.item);
                break;
            case SELECTION_TYPE.MULTIPLE:
                if (selectedRec.isselected) {
                    const selectedItem = this.SelectedItems.indexOf(item => item.id == selectedRec.id);
                    if (selectedItem) {
                        this.SelectedItems.splice(selectedItem, 1);
                    }
                } else {
                    this.SelectedItems.push(selectedRec.item);
                }
                selectedRec.isselected = !selectedRec.isselected;
                break;
        }
    }
    scrollToIndex(index: number) {
        this.catListViewComponent.listView.scrollToIndex(index);
    }
    setWidth(deviceWidth: number, numberOfItem: number) {
        if (numberOfItem > 0 && numberOfItem <= 3) {
            this.ItemWidth = (deviceWidth - 110) / numberOfItem;
        }
    }
}
