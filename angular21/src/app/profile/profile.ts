import { NgOptimizedImage } from '@angular/common';
import { Component, afterNextRender, NgZone, ChangeDetectorRef, OnInit,signal } from '@angular/core';
import { Profileservice } from '../services/profileservice';
import { SessionStorage } from '../services/session-storage';
import { FormGroup, FormControl, Validators, ReactiveFormsModule } from '@angular/forms';
import { Footer } from '../footer/footer';
declare var $: any;

@Component({
  selector: 'app-profile',
  imports: [NgOptimizedImage,ReactiveFormsModule,Footer],
  templateUrl: './profile.html',
  styleUrl: './profile.scss'
})

export class Profile implements OnInit  {  
  profileMsg = signal('');

  passwordChangeForm = new FormGroup({
    newpassword: new FormControl('', Validators.required),
    confnewpassword: new FormControl('', [Validators.required])
  });

  profileForm = new FormGroup({
    firstname: new FormControl('', Validators.required),
    lastname: new FormControl('', Validators.required),
    mobile: new FormControl('', Validators.required)
  });

  profileData: any;
  userId: number = 0;
  jwttoken: any;
  enableMfa: any = [];
  mfa: boolean = false;
  profilepic: string = '';
  email: any = '';
  qrcodeurl: any = null;
  userpicture: string = '';
  showSave: boolean = false;

  constructor(
    private profileService: Profileservice,
    private sessionStorageSevice: SessionStorage,
    private ngZone: NgZone, 
    private cdRef: ChangeDetectorRef    
  ) 
  {
    afterNextRender(() => {
      console.log('Window object is safe to use here:', window.location.href);
    });
  
    const uid = this.sessionStorageSevice.getItem('USERID');
    if (uid) {
     this.userId =  parseInt(uid);
    }

    const jwt = this.sessionStorageSevice.getItem('TOKEN');
    if (jwt) {
     this.jwttoken = jwt;
    }
    this.profileMsg.set("retrieving records...");
    this.profileService.getUserbyId(this.userId, this.jwttoken).subscribe({
     next: (res: any) => {

          if (res.errors) {
                this.profileMsg.set(res.errors[0].message);
                setTimeout(() => {
                  this.profileMsg.set('');
                }, 3000);
                return;
          } 

           const formData = {            
             firstname: res.data.getUserById.firstname,
             lastname: res.data.getUserById.lastname,
             mobile: res.data.getUserById.mobile
            }
            this.profileForm.setValue(formData);
            $("#email").val(res.data.getUserById.email);
            let userpicture: any = `http://127.0.0.1:8000/users/${res.data.getUserById.profilepic}`
            this.profilepic = userpicture;

            if (res.data.getUserById.qrcodeurl !== null) {
              this.qrcodeurl = res.data.getUserById.qrcodeurl;  
              this.mfa = true;
            } else {
              this.mfa = false;
              let qrcode: any = "/images/qrcode.png";
              this.qrcodeurl = qrcode;
            }

            setTimeout(() => {
              this.profileMsg.set('');
            }, 1000);
   
       },
       error: (err: any) => {
         this.profileMsg.set(err.errors[0].message);
         setTimeout(() => {
           this.profileMsg.set('');
         }, 3000);

       }
    });        
  }

  ngOnInit(): void {
    $("#cpwd").hide();
    $("#mfa1").hide();
    $("#mfa2").hide();  
  }

  changeProfilepic(event: any) {
    const file: File = event.target.files[0];
    if (file) {
        this.profileMsg.set('Uploading picture, please wait...');
        $("#pix").attr('src',URL.createObjectURL(file));    
        this.profileService.UploadPicture(this.userId ,file, this.jwttoken).subscribe({
          next: (res: any) => {

            if (res.errors) {
                  this.profileMsg.set(res.errors[0].message);
                  setTimeout(() => {
                    this.profileMsg.set('');
                  }, 3000);
                  return;
            } 

            this.profileMsg.set(res.data.uploadPicture.message);
            $('#twofactor').prop('checked', false);
            $('#changepwd').prop('checked', false);      
            setTimeout(() => {
              let userpicture = `http://127.0.0.1:8000/users/${res.data.uploadPicture.user.profilepic}`
              this.profilepic = userpicture;
              this.sessionStorageSevice.setItem('USERPIC', userpicture);
              this.profileMsg.set('');
              window.location.reload();
            }, 3000);
          },
          error: (err: any) => {  
              console.log("may error...");
              console.log(err.errors[0]);
              this.profileMsg.set(err.errors[0].message);
              setTimeout(() => {
                this.profileMsg.set('');
              }, 3000);

          }      
        });
    }     
  }

  changePassword() {
    if ($('#changepwd').is(":checked")) {
      this.showSave = true;
      $("#cpwd").show();
      $("#mfa1").hide();  
      $("#mfa2").hide();  
      $('#twofactor').prop('checked', false);
    } else {
      this.showSave = false;
      $("#cpwd").hide();
    }            
  }

