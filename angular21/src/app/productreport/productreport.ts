import { Component, inject, OnInit } from '@angular/core';
import { Productservice } from '../services/productservice';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser';
import { NgIf } from '@angular/common';

@Component({
  selector: 'app-productreport',
  imports: [],
  templateUrl: './productreport.html',
  styleUrl: './productreport.scss',
})

export class Productreport implements OnInit {
  private productsService = inject(Productservice);
  private sanitizer = inject(DomSanitizer);

  pdfUrl: SafeResourceUrl | null = null;
  private currentBlobUrl: string | null = null;

  ngOnInit() {
    this.onViewReport();
  }

  onViewReport() {
    this.productsService.showPdfReport().subscribe({
      next: (blob: Blob) => {
        if (this.currentBlobUrl) {
          URL.revokeObjectURL(this.currentBlobUrl);
        }

        this.currentBlobUrl = URL.createObjectURL(blob);
        this.pdfUrl = this.sanitizer.bypassSecurityTrustResourceUrl(this.currentBlobUrl);
      },
      error: (err: any) => console.error('Error downloading PDF:', err)
    });
  }

 ngOnDestroy() {
    if (this.currentBlobUrl) {
      URL.revokeObjectURL(this.currentBlobUrl);
    }
  }  

}
