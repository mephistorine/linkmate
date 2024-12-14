import {CdkMenu, CdkMenuItem, CdkMenuTrigger} from "@angular/cdk/menu"
import {ChangeDetectionStrategy, Component, inject} from "@angular/core"
import {RouterLink, RouterLinkActive, RouterOutlet} from "@angular/router"
import {AuthFacade} from "@linkmate/auth-domain"

@Component({
    selector: "lib-core-feature-shell",
    standalone: true,
    templateUrl: "./core-feature-shell.component.html",
    styleUrl: "./core-feature-shell.component.css",
    changeDetection: ChangeDetectionStrategy.OnPush,
    imports: [
        RouterOutlet,
        RouterLink,
        RouterLinkActive,
        CdkMenuTrigger,
        CdkMenu,
        CdkMenuItem,

    ],
})
export class CoreFeatureShellComponent {
    private readonly authFacade = inject(AuthFacade)

    logout() {
        this.authFacade.logout()
    }
}
