import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateFileFolderComponent } from './create-file-folder.component';

describe('CreateFileFolderComponent', () => {
  let component: CreateFileFolderComponent;
  let fixture: ComponentFixture<CreateFileFolderComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CreateFileFolderComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CreateFileFolderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
