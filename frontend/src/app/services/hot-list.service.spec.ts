import { TestBed } from '@angular/core/testing';

import { HotListService } from './hot-list.service';

describe('HotListService', () => {
  let service: HotListService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(HotListService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
