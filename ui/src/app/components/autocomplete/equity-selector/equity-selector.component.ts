import { Component, OnInit, EventEmitter, Output } from '@angular/core';
import { EquitycatalogService } from 'src/app/services/equitycatalog.service';
import { Observable } from 'rxjs';
import { FormControl } from '@angular/forms';
import { startWith, map } from 'rxjs/operators';
import { EquityCatalogItem } from 'src/data/equitycatalogitem';

@Component({
  selector: 'app-equity-selector',
  templateUrl: './equity-selector.component.html',
  styleUrls: ['./equity-selector.component.scss']
})
export class EquitySelectorComponent implements OnInit {
  equities: Observable<EquityCatalogItem[]>;
  options: EquityCatalogItem[] = [];
  equityName = new FormControl();
  filteredOptions: Observable<EquityCatalogItem[]>;

  @Output()
  equityChange = new EventEmitter<string>();

  constructor(private equityCatalogService: EquitycatalogService) { }

  ngOnInit() {
    this.equities = this.equityCatalogService.getAllEquityCatalogItems();
    this.equities.subscribe(
      list => list.forEach(item => this.options.push(item))
    );

    this.filteredOptions = this.equityName.valueChanges
      .pipe(
        startWith(''),
        map(value => typeof value === 'string' ? value : value.symbol),
        map(name => name ? this.filter(name) : this.options)
      );
  }

  private filter(name: string): EquityCatalogItem[] {
    const filterValue = name.toLowerCase();

    return this.options.filter(option => option.symbol.toLowerCase().indexOf(filterValue) === 0);
  }

  equityChanged() {
    this.equityChange.emit(this.equityName.value);
  }

  displayFn(item: EquityCatalogItem): string {
    return item.symbol;
  }
}
