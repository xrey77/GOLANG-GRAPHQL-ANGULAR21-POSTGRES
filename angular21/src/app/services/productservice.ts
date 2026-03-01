import { Injectable, inject } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import {  map, Observable } from 'rxjs';
import { Apollo, gql } from 'apollo-angular';

@Injectable({
  providedIn: 'root'
})

export class Productservice {
  private readonly apollo = inject(Apollo);
  private http = inject(HttpClient);
  private apiUrl = "";

  public sendSearchRequest(page: number, keyword: any): Observable<any>
  {
    const SEARCH_QUERY = gql`
      query ProductSearch($key: String!, $page: Int!) {
          productSearch(key: $key, page: $page) {
            page
            totalpage
            totalrecords
            products {
                id
                category
                descriptions
                qty
                unit
                sellprice
                productpicture
            }
          }
      }
    `
      return this.apollo.query({
        query: SEARCH_QUERY,
        variables: { 
          key: keyword,
          page: page,
        }
    });
  }

  public productDataRequest(): Observable<any> {
    return this.apollo.watchQuery<any>({
      query: gql`
        query ProductData {
          productData {
            id
            category
            descriptions
            qty
            unit
            sellprice
            productpicture
          }
        }
      `,
    }).valueChanges;
  }

  public sendProductRequest(page: number): Observable<any>
  {
    const LIST_QUERY = gql`
      query ProductList($page: Int!) {
          productList(page: $page) {
            page
            totalpage
            totalrecords
            products {
                id
                category
                descriptions
                qty
                unit
                costprice
                sellprice
                saleprice
                productpicture
                alertstocks
                criticalstocks
            }
          }
      }
    `
      return this.apollo.query({
        query: LIST_QUERY,
        variables: { 
          page: page,
        }
    });
  } 
  
public showPdfReport(): Observable<any> {
    const REPORT_QUERY = gql`
      query PdfQuery {
          pdfQuery {
            id
            category
            descriptions
            qty
            unit
            costprice
            sellprice
            saleprice
            productpicture
            alertstocks
            criticalstocks
          }
        }
      `;

    return this.apollo.query({
        query: REPORT_QUERY,
    });
}

  public showSalesGraph(): Observable<any> {
      const SALES_QUERY = gql`
        query GetSales {
          getSales {
            saleamount
            salesdate
          }
        }
      `;

      return this.apollo.query({
        query: SALES_QUERY,
    });
  }
}
