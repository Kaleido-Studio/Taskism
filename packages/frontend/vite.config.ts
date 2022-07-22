/* eslint-disable import/no-extraneous-dependencies */
import { defineConfig } from 'vite';
import solidPlugin from 'vite-plugin-solid';
import WindiCSS from 'vite-plugin-windicss'

export default defineConfig({
  plugins: [solidPlugin(),WindiCSS()],
  build: {
    target: 'esnext',
    outDir: '../backend/static',
    emptyOutDir: true,
  },
  
});
