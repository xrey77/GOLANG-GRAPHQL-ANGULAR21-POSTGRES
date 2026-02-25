import { Routes } from '@angular/router';
import { Home } from './home/home';
import { About } from './about/about';
import { Contact } from './contact/contact';
import { Profile } from './profile/profile';
import { Prodlist } from './prodlist/prodlist';
import { Prodcatalog } from './prodcatalog/prodcatalog';
import { Prodsearch } from './prodsearch/prodsearch';
import { ErrorPageComponent } from './error-page-component/error-page-component';
import { Productreport } from './productreport/productreport';
import { Saleschart } from './saleschart/saleschart';
import { ProductPdfReport } from './product-pdf-report/product-pdf-report';

export const routes: Routes = [
    { path: '', component: Home, title: 'Apple Inc.'},
    { path: 'static/', component: Home, title: 'Apple Inc.'},
    { path: 'about', component: About, title: 'About Us' },
    { path: 'contact', component: Contact, title: 'Contact Us' },
    { path: 'profile', component: Profile, title: 'User Profile' },
    { path: 'productlist', component: Prodlist, title: 'Products List' },
    { path: 'productcatalog', component: Prodcatalog,title: 'Products Catalog' },
    { path: 'productsearch', component: Prodsearch, title: 'Product Search' },
    { path: 'productreport', component: Productreport, title: 'Product Search' },
    { path: 'saleschart', component: Saleschart, title: 'Product Search' },
    { path: 'pdfreport', component: ProductPdfReport, title: 'Product Pdf Report' },
    { path: '**', component: ErrorPageComponent }    
];
