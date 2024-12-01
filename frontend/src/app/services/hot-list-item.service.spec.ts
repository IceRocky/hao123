import { TestBed } from '@angular/core/testing';

import { HotListItemService } from './hot-list-item.service';

describe('HotListItemService', () => {
  let service: HotListItemService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(HotListItemService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
