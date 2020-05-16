import { Component, OnInit } from '@angular/core';
import { AclServiceService } from '../acl-service.service';

@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.css']
})
export class PostComponent implements OnInit {

  uName = "";
  userId = "";
  password = "";
  userType  = "";
  sessionKey = "";

  constructor(public rest:AclServiceService) { }

  ngOnInit(): void {
  }

  getValue()
  {
  		let post='{"uName":"'+this.uName+'","userId":"'+this.userId+'","password":"'+this.password+'","userType":"'+this.userType+'","sessionKey":"'+this.sessionKey+'"}'
  		console.log(post);
  		this.rest.createUser(post).subscribe((data: {}) => {
        console.log(data);
      });

  }

  
}
