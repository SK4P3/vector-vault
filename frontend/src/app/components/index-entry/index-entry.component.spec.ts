import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IndexEntryComponent } from './index-entry.component';

describe('IndexEntryComponent', () => {
  let component: IndexEntryComponent;
  let fixture: ComponentFixture<IndexEntryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ IndexEntryComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(IndexEntryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
