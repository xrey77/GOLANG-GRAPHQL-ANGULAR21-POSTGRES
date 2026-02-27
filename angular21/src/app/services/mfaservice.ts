import { Injectable, inject } from '@angular/core';
import { HttpHeaders } from '@angular/common/http';
import {  Observable } from 'rxjs';
import { Apollo, gql } from 'apollo-angular';

@Injectable({
  providedIn: 'root'
})

export class Mfaservice {
  private readonly apollo = inject(Apollo);

  public sendMfaValidation(idno: number, userdtls: any, token: any): Observable<any> {
    const VERIFY_OTP = gql`
        mutation OtpVerification($id: Int!, $otp: String!) { 
          otpVerification(id: $id, otp: $otp) {
              message
              user {
                  username
              }
          }
        }
      `
      return this.apollo.mutate({
        mutation: VERIFY_OTP,
        variables: { 
          id: idno,
          otp: userdtls.otp
        },
        context: {
          headers: new HttpHeaders().set('Authorization', `Bearer ${token}`)
        }        
    });
  }
}
