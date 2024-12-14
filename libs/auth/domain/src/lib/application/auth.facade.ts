import {computed, inject, Injectable, signal} from "@angular/core"
import {Router} from "@angular/router"
import {WA_LOCAL_STORAGE, WA_SESSION_STORAGE} from "@ng-web-apis/common"
import {first, tap} from "rxjs"

import {AuthDataAccess} from "../infra/auth.data-access"

@Injectable({providedIn: "root"})
export class AuthFacade {
    private readonly router = inject(Router)
    private readonly authDataAccess = inject(AuthDataAccess)
    private readonly sessionStorage = inject(WA_SESSION_STORAGE)
    private readonly localStorage = inject(WA_LOCAL_STORAGE)
    readonly authData = signal({
        accessToken: "",
    })

    readonly isAuthorized = computed(() => this.authData().accessToken.length > 0)

    constructor() {
        const authData = this.sessionStorage.getItem("linkmate:authData") ?? this.localStorage.getItem("linkmate:authData")
        if (authData) {
            try {
                const parsedData = JSON.parse(authData) as {
                    readonly accessToken: string
                }
                this.authData.set({
                    accessToken: parsedData.accessToken,
                })
            } catch (err) {
                this.authData.set({accessToken: ""})
            }
        }
    }

    login(email: string, password: string, isRememberMe: boolean) {
        this.authDataAccess.login({email, password})
            .pipe(
                first(),
                tap((result) => {
                    const storage = isRememberMe ? this.localStorage : this.sessionStorage
                    this.authData.set({accessToken: result.accessToken})
                    storage.setItem("linkmate:authData", JSON.stringify({accessToken: result.accessToken}))
                }),
            )
            .subscribe()
    }

    logout() {
        this.reset()
        this.router.navigateByUrl("/login")
    }

    reset() {
        this.authData.set({accessToken: ""})
        this.localStorage.removeItem("linkmate:authData")
        this.sessionStorage.removeItem("linkmate:authData")
    }
}
