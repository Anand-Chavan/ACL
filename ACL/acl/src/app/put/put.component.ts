import { Component, OnInit } from '@angular/core';
import { AclServiceService } from '../acl-service.service';

@Component({
  selector: 'app-put',
  templateUrl: './put.component.html',
  styleUrls: ['./put.component.css']
})
export class PutComponent implements OnInit {

  userId = "";
  groupName = "";
  groupDescription = "";
  sessionKey = "";

  constructor(public rest:AclServiceService) { }

  ngOnInit(): void {
  }

   getValue()
  {
  		let post='{"userId":"'+this.userId+'","groupName":"'+this.groupName+'","groupDescription":"'+this.groupDescription+'","sessionKey":"'+this.sessionKey+'"}'
  		console.log(post);
  		this.rest.createGroup(post).subscribe((data: {}) => {
        console.log(data);
      });

  }

}
