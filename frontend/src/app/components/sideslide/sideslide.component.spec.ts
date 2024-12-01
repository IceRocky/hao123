import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SideslideComponent } from './sideslide.component';

describe('SideslideComponent', () => {
  let component: SideslideComponent;
  let fixture: ComponentFixture<SideslideComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [SideslideComponent]
    });
    fixture = TestBed.createComponent(SideslideComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
