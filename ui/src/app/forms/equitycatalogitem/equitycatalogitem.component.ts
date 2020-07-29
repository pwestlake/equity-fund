import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators } from '@angular/forms';
import { Location } from '@angular/common';
import { EquitycatalogService } from 'src/app/services/equitycatalog.service';
import { EquityCatalogItem } from 'src/data/equitycatalogitem';

@Component({
  selector: 'app-equitycatalogitem',
  templateUrl: './equitycatalogitem.component.html',
  encapsulation: ViewEncapsulation.None,
  host: {
    'class': 'router-flex'
  },
  styleUrls: ['./equitycatalogitem.component.scss']
})
export class EquitycatalogitemComponent implements OnInit {
  equityForm: FormGroup;

  constructor(private formBuilder: FormBuilder, private location: Location, private equityCatalogService: EquitycatalogService) { }

  ngOnInit(): void {
    this.equityForm = this.formBuilder.group({
      symbol: new FormControl('', [Validators.required, Validators.pattern("^[A-Z]*\.XLON$")]),
      lsetidm: new FormControl('', [Validators.required, Validators.pattern("^[A-Z]*$")]),
      lseissuername: new FormControl('', [Validators.required, Validators.maxLength(40)]),
      lsetabid: new FormControl('771b9c49-382e-4e74-bd94-e96af5c94285', [Validators.required, Validators.pattern("^[a-f0-9]{8}?-[a-f0-9]{4}?-[a-f0-9]{4}?-[a-f0-9]{4}?-[a-f0-9]{12}?$")]),
      lsecomponentid: new FormControl('eb11eb09-4797-469c-a6ca-a258d2a53d60', [Validators.required, Validators.pattern("^[a-f0-9]{8}?-[a-f0-9]{4}?-[a-f0-9]{4}?-[a-f0-9]{4}?-[a-f0-9]{12}?$")]),
    })
  }

  onSubmit() {
    const item: EquityCatalogItem = {
      id: undefined,
      symbol: this.equityForm.value.symbol,
      lsetidm: this.equityForm.value.lsetidm,
      lseissuername: this.equityForm.value.lseissuername,
      lsetabid: this.equityForm.value.lsetabid,
      lsecomponentid: this.equityForm.value.lsecomponentid,
      datecreated: undefined,
      lastmodified: undefined
    };

    this.equityCatalogService.addEquityCatalogItem(item);
    this.location.back();
  }
}
