import {HttpClient} from "@angular/common/http"
import {inject, Injectable} from "@angular/core"
import {Observable} from "rxjs"
import {AuthLoginReqDto, AuthLoginResDto} from "../entities/login"
import {AuthRegisterReqDto, AuthRegisterResDto} from "../entities/register"

@Injectable({
    providedIn: "root"
})
export class AuthDataAccess {
    private readonly httpClient = inject(HttpClient)

    login(dto: AuthLoginReqDto): Observable<AuthLoginResDto> {
        return this.httpClient.post<AuthLoginResDto>("http://localhost:9000/api/auth/login", dto)
    }

    register(dto: AuthRegisterReqDto): Observable<AuthRegisterResDto> {
        return this.httpClient.post<AuthRegisterResDto>("http://localhost:9000/api/auth/register", dto)
    }
}
