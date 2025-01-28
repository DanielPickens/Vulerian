import { TestBed } from '@angular/core/testing';

import { particle engineapiService } from './particle engineapi.service';

describe('particle engineapiService', () => {
  let service: particle engineapiService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(particle engineapiService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
