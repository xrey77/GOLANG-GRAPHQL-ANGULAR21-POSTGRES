import { Injectable, inject } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import {  Observable } from 'rxjs';
import { Apollo, gql } from 'apollo-angular';

@Injectable({
  providedIn: 'root'
})

export class Profileservice {
  private readonly apollo = inject(Apollo);
    
  
  public getUserbyId(idno: number, token: any): Observable<any> {
    const GETUSER_ID = gql`
        query GetUserID($id: Int!) {        
          getUserID(id: $id) {
              user {
                  id
                  firstname
                  lastname
                  email
                  mobile
                  username
                  qrcodeurl
                  userpicture
              }
          }
        }
      `
      return this.apollo.query({
        query: GETUSER_ID,
        variables: { 
          id: idno,
        },
        context: {
          headers: new HttpHeaders().set('Authorization', `Bearer ${token}`)
        }        
    });
  }

  public ActivateMFA(idno: number, isenabled: boolean, token: string) {
    const ACTIVATE_MFA = gql`
        mutation MfaActivation($id: Int!, $twofactorenabled: Boolean!) {        
          mfaActivation(id: $id, twofactorenabled: $twofactorenabled) {
              message
              user {
                  qrcodeurl
              }
          }
        }
      `
      return this.apollo.mutate({
        mutation: ACTIVATE_MFA,
        variables: { 
          id: idno,
          twofactorenabled: isenabled
        },
        context: {
          headers: new HttpHeaders().set('Authorization', `Bearer ${token}`)
        }        
    });

  }

 public UploadPicture(idno: number, file: File, token: string): Observable<any> {
    const UPLOAD_PICTURE = gql`
      mutation UploadPicture($id: Int!, $userpicture: Upload!) {
        uploadPicture(id: $id, userpicture: $userpicture) {
          message
          user {
            id
          }
        }
      }
    `;  

    return this.apollo.mutate({
      mutation: UPLOAD_PICTURE,
      variables: {
        id: idno,
        userpicture: file // Pass the File object directly
      },
      context: {
        useMultipart: true, // Required for file uploads
          headers: new HttpHeaders()
          .set('Authorization', `Bearer ${token}`) // Added closing backtick and parenthesis
          .set('Apollo-Require-Preflight', 'true')
      }
    });
  }

  // public async UploadPicture(idno: number, file: File, token: any): Promise<Observable<any>> {


    // const UPLOAD_PICTURE = gql`
    //     mutation UploadPicture($id : Int!, $userpicture: Upload!) {        
    //       uploadPicture(id: $id, userpicture: $userpicture) {
    //           message
    //           user {
    //               id
    //           }
    //       }
    //     }
    //   `

    //   return this.apollo.mutate({
    //     mutation: formData,
    //     variables: { 
    //       id: idno,
    //       userpicture: file
    //     },
    //     context: {
    //       useMultipart: true,          
        //   headers: new HttpHeaders()
        //   .set('Authorization', `Bearer ${token}`) // Added closing backtick and parenthesis
        //   .set('Apollo-Require-Preflight', 'true')
        // }        
    // });
  // }

  public sendProfileRequest(idno: number, userdtls: any, token: any): Observable<any> {
    const PROFILE_UPDATE = gql`
        mutation UpdateProfile($id: Int!,$firstname: String!, $lastname: String!, $mobile: String!) {        
          updateProfile(id: $id, firstname: $firstname, lastname: $lastname, mobile: $mobile) {
              message
              user {
                  id
              }
          }
        }
      `
      return this.apollo.mutate({
        mutation: PROFILE_UPDATE,
        variables: { 
          id: idno,
          firstname: userdtls.firstname,
          lastname: userdtls.lastname,
          mobile: userdtls.mobile
        },
        context: {
          headers: new HttpHeaders().set('Authorization', `Bearer ${token}`)
        }        
    });
     
  }  

  public sendNewpasswordRequest(idno: number, userdtls: any, token: any): Observable<any> {
    const UPDATE_PASSWORD = gql`
        mutation UpdatePassword($id: Int!, $password: String!) {        
          updatePassword(id: $id, password: $password) {
              message
              user {
                  id
              }
          }
        }
      `
      return this.apollo.mutate({
        mutation: UPDATE_PASSWORD,
        variables: { 
          id: idno,
          password: userdtls.password
        },
        context: {
          headers: new HttpHeaders().set('Authorization', `Bearer ${token}`)
        }        
    });
  }
}