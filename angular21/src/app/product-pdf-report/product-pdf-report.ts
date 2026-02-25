import { Component, inject, signal } from '@angular/core';
import { Productservice } from '../services/productservice';
import { TDocumentDefinitions } from 'pdfmake/interfaces';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser'; 
import "pdfmake/build/vfs_fonts"; 

@Component({
  selector: 'app-product-pdf-report',
  standalone: true,
  templateUrl: './product-pdf-report.html',
  styleUrl: './product-pdf-report.scss',
})
export class ProductPdfReport {
  private productsService = inject(Productservice);
  products = signal<any[]>([]);
  message = signal<string>('Loading report, please wait...');
  pdfUrl = signal<SafeResourceUrl | null>(null);
  private sanitizer = inject(DomSanitizer);

  constructor() {
    this.productDataList();
    
  }

async getBase64ImageFromUrl(imageUrl: string): Promise<string> {
  const res = await fetch(imageUrl);
  const blob = await res.blob();
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.addEventListener("load", () => resolve(reader.result as string), false);
    reader.onerror = () => reject(reader.error);
    reader.readAsDataURL(blob);
  });
}  

productDataList() {
  this.productsService.productDataRequest().subscribe({
    next: async (res) => {
      if (res.data) {

        const logoBase64 = await this.getBase64ImageFromUrl('/images/logo.png');
        const data = res.data.productData;
        this.products.set(data);

        const pdfMake = (await import('pdfmake/build/pdfmake')).default;
        const pdfFonts = (await import('pdfmake/build/vfs_fonts')).default;
        pdfMake.addVirtualFileSystem(pdfFonts);
        
        const docDefinition: TDocumentDefinitions = {
          content: [
            {
              image: logoBase64,
              width: 150,
              alignment: 'left',
              margin: [0, 0, 0, 20]
            },            
          {
            stack: [
              { text: 'Monthly Sales Report', style: 'header' },
              { 
                text: 'As of February 13, 2026', 
                alignment: 'left', 
                margin: [0, -10, 0, 0] // [left, top, right, bottom]
              }
            ]                        
          },
            {
              table: {
                headerRows: 1,
                widths: ['*', '*', '*'],
                body: [
                  ['ID', 'Descriptions', 'Price'],
                  ...data.map((p: any) => [
                    p.id.toString(),
                    p.descriptions,
                    `$${p.sellprice?.toFixed(2) || '0.00'}`
                  ])
                ]
              }
            }
          ],
          footer: (currentPage, pageCount) => {
            return {
              columns: [
                {
                  text: `Page ${currentPage} of ${pageCount}`,
                  alignment: 'right',
                  margin: [0, 0, 40, 0]
                }
              ],
              style: 'footerStyle'
            };
          },
          styles: {
            header: { fontSize: 18, bold: true, margin: [0, 0, 0, 10] },
            footerStyle: { fontSize: 10, color: '#555' }
          }
        };

        if (typeof window !== 'undefined') {
            const pdfDocGenerator = pdfMake.createPdf(docDefinition);
            pdfDocGenerator.download('sales-report.pdf');
            const dataUrl = await pdfDocGenerator.getDataUrl(); 

            const safeUrl = this.sanitizer.bypassSecurityTrustResourceUrl(dataUrl);
            this.pdfUrl.set(safeUrl);
        }
      }
    }
  });
 }
}