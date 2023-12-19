import { resolve } from "path"
import { defineConfig } from "vite"
import preact from "@preact/preset-vite"

export default defineConfig({
	plugins: [preact()],
	publicDir: resolve(__dirname, "public"),
	build: {
		outDir: "apple/Shared (Extension)/Resources/build",
		rollupOptions: {
			output: {
				entryFileNames: `src/[name].js`,
			},
			input: {
				popup: resolve(__dirname, "../../src/popup.html"),
				background: resolve(__dirname, "../../src/background.ts"),
			},
		},
	},
})
