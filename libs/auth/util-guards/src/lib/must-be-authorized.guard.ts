import {inject} from "@angular/core"
import {CanActivateFn, Router} from "@angular/router"
import {AuthFacade} from "@linkmate/auth-domain"

export const mustBeAuthorizedGuard: CanActivateFn = () => {
    const authFacade = inject(AuthFacade)
    const router = inject(Router)

    if (authFacade.isAuthorized()) {
        return true
    }

    return router.parseUrl("/login")
}
