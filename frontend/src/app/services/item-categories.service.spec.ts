import { TestBed } from '@angular/core/testing';

import { ItemCategoriesService } from './item-categories.service';

describe('ItemCategoriesService', () => {
  let service: ItemCategoriesService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ItemCategoriesService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
