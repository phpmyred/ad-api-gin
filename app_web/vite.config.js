import { defineConfig,loadEnv } from 'vite'
import { fileURLToPath, URL } from 'node:url'
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite'
import AutoImport from 'unplugin-auto-import/vite'
import { dirResolver, DirResolverHelper } from "vite-auto-import-resolvers";
import Pages from 'vite-plugin-pages'
import Layouts from 'vite-plugin-vue-layouts';


export default defineConfig((conditionalConfig) => {
  console.log(conditionalConfig);
  console.log(process.cwd(), __dirname);
  console.log(process.cwd() === __dirname);
  const { mode, command, isSsrBuild, isPreview } = conditionalConfig; // conditionalConfig对象包含4个字段
  const env = loadEnv(mode, __dirname); // 根据 mode 来判断当前是何种环境
  console.log(env);
  return {
    build: {
      rollupOptions: {
        output: {
          // 配置将console输出重新启用
          inlineDynamicImports: true
        }
      }
    },
    root: process.cwd(), // 项目根目录（index.html 文件所在的位置）
    base: env.VITE_MODE === 'production' ? './' : '/', // 开发或生产环境服务的公共基础路径。
    css: {
      preprocessorOptions: {
        less: {
          modifyVars: {
            // 'arcoblue-6': '#2AA846',
          },
          javascriptEnabled: true,
        },

      }
    },
    resolve: {
      alias: { // 定义项目路径别名，这样可以在引入文件时，以属性值为起点
        '@': fileURLToPath(new URL('./src', import.meta.url))
      }

    },
    plugins: [
      vue(),

      // 组件自动引入插件
      Components({
        resolvers: [

        ]
      }),

      // 布局插件
      Layouts(),
      // 本地api 自动引入插件
      DirResolverHelper(),
      // api自动引入插件
      AutoImport({
        imports: ["vue"],
        resolvers: [dirResolver(),
        ],
      }),
      // 文件路由插件
      Pages(),


    ],
  };
})



// https://vitejs.dev/config/
// export default defineConfig({


// })
