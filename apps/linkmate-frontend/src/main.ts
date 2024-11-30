import {ChangeDetectionStrategy, Component} from "@angular/core"
import {bootstrapApplication} from "@angular/platform-browser"
import {RouterOutlet} from "@angular/router"
import {appConfig} from "./app.config"

@Component({
    standalone: true,
    selector: "app-root",
    template: `
        <router-outlet></router-outlet>`,
    styles: [
        `:host {
            display: block;
            height: 100%;
        }`,
    ],
    imports: [
        RouterOutlet,
    ],
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AppComponent {
    constructor() {
        const theme = window.matchMedia("(prefers-color-scheme: dark)")
        const setTheme = (theme: "dark" | "light") => {
            document.body.dataset["bsTheme"] = theme
        }
        setTheme(theme.matches ? "dark" : "light")
        theme.addEventListener("change", (event) => setTheme(event.matches ? "dark" : "light"))
    }
}

bootstrapApplication(AppComponent, appConfig).catch((err) =>
    console.error(err),
)
