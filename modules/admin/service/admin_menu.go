package admin_service

import (
	"errors"
	"strings"
	"gin-modules/modules/admin/model"
	"encoding/json"
	"strconv"
)

type MenuMate struct {
	Title      string   `json:"title"`
	KeepAlive  bool     `json:"keepAlive"`
	Permission []string `json:"permission"`
	Icon       string   `json:"icon"`
	Target     string   `json:"target"`
}
type ChildMenu struct {
	Meta MenuMate `json:"meta"`
	Name string   `json:"name"`
	Path string   `json:"path"`
}
type Menus struct {
	Children []interface{} `json:"children"`
	Meta     MenuMate      `json:"meta"`
	Name     string        `json:"name"`
	Path     string        `json:"path"`
	Redirect string        `json:"redirect"`
}

//清除菜单缓存
func ClearMenuCache() bool {
	cacheKey := "card_admin_menu_list"
	cacheTreeKey := "card_admin_menu_tree"
	if redisCache.Delete(cacheKey) != nil || redisCache.Delete(cacheTreeKey) != nil {
		return false
	}
	return true
}

//获取菜单树的redis缓存数据
func GetMenuTree() ([]interface{}, error) {
	cacheKey := "card_admin_menu_tree"
	menuCache, err := redisCache.Get(cacheKey)
	if err != nil {
		return nil, errors.New("获取菜单树失败")
	}
	var menuCacheRet []interface{}
	json.Unmarshal([]byte(menuCache), &menuCacheRet)
	return menuCacheRet, nil
}

