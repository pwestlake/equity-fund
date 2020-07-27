import { Component, ViewEncapsulation, ChangeDetectorRef, ViewChild } from '@angular/core';
import { MediaMatcher } from '@angular/cdk/layout';
import { Router } from '@angular/router';
import { MatDrawer } from '@angular/material/sidenav';
import { HttpIndicatorService } from './services/http-indicator.service';
import { StyleService } from './services/style.service';
import { Subject } from 'rxjs';
import { UserService } from './services/user.service';
import { ContextProvider } from './providers/context.provider';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'angulara';
  mobileQuery: MediaQueryList;
  isLoading: Subject<boolean>;

  private mobileQueryListener: () => void;

  @ViewChild('snav') 
  private snav: MatDrawer; 

  theme: string = "";

  constructor(changeDetectorRef: ChangeDetectorRef, media: MediaMatcher, 
    private router: Router, private httpIndicatorService: HttpIndicatorService,
    private styleService: StyleService,
    private contextProvider: ContextProvider,
    private userService: UserService) { 
      
    this.mobileQuery = media.matchMedia('(max-width: 600px)');
    this.mobileQueryListener = () => changeDetectorRef.detectChanges();
    this.mobileQuery.addListener(this.mobileQueryListener);

    this.isLoading = this.httpIndicatorService.isLoading;
  }

  ngOnInit(): void {
    this.router.events.subscribe(event => {this.snav.close()});
    this.theme = this.contextProvider.getTheme();
    this.title = this.contextProvider.getTitle();
  }

  ngOnDestroy(): void {
    this.mobileQuery.removeListener(this.mobileQueryListener);
  }

  setTheme(theme: string) {
    this.styleService.setStyle("theme", `assets/themes/${theme}.css`);
    this.userService.setPreference(this.contextProvider.getAuthenticatedUser(), "theme", theme);
  }
}
