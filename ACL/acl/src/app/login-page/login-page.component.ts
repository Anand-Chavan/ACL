import { Component, OnInit } from '@angular/core';
import { AclServiceService } from '../acl-service.service';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent implements OnInit {

  loginSuccess=false;
  firstLogin = true;
  userId = "";
  password = "";
  user="anand";
  userArray = ['Anand','Chavan']
  dataSource = [];
  

  constructor(public rest:AclServiceService) { }

  ngOnInit(): void {
    this.rest.getAllUserInfo().subscribe((data: {}) => {
        console.log(data);
        this.getUserdata(data)
      });
  }
    
 
  getUserdata(data)
  {
      let array=Object.values(data.data);
      console.log(array);
      let MArray = [];
      for(var i=0;i<array.length;i++)
      {
         let field = {'userName':array[i][0],'userId':array[i][1],'password':array[i][2],'userType':array[i][3]};
         MArray.push(field);
      }
      this.dataSource = MArray;
      console.log(this.dataSource);
  }

  login()
  {
      
      let post = '{"userId":"'+this.userId+'","password":"'+this.password+'"}'
      this.rest.login(post).subscribe((data: {}) => {
        console.log(data);
        this.getLogindata(data)
      });

  }

  getUser(user)
  {
    this.user=user;
  }

  getLogindata(data)
  {
      let array=Object.values(data.data);
      let flag = 0;
      console.log(this.dataSource.length);
      if(this.dataSource.length>0)
      {
        for(var i=0;i<this.dataSource.length;i++)
        {
            // console.log(this.dataSource[i]['userId'],this.dataSource[i]['password']);
            if(this.dataSource[i]['userId'] == array[0] && this.dataSource[i]['password'] == this.password)
            {
                console.log(this.dataSource[i]['userId'],this.dataSource[i]['password']);
                flag = 1;
                break;
            }
        }
      }
    

      if(flag == 1)
      {
          console.log("Valid");
          this.loginSuccess = true;
          this.firstLogin = false;
      }
      else
      {
        console.log("Invalid");
        window.alert("Enter Valid Credential");
      }
  }

}
