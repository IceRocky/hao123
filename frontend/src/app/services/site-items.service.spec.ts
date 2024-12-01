import { TestBed } from '@angular/core/testing';

import { SiteItemsService } from './site-items.service';

describe('SiteItemsService', () => {
  let service: SiteItemsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SiteItemsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
