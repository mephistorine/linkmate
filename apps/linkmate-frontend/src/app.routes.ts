import {ActivatedRouteSnapshot, Route} from "@angular/router"
import {
    mustBeAuthorizedGuard,
    mustBeUnauthorizedGuard,
} from "@linkmate/auth-util-guards"
import {CoreFeatureShellComponent} from "@linkmate/core-feature-shell"
import {
    LinksFeatureLinksPageComponent,
} from "@linkmate/links-feature-links-page"

const createTitle = (title: string) => `${title} – Linkmate`

export const appRoutes: Route[] = [
    {
        path: "login",
        loadComponent: () => import("@linkmate/auth-feature-login-page").then(m => m.AuthFeatureLoginPageComponent),
        title: createTitle("Login"),
        canActivate: [mustBeUnauthorizedGuard],
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
            /*{
                path: "links/edit",
                component: LinksFeatureLinkUpsertPageComponent,
                title: createTitle("Edit link"),
            },
            {
                path: "links/create",
                component: LinksFeatureLinkUpsertPageComponent,
                title: createTitle("Create link"),
            },*/
            {
                path: "links/edit",
                loadComponent: () => import("@linkmate/links-feature-link-edit-page")
                    .then(m => m.LinksFeatureLinkEditPageComponent),
                title: createTitle("Edit"),
                canActivate: [
                    (route: ActivatedRouteSnapshot) => {
                        const id = route.queryParamMap.get("id")
                        return id !== null && id.length > 0
                    },
                ],
            },
            {
                path: "links/create",
                loadComponent: () => import("@linkmate/links-feature-link-create-page")
                    .then(m => m.LinksFeatureLinkCreatePageComponent),
                title: createTitle("Create link"),
            },
            {
                path: "",
                redirectTo: "/links",
                pathMatch: "full",
            },
        ],
    },
]
