import { Component, OnInit } from '@angular/core';

import { Input, OnDestroy, Inject, ViewEncapsulation} from '@angular/core';

@Component({
  selector: 'hkt-auth-layout',
  templateUrl: './auth-layout.component.html',
  styleUrls: ['./auth-layout.component.css']
})
export class AuthLayoutComponent implements OnInit {

  constructor(){}
//   public isAuthLayoutVisible = true;
//   public Spinkit = Spinkit;
//   @Input() public backgroundColor = 'rgba(255, 255, 255, 0.8)';
//   @Input() public authlayout = Spinkit.skLine;
//   constructor(private router: Router, @Inject(DOCUMENT) private document: Document) {
//       this.router.events.subscribe(event => {
//           if (event instanceof NavigationStart) {
//               this.isAuthLayoutVisible = true;
//           } else if ( event instanceof NavigationEnd || event instanceof NavigationCancel || event instanceof NavigationError) {
//               this.isAuthLayoutVisible = false;
//           }
//       }, () => {
//           this.isAuthLayoutVisible = false;
//       });
// }
  ngOnInit(){}
//   ngOnInit() : void {
//     this.isAuthLayoutVisible = false;
// }
}


