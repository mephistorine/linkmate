import { ChangeDetectionStrategy, Component } from "@angular/core";
import { CommonModule } from "@angular/common";
import {RouterLink, RouterLinkActive, RouterOutlet} from "@angular/router"

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
    ],
})
export class CoreFeatureShellComponent {}
