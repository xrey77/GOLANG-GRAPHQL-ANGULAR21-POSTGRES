import { Component, inject, OnInit, signal } from '@angular/core';
import { Productservice } from '../services/productservice';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser';
import PDFDocument from 'pdfkit-table';
import blobStream from 'blob-stream';

@Component({
  selector: 'app-productreport',
  standalone: true,
  templateUrl: './productreport.html',
  styleUrl: './productreport.scss',
})
export class Productreport implements OnInit {
  private productsService = inject(Productservice);
  private sanitizer = inject(DomSanitizer);
  public pdfUrl = signal<SafeResourceUrl | null>(null);

  public generatePdfFromGraphql() {
    this.productsService.showPdfReport().subscribe(({ data }) => {
      this.createPdf(data.pdfQuery);
    });
  }

  ngOnInit(): void {
    this.generatePdfFromGraphql();      
  }

  private async createPdf(items: any[]) {
    const doc = new PDFDocument({ margin: 30, size: 'A4', bufferPages: true });

    const stream = doc.pipe(blobStream());
    const logoUrl = 'assets/images/logo.png'; 
    const logoWidth = 100;
    const x = (doc.page.width - logoWidth) / 2;

    try {
      doc.image(logoUrl, x, doc.y, { fit: [logoWidth, 50] });
    } catch (e) {
      doc.text('Logo Missing', { align: 'center' });
    }

    doc.moveDown(3);
    doc.fontSize(16).text('Product List', { align: 'center' });
    doc.moveDown();

    const tableData = {
      headers: ["#", "Description", "Qty", "Unit", "Cost Price", "Sell Price"],
      rows: items.map((p: any, i: number) => [
        i + 1, p.descriptions, p.qty, p.unit, p.costPrice, p.sellPrice
      ])
    };

    await doc.table(tableData, {
      prepareHeader: () => doc.font("Helvetica-Bold").fontSize(10),
      prepareRow: () => doc.font("Helvetica").fontSize(10),
      columnsSize: [30, 220, 50, 50, 70, 70],
    });

    doc.end();
    stream.on('finish', () => {
      const blobUrl = stream.toBlobURL('application/pdf');
      this.pdfUrl.set(this.sanitizer.bypassSecurityTrustResourceUrl(blobUrl));
    });
  }
}
