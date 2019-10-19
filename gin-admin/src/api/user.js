import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/api/v1/signIn',
    method: 'post',
    data
  })
}

export function getInfo() {
  return request({
    url: '/api/v1/user/info',
    method: 'get'
  })
}

export function logout(data) {
  return request({
    url: '/api/v1/signOut',
    method: 'post',
    data
  })
}
