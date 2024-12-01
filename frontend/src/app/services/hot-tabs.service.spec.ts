import { TestBed } from '@angular/core/testing';

import { HotTabsService } from './hot-tabs.service';

describe('HotTabsService', () => {
  let service: HotTabsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(HotTabsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
