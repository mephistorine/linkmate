import {HttpClient, HttpParams} from "@angular/common/http"
import {inject, Injectable} from "@angular/core"
import {APP_CONFIG} from "@linkmate/shared-util-app-config"
import {map, Observable} from "rxjs"

type LinkListItemResDto = {
    readonly id: number
    readonly createTime: string
    readonly key: string
    readonly url: string
    readonly userId: string
    readonly tagIds: readonly string[]
}

type CreateLinkReqDto = {
    readonly url: string
    readonly key: string
}

type CreateLinkResDto = {
    readonly id: string
    readonly key: string
    readonly url: string
    readonly createTime: string
}

@Injectable({
    providedIn: "root",
})
export class LinksDataAccessService {
    private readonly httpClient = inject(HttpClient)
    private readonly appConfig = inject(APP_CONFIG)

    getLinks(): Observable<LinkListItemResDto[]> {
        return this.httpClient.get<LinkListItemResDto[]>(`${this.appConfig.apiUrl}/links`)
    }

    getLink(id: number): Observable<LinkListItemResDto | null> {
        // TODO: Добавить отдельную ручку для получения одной ссылки
        return this.httpClient.get<LinkListItemResDto[]>(`${this.appConfig.apiUrl}/links`, {
            params: new HttpParams()
                .append("id", id),
        }).pipe(
            map((links) => links.find((link) => link.id === id) ?? null),
        )
    }

    create(link: CreateLinkReqDto): Observable<CreateLinkResDto> {
        return this.httpClient.post<CreateLinkResDto>(`${this.appConfig.apiUrl}/links`, link)
    }

    deleteById(id: number): Observable<void> {
        return this.httpClient.delete(`${this.appConfig.apiUrl}/links`, {
            params: new HttpParams({fromObject: {id}}),
        }).pipe(
            map(() => undefined),
        )
    }
}
