/**
 * Created by xuwei on 2017/7/12.
 *
 */
import ajax from '@service/ajax'

const urlCategoryList = '/site/cates'

export default {
  getCategoryList: (data = {}) => ajax({
    url: urlCategoryList,
    data: data,
    method: 'get'
  })
}
