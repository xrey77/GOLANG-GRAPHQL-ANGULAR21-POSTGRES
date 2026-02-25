import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {  Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})

export class Registerservice {

    private http = inject(HttpClient);
    private apiUrl = "http://127.0.0.1:8000/graphql";
  
  public sendRegistrationRequest(userdtls: any): Observable<any> {

    const REGISTER_MUTATION = `
      mutation CreateUser($input: SignupInput!) {
        registerUser(input: $input) { 
          message 
          user {            
              id
              firstname
              lastname
              email
              mobile
              username
          }
        }
      }
    `
     return this.http.post(this.apiUrl, {
      query: REGISTER_MUTATION,
      variables: { 
        input: {
          firstname: userdtls.firstname,
          lastname: userdtls.lastname,
          email: userdtls.email,
          mobile: userdtls.mobile,
          username: userdtls.username,
          password: userdtls.password
        }
       }
    });

  }    

}
