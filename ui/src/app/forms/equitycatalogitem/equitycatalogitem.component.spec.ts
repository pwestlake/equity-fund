import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EquitycatalogitemComponent } from './equitycatalogitem.component';

describe('EquitycatalogitemComponent', () => {
  let component: EquitycatalogitemComponent;
  let fixture: ComponentFixture<EquitycatalogitemComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EquitycatalogitemComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EquitycatalogitemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
