import { Injectable, inject } from '@angular/core';
import {  map, Observable } from 'rxjs';
import { Apollo, gql } from 'apollo-angular';


@Injectable({
  providedIn: 'root'
})

export class Loginservice {
    private readonly apollo = inject(Apollo);

  public sendLoginRequest(userdtls: any): Observable<any> {

    const SIGNIN_USER = gql`
        mutation SigninUser($username: String!, $password: String!) {
          signinUser(username: $username, password: $password) {
              token
              rolename
              message
              user {
                  id
                  firstname
                  lastname
                  email
                  mobile
                  username
                  roles
                  userpicture
                  qrcodeurl
              }
          }
        }
      `
      return this.apollo.mutate({
        mutation: SIGNIN_USER,
        variables: { 
          username: userdtls.username,
          password: userdtls.password },
      })
  }
}