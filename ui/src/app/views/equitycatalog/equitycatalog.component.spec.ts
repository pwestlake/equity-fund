import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EquitycatalogComponent } from './equitycatalog.component';

describe('EquitycatalogComponent', () => {
  let component: EquitycatalogComponent;
  let fixture: ComponentFixture<EquitycatalogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EquitycatalogComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EquitycatalogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
