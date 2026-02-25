import { Component, signal, afterNextRender} from '@angular/core';
import { Productservice } from '../services/productservice';
import { Footer } from '../footer/footer';

@Component({
  selector: 'app-prodlist',
  imports: [Footer],
  templateUrl: './prodlist.html',
  styleUrl: './prodlist.scss'
})
export class Prodlist {
  message: string = '';
  page: number = 1;
  totpage: any;
  products: any = [];
  totalrecs: any;

  constructor(
    private productsService: Productservice
  ) { 
    this.message = "please wait..";
    this.productList(this.page);

    afterNextRender(() => {
      console.log('Window object is safe to use here:', window.location.href);
    });  
  }

  ngOnInit(): void {
  }

  async productList(page: number) {
    this.productsService.sendProductRequest(page).subscribe({
      next: (res: any) => {
        
          if (res.data.errors) {
            this.message = res.data.errors[0].message;
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
          this.message = '';
        }, 3000);
      }
    });
  }

  lastPage(event: any) {
    event.preventDefault();    
    this.page = this.totpage;
    this.productList(this.page);
    return;    
  }

  nextPage(event: any) {
    event.preventDefault();    
    if (this.page == this.totpage) {
      this.page = this.totpage;
      return;
    }
    this.page++;
    return this.productList(this.page);
  }

  prevPage(event: any) {
    event.preventDefault();    
    if (this.page == 1) {
      return;
    }
    this.page = this.page - 1;
    this.productList(this.page);
    return;    

  }

  firstPage(event: any) {
    event.preventDefault();    
    this.page = 1;
    this.productList(this.page);
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
