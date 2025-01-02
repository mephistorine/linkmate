import {CdkCopyToClipboard} from "@angular/cdk/clipboard"
import {
    ChangeDetectionStrategy,
    Component,
    inject, OnInit,
    ResourceStatus,
} from "@angular/core"
import {RouterLink} from "@angular/router"
import {LinksFacade} from "@linkmate/links-domain"

@Component({
    selector: "lib-links-feature-links-page",
    templateUrl: "./links-feature-links-page.component.html",
    styleUrl: "./links-feature-links-page.component.css",
    changeDetection: ChangeDetectionStrategy.OnPush,
    imports: [
        RouterLink,
        CdkCopyToClipboard,
    ],
})
export class LinksFeatureLinksPageComponent implements OnInit {
    readonly ResourceStatus = ResourceStatus
    linksFacade = inject(LinksFacade)

    ngOnInit(): void {
        this.linksFacade.loadLinks()
    }

    originalUrlIcon(urlString: string) {
        const url = new URL(urlString)
        // TODO: Вынести в переменную окружения
        return `https://www.google.com/s2/favicons?domain=${url.host}&sz=128`
    }
}
