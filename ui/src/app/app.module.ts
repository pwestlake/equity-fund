import { BrowserModule } from '@angular/platform-browser';
import { NgModule, APP_INITIALIZER } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ContextProvider } from './providers/context.provider';
import { FlexLayoutModule } from '@angular/flex-layout';
import { FormsModule} from '@angular/forms';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatListModule } from '@angular/material/list';
import { MatMenuModule } from '@angular/material/menu';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { MatSelectModule } from '@angular/material/select';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatSortModule } from '@angular/material/sort';
import { MatTableModule } from '@angular/material/table';
import { MatTabsModule } from '@angular/material/tabs';
import { MatToolbarModule } from '@angular/material/toolbar';
import { ReactiveFormsModule } from '@angular/forms';
import { ScrollingModule } from '@angular/cdk/scrolling';
import { SummaryComponent } from './views/summary/summary.component';
import { DateValueComponent } from './components/charts/date-value/date-value.component';
import { EquitycatalogComponent } from './views/equitycatalog/equitycatalog.component';
import { EquitycatalogitemComponent } from './forms/equitycatalogitem/equitycatalogitem.component';
import { EquitySelectorComponent } from './components/autocomplete/equity-selector/equity-selector.component';

@NgModule({
  declarations: [
    AppComponent,
    SummaryComponent,
    DateValueComponent,
    EquitycatalogComponent,
    EquitycatalogitemComponent,
    EquitySelectorComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    HttpClientModule,
    FlexLayoutModule,
    FormsModule,
    MatAutocompleteModule,
    MatButtonModule,
    MatCardModule,
    MatCheckboxModule,
    MatIconModule,
    MatInputModule,
    MatListModule,
    MatMenuModule,
    MatProgressBarModule,
    MatSelectModule,
    MatSidenavModule,
    MatSortModule,
    MatTableModule,
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