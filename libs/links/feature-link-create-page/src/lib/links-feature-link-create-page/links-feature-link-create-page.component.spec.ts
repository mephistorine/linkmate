import { ComponentFixture, TestBed } from "@angular/core/testing";
import { LinksFeatureLinkCreatePageComponent } from "./links-feature-link-create-page.component";

describe("LinksFeatureLinkCreatePageComponent", () => {
    let component: LinksFeatureLinkCreatePageComponent;
    let fixture: ComponentFixture<LinksFeatureLinkCreatePageComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [LinksFeatureLinkCreatePageComponent],
        }).compileComponents();

        fixture = TestBed.createComponent(LinksFeatureLinkCreatePageComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it("should create", () => {
        expect(component).toBeTruthy();
    });
});
