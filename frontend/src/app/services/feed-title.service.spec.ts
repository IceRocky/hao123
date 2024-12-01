import { TestBed } from '@angular/core/testing';

import { FeedTitleService } from './feed-title.service';

describe('FeedTitleService', () => {
  let service: FeedTitleService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(FeedTitleService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