//获取所有菜单列表
func GetMenuList() ([]interface{}, error) {
	cacheKey := "card_admin_menu_list"
	cacheTreeKey := "card_admin_menu_tree"
	menuCache, err := redisCache.Get(cacheKey)
	if err == nil {
		var menuCacheRet []interface{}
		json.Unmarshal([]byte(menuCache), &menuCacheRet)
		return menuCacheRet, nil
	}

	menuFunc := func(pid int, level int) []admin_model.AdminMenu {
		menus, has := admin_model.GetMenuList(
			"id,pid,name,icon,`key`,path,is_default,level,sort",
			"sort DESC",
			1,
			100,
			"pid = ? AND level = ? AND status = ? AND is_del = ?",
			pid, level, 0, 0,
		)
		if !has {
			return []admin_model.AdminMenu{}
		}
		return menus
	}

	menus := menuFunc(0, 1)
	if len(menus) == 0 {
		return []interface{}{}, errors.New("菜单获取失败")
	}
	var menuList []interface{}
	var menuTree []map[string]interface{}
	masterMaxIndex := len(menus)
	for masterIndex, menu := range menus {
		subMenus := menuFunc(menu.ID, 2)
		meta := MenuMate{
			Title:      menu.Name,
			Icon:       menu.Icon,
			KeepAlive:  true,
			Permission: []string{menu.Key},
		}
		if len(subMenus) == 0 {
			if strings.Contains(menu.Path, "http") {
				meta.Target = "_blank"
			}
			menuList = append(menuList, ChildMenu{
				Name: menu.Key,
				Path: menu.Path,
				Meta: meta,
			})
			continue
		}
		var childMenu []interface{}
		var childrenMenuTree []map[string]interface{}
		var menuRedirect string
		subMaxIndex := len(subMenus)
		for subIndex, subMenu := range subMenus {
			if menuRedirect == "" {
				menuRedirect = subMenu.Path
			}
			menuItem := menuFunc(subMenu.ID, 3)
			subMeta := MenuMate{
				Title:      subMenu.Name,
				Icon:       subMenu.Icon,
				KeepAlive:  true,
				Permission: []string{menu.Key},
			}
			if len(menuItem) == 0 {
				if strings.Contains(subMenu.Path, "http") {
					subMeta.Target = "_blank"
				}
				childMenu = append(childMenu, ChildMenu{
					Meta: subMeta,
					Name: subMenu.Key,
					Path: subMenu.Path,
				})
				sort := subMenu.Sort
				if subIndex == 0 {
					sort = 0
				} else if subIndex == subMaxIndex-1 {
					sort = -1
				}
				print("subIndex:", subIndex, "p:", subMaxIndex, "，", "sort:", sort, "\n")
				childrenMenuTree = append(childrenMenuTree, map[string]interface{}{
					"icon":  subMenu.Icon,
					"key":   subMenu.Key,
					"title": subMenu.Name,
					"level": subMenu.Level,
					"sort":  sort,
				})
				continue
			}
			var childMenuItem []interface{}
			var childMenuItemTree []map[string]interface{}
			var subRedirect string
			itemMaxIndex := len(menuItem)
			for itemIndex, subMenuItem := range menuItem {
				if subRedirect == "" {
					subRedirect = subMenuItem.Path
				}
				itemMeta := MenuMate{
					Title:      subMenuItem.Name,
					Icon:       subMenuItem.Icon,
					KeepAlive:  true,
					Permission: []string{menu.Key},
				}
				if strings.Contains(subMenuItem.Path, "http") {
					itemMeta.Target = "_blank"
				}
				childMenuItem = append(childMenuItem, ChildMenu{
					Meta: itemMeta,
					Name: subMenuItem.Key,
					Path: subMenuItem.Path,
				})
				sort := subMenuItem.Sort
				if itemIndex == 0 {
					sort = 0
				} else if itemIndex == itemMaxIndex-1 {
					sort = -1
				}
				print("itemIndex:", itemIndex, "p:", itemMaxIndex, "，", "sort:", sort, "\n")
				childMenuItemTree = append(childMenuItemTree, map[string]interface{}{
					"icon":  subMenuItem.Icon,
					"key":   subMenuItem.Key,
					"title": subMenuItem.Name,
					"level": subMenuItem.Level,
					"sort":  sort,
				})
			}
			childMenu = append(childMenu, Menus{
				Meta:     subMeta,
				Name:     subMenu.Key,
				Path:     subMenu.Path,
				Children: childMenuItem,
				Redirect: subRedirect,
			})
			sort := subMenu.Sort
			if subIndex == 0 {
				sort = 0
			} else if subIndex == subMaxIndex-1 {
				sort = -1
			}
			print("subIndex:", subIndex, "p:", subMaxIndex, "，", "sort:", sort, "\n")
			subRes := map[string]interface{}{
				"icon":  subMenu.Icon,
				"key":   subMenu.Key,
				"title": subMenu.Name,
				"level": subMenu.Level,
				"sort":  sort,
			}
			if childMenuItemTree != nil {
				subRes["group"] = true
				subRes["children"] = childMenuItemTree
			}
			childrenMenuTree = append(childrenMenuTree, subRes)
		}
		menuList = append(menuList, Menus{
			Name:     menu.Key,
			Path:     menu.Path,
			Redirect: menuRedirect,
			Meta:     meta,
			Children: childMenu,
		})
		sort := menu.Sort
		if masterIndex == 0 {
			sort = 0
		} else if masterIndex == masterMaxIndex-1 {
			sort = -1
		}
		print("masterIndex:", masterIndex, "，p:", masterMaxIndex, "，", "sort:", sort, "\n")
		masterRes := map[string]interface{}{
			"icon":  menu.Icon,
			"key":   menu.Key,
			"title": menu.Name,
			"level": menu.Level,
			"sort":  sort,
		}
		if childrenMenuTree != nil {
			masterRes["children"] = childrenMenuTree
		}
		menuTree = append(menuTree, masterRes)
	}
	menuListJson, err := json.Marshal(menuList)
	menuTreeJson, err := json.Marshal(menuTree)
	err = redisCache.Set(cacheKey, menuListJson, 0)
	if err != nil {
		print(err.Error())
	}
	err = redisCache.Set(cacheTreeKey, menuTreeJson, 0)
	if err != nil {
		print(err.Error())
	}
	return menuList, nil
}

