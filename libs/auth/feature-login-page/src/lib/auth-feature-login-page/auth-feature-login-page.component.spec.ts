import { ComponentFixture, TestBed } from "@angular/core/testing";
import { AuthFeatureLoginPageComponent } from "./auth-feature-login-page.component";

describe("AuthFeatureLoginPageComponent", () => {
    let component: AuthFeatureLoginPageComponent;
    let fixture: ComponentFixture<AuthFeatureLoginPageComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [AuthFeatureLoginPageComponent],
        }).compileComponents();

        fixture = TestBed.createComponent(AuthFeatureLoginPageComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it("should create", () => {
        expect(component).toBeTruthy();
    });
});
