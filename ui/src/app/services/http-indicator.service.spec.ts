import { TestBed } from '@angular/core/testing';

import { HttpIndicatorService } from './http-indicator.service';

describe('HttpIndicatorService', () => {
  let service: HttpIndicatorService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(HttpIndicatorService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
