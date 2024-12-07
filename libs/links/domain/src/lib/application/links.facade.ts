import {computed, inject, Injectable, signal} from "@angular/core"
import {Router} from "@angular/router"
import {injectLogger} from "@linkmate/shared-util-logger"
import {produce} from "immer"
import {catchError, EMPTY, take, tap} from "rxjs"
import {LinksDataAccessService} from "../infra/links.data-access.service"

const enum LoadingState {
    Loading = "LOADING",
    Error = "ERROR",
    Success = "SUCCESS",
}

type LinksFacadeState = {
    links: {
        loadingState: LoadingState
        items: any[]
    }
}

@Injectable({providedIn: "root"})
export class LinksFacade {
    private readonly linksDataAccessService = inject(LinksDataAccessService)
    private readonly logger = injectLogger()
    private readonly router = inject(Router)

    private state = signal<LinksFacadeState>({
        links: {
            loadingState: LoadingState.Loading,
            items: [],
        },
    })

    linksLoadingState = computed(() => this.state().links.loadingState)

    links = computed(() => this.state().links.items)

    loadLinks(): void {
        this.linksDataAccessService
            .getLinks()
            .pipe(
                tap((links) => {
                    this.state.update(baseState => produce(baseState, s => {
                        s.links.loadingState = LoadingState.Success
                        s.links.items = links
                    }))
                }),
                catchError((error) => {
                    this.state.update(baseState => produce(baseState, s => {
                        s.links.loadingState = LoadingState.Error
                    }))
                    this.logger.warn(error)
                    return EMPTY
                }),
                take(1),
            )
            .subscribe()
    }

    createLink(link: any): void {
        this.linksDataAccessService
            .createLink(link)
            .pipe(
                tap(() => {
                    this.router.navigateByUrl("/links")
                }),
                catchError((error) => {
                    this.logger.warn(error)
                    return EMPTY
                }),
                take(1),
            )
            .subscribe()
    }
}
