import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Saleschart } from './saleschart';

describe('Saleschart', () => {
  let component: Saleschart;
  let fixture: ComponentFixture<Saleschart>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Saleschart]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Saleschart);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
