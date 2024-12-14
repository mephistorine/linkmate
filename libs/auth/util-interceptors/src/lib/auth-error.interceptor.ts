import {
    HttpErrorResponse,
    HttpHandlerFn,
    HttpInterceptorFn,
    HttpRequest,
} from "@angular/common/http"
import {inject} from "@angular/core"
import {AuthFacade} from "@linkmate/auth-domain"
import {tap} from "rxjs"

export const authErrorInterceptor: HttpInterceptorFn = (req: HttpRequest<unknown>, next: HttpHandlerFn) => {
    const authFacade = inject(AuthFacade)

    return next(req).pipe(
        tap({
            error: (error) => {
                if (error instanceof HttpErrorResponse && error.status === 401) {
                    authFacade.logout()
                }
            },
        }),
    )
}
