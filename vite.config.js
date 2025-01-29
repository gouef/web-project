import { defineConfig } from 'vite'

export default defineConfig({
    build: {
        //assetsDir: 'assets',
        rollupOptions: {
           input:  {
               main: 'src/index.js',
               style: 'src/styles.scss'
           },
            output: {
                entryFileNames: 'js/main.js',
                chunkFileNames: 'js/chunks.js',
                assetFileNames: ({ name: names }) => {
                    if (names.endsWith('.css')) {
                        return 'css/[name].css';
                    } else if (names.startsWith('images/')) {
                        return 'assets/images/[name]'; // Obrázky z `src/images/` do `assets/images/`
                    } else if (/\.(jpe?g|png|gif|svg|webp|ttf|woff|woff2|eot|otf|mp4|webm)$/.test(names)) {
                        return 'assets/[name]';
                    }
                    return 'assets/[name]';
                }
            }
        },
        outDir: 'static',
        emptyOutDir: false
    },
})