import { TestBed } from '@angular/core/testing';

import { GovSitesService } from './gov-sites.service';

describe('GovSitesService', () => {
  let service: GovSitesService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GovSitesService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
