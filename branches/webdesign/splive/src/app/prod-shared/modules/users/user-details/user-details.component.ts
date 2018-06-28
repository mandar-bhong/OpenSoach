import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-user-details',
  templateUrl: './user-details.component.html',
  styleUrls: ['./user-details.component.css']
})
export class UserDetailsComponent implements OnInit {
  routeSubscription: Subscription;
  callbackUrl;
  constructor(
    private route: ActivatedRoute,
    private router: Router) {}
  ngOnInit() {
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      this.callbackUrl = params['callbackurl'];
    });
  }
  closeForm() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }
}
