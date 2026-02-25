import { NgbModal, NgbModalConfig } from '@ng-bootstrap/ng-bootstrap';
import { Component, signal, afterNextRender,Input, TemplateRef, AfterViewInit } from '@angular/core';
import { Mfaservice } from '../services/mfaservice';
import { Logoutservice } from '../services/logoutservice';
import { SessionStorage } from '../services/session-storage';
import { FormGroup, FormControl, Validators, ReactiveFormsModule } from '@angular/forms';
declare var $: any;

@Component({
  selector: 'app-mfa',
  imports: [ReactiveFormsModule],
  templateUrl: './mfa.html',
  styleUrl: './mfa.scss',
	providers: [NgbModalConfig, NgbModal],  
})

export class Mfa {
  @Input() templateRefMfa?: TemplateRef<any>;
  message = signal('');
  isDisabled: boolean = false;
  userId: number = 0;
  userToken: any = '';

  constructor(    
    config: NgbModalConfig,
		private modalService: NgbModal,
    private mfaService: Mfaservice,
    private logoutService: Logoutservice,
    private sessionStorageService: SessionStorage
  ) { 
      config.backdrop = 'static';
      config.keyboard = false;        
      afterNextRender(() => {
        // This code runs only in the browser, after the next render cycle
        console.log('Window object is safe to use here:', window.location.href);
      });  
  }

  mfaForm = new FormGroup({
    otpcode: new FormControl('', [Validators.required]),
  });


  submitMfaForm(event: any) {
    event.preventDefault();
    const uid = this.sessionStorageService.getItem('USERID');
    if (uid) {
      this.userId = parseInt(uid);
    }

    const uToken = this.sessionStorageService.getItem('TOKEN');
    if (uToken) {
      this.userToken = uToken;
    }

    if(this.mfaForm.valid)
    {
       this.message.set('please wait...');
       let otpcode: any | null = this.mfaForm.get('otpcode')?.value;
       const formData = {
        'otp': otpcode
       }
       this.mfaService.sendMfaValidation(this.userId, formData, this.userToken).subscribe({
         next: (res: any) => {

              if (res.errors) {
                this.message.set(res.errors[0].message);
                setTimeout(() => {
                  this.message.set('');
                  return true;
                }, 3000);
              } else {
                this.message.set(res.data.mfaVerification.message);
                this.sessionStorageService.setItem("USERNAME", res.data.mfaVerification.user.username);  
                setTimeout(() => {
                  this.message.set('');
                  $("#reset").trigger('click')
                  return location.reload();
                }, 6000);
              }
          },
          error: (err: any) => {
            this.message.set(err.errors[0].message);
            setTimeout(() => {
              $("#reset").trigger('click')
              this.message.set('');
              this.isDisabled = false;
            }, 3000);
            return;
          }
      });
    
    }
  }    

  closeMfa() {
    $("#reset").trigger('click');
    this.sessionStorageService.removeItem('USERNAME');
    this.sessionStorageService.removeItem('USERPIC');
    this.sessionStorageService.removeItem('USERID');
    this.sessionStorageService.removeItem('TOKEN');
    this.sessionStorageService.clear();  
    location.href="/";
    location.reload();
  }

  public mfaOpen(mfaTemplate: any): void {
		this.modalService.open(mfaTemplate, { size: 'sm', centered: true });
	}


}
