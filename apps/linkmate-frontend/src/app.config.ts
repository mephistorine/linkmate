import {
    provideHttpClient,
    withFetch,
    withInterceptors,
} from "@angular/common/http"
import {ApplicationConfig, provideZoneChangeDetection} from "@angular/core"
import {provideRouter} from "@angular/router"
import {
    authErrorInterceptor,
    authTokenInterceptor,
} from "@linkmate/auth-util-interceptors"
import {appRoutes} from "./app.routes"

export const appConfig: ApplicationConfig = {
    providers: [
        provideZoneChangeDetection({eventCoalescing: true}),
        provideRouter(appRoutes),
        provideHttpClient(
            withFetch(),
            withInterceptors([
                authTokenInterceptor,
                authErrorInterceptor,
            ]),
        ),
    ],
}
