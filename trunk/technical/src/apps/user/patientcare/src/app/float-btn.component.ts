import { Component } from "@angular/core";
import { RouterExtensions } from "nativescript-angular/router";
@Component({
    moduleId: module.id,
    selector: "float-btn",
    template: `
    <RL:Ripple rippleColor="#fff" borderRadius="28" height="56" width="56" (tap)="details()">
    <StackLayout class="float-btn">
        <Label class="float-btn-text" text="+"></Label>
    </StackLayout>
    </RL:Ripple>`,
    styles: [
        ` 
            .float-btn
            {
                background-color: #FF8910;
                border-radius:28;
                width:56;
                height: 56;
                text-align: center;
                vertical-align: middle;
                box-shadow: 2px 2px 3px #999;
            }
            .float-btn-text
            {
                color: #ffffff;
                font-size:36;
            }
        `
    ]

})
export class FloatBtnComponent {
    constructor(private routerExtensions: RouterExtensions) {
    }
    details() {
		this.routerExtensions.navigate(["/list/cameras"], { clearHistory: true });
	}
}
