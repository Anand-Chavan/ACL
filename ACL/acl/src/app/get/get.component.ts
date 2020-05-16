import { Component, OnInit } from '@angular/core';
import { AclServiceService } from '../acl-service.service';

@Component({
  selector: 'app-get',
  templateUrl: './get.component.html',
  styleUrls: ['./get.component.css']
})
export class GetComponent implements OnInit {

  

  displayedColumns=['userName','userId','passsword','userType'];
  dataSource=[];


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
      let MArray=[];
    
      for(var i=0;i<array.length;i++)
      {
         let field = {'userName':array[i][0],'userId':array[i][1],'passsword':array[i][2],'userType':array[i][3]};
         MArray.push(field);
      }
      this.dataSource = MArray;
      console.log(this.dataSource[0]);
      console.log(typeof(this.dataSource));
  }

}
