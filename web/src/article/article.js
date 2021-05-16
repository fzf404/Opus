$(function () {
  // URL解析
  $.getUrlParam = function (name) {
    var reg = new RegExp('(^|&)' + name + '=([^&]*)(&|$)')
    var r = window.location.search.substr(1).match(reg)
    if (r != null) return unescape(r[2])
    return null
  }

  var artid = $.getUrlParam('artid')

  console.log(artid)
  if (artid == null) {
    window.location.href = '/'
    return
  }
  // 发起请求
  infoURL = 'http://127.0.0.1:8080/getart'
  $.ajax(infoURL, {
    method: 'POST',
    data: {
      artid: artid,
    },
    // 得到响应的处理
    success: (data) => {
      if (data.code != '200') {
        $('#userNULL').removeClass('d-none')
      }
      $('#headimg').attr('src', data.data.user.headimg)
      $('#username').text(data.data.user.name)
      $('#artid').text(data.data.article.artid)
      $('#arttype').text(data.data.article.arttype)
      $('#createTime').text(data.data.article.create)
      $('#updateTime').text(data.data.article.update)
      $('#likesg').text(data.data.article.likes)
      $('#share').text(data.data.article.share)
      $('#title').text(data.data.article.title)
      $('#artimg').attr('src', data.data.article.headimg)
      $('#content').html(data.data.article.content.replace('"', '"'))
    },
    error: () => {
      alert('服务器可能正在维护，请刷新重试或联系管理员')
    },
  })
})
