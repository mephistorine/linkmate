import {ChangeDetectionStrategy, Component} from "@angular/core"
import {RouterOutlet} from "@angular/router"

@Component({
    selector: "lib-home-feature-main-page",
    standalone: true,
    templateUrl: "./home-feature-main-page.component.html",
    styleUrl: "./home-feature-main-page.component.css",
    imports: [
        RouterOutlet,
    ],
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class HomeFeatureMainPageComponent {}
