import {
    HttpHandlerFn,
    HttpInterceptorFn,
    HttpRequest,
} from "@angular/common/http"
import {inject} from "@angular/core"
import {AuthFacade} from "@linkmate/auth-domain"
import {APP_CONFIG} from "@linkmate/shared-util-app-config"

export const authTokenInterceptor: HttpInterceptorFn = (req: HttpRequest<unknown>, next: HttpHandlerFn) => {
    const authFacade = inject(AuthFacade)
    const appConfig = inject(APP_CONFIG)

    if (!authFacade.isAuthorized()) {
        return next(req)
    }

    // TODO: Вынести в переменные окружения
    if (!req.url.startsWith(appConfig.apiUrl)) {
        return next(req)
    }

    return next(req.clone({
        headers: req.headers.append("Authorization", `Bearer ${authFacade.authData().accessToken}`),
    }))
}
