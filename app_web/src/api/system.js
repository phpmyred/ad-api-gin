import request from '@/modules/request'
// @Router /authority/getAuthorityList [post]



export const updatePassword = (data) => {
    return request({
        url: '/update_password',
        method: 'post',
        data
    })
}

export const register =  (data) => {
    return  request({
        url: '/register',
        method: 'post',
        data
    })
}

export const get_email_code =  (data) => {
    return  request({
        url: '/get_email_code',
        method: 'post',
        data
    })
}

export const getAppName =  (data) => {
    return  request({
        url: '/getAppName',
        method: 'post',
        data
    })
}