import {Route} from "@angular/router"
import {CoreFeatureShellComponent} from "@linkmate/core-feature-shell"
import {
    LinksFeatureLinksPageComponent,
} from "@linkmate/links-feature-links-page"
import {
    mustBeAuthorizedGuard, mustBeUnauthorizedGuard,
} from "@linkmate/auth-util-guards"

const createTitle = (title: string) => `${title} – Linkmate`

export const appRoutes: Route[] = [
    {
        path: "login",
        loadComponent: () => import("@linkmate/auth-feature-login-page").then(m => m.AuthFeatureLoginPageComponent),
        title: createTitle("Title"),
        canActivate: [mustBeUnauthorizedGuard]
    },
    {
        path: "",
        component: CoreFeatureShellComponent,
        canActivate: [mustBeAuthorizedGuard],
        canActivateChild: [mustBeAuthorizedGuard],
        children: [
            {
                path: "links",
                component: LinksFeatureLinksPageComponent,
                title: createTitle("Links"),
            },
            {
                path: "",
                redirectTo: "/links",
                pathMatch: "full"
            },
        ],
    },
]
