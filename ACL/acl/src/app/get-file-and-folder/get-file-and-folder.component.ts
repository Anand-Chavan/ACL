import { Component, OnInit } from '@angular/core';
import { AclServiceService } from '../acl-service.service';

@Component({
  selector: 'app-get-file-and-folder',
  templateUrl: './get-file-and-folder.component.html',
  styleUrls: ['./get-file-and-folder.component.css']
})
export class GetFileAndFolderComponent implements OnInit {

  displayedColumns=['filePath','fileName','FileOrDirectory','Id','permission','operation'];
  dataSource=[];


  constructor(public rest:AclServiceService) { }

  ngOnInit(): void {

  	let post = '{"filefolderPath":"/" ,"userId":"u2"}';
  	this.rest.getFileFolder(post).subscribe((data: {}) => {
        console.log(data);
        this.getUserdata(data)
      });
  }

   getUserdata(data)
  {
      let array=Object.values(data.data);
      console.log(array);
      let MArray=[];

    
      for(var i=0;i<array.length;i++)
      {
         let doSome = ""
         if(array[i][4] == 'r')
         {
            doSome = 'visibility';
         }
         else
         {
            doSome = 'edit';
         }

         let field = {'filePath':array[i][0],'fileName':array[i][1],'FileOrDirectory':array[i][2],'Id':array[i][3],'permission':array[i][4],'operation':doSome};
         MArray.push(field)
      }
      this.dataSource = MArray;
      console.log(this.dataSource[0]);
      console.log(typeof(this.dataSource));
  }

  performOperation(index)
  {
    console.log(typeof(this.dataSource[index]['fileName']));
    window.open('/home/anand/check.c');
  }

}
