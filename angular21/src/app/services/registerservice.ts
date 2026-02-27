import { inject, Injectable } from '@angular/core';
import {  Observable } from 'rxjs';
import { Apollo, gql } from 'apollo-angular';

@Injectable({
  providedIn: 'root'
})

export class Registerservice {
  private readonly apollo = inject(Apollo);
  
  public sendRegistrationRequest(userdtls: any): Observable<any> {

    const SIGNUP_USER = gql`
        mutation SignupUser(
          $firstname: String!,
          $lastname: String!,
          $email: String!,
          $mobile: String!,
          $username: String!,
          $password: String!) {        
          signupUser(
            firstname: $firstname,
            lastname: $lastname,
            email: $email,
            mobile: $mobile,
            username: $username,
            password: $password          
          ) {
              message
              user {
                  id
              }
          }
        }
      `
      return this.apollo.mutate({
        mutation: SIGNUP_USER,
        variables: { 
          firstname: userdtls.firstname,
          lastname: userdtls.lastname,
          email: userdtls.email,
          mobile: userdtls.mobile,
          username: userdtls.username,
          password: userdtls.password },
      })
  }    

}
