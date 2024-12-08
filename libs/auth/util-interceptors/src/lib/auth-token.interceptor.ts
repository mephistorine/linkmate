import {
    HttpHandlerFn,
    HttpInterceptorFn,
    HttpRequest,
} from "@angular/common/http"
import {inject} from "@angular/core"
import {AuthFacade} from "@linkmate/auth-domain"

export const authTokenInterceptor: HttpInterceptorFn = (req: HttpRequest<unknown>, next: HttpHandlerFn) => {
    const authFacade = inject(AuthFacade)

    if (!authFacade.isAuthorized()) {
        return next(req)
    }

    // TODO: Вынести в переменные окружения
    if (!req.url.startsWith("http://localhost:9000")) {
        return next(req)
    }

    return next(req.clone({
        headers: req.headers.append("Authorization", `Bearer ${authFacade.authData().accessToken}`),
    }))
}