//通过权限ID获取到指定权限的菜单列表
func GetMyRoleMenu(roleId int) ([]map[string]interface{}, error) {
	cacheKey := "card_admin_role::" + strconv.Itoa(roleId)
	menuCache, err := redisCache.Get(cacheKey)
	if err == nil {
		var menuCacheRet []map[string]interface{}
		json.Unmarshal([]byte(menuCache), &menuCacheRet)
		return menuCacheRet, nil
	}

	roleInfo, has := admin_model.GetRoleInfo("*", "id = ? AND is_del = ?", roleId, 0)
	if !has {
		return nil, errors.New("获取权限菜单失败")
	}
	var menus []map[string]interface{}
	if json.Unmarshal([]byte(roleInfo.MenuList), &menus) != nil {
		return nil, errors.New("权限菜单解析失败")
	}
	masterMenu := make(map[string]interface{})
	var menuList []map[string]interface{}
	var meta map[string]interface{}
	var subMeta map[string]interface{}
	var itemMeta map[string]interface{}
	for _, menu := range menus {
		masterMenu = map[string]interface{}{}
		id, _ := strconv.Atoi(menu["id"].(string))
		menuInfo, has := admin_model.GetMenuInfo("*", "id = ? AND status = ? AND is_del = ?", id, 0, 0)
		if !has {
			continue
		}
		meta = map[string]interface{}{
			"title":      menuInfo.Name,
			"icon":       menuInfo.Icon,
			"keepAlive":  true,
			"permission": []string{menuInfo.Key},
		}
		if strings.Contains(menuInfo.Path, "http") {
			meta["target"] = "_blank"
		}
		masterMenu["name"] = menuInfo.Key
		masterMenu["path"] = menuInfo.Path
		masterMenu["meta"] = meta
		if _, ok := menu["children"]; ok {
			masterMenu["children"] = []map[string]interface{}{}
			for _, subMenu := range menu["children"].([]interface{}) {
				subMenuItem := subMenu.(map[string]interface{})
				subMenuInfo, has := admin_model.GetMenuInfo("*", "id = ? AND status = ? AND is_del = ?", subMenuItem["id"], 0, 0)
				if !has {
					continue
				}
				if _, ok := masterMenu["redirect"]; !ok {
					masterMenu["redirect"] = subMenuInfo.Path
				}
				subMeta = map[string]interface{}{
					"title":      subMenuInfo.Name,
					"icon":       subMenuInfo.Icon,
					"keepAlive":  true,
					"permission": []string{menuInfo.Key},
				}
				if strings.Contains(subMenuInfo.Path, "http") {
					subMeta["target"] = "_blank"
				}
				subMenuItem["name"] = subMenuInfo.Key
				subMenuItem["path"] = subMenuInfo.Path
				subMenuItem["meta"] = subMeta
				delete(subMenuItem, "id")
				if _, ok := subMenuItem["children"]; ok {
					subChild := []map[string]interface{}{}
					for _, itemMenu := range subMenuItem["children"].([]interface{}) {
						itemMenuItem := itemMenu.(map[string]interface{})
						itemMenuInfo, has := admin_model.GetMenuInfo("*", "id = ? AND status = ? AND is_del = ?", itemMenuItem["id"], 0, 0)
						if !has {
							continue
						}
						if _, ok := subMenuItem["redirect"]; !ok {
							subMenuItem["redirect"] = subMenuInfo.Path
						}
						itemMeta = map[string]interface{}{
							"title":      itemMenuInfo.Name,
							"icon":       itemMenuInfo.Icon,
							"keepAlive":  true,
							"permission": []string{menuInfo.Key},
						}
						if strings.Contains(itemMenuInfo.Path, "http") {
							itemMeta["target"] = "_blank"
						}
						itemMenuItem["name"] = itemMenuInfo.Key
						itemMenuItem["path"] = itemMenuInfo.Path
						itemMenuItem["meta"] = itemMeta
						delete(itemMenuItem, "id")
						subChild = append(subChild, itemMenuItem)
					}
					subMenuItem["children"] = subChild
				}
				masterMenu["children"] = append(masterMenu["children"].([]map[string]interface{}), subMenuItem)
			}
		}
		menuList = append(menuList, masterMenu)
	}
	menuListJson, err := json.Marshal(menuList)
	err = redisCache.Set(cacheKey, menuListJson, 0)
	if err != nil {
		print(err.Error())
	}
	return menuList, nil
}
