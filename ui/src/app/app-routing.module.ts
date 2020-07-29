import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { SummaryComponent } from './views/summary/summary.component';
import { EquitycatalogComponent } from './views/equitycatalog/equitycatalog.component';
import { EquitycatalogitemComponent } from './forms/equitycatalogitem/equitycatalogitem.component';

const routes: Routes = [
  { path: 'summary', component: SummaryComponent },
  { path: 'equitycatalog', component: EquitycatalogComponent },
  { path: 'equitycatalog/equitycatalogitemform', component: EquitycatalogitemComponent },
  { path: 'equity-fund/ui',
    redirectTo: 'summary',
    pathMatch: 'prefix'
  },
];
@NgModule({
  imports: [RouterModule.forRoot(routes, { useHash: true })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
