import { NgbModal, NgbModalConfig } from '@ng-bootstrap/ng-bootstrap';
import { Component, signal, afterNextRender,Input, TemplateRef, AfterViewInit } from '@angular/core';
import { Registerservice } from '../services/registerservice';
import { FormGroup, FormControl, Validators, ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';
declare var $: any;


@Component({
  selector: 'app-register',
  imports: [ReactiveFormsModule],
  templateUrl: './register.html',
  styleUrl: './register.scss'
})

export class Register {
  @Input() templateRefReg?: TemplateRef<any>;

  registrationMessage = signal('');
  
  registrationForm = new FormGroup({
    firstname: new FormControl('', Validators.required),
    lastname: new FormControl('', Validators.required),
    email: new FormControl('', Validators.required),
    mobile: new FormControl('', Validators.required),
    username: new FormControl('', Validators.required),
    password: new FormControl('', [Validators.required]),
  });

  constructor(
    private router: Router,
    config: NgbModalConfig,
		private modalService: NgbModal,
    private registrationService: Registerservice,
  ) { 
      config.backdrop = 'static';
		  config.keyboard = false;    
      afterNextRender(() => {
        console.log('Window object is safe to use here:', window.location.href);
      });  
  }

  submitRegistrationForm() {
    if(this.registrationForm.valid)
    {
      this.registrationMessage.set('please wait...');
       this.registrationService.sendRegistrationRequest(this.registrationForm.value).subscribe({
        next: (res: any) => {

          if (res.errors) {
                this.registrationMessage.set(res.errors[0].message);          
                setTimeout(() => {
                  this.registrationMessage.set('');
                }, 3000);
                return;
          } else {
            this.registrationMessage.set(res.data.registerUser.message);
          }

          setTimeout(() => {
            this.registrationMessage.set('');
          }, 3000);

        },
        error: (err: any) => {
          this.registrationMessage.set(err.errors.message);
          setTimeout(() => {
            this.registrationMessage.set('');
          }, 3000);

        }        
      });     
    }
  }

  closeRegistration() {
    $("#reset").trigger('click');
    $("#hideRegister").trigger('click');
    this.router.navigate(['/']); 
  }

  public registrationOpen(registrationTemplate: any): void {
		this.modalService.open(registrationTemplate, { centered: true });
	}


}
