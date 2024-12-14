import {HttpClient} from "@angular/common/http"
import {inject, Injectable} from "@angular/core"
import {APP_CONFIG} from "@linkmate/shared-util-app-config"
import {Observable} from "rxjs"
import {AuthLoginReqDto, AuthLoginResDto} from "../entities/login"
import {AuthRegisterReqDto, AuthRegisterResDto} from "../entities/register"

@Injectable({
    providedIn: "root"
})
export class AuthDataAccess {
    private readonly httpClient = inject(HttpClient)
    private readonly appConfig = inject(APP_CONFIG)

    login(dto: AuthLoginReqDto): Observable<AuthLoginResDto> {
        return this.httpClient.post<AuthLoginResDto>(`${this.appConfig.apiUrl}/auth/login`, dto)
    }

    register(dto: AuthRegisterReqDto): Observable<AuthRegisterResDto> {
        return this.httpClient.post<AuthRegisterResDto>(`${this.appConfig.apiUrl}/auth/register`, dto)
    }
}
