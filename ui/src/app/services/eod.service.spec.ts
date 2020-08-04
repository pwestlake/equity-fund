import { TestBed } from '@angular/core/testing';

import { EodService } from './eod.service';

describe('EodService', () => {
  let service: EodService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(EodService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
