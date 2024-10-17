import {useRouter} from "vue-router";
import {ElMessage} from "element-plus";

const r = useRouter()
export const setToken = (obj) => {
    const str = "root-"

    window.localStorage.setItem(str+'token', obj.token)
    window.localStorage.setItem(str+'userId', obj.userId)
    window.localStorage.setItem(str+'name', obj.userInfo.name)

    window.localStorage.setItem(str+'userType', obj.type)

    window.localStorage.setItem(str+'userInfo',  JSON.stringify(obj.userInfo))

}

export const getToken = () => {
    const str = "root-"

    return  window.localStorage.getItem(str+'token');
}


export const longOut = () =>{
    const str = "root-"

    window.localStorage.removeItem(str+"token")
}

export  const handleRes = (res) =>{
    if (res.code === 2000) {

        ElMessage.success(res.msg)
        return true
    } else {
        ElMessage.warning(res.msg)


        return  false
    }
}

export function formatTimeToStr(time) {
    if (!time) return "--";
    var date = new Date(time)

    let year = date.getFullYear();
    let month = ('0' + (date.getMonth() + 1)).slice(-2);
    let day = ('0' + date.getDate()).slice(-2);
// 获取小时和分钟
    var hours = date.getHours();
    var minutes = date.getMinutes();

// 格式化小时，如果小于10则在前边加0
    hours = hours < 10 ? '0' + hours : hours;
// 格式化分钟，如果小于10则在前边加0
    minutes = minutes < 10 ? '0' + minutes : minutes;
    return `${year}-${month}-${day}-${hours}:${minutes}`;
}

export function isTimeGreaterThanToday(timeStr) {

    var today = new Date()
    today.setHours(0,0,0,0)//将时、分、秒、毫秒设置为日，只比较日期
    var selectedDate=new Date(timeStr)
    selectedDate.setHours(0,0,0,0)

    return selectedDate.getTime() > today.getTime();

}

export function daysBetween(date1, date2) {
    const oneDay = 24 * 60 * 60 * 1000; // 每天的毫秒数


    const firstTime = new Date(date1).getTime();
    const secondTime = new Date(date2).getTime();

    return Math.round(Math.abs((firstTime - secondTime) / oneDay));
}