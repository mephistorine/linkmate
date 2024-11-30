import { ComponentFixture, TestBed } from "@angular/core/testing";
import { HomeFeatureMainPageComponent } from "./home-feature-main-page.component";

describe("HomeFeatureHomePageComponent", () => {
    let component: HomeFeatureMainPageComponent;
    let fixture: ComponentFixture<HomeFeatureMainPageComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [HomeFeatureMainPageComponent],
        }).compileComponents();

        fixture = TestBed.createComponent(HomeFeatureMainPageComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it("should create", () => {
        expect(component).toBeTruthy();
    });
});
