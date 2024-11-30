import { ComponentFixture, TestBed } from "@angular/core/testing";
import { CoreFeatureShellComponent } from "./core-feature-shell.component";

describe("CoreFeatureShellComponent", () => {
    let component: CoreFeatureShellComponent;
    let fixture: ComponentFixture<CoreFeatureShellComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [CoreFeatureShellComponent],
        }).compileComponents();

        fixture = TestBed.createComponent(CoreFeatureShellComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it("should create", () => {
        expect(component).toBeTruthy();
    });
});
