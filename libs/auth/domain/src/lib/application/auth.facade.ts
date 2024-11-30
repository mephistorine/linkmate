import {inject, Injectable, signal} from "@angular/core"
import {WA_WINDOW} from "@ng-web-apis/common"
import {first, tap} from "rxjs"

import {AuthDataAccess} from "../infra/auth.data-access"

@Injectable({providedIn: "root"})
export class AuthFacade {
    private readonly authDataAccess = inject(AuthDataAccess)
    private readonly window = inject(WA_WINDOW)
    readonly authData = signal({
        accessToken: "",
    })

    login(email: string, password: string, isRememberMe: boolean) {
        this.authDataAccess.login({email, password})
            .pipe(
                first(),
                tap((result) => {
                    const storage = isRememberMe ? this.window.localStorage : this.window.sessionStorage
                    this.authData.set({accessToken: result.accessToken})
                    storage.setItem("linkmate:authData", JSON.stringify({accessToken: result.accessToken}))
                }),
            )
            .subscribe()
    }
}
