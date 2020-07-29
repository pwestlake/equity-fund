import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { UserService } from '../services/user.service';
import { mergeMap, map, catchError } from 'rxjs/operators';
import { StyleService } from '../services/style.service';
import { ContextService } from '../services/context.service';
import { of, throwError } from 'rxjs';

@Injectable()
export class ContextProvider {
    private theme: string;
    private title: string;
    private authenticatedUser: string;

    constructor(private http: HttpClient, 
        private userService: UserService, 
        private styleService: StyleService,
        private contextService: ContextService) { }

    public getAuthenticatedUser(): string {
        return this.authenticatedUser;
    }

    public getTheme(): string {
        return this.theme;
    }

    public getTitle(): string {
        return this.title;
    }

    load() {
        return this.userService.getAuthenticatedUserID().pipe(
            mergeMap(user => {
                this.authenticatedUser = user.userid;
                return this.userService.getPreference(user.userid, "theme")
            })
        ).pipe(
            mergeMap(theme => {
                this.theme = theme.value; 
                this.styleService.setStyle("theme", `assets/themes/${theme.value}.css`);
                return this.contextService.getTitle();
            })
        ).pipe(
            map(title => this.title = title.value)
        ).toPromise();
    }
}
