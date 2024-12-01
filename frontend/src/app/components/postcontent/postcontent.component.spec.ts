import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PostcontentComponent } from './postcontent.component';

describe('PostcontentComponent', () => {
  let component: PostcontentComponent;
  let fixture: ComponentFixture<PostcontentComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [PostcontentComponent]
    });
    fixture = TestBed.createComponent(PostcontentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
