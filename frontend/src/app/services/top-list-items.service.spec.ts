import { TestBed } from '@angular/core/testing';

import { TopListItemsService } from './top-list-items.service';

describe('TopListItemsService', () => {
  let service: TopListItemsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TopListItemsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
