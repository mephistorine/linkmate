import {
    ChangeDetectionStrategy,
    Component,
    effect,
    inject,
    ResourceStatus,
} from "@angular/core"
import {rxResource} from "@angular/core/rxjs-interop"
import {
    FormsModule,
    NonNullableFormBuilder,
    ReactiveFormsModule,
    Validators,
} from "@angular/forms"
import {LinksFacade} from "@linkmate/links-domain"
import {injectQueryParams} from "ngxtension/inject-query-params"

@Component({
    selector: "lib-links-feature-link-edit-page",
    templateUrl: "./links-feature-link-edit-page.component.html",
    styleUrl: "./links-feature-link-edit-page.component.css",
    changeDetection: ChangeDetectionStrategy.OnPush,
    imports: [
        FormsModule,
        ReactiveFormsModule,
    ],
})
export class LinksFeatureLinkEditPageComponent {
    private readonly linksFacade = inject(LinksFacade)
    private readonly fb = inject(NonNullableFormBuilder)
    private readonly id = injectQueryParams("id", {
        transform: (val) => Number(val),
    })

    private link = rxResource({
        request: this.id,
        loader: ({request: id}) => {
            return this.linksFacade.loadLink(id!)
        },
    })

    form = this.fb.group({
        url: ["", [Validators.required, Validators.minLength(1)]]
    })

    constructor() {
        effect(() => {
            if (this.link.status() === ResourceStatus.Resolved) {
                this.form.patchValue({
                    url: this.link.value().url
                })
            }
        })
    }
}
