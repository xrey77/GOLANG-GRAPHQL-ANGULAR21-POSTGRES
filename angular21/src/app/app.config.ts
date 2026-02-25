import { NgbModal, NgbModalConfig } from '@ng-bootstrap/ng-bootstrap';
import { ApplicationConfig, provideBrowserGlobalErrorListeners, provideZoneChangeDetection, inject } from '@angular/core';
import { provideRouter } from '@angular/router';
import { provideHttpClient, withFetch } from '@angular/common/http';
import { routes } from './app.routes';
import { provideClientHydration, withEventReplay } from '@angular/platform-browser';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { provideApollo } from 'apollo-angular';
import { HttpLink } from 'apollo-angular/http';
import { InMemoryCache } from '@apollo/client';

export const appConfig: ApplicationConfig = {
  providers: [
    provideHttpClient(withFetch()),
    provideBrowserGlobalErrorListeners(),
    NgbModule,
    NgbModalConfig, NgbModal,
    provideZoneChangeDetection({ eventCoalescing: true }),
    provideRouter(routes), provideClientHydration(withEventReplay()), provideHttpClient(), provideHttpClient(), provideApollo(() => {
      const httpLink = inject(HttpLink);

      return {
        link: httpLink.create({
          uri: 'http://127.0.0.1:8000/graphql',
        }),
        cache: new InMemoryCache(),
      };
    })    
  ]
};