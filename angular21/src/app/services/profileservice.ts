import { Injectable, inject } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import {  Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})

export class Profileservice {
  private apiUrl = "http://127.0.0.1:8000/graphql";
  private http = inject(HttpClient)
    

  public getUserbyId(idno: number, token: any): Observable<any> {

    const GETUSERID_QUERY = `
      query getUserById($id: Int!) {
          getUserById(id: $id) {
              id
              firstname
              lastname
              email
              mobile
              username
              profilepic
              qrcodeurl
          }
      }
    `
      const headers = new HttpHeaders({
          'Authorization': `Bearer ${token}`
        });

      return this.http.post(
        this.apiUrl, 
        {
          query: GETUSERID_QUERY,
          variables: { id: idno }
        },
        { headers }
      );        
  }

  public ActivateMFA(idno: number, isenabled: boolean, token: string) {
    const MFAACTIVATIOIN_QUERY = `
        mutation MfaActivation($input: MfaInput!) {
            mfaActivation(input: $input) {
                message
                user {
                    id
                    qrcodeurl
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
          query: MFAACTIVATIOIN_QUERY,
          variables: { 
            input: {
              "id": idno,
              "twofactorenabled": isenabled,
            }
          }
        },
        { headers }
      );        
  }

  public UploadPicture(idno: number, file: File, token: any): Observable<any> {
      const operations = {
          query: `
            mutation UploadPicture($id: Int!, $file: Upload!) {
              uploadPicture(id: $id, profilepic: $file) {
                message
                user {
                  id
                  profilepic
                }
              }
            }
          `,
          variables: {
              id: idno,
              file: null
          }
        };

        const map = {
          '0': ['variables.file']
        };

        const formData = new FormData();
        formData.append('operations', JSON.stringify(operations));
        formData.append('map', JSON.stringify(map));
        formData.append('0', file);

        const headers = new HttpHeaders({
          'Authorization': `Bearer ${token}`
        });

        return this.http.post(this.apiUrl, formData, { headers });        
  }

  public sendProfileRequest(idno: number, userdtls: any, token: any): Observable<any> {
    const PROFILEUPDATE_QUERY = `
        mutation UpdateProfile($input: ProfileInput!) {
            updateProfile(input: $input) {
                message
                user {
                    id
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
          query: PROFILEUPDATE_QUERY,
          variables: { 
            input: {
              "id": idno,
              "firstname": userdtls.firstname,
              "lastname": userdtls.lastname,
              "mobile": userdtls.mobile
            }
          }
        },
        { headers }
      );        
  }  

  public sendNewpasswordRequest(idno: number, userdtls: any, token: any): Observable<any> {
    const CHANGEPASSWORD_QUERY = `
        mutation ChangePassword($input: PasswordInput!) {
            changePassword(input: $input) {
                message
                user {
                    id
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
          query: CHANGEPASSWORD_QUERY,
          variables: { 
            input: {
              "id": idno,
              "password": userdtls.password,
            }
          }
        },
        { headers }
      );        
  }  
  
}
