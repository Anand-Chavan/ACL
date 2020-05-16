import { Component, OnInit } from '@angular/core';
import { AclServiceService } from '../acl-service.service';

@Component({
  selector: 'app-delete',
  templateUrl: './delete.component.html',
  styleUrls: ['./delete.component.css']
})
export class DeleteComponent implements OnInit {

  userId = 0;

  constructor(public rest:AclServiceService) { }

  ngOnInit(): void {

  		
  }


  getValue()
  {
  		    this.rest.deleteUser(this.userId).subscribe((data: {}) => {
        console.log(data);
      });
  }


}
