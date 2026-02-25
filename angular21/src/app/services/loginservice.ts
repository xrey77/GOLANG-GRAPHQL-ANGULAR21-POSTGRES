import { Injectable, inject } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
// import { User } from '../interface/user';
import {  map, Observable } from 'rxjs';
// import { LoginResponse } from '../models/auth.types';


@Injectable({
  providedIn: 'root'
})

export class Loginservice {
    private http = inject(HttpClient);
    private apiUrl = "http://127.0.0.1:8000/graphql";

  public sendLoginRequest(userdtls: any): Observable<any> {

    const LOGIN_MUTATION = `
      mutation LoginUser($input: LoginInput!) {
        signIn(input: $input) {
          token
          message
          user {
            id
            firstname
            lastname
            email
            mobile
            username
            isactivated
            isblocked
            mailtoken
            profilepic
            qrcodeurl
          }
        }
      }
    `

     return this.http.post(this.apiUrl, {
      query: LOGIN_MUTATION,
      variables: { 
        input: {
          username: userdtls.username, password: userdtls.password
        }
       }
    });
  };
}
