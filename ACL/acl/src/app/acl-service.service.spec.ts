import { TestBed } from '@angular/core/testing';

import { AclServiceService } from './acl-service.service';

describe('AclServiceService', () => {
  let service: AclServiceService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AclServiceService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
