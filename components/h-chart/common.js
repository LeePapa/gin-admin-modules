/**
 * Created by joylee on 2019/3/6.
 */
export default {
  color: ['#1890FF', '#F0CA2C', '#57C164', '#935ADF', '#E48746', '#E251A5', '#3B54B0', '#413D92'],
  getRandomArray (len = 24, max = 500) {
    return Array.from({ length: len }, () => Number((Math.random() * max).toFixed(2)))
  },
  getMockXAxis (len = 24) {
    let result = []
    for (let i = 0; i < len; i++) {
      result.push(i + ':00')
    }
    return result
  }
}