  onetimePassword() {
    if ($('#twofactor').is(":checked")) {
      this.showSave = true;
      $("#cpwd").hide();
      $("#mfa1").show();
      $("#mfa2").show();
      $('#changepwd').prop('checked', false);
    } else {
      $("#mfa1").hide();  
      $("#mfa2").hide();  
      this.showSave = false;
    }            
  }

  passwordChange() {
    this.ngZone.run(() => {
      if (this.passwordChangeForm.get('newpassword')?.value === '') {
        this.profileMsg.set('Please enter New password...')
        setTimeout(() => {
          this.profileMsg.set('');
        }, 3000)
        return;
      }

      if (this.passwordChangeForm.get('confnewpassword')?.value === '') {
        this.profileMsg.set('Please confirm New password...')
        setTimeout(() => {
          this.profileMsg.set('');
        }, 3000)
        return;
      }

      if (this.passwordChangeForm.get('newpassword')?.value != this.passwordChangeForm.get('confnewpassword')?.value) {
        this.profileMsg.set('New password does not mactched...')
        setTimeout(() => {
          this.profileMsg.set('');
        }, 3000)
        return;
      }
      const formData = {
        'password': this.passwordChangeForm.get('newpassword')?.value
      }
      
      this.profileService.sendNewpasswordRequest(this.userId, formData, this.jwttoken).subscribe({
        next: (res: any) => {


          if (res.errors) {
                this.profileMsg.set(res.errors[0].message);
                setTimeout(() => {
                  this.profileMsg.set('');
                }, 3000);
                return;
          } 

          this.profileMsg.set(res.data.changePassword.message);
          setTimeout(() => {
            this.profileMsg.set('');
          }, 3000);

      },
      error: (err: any) => {
        this.profileMsg.set(err.errors[0].message);
        setTimeout(() => {
          this.profileMsg.set('');
        }, 3000);

      }

    });      

    }); //END-ngZone
  }

  enableMFA(event: any) {
    event.preventDefault();    
    this.profileMsg.set("activating...");
    this.profileService.ActivateMFA(this.userId, true, this.jwttoken).subscribe({
      next: (res: any) => {
          if (res.data.errors) {
            this.profileMsg.set(res.errors[0].message);
            setTimeout(() => {
               this.profileMsg.set('');
            }, 3000);
          } 
          this.profileMsg.set(res.data.mfaActivation.message);
          this.qrcodeurl = res.data.mfaActivation.user.qrcodeurl;
          setTimeout(() => {
              this.profileMsg.set('');
          }, 3000);
          
        },
        error: (err: any) => {
            const errorMsg = err.response?.data?.errors?.[0]?.message || err.message;          
            this.profileMsg.set(errorMsg);
            setTimeout(() => {
              this.profileMsg.set('');
              return;
              // this.qrcodeurl = null;
            }, 3000);
  
        }  
    });
  }

  disableMFA(event: any) {
    event.preventDefault();      
    this.profileMsg.set("de-activating...");
    this.profileService.ActivateMFA(this.userId, false, this.jwttoken).subscribe({
      next: (res: any) => {

          if (res.data.errors) {
            this.profileMsg.set(res.errors[0].message);
            setTimeout(() => {
              this.profileMsg.set('');
            }, 3000);
            return;
          } 

          this.profileMsg.set(res.data.mfaActivation.message);

        let qrcode: any = '/images/qrcode.png';
        this.qrcodeurl = qrcode ;

      },
      error: (err: any) => {

        this.profileMsg.set(err.errors[0].message);
        setTimeout(() => {
          this.profileMsg.set('');
        }, 3000);

      }

    });
    setTimeout(() => {
      this.profileMsg.set('');
    }, 3000);

  }

  submitProfileForm() {
    this.ngZone.run(() => {
        this.profileMsg.set("please wait..");
        const jsonData = { 
          'firstname': this.profileForm.get('firstname')?.value,
          'lastname': this.profileForm.get('lastname')?.value, 
          'mobile': this.profileForm.get('mobile')?.value};
        this.profileService.sendProfileRequest(this.userId,jsonData, this.jwttoken).subscribe({
          next: (res: any) => {

            if (res.errors) {
              this.profileMsg.set(res.errors[0].message);
              setTimeout(() => {
                this.profileMsg.set('');
              }, 3000);
              return;
            } 

            this.profileMsg.set(res.data.updateProfile.message);
            setTimeout(() => {
              this.profileMsg.set('');
            }, 3000);

          },
          error: (err: any) => {
            this.profileMsg.set(err.errors[0].message);
            setTimeout(() => {
              this.profileMsg.set('');
            }, 3000);

          }
          
      });      
    });
  }  

}
