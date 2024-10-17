import axios   from 'axios'
import {getToken} from "@/modules/util.js";
import router from "@/modules/router.js";
import {ElMessage} from "element-plus";

const env = import.meta.env;

const request = axios.create({
    baseURL:env.VITE_BASE_API ,  // 注意！！ 这里是全局统一加上了 '/api' 前缀，也就是说所有接口都会加上'/api'前缀在，页面里面写接口的时候就不要加 '/api'了，否则会出现2个'/api'，类似 '/api/api/user'这样的报错，切记！！！
    timeout: 5000
})

// 统一加解密
// const Unify = {
//     // 统一加密方法
//     en(data, key) {
//         // 1.aes加密
//         console.log("aes加密",key);
//         let aesStr = aes.en(JSON.stringify(data), key);
//         return aesStr;
//     },
//     // 统一解密
//     de(aesStr, key) {
//         // 1.aes解密
//         let dataStr = aes.de(aesStr, key);
//         // console.log(dataStr);
//         // 3.转json对象
//         let data = JSON.parse(dataStr);
//         return data;
//     },
// };

// request 拦截器
// 可以自请求发送前对请求做一些处理
// 比如统一加token，对请求参数统一加密
request.interceptors.request.use(config => {
    config.headers['Content-Type'] = 'application/json;charset=utf-8';
    let tokenObj =  getToken()
    config.headers['Authorization'] = 'Bearer ' + tokenObj;  // 设置请求头
    // console.log("this is key",key);
    // config.headers['AseKey'] =  key.code;  // 设置请求头

    //参数加密
    // if (config.method === "post" || config.method === "put") {
    //     if (config.data) {
    //         config.headers["crypto"] = true;
    //         config.headers["content-type"] = "application/json";
    //         let data = config.data;
    //
    //         // 加密 post 请求参数
    //         config.data = Unify.en(data, key.key);
    //     }
    // }
    // console.log(config.data);
    return config
}, error => {
    return Promise.reject(error)
});

// response 拦截器
// 可以在接口响应后统一处理结果
request.interceptors.response.use(
    response => {

        let res = response.data;
        // console.log(res);

        // 如果是返回的文件
        if (response.config.responseType === 'blob') {
            return res
        }
        // 兼容服务端返回的字符串数据
        if (typeof res === 'string') {
            res = res ? JSON.parse(res) : res
        }
        if (res.code === 4001) {
            router.push("/login")
            return res;
        }


        try {

            // let key = getAesKey(response.headers['asekey']); /// 拿到公钥加密字符串
            // if (!key) {
            //     // ElMessage.error('获取密钥异常！！！')
            //     // console.log("获取解密密钥异常！！！")
            //     return Promise.reject(res);
            // }
            // //有返回数据才去解密
            // if (res.data) {
            //     res.data = Unify.de(res.data, key);
            // }


        } catch (err) {
            console.log("请求系统异常！！！")
            ElMessage.error("请求系统异常！！！")
            return Promise.reject(err);
        }
        // console.log("接口解密后的数据",res);
        return res;
    },
    error => {
        ElMessage.error(error)
        console.log('' + error) // for debug
        return Promise.reject(error)
    }
)

export default  request