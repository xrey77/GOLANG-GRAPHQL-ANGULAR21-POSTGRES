import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProductPdfReport } from './product-pdf-report';

describe('ProductPdfReport', () => {
  let component: ProductPdfReport;
  let fixture: ComponentFixture<ProductPdfReport>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ProductPdfReport]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ProductPdfReport);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
