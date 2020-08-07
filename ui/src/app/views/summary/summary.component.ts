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

  constructor(private equityCatalogService: EquitycatalogService, private endOfDayService: EodService) { }

  @ViewChild(MatSort, {static: true}) 
  sort: MatSort;

  ngOnInit(): void {
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
}
