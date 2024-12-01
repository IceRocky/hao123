import { TestBed } from '@angular/core/testing';

import { CommonSitesService } from './common-sites.service';

describe('CommonSitesService', () => {
  let service: CommonSitesService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(CommonSitesService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
