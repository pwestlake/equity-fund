import { TestBed } from '@angular/core/testing';

import { EquitycatalogService } from './equitycatalog.service';

describe('EquitycatalogService', () => {
  let service: EquitycatalogService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(EquitycatalogService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
