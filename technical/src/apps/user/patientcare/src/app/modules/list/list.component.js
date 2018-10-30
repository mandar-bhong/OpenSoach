"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var core_1 = require("@angular/core");
var router_1 = require("nativescript-angular/router");
var platform_1 = require("platform");
// import { RadSideDrawerComponent, SideDrawerType } from "nativescript-ui-sidedrawer/angular";
// import { ViewChild } from "@angular/core";
// import { RadSideDrawer } from "nativescript-ui-sidedrawer";
var ListComponent = /** @class */ (function () {
    // @ViewChild(RadSideDrawerComponent) public drawerComponent: RadSideDrawerComponent;
    function ListComponent(routerExtensions) {
        this.routerExtensions = routerExtensions;
        this.data = [];
        this.searchshow = false;
        this.searchiocn = true;
    }
    ListComponent.prototype.ngOnInit = function () {
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
    };
    ListComponent.prototype.sBLoaded = function (args) {
        var searchbar = args.object;
        if (platform_1.isAndroid) {
            searchbar.android.clearFocus();
        }
    };
    ListComponent.prototype.goBackPage = function () {
        this.routerExtensions.navigate(["/home"], { clearHistory: true });
    };
    ListComponent.prototype.details = function () {
        this.routerExtensions.navigate(["/list/details"], { clearHistory: true });
    };
    // onOpenDrawerTap() {
    //     this.drawerComponent.sideDrawer.showDrawer();
    // }
    // onCloseDrawerTap() {
    //     this.drawerComponent.sideDrawer.closeDrawer();
    // }
    ListComponent.prototype.searchTab = function () {
        this.searchshow = true;
        this.searchiocn = false;
    };
    ListComponent.prototype.searchTabClose = function () {
        this.searchshow = false;
        this.searchiocn = true;
    };
    ListComponent.prototype.onSubmit = function () {
    };
    ListComponent.prototype.searchBarLoaded = function () {
    };
    ListComponent.prototype.onTextChange = function () {
    };
    ListComponent = __decorate([
        core_1.Component({
            moduleId: module.id,
            selector: 'list',
            templateUrl: './list.component.html',
            styleUrls: ['./list.component.css']
        }),
        __metadata("design:paramtypes", [router_1.RouterExtensions])
    ], ListComponent);
    return ListComponent;
}());
exports.ListComponent = ListComponent;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoibGlzdC5jb21wb25lbnQuanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyJsaXN0LmNvbXBvbmVudC50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiOztBQUFBLHNDQUFrRDtBQUNsRCxzREFBK0Q7QUFFL0QscUNBQXFDO0FBRXJDLCtGQUErRjtBQUMvRiw2Q0FBNkM7QUFDN0MsOERBQThEO0FBUTlEO0lBSUMscUZBQXFGO0lBQ3JGLHVCQUFvQixnQkFBa0M7UUFBbEMscUJBQWdCLEdBQWhCLGdCQUFnQixDQUFrQjtRQUp0RCxTQUFJLEdBQUcsRUFBRSxDQUFDO1FBQ1YsZUFBVSxHQUFHLEtBQUssQ0FBQztRQUNuQixlQUFVLEdBQUcsSUFBSSxDQUFDO0lBRXdDLENBQUM7SUFFM0QsZ0NBQVEsR0FBUjtRQUNDLGtEQUFrRDtRQUNsRCxJQUFJLENBQUMsSUFBSSxDQUFDLElBQUksQ0FBQyxFQUFFLElBQUksRUFBRSxRQUFRLEVBQUUsSUFBSSxFQUFFLGdCQUFnQixFQUFFLE1BQU0sRUFBRSxZQUFZLEVBQUUsQ0FBQyxDQUFDO1FBQ2pGLElBQUksQ0FBQyxJQUFJLENBQUMsSUFBSSxDQUFDLEVBQUUsSUFBSSxFQUFFLFFBQVEsRUFBRSxJQUFJLEVBQUUsWUFBWSxFQUFFLE1BQU0sRUFBRSxhQUFhLEVBQUUsQ0FBQyxDQUFDO1FBQzlFLGdGQUFnRjtRQUNoRixrRkFBa0Y7UUFDbEYsNEVBQTRFO1FBQzVFLGtGQUFrRjtRQUNsRix5RUFBeUU7UUFDekUsNEVBQTRFO1FBQzVFLDJFQUEyRTtRQUMzRSw0RUFBNEU7UUFDNUUsa0ZBQWtGO1FBQ2xGLHlFQUF5RTtRQUN6RSw0RUFBNEU7UUFFNUUsdUNBQXVDO1FBQ3ZDLGtDQUFrQztRQUNsQyxrRkFBa0Y7UUFDbEYsNEJBQTRCO1FBQzVCLElBQUk7SUFDTCxDQUFDO0lBQ00sZ0NBQVEsR0FBZixVQUFnQixJQUFJO1FBQ2IsSUFBSSxTQUFTLEdBQXdCLElBQUksQ0FBQyxNQUFNLENBQUM7UUFDakQsRUFBRSxDQUFBLENBQUMsb0JBQVMsQ0FBQyxDQUFBLENBQUM7WUFDVixTQUFTLENBQUMsT0FBTyxDQUFDLFVBQVUsRUFBRSxDQUFDO1FBQ25DLENBQUM7SUFDTCxDQUFDO0lBQ0osa0NBQVUsR0FBVjtRQUNDLElBQUksQ0FBQyxnQkFBZ0IsQ0FBQyxRQUFRLENBQUMsQ0FBQyxPQUFPLENBQUMsRUFBRSxFQUFFLFlBQVksRUFBRSxJQUFJLEVBQUUsQ0FBQyxDQUFDO0lBQ25FLENBQUM7SUFDRCwrQkFBTyxHQUFQO1FBQ0MsSUFBSSxDQUFDLGdCQUFnQixDQUFDLFFBQVEsQ0FBQyxDQUFDLGVBQWUsQ0FBQyxFQUFFLEVBQUUsWUFBWSxFQUFFLElBQUksRUFBRSxDQUFDLENBQUM7SUFDM0UsQ0FBQztJQUNELHNCQUFzQjtJQUNuQixvREFBb0Q7SUFDcEQsSUFBSTtJQUNKLHVCQUF1QjtJQUN2QixxREFBcUQ7SUFDckQsSUFBSTtJQUVQLGlDQUFTLEdBQVQ7UUFDQyxJQUFJLENBQUMsVUFBVSxHQUFHLElBQUksQ0FBQztRQUN2QixJQUFJLENBQUMsVUFBVSxHQUFJLEtBQUssQ0FBQztJQUMxQixDQUFDO0lBQ0Qsc0NBQWMsR0FBZDtRQUNDLElBQUksQ0FBQyxVQUFVLEdBQUcsS0FBSyxDQUFDO1FBQ3hCLElBQUksQ0FBQyxVQUFVLEdBQUksSUFBSSxDQUFDO0lBQ3pCLENBQUM7SUFDRCxnQ0FBUSxHQUFSO0lBRUEsQ0FBQztJQUNELHVDQUFlLEdBQWY7SUFFQSxDQUFDO0lBQ0Qsb0NBQVksR0FBWjtJQUVBLENBQUM7SUFoRVcsYUFBYTtRQVB6QixnQkFBUyxDQUFDO1lBQ1YsUUFBUSxFQUFFLE1BQU0sQ0FBQyxFQUFFO1lBQ25CLFFBQVEsRUFBRSxNQUFNO1lBQ2hCLFdBQVcsRUFBRSx1QkFBdUI7WUFDcEMsU0FBUyxFQUFFLENBQUMsc0JBQXNCLENBQUM7U0FDbkMsQ0FBQzt5Q0FPcUMseUJBQWdCO09BTDFDLGFBQWEsQ0FpRXpCO0lBQUQsb0JBQUM7Q0FBQSxBQWpFRCxJQWlFQztBQWpFWSxzQ0FBYSIsInNvdXJjZXNDb250ZW50IjpbImltcG9ydCB7IENvbXBvbmVudCwgT25Jbml0IH0gZnJvbSAnQGFuZ3VsYXIvY29yZSc7XG5pbXBvcnQgeyBSb3V0ZXJFeHRlbnNpb25zIH0gZnJvbSBcIm5hdGl2ZXNjcmlwdC1hbmd1bGFyL3JvdXRlclwiO1xuaW1wb3J0IHsgU2VhcmNoQmFyIH0gZnJvbSBcInRucy1jb3JlLW1vZHVsZXMvdWkvc2VhcmNoLWJhclwiO1xuaW1wb3J0IHsgaXNBbmRyb2lkIH0gZnJvbSBcInBsYXRmb3JtXCI7XG5pbXBvcnQgKiBhcyBhcHAgZnJvbSBcImFwcGxpY2F0aW9uXCI7XG4vLyBpbXBvcnQgeyBSYWRTaWRlRHJhd2VyQ29tcG9uZW50LCBTaWRlRHJhd2VyVHlwZSB9IGZyb20gXCJuYXRpdmVzY3JpcHQtdWktc2lkZWRyYXdlci9hbmd1bGFyXCI7XG4vLyBpbXBvcnQgeyBWaWV3Q2hpbGQgfSBmcm9tIFwiQGFuZ3VsYXIvY29yZVwiO1xuLy8gaW1wb3J0IHsgUmFkU2lkZURyYXdlciB9IGZyb20gXCJuYXRpdmVzY3JpcHQtdWktc2lkZWRyYXdlclwiO1xuQENvbXBvbmVudCh7XG5cdG1vZHVsZUlkOiBtb2R1bGUuaWQsXG5cdHNlbGVjdG9yOiAnbGlzdCcsXG5cdHRlbXBsYXRlVXJsOiAnLi9saXN0LmNvbXBvbmVudC5odG1sJyxcblx0c3R5bGVVcmxzOiBbJy4vbGlzdC5jb21wb25lbnQuY3NzJ11cbn0pXG5cbmV4cG9ydCBjbGFzcyBMaXN0Q29tcG9uZW50IGltcGxlbWVudHMgT25Jbml0IHtcblx0ZGF0YSA9IFtdO1xuXHRzZWFyY2hzaG93ID0gZmFsc2U7XG5cdHNlYXJjaGlvY24gPSB0cnVlO1xuXHQvLyBAVmlld0NoaWxkKFJhZFNpZGVEcmF3ZXJDb21wb25lbnQpIHB1YmxpYyBkcmF3ZXJDb21wb25lbnQ6IFJhZFNpZGVEcmF3ZXJDb21wb25lbnQ7XG5cdGNvbnN0cnVjdG9yKHByaXZhdGUgcm91dGVyRXh0ZW5zaW9uczogUm91dGVyRXh0ZW5zaW9ucykgeyB9XG5cblx0bmdPbkluaXQoKSB7XG5cdFx0Ly8gdGhpcy5kYXRhLnB1c2goeyB0ZXh0OiBcIkJ1bGJhc2F1clwiLCBzcmM6IFwiXCIgfSk7XG5cdFx0dGhpcy5kYXRhLnB1c2goeyB3YXJkOiBcIjNBLzMxMlwiLCBuYW1lOiBcIlN1bWVldCBrYXJhbmRlXCIsIG1vYmlsZTogXCI5ODc4OTc4OTgwXCIgfSk7XG5cdFx0dGhpcy5kYXRhLnB1c2goeyB3YXJkOiBcIjNCLzMyM1wiLCBuYW1lOiBcIkFtb2wgUGF0aWxcIiwgbW9iaWxlOiBcIjMyNDIzNDMyNDMyXCIgfSk7XG5cdFx0Ly8gdGhpcy5kYXRhLnB1c2goeyB3YXJkOiBcIjJBLzY0M1wiLCBuYW1lOiBcIlNhbmpheSBNb2hhblwiLCBtb2JpbGU6IFwiNDUzNTQ1MzUyXCJ9KTtcblx0XHQvLyB0aGlzLmRhdGEucHVzaCh7IHdhcmQ6IFwiNEEvNTE1XCIsIG5hbWU6IFwiU2FuamF5IFNhd2FudFwiLCBtb2JpbGU6IFwiOTg3ODk3ODk4MFwifSk7XG5cdFx0Ly8gdGhpcy5kYXRhLnB1c2goeyB3YXJkOiBcIjVBLzYxNlwiLCBuYW1lOiBcIlBhcnZlZW5cIiwgbW9iaWxlOiBcIjU2NDU2NDU2NjVcIn0pO1xuXHRcdC8vIHRoaXMuZGF0YS5wdXNoKHsgd2FyZDogXCI2QS84OTdcIiwgbmFtZTogXCJNYW5kYXIgYmhvbmdcIiwgbW9iaWxlOiBcIjk4Nzg5OTA5MDkwXCJ9KTtcblx0XHQvLyB0aGlzLmRhdGEucHVzaCh7IHdhcmQ6IFwiN0EvMjQ0XCIsIG5hbWU6IFwiT21cIiwgbW9iaWxlOiBcIjMyMzQzNDM1NTQ0NVwifSk7XG5cdFx0Ly8gdGhpcy5kYXRhLnB1c2goeyB3YXJkOiBcIjNCLzMyNFwiLCBuYW1lOiBcIkFodWJoYW1cIiwgbW9iaWxlOiBcIjk4MDk4Nzg2NzlcIn0pO1xuXHRcdC8vIHRoaXMuZGF0YS5wdXNoKHsgd2FyZDogXCIyQS80NTRcIiwgbmFtZTogXCJTdXJhalwiLCBtb2JpbGU6IFwiNzY1Njg3Njg3NzhcIn0pO1xuXHRcdC8vIHRoaXMuZGF0YS5wdXNoKHsgd2FyZDogXCI1QS82MTZcIiwgbmFtZTogXCJQYXJ2ZWVuXCIsIG1vYmlsZTogXCI1NjQ1NjQ1NjY1XCJ9KTtcblx0XHQvLyB0aGlzLmRhdGEucHVzaCh7IHdhcmQ6IFwiNkEvODk3XCIsIG5hbWU6IFwiTWFuZGFyIGJob25nXCIsIG1vYmlsZTogXCI5ODc4OTkwOTA5MFwifSk7XG5cdFx0Ly8gdGhpcy5kYXRhLnB1c2goeyB3YXJkOiBcIjdBLzI0NFwiLCBuYW1lOiBcIk9tXCIsIG1vYmlsZTogXCIzMjM0MzQzNTU0NDVcIn0pO1xuXHRcdC8vIHRoaXMuZGF0YS5wdXNoKHsgd2FyZDogXCIzQi8zMjRcIiwgbmFtZTogXCJBaHViaGFtXCIsIG1vYmlsZTogXCI5ODA5ODc4Njc5XCJ9KTtcblxuXHRcdC8vIGNvbnNvbGUubG9nKCd0aGlzLmRhdGEnLCB0aGlzLmRhdGEpO1xuXHRcdC8vIGZvciAobGV0IGkgPSAxOyBpIDwgMTAwOyBpKyspIHtcblx0XHQvLyBcdGxldCBuZXdOYW1lID0geyB3YXJkOiBcIjNBLzMxMlwiLCBuYW1lOiBcIlN1bWVldCBrYXJhbmRlXCIsIG1vYmlsZTogXCI5ODc4OTc4OTgwXCJ9O1xuXHRcdC8vIFx0dGhpcy5kYXRhLnB1c2gobmV3TmFtZSk7XG5cdFx0Ly8gfVxuXHR9XG5cdHB1YmxpYyBzQkxvYWRlZChhcmdzKXtcbiAgICAgICAgdmFyIHNlYXJjaGJhcjpTZWFyY2hCYXIgPSA8U2VhcmNoQmFyPmFyZ3Mub2JqZWN0O1xuICAgICAgICBpZihpc0FuZHJvaWQpeyAgICBcbiAgICAgICAgICAgIHNlYXJjaGJhci5hbmRyb2lkLmNsZWFyRm9jdXMoKTtcbiAgICAgICAgfVxuICAgIH1cblx0Z29CYWNrUGFnZSgpIHtcblx0XHR0aGlzLnJvdXRlckV4dGVuc2lvbnMubmF2aWdhdGUoW1wiL2hvbWVcIl0sIHsgY2xlYXJIaXN0b3J5OiB0cnVlIH0pO1xuXHR9XG5cdGRldGFpbHMoKSB7XG5cdFx0dGhpcy5yb3V0ZXJFeHRlbnNpb25zLm5hdmlnYXRlKFtcIi9saXN0L2RldGFpbHNcIl0sIHsgY2xlYXJIaXN0b3J5OiB0cnVlIH0pO1xuXHR9XG5cdC8vIG9uT3BlbkRyYXdlclRhcCgpIHtcbiAgICAvLyAgICAgdGhpcy5kcmF3ZXJDb21wb25lbnQuc2lkZURyYXdlci5zaG93RHJhd2VyKCk7XG4gICAgLy8gfVxuICAgIC8vIG9uQ2xvc2VEcmF3ZXJUYXAoKSB7XG4gICAgLy8gICAgIHRoaXMuZHJhd2VyQ29tcG9uZW50LnNpZGVEcmF3ZXIuY2xvc2VEcmF3ZXIoKTtcbiAgICAvLyB9XG5cblx0c2VhcmNoVGFiKCkge1xuXHRcdHRoaXMuc2VhcmNoc2hvdyA9IHRydWU7XG5cdFx0dGhpcy5zZWFyY2hpb2NuID0gIGZhbHNlO1xuXHR9XG5cdHNlYXJjaFRhYkNsb3NlKCl7XG5cdFx0dGhpcy5zZWFyY2hzaG93ID0gZmFsc2U7XG5cdFx0dGhpcy5zZWFyY2hpb2NuID0gIHRydWU7XG5cdH1cblx0b25TdWJtaXQoKSB7XG5cblx0fVxuXHRzZWFyY2hCYXJMb2FkZWQoKSB7XG5cblx0fVxuXHRvblRleHRDaGFuZ2UoKSB7XG5cblx0fVxufSJdfQ==