import {ChangeDetectionStrategy, Component, inject} from "@angular/core"
import {
    NonNullableFormBuilder,
    ReactiveFormsModule,
    Validators,
} from "@angular/forms"
import {LinksFacade} from "@linkmate/links-domain"

@Component({
    selector: "lib-links-feature-link-create-page",
    imports: [ReactiveFormsModule],
    templateUrl: "./links-feature-link-create-page.component.html",
    styleUrl: "./links-feature-link-create-page.component.css",
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class LinksFeatureLinkCreatePageComponent {
    private readonly fb = inject(NonNullableFormBuilder)
    private readonly linksFacade = inject(LinksFacade)

    form = this.fb.group({
        key: [""],
        url: ["", [Validators.required, Validators.minLength(1)]],
    })

    createShortUrl(): void {
        if (this.form.invalid) {
            return
        }

        const value = this.form.getRawValue()

        this.linksFacade.createLink(value)
    }
}
