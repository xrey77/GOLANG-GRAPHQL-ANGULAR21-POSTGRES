import { Component } from '@angular/core';
import { Productservice } from '../services/productservice';
import { Footer } from '../footer/footer';

@Component({
  selector: 'app-prodcatalog',
  imports: [Footer],
  templateUrl: './prodcatalog.html',
  styleUrl: './prodcatalog.scss'
})
export class Prodcatalog {

  message: any;
  page: number = 1;
  totpage: any = [];
  totalrecs: any;
  products: any;

  constructor(
    private productsService: Productservice
  ) { 
    this.message = "please wait..";
    this.productsCatalog(this.page);
  }

  async productsCatalog(page: any) {
    this.productsService.sendProductRequest(page).subscribe({
      next: (res: any) => {

          if (res.data.errors) {
            this.message = res.errors[0].message;
            setTimeout(() => {
               this.message = '';
            }, 3000);
          } 


          this.page = res.data.productsList.paginatorInfo.currentPage;
          this.totpage = res.data.productsList.paginatorInfo.lastPage
          this.products = res.data.productsList.data;
          this.totalrecs = res.data.productsList.paginatorInfo.total

          window.setTimeout(() => {
            this.message = '';
          }, 1000);

    },
    error: (err: any) => {
        const errorMsg = err.response?.data?.errors?.[0]?.message || err.message;      
        this.message = errorMsg;
        window.setTimeout(() => {
          this.products = [];
          this.message = '';
        }, 3000);
      }
    });
  }

  lastPage(event: any) {
    event.preventDefault();    
    this.page = this.totpage;
    this.productsCatalog(this.page);
    return;    
  }

  nextPage(event: any) {
    event.preventDefault();    
    if (this.page == this.totpage) {
      return;
    }
    this.page++;
    return this.productsCatalog(this.page);
  }

  prevPage(event: any) {
    event.preventDefault();    
    if (this.page == 1) {
      return;
    }
    this.page = this.page - 1;
    this.productsCatalog(this.page);
    return;    

  }

  firstPage(event: any) {
    event.preventDefault();    
    this.page = 1;
    this.productsCatalog(this.page);
    return;    
  }  

  toDecimal(nos: any) {
    const formatter = new Intl.NumberFormat('en-US', {
      minimumFractionDigits: 2,
      maximumFractionDigits: 2,
    });
    return formatter.format(nos);
  }


}
