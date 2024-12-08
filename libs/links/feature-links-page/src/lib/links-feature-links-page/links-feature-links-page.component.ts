import {CdkCopyToClipboard} from "@angular/cdk/clipboard"
import {ChangeDetectionStrategy, Component, inject, OnInit} from "@angular/core"
import {RouterLink} from "@angular/router"
import {LinksFacade} from "@linkmate/links-domain"

@Component({
    selector: "lib-links-feature-links-page",
    standalone: true,
    templateUrl: "./links-feature-links-page.component.html",
    styleUrl: "./links-feature-links-page.component.css",
    changeDetection: ChangeDetectionStrategy.OnPush,
    imports: [
        RouterLink,
        CdkCopyToClipboard,
    ],
})
export class LinksFeatureLinksPageComponent implements OnInit {
    linksFacade = inject(LinksFacade)

    ngOnInit(): void {
        this.linksFacade.loadLinks()
    }

    originalUrlIcon(urlString: string) {
        const url = new URL(urlString)
        // TODO: Вынести в переменные окружения
        return `https://www.google.com/s2/favicons?domain=${url.host}&sz=128`
    }
}
