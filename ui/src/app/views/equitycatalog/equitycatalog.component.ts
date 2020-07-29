import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { EquitycatalogService } from 'src/app/services/equitycatalog.service'
import { EquityCatalogItem } from 'src/data/equitycatalogitem';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-equitycatalog',
  templateUrl: './equitycatalog.component.html',
  encapsulation: ViewEncapsulation.None,
  host: {
    'class': 'router-flex'
  },
  styleUrls: ['./equitycatalog.component.scss']
})
export class EquitycatalogComponent implements OnInit {
  equityCatalog: Observable<EquityCatalogItem>;

  constructor(private equityCatalogService: EquitycatalogService) { }

  ngOnInit(): void {
    this.equityCatalog = this.equityCatalogService.getAllEquityCatalogItems();
  }

  addItem() {
    const item: EquityCatalogItem = {
      id: "",
      symbol: "asdad",
      lsetidm: "asdad",
      lseissuername: "asda",
      lsetabid: "asdad",
      lsecomponentid: "string",
      datecreated: undefined,
      lastmodified: undefined
    }

    this.equityCatalogService.addEquityCatalogItem(item);
  }
}
