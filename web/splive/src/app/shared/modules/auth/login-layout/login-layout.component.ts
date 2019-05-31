import { Component, OnInit, OnDestroy } from '@angular/core';
import { Subscription } from 'rxjs';
import { AuthService } from '../../../services/auth.service';

@Component({
  selector: 'hkt-login-layout',
  templateUrl: './login-layout.component.html',
  styleUrls: ['./login-layout.component.css']
})
export class LoginLayoutComponent implements OnInit, OnDestroy {
  flipped = false;
  imageVisibilitySubscription: Subscription;
  isVisible = true;
  constructor(public authService: AuthService) { }

  ngOnInit() {
    // subscription for setting visibility of login layout footer
    this.imageVisibilitySubscription = this.authService.removeImageAfterSuccessSubscription.subscribe((value) => {
      this.isVisible = value;
    });
  }
  
  flipIt() {
    this.flipped = !this.flipped;
  }

  ngOnDestroy() {
    if (this.imageVisibilitySubscription) {
      this.imageVisibilitySubscription.unsubscribe();
    }
  }
}
