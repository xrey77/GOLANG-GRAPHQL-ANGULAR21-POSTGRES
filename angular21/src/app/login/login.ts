import { NgbModal, NgbModalConfig } from '@ng-bootstrap/ng-bootstrap';
import { Component, signal, afterNextRender,Input, TemplateRef, AfterViewInit } from '@angular/core';
import { Mfa } from "../mfa/mfa";
import { Loginservice } from '../services/loginservice';
import { Logoutservice } from '../services/logoutservice';
import { SessionStorage } from '../services/session-storage';
import { FormGroup, FormControl, Validators, ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';

declare var $: any;

@Component({
  selector: 'app-login',
  imports: [Mfa, ReactiveFormsModule],
  templateUrl: './login.html',
  styleUrl: './login.scss',
	providers: [NgbModalConfig, NgbModal],  
})

export class Login implements AfterViewInit{
  @Input() templateRef?: TemplateRef<any>;

  loginMessage = signal('');
  isDisabled: boolean = false;

  loginForm = new FormGroup({
    username: new FormControl('', [Validators.required]),
    password: new FormControl('', [Validators.required])
  });


  constructor(  
    private router: Router,    
    config: NgbModalConfig,
		private modalService: NgbModal,
    private loginService: Loginservice,
    private logoutService: Logoutservice,
    private sessionStorageService: SessionStorage
  ) { 
    config.backdrop = 'static';
		config.keyboard = false;    
      afterNextRender(() => {
        console.log('Window object is safe to use here:', window.location.href);
      });  
  }
  ngAfterViewInit(): void {
  }

  submitLoginForm() {
    if(this.loginForm.valid)
    {
       this.loginMessage.set('please wait...');
       this.isDisabled = true;
       this.loginService.sendLoginRequest(this.loginForm.value).subscribe({
         next: (res: any) => {
            if (res.errors) {
                  this.loginMessage.set(res.errors[0].message);
                  setTimeout(() => {
                    this.loginMessage.set('');
                    this.isDisabled = false;
                  }, 3000);
                  return;
            } 
            this.loginMessage.set(res.data.signinUser.message);
            this.sessionStorageService.setItem("USERID", res.data.signinUser.user.id);
            this.sessionStorageService.setItem("TOKEN", res.data.signinUser.token);
            this.sessionStorageService.setItem("ROLES", res.data.signinUser.roles);
            let userpicture = `http://localhost:5000/assets/users/${res.data.signinUser.user.userpicture}`
            this.sessionStorageService.setItem("USERPIC", userpicture);

            if (res.data.signinUser.user.qrcodeurl !== null) {
              this.isDisabled = false;
              $("#reset").trigger('click'); 
              $("#hideLogin").trigger('click');
              $("#showMfa").trigger('click');
              this.loginMessage.set('');                          
            } else {
              this.sessionStorageService.setItem("USERNAME", res.data.signinUser.user.username);  
              $("#reset").trigger('click');
              $("#hideLogin").trigger('click');
              this.router.navigate(['/']); 
              setTimeout(() => {
              this.loginMessage.set('');
              this.isDisabled = false;
                location.reload();
              }, 1000);
            }
          
          },
          error: (err: any) => {
            alert("error 2");
            console.log(err.errors);
            this.loginMessage.set(err.errors.message);
            setTimeout(() => {
              this.loginMessage.set('');
              this.isDisabled = false;
            }, 3000);

          }

      });
    }
  }

  public loginOpen(loginTemplate: any): void {
		this.modalService.open(loginTemplate, { size: 'sm', centered: true });
	}

  closeLogin() {
    $("#reset").trigger('click');
    $("#hideLogin").trigger('click');
    this.router.navigate(['/']); 
  }

}
