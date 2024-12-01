import { ChangeDetectionStrategy, Component } from "@angular/core";
import { CommonModule } from "@angular/common";
import {RouterLink} from "@angular/router"

@Component({
    selector: "lib-links-feature-links-page",
    standalone: true,
    templateUrl: "./links-feature-links-page.component.html",
    styleUrl: "./links-feature-links-page.component.css",
    changeDetection: ChangeDetectionStrategy.OnPush,
    imports: [
        RouterLink,
    ],
})
export class LinksFeatureLinksPageComponent {}
