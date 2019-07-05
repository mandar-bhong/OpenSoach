import { Component } from "@angular/core";
import { RouterExtensions } from "nativescript-angular/router";
@Component({
    moduleId: module.id,
    selector: "float-btn",
    template: `
    <RL:Ripple rippleColor="#fff" borderRadius="28" height="30" width="30" (tap)="monitor()">
    <StackLayout class="float-btn" (tap)="monitor()" >
        <Label class="float-btn-text fas" text="&#xf201;"></Label>
    </StackLayout>
    </RL:Ripple>
    <RL:Ripple rippleColor="#fff" borderRadius="28" height="30" width="30" (tap)="medicine()">
    <StackLayout class="float-btn">
        <Label class="float-btn-text fas" text="&#xf484;"></Label>
    </StackLayout>
    </RL:Ripple>
    <RL:Ripple rippleColor="#fff" borderRadius="28" height="30" width="30" (tap)="intake()">
    <StackLayout class="float-btn">
        <Label class="float-btn-text fas" text="&#xf48e;"></Label>
    </StackLayout>
    </RL:Ripple>
    `,
    styles: [
        ` 
            .float-btn
            {
                background-color: #FF8910;
                border-radius:28;
                width:30;
                height: 30;
                text-align: center;
                vertical-align: middle;
                box-shadow: 2px 2px 3px #999;
            }
            .float-btn-text
            {
                color: #ffffff;
                font-size:15;
            }
        `
    ]

})
export class FloatBtnComponent {
    constructor(private routerExtensions: RouterExtensions) {
    }

    monitor() {
		this.routerExtensions.navigate(["/patientmgnt/monitor-chart"], { clearHistory: true });	
    }
    medicine() {
		this.routerExtensions.navigate(["/patientmgnt/medicine-chart"], { clearHistory: true });			
    }
    intake() {
		this.routerExtensions.navigate(["/patientmgnt/intake-chart"], { clearHistory: true });		
    }
}
