import {axios} from '@/utils/request'

export function getMyMenuList(parameter) {
  return axios({
    url: "/admin/menu/getMy",
    method: 'get',
    params: parameter
  })
}

export function getMenuTree(parameter) {
  return axios({
    url: "/admin/menu/getTree",
    method: 'get',
    params: parameter
  })
}

export function checkMenuKey(parameter) {
  return axios({
    url: "/admin/menu/checkKey",
    method: 'post',
    data: parameter
  })
}

export function getMenuDetail(parameter) {
  return axios({
    url: "/admin/menu/detail",
    method: 'get',
    params: parameter,
  })
}

export function SaveMenu(parameter) {
  return axios({
    url: "/admin/menu/save",
    method: 'post',
    data: parameter,
  })
}

export function SetMenuSort(parameter) {
  return axios({
    url: "/admin/menu/setSort",
    method: 'post',
    data: parameter,
  })
}

export function DeleteMenu(parameter) {
  return axios({
    url: "/admin/menu/delete",
    method: 'post',
    data: parameter,
  })
}