import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		proxy: {
			'/api': {
				target: 'https://course-scope-cmu.alwaysdata.net/',
				changeOrigin: true,
				rewrite: (path) => path.replace(/^\/api/, ''),
				secure: true,
			},
		}
	}
});
