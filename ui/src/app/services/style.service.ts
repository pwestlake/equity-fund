import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class StyleService {

  constructor() { }
  /** 
   * Set the stylesheet with the specified key. */ 
  setStyle(key: string, href: string) { 
    getLinkElementForKey(key).setAttribute('href', href); 
  }
}

function getLinkElementForKey(key: string) {
  return document.head.querySelector(`link[rel="stylesheet"].${key}`);
}
