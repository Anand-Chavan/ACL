import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ChangePermissionComponent } from './change-permission.component';

describe('ChangePermissionComponent', () => {
  let component: ChangePermissionComponent;
  let fixture: ComponentFixture<ChangePermissionComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ChangePermissionComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ChangePermissionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
