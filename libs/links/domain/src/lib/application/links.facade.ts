import {Dialog} from "@angular/cdk/dialog"
import {
    inject,
    Injectable,
    Injector,
    ResourceRef,
    runInInjectionContext,
} from "@angular/core"
import {rxResource} from "@angular/core/rxjs-interop"
import {Router} from "@angular/router"
import {APP_CONFIG} from "@linkmate/shared-util-app-config"
import {LOGGER} from "@linkmate/shared-util-logger"
import {
    catchError,
    EMPTY,
    filter,
    map,
    Observable,
    switchMap,
    take,
    tap,
} from "rxjs"
// TODO: Вынести в отдельный сервис ConfirmService
import {
    UiConfirmDialogComponent,
} from "@linkmate/ui-confirm-dialog"
import {LinksDataAccessService} from "../infra/links.data-access.service"

@Injectable({
    providedIn: "root",
})
export class LinksFacade {
    private readonly linksDataAccessService = inject(LinksDataAccessService)
    private readonly logger = inject(LOGGER)
    private readonly appConfig = inject(APP_CONFIG)
    private readonly router = inject(Router)
    private readonly dialogService = inject(Dialog)
    private readonly injector = inject(Injector)

    links: ResourceRef<any[]> | null = null

    loadLinks() {
        runInInjectionContext(this.injector, () => {
            this.links = rxResource({
                loader: () => this.linksDataAccessService.getLinks().pipe(
                    map((links) =>
                        links.map((link) => ({
                            ...link,
                            shortUrl: this.appConfig.createShortUrl(link.key),
                        }))),
                ),
            })
        })
    }

    loadLink(id: number): Observable<any> {
        return this.linksDataAccessService.getLink(id).pipe(
            take(1),
        )
    }

    deleteLink(id: number): void {
        this.dialogService
            .open(UiConfirmDialogComponent, {
                width: "300px",
                data: {
                    title: "Delete link?",
                    okText: "Delete",
                },
            })
            .closed
            .pipe(
                take(1),
                filter(Boolean),
                switchMap(() => {
                    return this.linksDataAccessService
                        .deleteById(id)
                        .pipe(
                            take(1),
                            tap(() => this.links?.reload()),
                            catchError((error) => {
                                this.logger.warn(error)
                                return EMPTY
                            })
                        )
                })
            ).subscribe()
    }

    createLink(link: any): void {
        this.linksDataAccessService
            .create(link)
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
