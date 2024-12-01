import { TestBed } from '@angular/core/testing';

import { HotSearchsService } from './hot-searchs.service';

describe('HotSearchsService', () => {
  let service: HotSearchsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(HotSearchsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
