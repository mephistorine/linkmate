export default {
    displayName: "home-feature-home-page",
    preset: "../../../jest.preset.js",
    setupFilesAfterEnv: ["<rootDir>/src/test-setup.ts"],
    coverageDirectory: "../../../coverage/libs/home/feature-home-page",
    transform: {
        "^.+\\.(ts|mjs|js|html)$": [
            "jest-preset-angular",
            {
                tsconfig: "<rootDir>/tsconfig.spec.json",
                stringifyContentPathRegex: "\\.(html|svg)$",
            },
        ],
    },
    transformIgnorePatterns: ["node_modules/(?!.*\\.mjs$)"],
    snapshotSerializers: [
        "jest-preset-angular/build/serializers/no-ng-attributes",
        "jest-preset-angular/build/serializers/ng-snapshot",
        "jest-preset-angular/build/serializers/html-comment",
    ],
};