import {DIALOG_DATA, DialogRef} from "@angular/cdk/dialog"
import {ChangeDetectionStrategy, Component, inject} from "@angular/core"

@Component({
    selector: "lib-ui-confirm-dialog",
    templateUrl: "./ui-confirm-dialog.component.html",
    styleUrl: "./ui-confirm-dialog.component.css",
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class UiConfirmDialogComponent {
    readonly dialogData = inject(DIALOG_DATA)
    readonly dialogRef = inject(DialogRef)
}
