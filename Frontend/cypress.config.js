import { defineConfig } from "cypress";

export default defineConfig({
  e2e: {
    setupNodeEvents(on, config) {
      // implement node event listeners here
    },
    reporter: 'junit',
    reporterOptions: {
      mochaFile: 'cypress/results/results-[hash].xml',
      toConsole: true,
    },
  },
  experimentalStudio: true,
});
