import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { Observable } from 'rxjs';
import { EquityCatalogItem } from 'src/data/equitycatalogitem';
import { EquitycatalogService } from 'src/app/services/equitycatalog.service';
import { EodService } from 'src/app/services/eod.service';
import { mergeMap, map } from 'rxjs/operators';
import { DateValueModel } from 'src/data/date-value.model';

@Component({
  selector: 'app-summary',
  templateUrl: './summary.component.html',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./summary.component.scss']
})
export class SummaryComponent implements OnInit {
  catalog: Observable<EquityCatalogItem[]>
  chartData: Map<String, Observable<DateValueModel[]>> = new Map<String, Observable<DateValueModel[]>>()

  constructor(private equityCatalogService: EquitycatalogService, private endOfDayService: EodService) { }

  ngOnInit(): void {
    this.catalog = this.equityCatalogService.getAllEquityCatalogItems();
  
    this.catalog.pipe(
      map(catalog => {
        catalog.forEach(item => {
          this.chartData.set(item.id, this.endOfDayService.getClosePriceTimeSeries(item.id))
        })
      })
      
      ).subscribe()
  }

  closePriceChart(id: string): Observable<DateValueModel[]>{
    return this.chartData.get(id);
  }
}
