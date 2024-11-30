const nx = require("@nx/eslint-plugin");

module.exports = [
    ...nx.configs["flat/base"],
    ...nx.configs["flat/typescript"],
    ...nx.configs["flat/javascript"],
    {
        ignores: ["**/dist"],
    },
    {
        files: ["**/*.ts", "**/*.tsx", "**/*.js", "**/*.jsx"],
        rules: {
            "@nx/enforce-module-boundaries": [
                "error",
                {
                    enforceBuildableLibDependency: true,
                    allow: ["^.*/eslint(\\.base)?\\.config\\.[cm]?js$"],
                    depConstraints: [
                        {
                            sourceTag: "*",
                            onlyDependOnLibsWithTags: ["*"]
                        },
                        {
                            sourceTag: "type:util",
                            onlyDependOnLibsWithTags: [""]
                        },
                        {
                            sourceTag: "type:ui",
                            onlyDependOnLibsWithTags: ["type:util"],
                        },
                        {
                            sourceTag: "type:domain",
                            onlyDependOnLibsWithTags: ["type:util"],
                        },
                        {
                            sourceTag: "type:feature",
                            onlyDependOnLibsWithTags: ["type:domain", "type:ui", "type:util"],
                        },
                        {
                            sourceTag: "type:app",
                            onlyDependOnLibsWithTags: ["type:domain", "type:ui", "type:util", "type:feature"],
                        },
                    ],
                },
            ],
        },
    },
    {
        files: ["**/*.ts", "**/*.tsx", "**/*.js", "**/*.jsx"],
        // Override or add rules here
        rules: {},
    },
];
