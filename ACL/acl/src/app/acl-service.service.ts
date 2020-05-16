import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { map, catchError, tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AclServiceService {

	endpoint = 'http://localhost:9090/webapi/v1';
	httpOptions = {
  	headers: new HttpHeaders({
    	'Content-Type':  'application/json'
  	})
};

  constructor(private http: HttpClient) { }

    getAllUserInfo() 
    {
        return this.http.get(this.endpoint+'/acl');
    }

    createUser(poststring)
    {
      return this.http.post(this.endpoint+'/acl',poststring);
    }

    getFileFolder(poststring)
    {
       console.log(poststring);
       return this.http.post(this.endpoint+'/getfilefolder',poststring); 
    }

    createGroup(poststring)
    {
       console.log(poststring);
       return this.http.post(this.endpoint+'/creategroup',poststring);
    }

    addUserIntoGroup(poststring)
    {
       console.log(poststring);
       return this.http.post(this.endpoint+'/adduserintogroup',poststring);
    }

    deleteUser(id)
    {
      // console.log(poststring);
      return this.http.delete(this.endpoint+'/acl/'+id);
    }

    update(poststring,id)
    {
      console.log(poststring);
      return this.http.post(this.endpoint+'/acl/'+id,poststring);
    }

    logout()
    {
      let post='{"userId":"abc","sessionKey":"f71c9f90969411eaa5b0525400d045b7"}';
      return this.http.post(this.endpoint+'/logout',post);
    }
    createFileFolder(poststring)
    {
        return this.http.post(this.endpoint+'/createfilefolder',poststring);
    }

    login(poststring)
    {
        return this.http.post(this.endpoint+'/auth',poststring);
    }
}