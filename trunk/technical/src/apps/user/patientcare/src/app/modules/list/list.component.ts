import { Component, OnInit } from '@angular/core';
import { RouterExtensions } from "nativescript-angular/router";
import { SearchBar } from "tns-core-modules/ui/search-bar";
import { isAndroid } from "platform";
import * as app from "application";
// import { RadSideDrawerComponent, SideDrawerType } from "nativescript-ui-sidedrawer/angular";
// import { ViewChild } from "@angular/core";
// import { RadSideDrawer } from "nativescript-ui-sidedrawer";
@Component({
	moduleId: module.id,
	selector: 'list',
	templateUrl: './list.component.html',
	styleUrls: ['./list.component.css']
})

export class ListComponent implements OnInit {
	data = [];
	searchshow = false;
	searchiocn = true;
	// @ViewChild(RadSideDrawerComponent) public drawerComponent: RadSideDrawerComponent;
	constructor(private routerExtensions: RouterExtensions) { }

	ngOnInit() {
		// this.data.push({ text: "Bulbasaur", src: "" });
		this.data.push({ ward: "3A/312", name: "Sumeet karande", mobile: "9878978980" });
		this.data.push({ ward: "3B/323", name: "Amol Patil", mobile: "32423432432" });
		// this.data.push({ ward: "2A/643", name: "Sanjay Mohan", mobile: "453545352"});
		// this.data.push({ ward: "4A/515", name: "Sanjay Sawant", mobile: "9878978980"});
		// this.data.push({ ward: "5A/616", name: "Parveen", mobile: "5645645665"});
		// this.data.push({ ward: "6A/897", name: "Mandar bhong", mobile: "98789909090"});
		// this.data.push({ ward: "7A/244", name: "Om", mobile: "323434355445"});
		// this.data.push({ ward: "3B/324", name: "Ahubham", mobile: "9809878679"});
		// this.data.push({ ward: "2A/454", name: "Suraj", mobile: "76568768778"});
		// this.data.push({ ward: "5A/616", name: "Parveen", mobile: "5645645665"});
		// this.data.push({ ward: "6A/897", name: "Mandar bhong", mobile: "98789909090"});
		// this.data.push({ ward: "7A/244", name: "Om", mobile: "323434355445"});
		// this.data.push({ ward: "3B/324", name: "Ahubham", mobile: "9809878679"});

		// console.log('this.data', this.data);
		// for (let i = 1; i < 100; i++) {
		// 	let newName = { ward: "3A/312", name: "Sumeet karande", mobile: "9878978980"};
		// 	this.data.push(newName);
		// }
	}
	public sBLoaded(args){
        var searchbar:SearchBar = <SearchBar>args.object;
        if(isAndroid){    
            searchbar.android.clearFocus();
        }
    }
	goBackPage() {
		this.routerExtensions.navigate(["/home"], { clearHistory: true });
	}
	details() {
		this.routerExtensions.navigate(["/list/details"], { clearHistory: true });
	}
	// onOpenDrawerTap() {
    //     this.drawerComponent.sideDrawer.showDrawer();
    // }
    // onCloseDrawerTap() {
    //     this.drawerComponent.sideDrawer.closeDrawer();
    // }

	searchTab() {
		this.searchshow = true;
		this.searchiocn =  false;
	}
	searchTabClose(){
		this.searchshow = false;
		this.searchiocn =  true;
	}
	onSubmit() {

	}
	searchBarLoaded() {

	}
	onTextChange() {

	}
}