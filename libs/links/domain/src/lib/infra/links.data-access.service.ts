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

@Injectable({
    providedIn: "root",
})
export class LinksDataAccessService {
    private readonly httpClient = inject(HttpClient)

    getLinks(): Observable<readonly LinkListItemResDto[]> {
        return this.httpClient.get<readonly LinkListItemResDto[]>("http://localhost:9000/api/links")
    }
}
