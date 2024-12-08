import {
    HttpErrorResponse,
    HttpHandlerFn,
    HttpInterceptorFn,
    HttpRequest,
} from "@angular/common/http"
import {inject} from "@angular/core"
import {Router} from "@angular/router"
import {AuthFacade} from "@linkmate/auth-domain"
import {tap} from "rxjs"

export const authErrorInterceptor: HttpInterceptorFn = (req: HttpRequest<unknown>, next: HttpHandlerFn) => {
    const router = inject(Router)
    const authFacade = inject(AuthFacade)

    return next(req).pipe(
        tap({
            error: (error) => {
                if (error instanceof HttpErrorResponse && error.status === 401) {
                    authFacade.reset()
                    router.navigateByUrl("/login")
                }
            },
        }),
    )
}
