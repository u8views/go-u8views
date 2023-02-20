require("esbuild").buildSync({
    entryPoints: [
        // status: outdated
        "./src/index-app.ts",
        "./src/github-profile-app.ts",
        "./src/stats-app.ts",
    ],
    bundle: true,
    minify: process.env.MINIFY === "true",
    outdir: "../public/assets/js",
});
