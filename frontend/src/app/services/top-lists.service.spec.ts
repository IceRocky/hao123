import { TestBed } from '@angular/core/testing';

import { TopListsService } from './top-lists.service';

describe('TopListsService', () => {
  let service: TopListsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TopListsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
