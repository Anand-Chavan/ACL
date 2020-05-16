import { Component, OnInit } from '@angular/core';
import { AclServiceService } from '../acl-service.service';

@Component({
  selector: 'app-create-file-folder',
  templateUrl: './create-file-folder.component.html',
  styleUrls: ['./create-file-folder.component.css']
})
export class CreateFileFolderComponent implements OnInit {

  filefolderPath = "";
  filefolderName = "";
  filesOrFolderId = "";
  userId = "";
  sessionKey = "";

  constructor(public rest:AclServiceService) { }

  ngOnInit(): void {
  }

  getValue()
  {
  	let post ='{"filefolderPath":"'+this.filefolderPath+'","filefolderName":"'+this.filefolderName+'","filesOrFolderId":"'+this.filesOrFolderId+'" ,"userId":"'+this.userId+'","sessionKey":"'+this.sessionKey+'"}';
  	 this.rest.createFileFolder(post).subscribe((data: {}) => {
        console.log(data);
        this.getUserdata(data)
      });
  }
  getUserdata(data)
  {

  }
}
