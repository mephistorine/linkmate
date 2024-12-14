const myOrgEnvRegex = /^LNKM_/i;

const envVarPlugin = {
    name: 'env-var-plugin',
    setup(build) {
        const options = build.initialOptions;

        const envVars = {};
        for (const key in process.env) {
            if (myOrgEnvRegex.test(key)) {
                envVars[key] = process.env[key];
            }
        }

        options.define['envs'] = JSON.stringify(envVars);
    },
};

module.exports = envVarPlugin;
