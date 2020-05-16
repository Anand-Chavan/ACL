import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GetFileAndFolderComponent } from './get-file-and-folder.component';

describe('GetFileAndFolderComponent', () => {
  let component: GetFileAndFolderComponent;
  let fixture: ComponentFixture<GetFileAndFolderComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GetFileAndFolderComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GetFileAndFolderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
