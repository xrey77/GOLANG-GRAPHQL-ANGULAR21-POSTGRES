export interface User {
  id: string;
  firstname: string;
  lastname: string;
  email: string;
  mobile?: string;
  username: string;
  isactivated: boolean;
  isblocked: boolean;
  mailtoken?: string;
  profilepic?: string;
  qrcodeurl?: string;
}

export interface LoginResponse {
  signIn: {
    token: string;
    message: string;
    user: User;
  };
}
