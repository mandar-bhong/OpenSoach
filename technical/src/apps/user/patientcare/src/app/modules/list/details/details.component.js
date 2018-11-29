"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var core_1 = require("@angular/core");
var router_1 = require("nativescript-angular/router");
var DetailsComponent = /** @class */ (function () {
    function DetailsComponent(routerExtensions) {
        this.routerExtensions = routerExtensions;
        this.selectedfirst = true;
        this.selectedsecond = false;
        this.tabSelectedIndex = 0;
    }
    DetailsComponent.prototype.ngOnInit = function () { };
    DetailsComponent.prototype.changeTab = function () {
        if (this.tabSelectedIndex === 0) {
            this.tabSelectedIndex = 1;
        }
        else if (this.tabSelectedIndex === 1) {
            this.tabSelectedIndex = 2;
        }
        else if (this.tabSelectedIndex === 2) {
            this.tabSelectedIndex = 0;
        }
    };
    DetailsComponent.prototype.goBackPage = function () {
        this.routerExtensions.navigate(["/list"], { clearHistory: true });
    };
    DetailsComponent.prototype.firstTab = function () {
        this.selectedfirst = true;
        this.selectedsecond = false;
    };
    DetailsComponent.prototype.secondTab = function () {
        this.selectedfirst = false;
        this.selectedsecond = true;
    };
    DetailsComponent.prototype.test = function () {
        var coll = document.getElementsByClassName("collapsible");
        var i;
        for (i = 0; i < coll.length; i++) {
            coll[i].addEventListener("click", function () {
                this.classList.toggle("active");
                var content = this.nextElementSibling;
                if (content.style.display === "block") {
                    content.style.display = "none";
                }
                else {
                    content.style.display = "block";
                }
            });
        }
    };
    DetailsComponent = __decorate([
        core_1.Component({
            moduleId: module.id,
            selector: 'details',
            templateUrl: './details.component.html',
            styleUrls: ['./details.component.css']
        }),
        __metadata("design:paramtypes", [router_1.RouterExtensions])
    ], DetailsComponent);
    return DetailsComponent;
}());
exports.DetailsComponent = DetailsComponent;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiZGV0YWlscy5jb21wb25lbnQuanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyJkZXRhaWxzLmNvbXBvbmVudC50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiOztBQUFBLHNDQUFrRDtBQUNsRCxzREFBK0Q7QUFXL0Q7SUFPQywwQkFBb0IsZ0JBQWtDO1FBQWxDLHFCQUFnQixHQUFoQixnQkFBZ0IsQ0FBa0I7UUFKdEQsa0JBQWEsR0FBRyxJQUFJLENBQUM7UUFDckIsbUJBQWMsR0FBRyxLQUFLLENBQUM7UUFJdEIsSUFBSSxDQUFDLGdCQUFnQixHQUFHLENBQUMsQ0FBQztJQUMzQixDQUFDO0lBRUQsbUNBQVEsR0FBUixjQUFhLENBQUM7SUFDZCxvQ0FBUyxHQUFUO1FBQ0MsSUFBSSxJQUFJLENBQUMsZ0JBQWdCLEtBQUssQ0FBQyxFQUFFO1lBQ2hDLElBQUksQ0FBQyxnQkFBZ0IsR0FBRyxDQUFDLENBQUM7U0FDMUI7YUFBTSxJQUFJLElBQUksQ0FBQyxnQkFBZ0IsS0FBSyxDQUFDLEVBQUU7WUFDdkMsSUFBSSxDQUFDLGdCQUFnQixHQUFHLENBQUMsQ0FBQztTQUMxQjthQUFNLElBQUksSUFBSSxDQUFDLGdCQUFnQixLQUFLLENBQUMsRUFBRTtZQUN2QyxJQUFJLENBQUMsZ0JBQWdCLEdBQUcsQ0FBQyxDQUFDO1NBQzFCO0lBQ0YsQ0FBQztJQUNELHFDQUFVLEdBQVY7UUFDQyxJQUFJLENBQUMsZ0JBQWdCLENBQUMsUUFBUSxDQUFDLENBQUMsT0FBTyxDQUFDLEVBQUUsRUFBRSxZQUFZLEVBQUUsSUFBSSxFQUFFLENBQUMsQ0FBQztJQUNuRSxDQUFDO0lBQ0QsbUNBQVEsR0FBUjtRQUNDLElBQUksQ0FBQyxhQUFhLEdBQUcsSUFBSSxDQUFDO1FBQzFCLElBQUksQ0FBQyxjQUFjLEdBQUcsS0FBSyxDQUFDO0lBQzdCLENBQUM7SUFDRCxvQ0FBUyxHQUFUO1FBQ0MsSUFBSSxDQUFDLGFBQWEsR0FBRyxLQUFLLENBQUM7UUFDM0IsSUFBSSxDQUFDLGNBQWMsR0FBRyxJQUFJLENBQUM7SUFDNUIsQ0FBQztJQUNELCtCQUFJLEdBQUo7UUFDQyxJQUFJLElBQUksR0FBRyxRQUFRLENBQUMsc0JBQXNCLENBQUMsYUFBYSxDQUFDLENBQUM7UUFDMUQsSUFBSSxDQUFDLENBQUM7UUFFTixLQUFLLENBQUMsR0FBRyxDQUFDLEVBQUUsQ0FBQyxHQUFHLElBQUksQ0FBQyxNQUFNLEVBQUUsQ0FBQyxFQUFFLEVBQUU7WUFDakMsSUFBSSxDQUFDLENBQUMsQ0FBQyxDQUFDLGdCQUFnQixDQUFDLE9BQU8sRUFBRTtnQkFDakMsSUFBSSxDQUFDLFNBQVMsQ0FBQyxNQUFNLENBQUMsUUFBUSxDQUFDLENBQUM7Z0JBQ2hDLElBQUksT0FBTyxHQUFHLElBQUksQ0FBQyxrQkFBa0IsQ0FBQztnQkFDdEMsSUFBSSxPQUFPLENBQUMsS0FBSyxDQUFDLE9BQU8sS0FBSyxPQUFPLEVBQUU7b0JBQ3RDLE9BQU8sQ0FBQyxLQUFLLENBQUMsT0FBTyxHQUFHLE1BQU0sQ0FBQztpQkFDL0I7cUJBQU07b0JBQ04sT0FBTyxDQUFDLEtBQUssQ0FBQyxPQUFPLEdBQUcsT0FBTyxDQUFDO2lCQUNoQztZQUNGLENBQUMsQ0FBQyxDQUFDO1NBQ0g7SUFFRixDQUFDO0lBaERXLGdCQUFnQjtRQVA1QixnQkFBUyxDQUFDO1lBQ1YsUUFBUSxFQUFFLE1BQU0sQ0FBQyxFQUFFO1lBQ25CLFFBQVEsRUFBRSxTQUFTO1lBQ25CLFdBQVcsRUFBRSwwQkFBMEI7WUFDdkMsU0FBUyxFQUFFLENBQUMseUJBQXlCLENBQUM7U0FDdEMsQ0FBQzt5Q0FTcUMseUJBQWdCO09BUDFDLGdCQUFnQixDQXVENUI7SUFBRCx1QkFBQztDQUFBLEFBdkRELElBdURDO0FBdkRZLDRDQUFnQiIsInNvdXJjZXNDb250ZW50IjpbImltcG9ydCB7IENvbXBvbmVudCwgT25Jbml0IH0gZnJvbSAnQGFuZ3VsYXIvY29yZSc7XG5pbXBvcnQgeyBSb3V0ZXJFeHRlbnNpb25zIH0gZnJvbSBcIm5hdGl2ZXNjcmlwdC1hbmd1bGFyL3JvdXRlclwiO1xuaW1wb3J0IHsgU2VsZWN0ZWRJbmRleENoYW5nZWRFdmVudERhdGEgfSBmcm9tIFwidG5zLWNvcmUtbW9kdWxlcy91aS90YWItdmlld1wiO1xuaW1wb3J0IHsgU2VnbWVudGVkQmFyLCBTZWdtZW50ZWRCYXJJdGVtIH0gZnJvbSBcInRucy1jb3JlLW1vZHVsZXMvdWkvc2VnbWVudGVkLWJhclwiO1xuXG5AQ29tcG9uZW50KHtcblx0bW9kdWxlSWQ6IG1vZHVsZS5pZCxcblx0c2VsZWN0b3I6ICdkZXRhaWxzJyxcblx0dGVtcGxhdGVVcmw6ICcuL2RldGFpbHMuY29tcG9uZW50Lmh0bWwnLFxuXHRzdHlsZVVybHM6IFsnLi9kZXRhaWxzLmNvbXBvbmVudC5jc3MnXVxufSlcblxuZXhwb3J0IGNsYXNzIERldGFpbHNDb21wb25lbnQgaW1wbGVtZW50cyBPbkluaXQge1xuXHRwdWJsaWMgdGFiU2VsZWN0ZWRJbmRleDogbnVtYmVyO1xuXHRiMTtcblx0c2VsZWN0ZWRmaXJzdCA9IHRydWU7XG5cdHNlbGVjdGVkc2Vjb25kID0gZmFsc2U7XG5cdHB1YmxpYyBTZWxlY3RlZEluZGV4OiBudW1iZXI7XG5cblx0Y29uc3RydWN0b3IocHJpdmF0ZSByb3V0ZXJFeHRlbnNpb25zOiBSb3V0ZXJFeHRlbnNpb25zKSB7XG5cdFx0dGhpcy50YWJTZWxlY3RlZEluZGV4ID0gMDtcblx0fVxuXG5cdG5nT25Jbml0KCkgeyB9XG5cdGNoYW5nZVRhYigpIHtcblx0XHRpZiAodGhpcy50YWJTZWxlY3RlZEluZGV4ID09PSAwKSB7XG5cdFx0XHR0aGlzLnRhYlNlbGVjdGVkSW5kZXggPSAxO1xuXHRcdH0gZWxzZSBpZiAodGhpcy50YWJTZWxlY3RlZEluZGV4ID09PSAxKSB7XG5cdFx0XHR0aGlzLnRhYlNlbGVjdGVkSW5kZXggPSAyO1xuXHRcdH0gZWxzZSBpZiAodGhpcy50YWJTZWxlY3RlZEluZGV4ID09PSAyKSB7XG5cdFx0XHR0aGlzLnRhYlNlbGVjdGVkSW5kZXggPSAwO1xuXHRcdH1cblx0fVxuXHRnb0JhY2tQYWdlKCkge1xuXHRcdHRoaXMucm91dGVyRXh0ZW5zaW9ucy5uYXZpZ2F0ZShbXCIvbGlzdFwiXSwgeyBjbGVhckhpc3Rvcnk6IHRydWUgfSk7XG5cdH1cblx0Zmlyc3RUYWIoKSB7XG5cdFx0dGhpcy5zZWxlY3RlZGZpcnN0ID0gdHJ1ZTtcblx0XHR0aGlzLnNlbGVjdGVkc2Vjb25kID0gZmFsc2U7XG5cdH1cblx0c2Vjb25kVGFiKCkge1xuXHRcdHRoaXMuc2VsZWN0ZWRmaXJzdCA9IGZhbHNlO1xuXHRcdHRoaXMuc2VsZWN0ZWRzZWNvbmQgPSB0cnVlO1xuXHR9XG5cdHRlc3QoKSB7XG5cdFx0dmFyIGNvbGwgPSBkb2N1bWVudC5nZXRFbGVtZW50c0J5Q2xhc3NOYW1lKFwiY29sbGFwc2libGVcIik7XG5cdFx0dmFyIGk7XG5cblx0XHRmb3IgKGkgPSAwOyBpIDwgY29sbC5sZW5ndGg7IGkrKykge1xuXHRcdFx0Y29sbFtpXS5hZGRFdmVudExpc3RlbmVyKFwiY2xpY2tcIiwgZnVuY3Rpb24gKCkge1xuXHRcdFx0XHR0aGlzLmNsYXNzTGlzdC50b2dnbGUoXCJhY3RpdmVcIik7XG5cdFx0XHRcdHZhciBjb250ZW50ID0gdGhpcy5uZXh0RWxlbWVudFNpYmxpbmc7XG5cdFx0XHRcdGlmIChjb250ZW50LnN0eWxlLmRpc3BsYXkgPT09IFwiYmxvY2tcIikge1xuXHRcdFx0XHRcdGNvbnRlbnQuc3R5bGUuZGlzcGxheSA9IFwibm9uZVwiO1xuXHRcdFx0XHR9IGVsc2Uge1xuXHRcdFx0XHRcdGNvbnRlbnQuc3R5bGUuZGlzcGxheSA9IFwiYmxvY2tcIjtcblx0XHRcdFx0fVxuXHRcdFx0fSk7XG5cdFx0fVxuXG5cdH1cblxuXG5cblxuXG5cdFxufSJdfQ==