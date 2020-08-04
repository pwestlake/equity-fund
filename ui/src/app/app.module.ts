import { BrowserModule } from '@angular/platform-browser';
import { NgModule, APP_INITIALIZER } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ContextProvider } from './providers/context.provider';
import { FlexLayoutModule } from '@angular/flex-layout';
import { FormsModule} from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatListModule } from '@angular/material/list';
import { MatMenuModule } from '@angular/material/menu';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatTabsModule } from '@angular/material/tabs';
import { MatToolbarModule } from '@angular/material/toolbar';
import { ReactiveFormsModule } from '@angular/forms';
import { ScrollingModule } from '@angular/cdk/scrolling';
import { SummaryComponent } from './views/summary/summary.component';
import { DateValueComponent } from './components/charts/date-value/date-value.component';
import { EquitycatalogComponent } from './views/equitycatalog/equitycatalog.component';
import { EquitycatalogitemComponent } from './forms/equitycatalogitem/equitycatalogitem.component';

@NgModule({
  declarations: [
    AppComponent,
    SummaryComponent,
    DateValueComponent,
    EquitycatalogComponent,
    EquitycatalogitemComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    HttpClientModule,
    FlexLayoutModule,
    FormsModule,
    MatButtonModule,
    MatCardModule,
    MatIconModule,
    MatInputModule,
    MatListModule,
    MatMenuModule,
    MatProgressBarModule,
    MatSidenavModule,
    MatTabsModule,
    MatToolbarModule,
    ReactiveFormsModule,
    ScrollingModule
  ],
  providers: [
    ContextProvider,
    {
      provide: APP_INITIALIZER,
      useFactory: themeProviderFactory,
      multi: true,
      deps: [ContextProvider]
    }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }

export function themeProviderFactory(provider: ContextProvider) {
  return () => provider.load();
}