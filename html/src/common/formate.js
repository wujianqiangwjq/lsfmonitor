function baseFormateDate(date, fmt){
    var o = {
        "M+": date.getMonth() +1,
        "d+": date.getDate(),
        "h+": date.getHours(),
        "m+": date.getMinutes(),
        "s+": date.getSeconds(),
        "q+": Math.floor((date.getMonth() + 3) / 3),
        "S": date.getMilliseconds()

    };
    if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (date.getFullYear() + "").substr(4 - RegExp.$1.length));
  for (var k in o)
    if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
  return fmt;
}

function formatDateTime(date, format) {
    if(!format)
      format = 'yyyy-MM-dd hh:mm';
    if(date) {
      return baseFormateDate(new Date(date*1000), format);
    }
    return '-';
  }
  export default {
      formatDateTime,
      baseFormateDate,
  }