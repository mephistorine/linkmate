import {Component, effect, inject} from "@angular/core"
import {
    NonNullableFormBuilder,
    ReactiveFormsModule,
    Validators,
} from "@angular/forms"
import {Router} from "@angular/router"
import {AuthFacade} from "@linkmate/auth-domain"
import {NgxControlError} from "ngxtension/control-error"

@Component({
    selector: "lib-auth-feature-login-page",
    standalone: true,
    imports: [NgxControlError, ReactiveFormsModule],
    templateUrl: "./auth-feature-login-page.component.html",
    styleUrl: "./auth-feature-login-page.component.css",
})
export class AuthFeatureLoginPageComponent {
    private fb = inject(NonNullableFormBuilder)
    private readonly authFacade = inject(AuthFacade)
    private readonly router = inject(Router)

    private readonly homeRedirect = effect(() => {
        if (this.authFacade.authData().accessToken.length > 0) {
            this.router.navigateByUrl("/")
        }
    })

    form = this.fb.group({
        email: ["", [Validators.required, Validators.email, Validators.minLength(1)]],
        password: ["", [Validators.required, Validators.minLength(1)]],
        rememberMe: [true, [Validators.required]],
    })

    sendLogin(event: SubmitEvent) {
        event.preventDefault()

        if (this.form.invalid) {
            return
        }

        const value = this.form.getRawValue()

        this.authFacade.login(
            value.email,
            value.password,
            value.rememberMe,
        )
    }
}
