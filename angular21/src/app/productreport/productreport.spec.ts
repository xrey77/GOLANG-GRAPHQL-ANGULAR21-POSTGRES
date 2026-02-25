import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Productreport } from './productreport';

describe('Productreport', () => {
  let component: Productreport;
  let fixture: ComponentFixture<Productreport>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Productreport]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Productreport);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
