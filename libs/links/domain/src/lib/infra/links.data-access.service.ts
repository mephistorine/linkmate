import {HttpClient} from "@angular/common/http"
import {inject, Injectable} from "@angular/core"
import {Observable} from "rxjs"

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

    getLinks(): Observable<LinkListItemResDto[]> {
        return this.httpClient.get<LinkListItemResDto[]>("http://localhost:9000/api/links")
    }

    createLink(link: CreateLinkReqDto): Observable<CreateLinkResDto> {
        return this.httpClient.post<CreateLinkResDto>("http://localhost:9000/api/links", link)
    }
}
