import { Injectable, inject } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import {  Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})

export class Mfaservice {
  private apiUrl = "http://127.0.0.1:8000/graphql";  
  private http = inject(HttpClient);

public sendMfaValidation(idno: number, userdtls: any, token: any): Observable<any> {

      const TOTP_QUERY = `
        mutation MfaVerification($input: VerificationInput!) {
            mfaVerification(input: $input) {
                message
                user {
                    id
                    username
                }
            }
        }
      `

      const headers = new HttpHeaders({
          'Authorization': `Bearer ${token}`
        });

      return this.http.post(
        this.apiUrl, 
        {
          query: TOTP_QUERY,
          variables: { 
            input: {
              "id": idno,
              "otp": userdtls.otp,
            }
          }
        },
        { headers }
      );        
  };  
}
