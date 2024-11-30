import {Route} from "@angular/router"
import {CoreFeatureShellComponent} from "@linkmate/core-feature-shell"
import {HomeFeatureMainPageComponent} from "@linkmate/home-feature-home-page"

export const appRoutes: Route[] = [
    {
        path: "login",
        loadComponent: () => import("@linkmate/auth-feature-login-page").then(m => m.AuthFeatureLoginPageComponent),
        title: "Login – Linkmate",
    },
    {
        path: "",
        component: CoreFeatureShellComponent,
        children: [
            {
                path: "",
                component: HomeFeatureMainPageComponent,
                title: "Home — Linkmate"
            }
        ],
    },
]
