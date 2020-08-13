import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EquitySelectorComponent } from './equity-selector.component';

describe('EquitySelectorComponent', () => {
  let component: EquitySelectorComponent;
  let fixture: ComponentFixture<EquitySelectorComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EquitySelectorComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EquitySelectorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
