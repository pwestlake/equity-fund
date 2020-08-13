import { Component, OnInit, ViewEncapsulation, ViewChild } from '@angular/core';
import { Observable } from 'rxjs';
import { EquityCatalogItem } from 'src/data/equitycatalogitem';
import { EquitycatalogService } from 'src/app/services/equitycatalog.service';
import { EodService } from 'src/app/services/eod.service';
import { mergeMap, map } from 'rxjs/operators';
import { DateValueModel } from 'src/data/date-value.model';
import { EndOfDayItem } from 'src/data/eoditem';
import { MatTableDataSource } from '@angular/material/table';
import { MatSort } from '@angular/material/sort';
import { NewsDataSource } from './newsdatasource';
import { NewsService } from 'src/app/services/news.service';

@Component({
  selector: 'app-summary',
  templateUrl: './summary.component.html',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./summary.component.scss']
})
export class SummaryComponent implements OnInit {
  catalog: Observable<EquityCatalogItem[]>;
  chartData: Map<String, Observable<DateValueModel[]>> = new Map<String, Observable<DateValueModel[]>>();
  displayedColumns: string[] = ['symbol', 'open', 'high', 'low', 'close', 'close_chg'];
  tableData: MatTableDataSource<EndOfDayItem>;
  latest: Observable<EndOfDayItem>;
  newsData: NewsDataSource;

  constructor(private equityCatalogService: EquitycatalogService, 
    private endOfDayService: EodService,
    private newsService: NewsService) { }

  @ViewChild(MatSort, {static: true}) 
  sort: MatSort;

  ngOnInit(): void {
    this.newsData = new NewsDataSource(this.newsService);
    this.latest = this.endOfDayService.getLatestEndOfDayItem();
    this.endOfDayService.getLatestEndOfDayItems().subscribe(
      items => {
        this.tableData = new MatTableDataSource(items);
        this.tableData.sort = this.sort;
      }
    )
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

  selectedTabChanged(index: number) {
    if (index == 1) {
      this.newsData.loadData("");
    }
  }

  equityChanged(equity: EquityCatalogItem) {
    if (typeof equity.id === "undefined") {
      this.newsData.setSearchId("")
    } else {
      this.newsData.setSearchId(equity.id)
    }
  }
}
