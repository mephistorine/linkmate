import {inject} from "@angular/core"
import {CanActivateFn, Router} from "@angular/router"
import {AuthFacade} from "@linkmate/auth-domain"

export const mustBeUnauthorizedGuard: CanActivateFn = () => {
    const authFacade = inject(AuthFacade)
    const router = inject(Router)

    if (authFacade.isAuthorized()) {
        return router.parseUrl("/")
    }

    return true
}
